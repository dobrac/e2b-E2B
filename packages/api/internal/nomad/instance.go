package nomad

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/e2b-dev/infra/packages/api/internal/api"
	"github.com/e2b-dev/infra/packages/api/internal/utils"
	"github.com/e2b-dev/infra/packages/shared/pkg/telemetry"
	"github.com/google/uuid"

	nomadAPI "github.com/hashicorp/nomad/api"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

const (
	instanceJobName          = "env-instance"
	instanceJobNameWithSlash = instanceJobName + "/"
	instanceIDPrefix         = "i"

	instanceStartTimeout = time.Second * 20

	envIDMetaKey      = "ENV_ID"
	instanceIDMetaKey = "INSTANCE_ID"
	teamIDMetaKey     = "TEAM_ID"
	metaKey           = "META"
)

var (
	logsProxyAddress = os.Getenv("LOGS_PROXY_ADDRESS")
	consulToken      = os.Getenv("CONSUL_TOKEN")
)

//go:embed env-instance.hcl
var envInstanceFile string
var envInstanceTemplate = template.Must(template.New(instanceJobName).Parse(envInstanceFile))

func (n *NomadClient) GetInstances() ([]*InstanceInfo, *api.APIError) {
	jobs, _, err := n.client.Jobs().ListOptions(&nomadAPI.JobListOptions{
		Fields: &nomadAPI.JobListFields{
			Meta: true,
		},
	}, &nomadAPI.QueryOptions{
		Filter: fmt.Sprintf("ID contains \"%s\" and Status == \"%s\"", instanceJobNameWithSlash, jobRunningStatus),
	})
	if err != nil {
		errMsg := fmt.Errorf("failed to get jobs from Nomad %w", err)

		return nil, &api.APIError{
			Err:       errMsg,
			ClientMsg: "Cannot retrieve instances right now",
			Code:      http.StatusInternalServerError,
		}
	}

	allocations, _, err := n.client.Allocations().List(&nomadAPI.QueryOptions{
		Filter: fmt.Sprintf("JobID contains \"%s\"", instanceJobNameWithSlash),
	})
	if err != nil {
		errMsg := fmt.Errorf("failed to get allocations from Nomad %w", err)

		return nil, &api.APIError{
			Err:       errMsg,
			ClientMsg: "Cannot retrieve instances right now",
			Code:      http.StatusInternalServerError,
		}
	}

	nodeMap := make(map[string]string, len(allocations))
	for _, alloc := range allocations {
		nodeMap[alloc.JobID] = alloc.NodeID[:shortNodeIDLength]
	}

	instances := make([]*InstanceInfo, 0, len(jobs))

	for _, job := range jobs {
		instanceID := job.Meta[instanceIDMetaKey]
		envID := job.Meta[envIDMetaKey]
		teamID := job.Meta[teamIDMetaKey]
		metadataRaw := job.Meta[metaKey]
		metadata := make(map[string]string)

		err = json.Unmarshal([]byte(metadataRaw), &metadata)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to unmarshal metadata for job '%s': %v\n", job.ID, err)
		}

		var teamUUID *uuid.UUID

		if teamID != "" {
			parsedTeamID, parseErr := uuid.Parse(teamID)
			if parseErr != nil {
				fmt.Fprintf(os.Stderr, "failed to parse team ID '%s' for job '%s': %v\n", teamID, job.ID, parseErr)
			} else {
				teamUUID = &parsedTeamID
			}
		}

		clientID, ok := nodeMap[job.ID]
		if !ok {
			fmt.Fprintf(os.Stderr, "failed to get client ID for job '%s'\n", job.ID)
		}

		instances = append(instances, &InstanceInfo{
			Instance: &api.Instance{
				InstanceID: instanceID,
				EnvID:      envID,
				ClientID:   clientID,
			},
			TeamID:   teamUUID,
			Metadata: metadata,
		})
	}

	return instances, nil
}

func (n *NomadClient) CreateInstance(
	t trace.Tracer,
	ctx context.Context,
	envID, teamID string,
	metadata map[string]string,
) (*api.Instance, *api.APIError) {
	childCtx, childSpan := t.Start(ctx, "create-instance",
		trace.WithAttributes(
			attribute.String("env.id", envID),
		),
	)
	defer childSpan.End()

	instanceID := instanceIDPrefix + utils.GenerateID()

	traceID := childSpan.SpanContext().TraceID().String()
	spanID := childSpan.SpanContext().SpanID().String()

	metadataSerialized, err := json.Marshal(metadata)
	if err != nil {
		errMsg := fmt.Errorf("failed to marshal metadata: %w", err)

		return nil, &api.APIError{
			Err:       errMsg,
			ClientMsg: "Cannot create a environment instance right now",
			Code:      http.StatusInternalServerError,
		}
	}

	telemetry.SetAttributes(
		childCtx,
		attribute.String("passed_trace_id_hex", traceID),
		attribute.String("passed_span_id_hex", spanID),
	)

	var jobDef bytes.Buffer

	jobVars := struct {
		SpanID           string
		ConsulToken      string
		TraceID          string
		EnvID            string
		InstanceID       string
		LogsProxyAddress string
		TaskName         string
		JobName          string
		EnvsDisk         string
		TeamID           string
		EnvIDKey         string
		InstanceIDKey    string
		TeamIDKey        string
		MetadataKey      string
		Metadata         string
	}{
		TeamIDKey:        teamIDMetaKey,
		EnvIDKey:         envIDMetaKey,
		InstanceIDKey:    instanceIDMetaKey,
		MetadataKey:      metaKey,
		SpanID:           spanID,
		TeamID:           teamID,
		TraceID:          traceID,
		LogsProxyAddress: logsProxyAddress,
		ConsulToken:      consulToken,
		EnvID:            envID,
		InstanceID:       instanceID,
		TaskName:         defaultTaskName,
		JobName:          instanceJobName,
		EnvsDisk:         envsDisk,
		Metadata:         strings.ReplaceAll(string(metadataSerialized), "\"", "\\\""),
	}

	err = envInstanceTemplate.Execute(&jobDef, jobVars)
	if err != nil {
		errMsg := fmt.Errorf("failed to `envInstanceJobTemp.Execute()`: %w", err)

		return nil, &api.APIError{
			Err:       errMsg,
			ClientMsg: "Cannot create a environment instance right now",
			Code:      http.StatusInternalServerError,
		}
	}

	job, err := n.client.Jobs().ParseHCL(jobDef.String(), false)
	if err != nil {
		errMsg := fmt.Errorf("failed to parse the HCL job file %+s: %w", jobDef.String(), err)

		return nil, &api.APIError{
			Err:       errMsg,
			ClientMsg: "Cannot create a environment instance right now",
			Code:      http.StatusInternalServerError,
		}
	}

	sub := n.newSubscriber(*job.ID, taskRunningState, defaultTaskName)
	defer sub.close()

	_, _, err = n.client.Jobs().Register(job, nil)
	if err != nil {
		errMsg := fmt.Errorf("failed to register '%s%s' job: %w", instanceJobNameWithSlash, jobVars.InstanceID, err)

		return nil, &api.APIError{
			Err:       errMsg,
			ClientMsg: "Cannot create a environment instance right now",
			Code:      http.StatusInternalServerError,
		}
	}

	telemetry.ReportEvent(childCtx, "Started waiting for job to start")

	select {
	case <-ctx.Done():
		errMsg := fmt.Errorf("error waiting for env '%s' instance: %w", envID, ctx.Err())

		delErr := n.DeleteInstance(instanceID, false)
		if delErr != nil {
			cleanupErr := fmt.Errorf("error in cleanup after failing to create instance of environment error: %w: %w", delErr.Err, errMsg)

			return nil, &api.APIError{
				Err:       cleanupErr,
				ClientMsg: "Cannot create a environment instance right now",
				Code:      http.StatusInternalServerError,
			}
		}

		return nil, &api.APIError{
			Err:       errMsg,
			ClientMsg: "Cannot create a environment instance right now",
			Code:      http.StatusInternalServerError,
		}
	case <-time.After(instanceStartTimeout):
		errMsg := fmt.Errorf("failed to create instance of environment '%s'", envID)

		delErr := n.DeleteInstance(instanceID, false)
		if delErr != nil {
			cleanupErr := fmt.Errorf("error in cleanup after failing to create instance of environment error: %w: %w", delErr.Err, errMsg)

			return nil, &api.APIError{
				Err:       cleanupErr,
				ClientMsg: "Cannot create a environment instance right now",
				Code:      http.StatusInternalServerError,
			}
		}

		return nil, &api.APIError{
			Err:       errMsg,
			ClientMsg: "Cannot create a environment instance right now",
			Code:      http.StatusInternalServerError,
		}
	case alloc := <-sub.wait:
		var allocErr error

		defer func() {
			if allocErr != nil {
				cleanupErr := n.DeleteInstance(*job.ID, false)
				if cleanupErr != nil {
					errMsg := fmt.Errorf("error in cleanup after failing to create instance of environment '%s': %w", envID, cleanupErr.Err)
					telemetry.ReportError(childCtx, errMsg)
				} else {
					telemetry.ReportEvent(childCtx, "cleaned up env instance job", attribute.String("env.id", envID), attribute.String("instance.id", instanceID))
				}
			}
		}()

		if alloc.TaskStates == nil {
			allocErr = fmt.Errorf("task states are nil")

			telemetry.ReportCriticalError(childCtx, allocErr)

			return nil, &api.APIError{
				Err:       allocErr,
				ClientMsg: "Cannot create a environment instance right now",
				Code:      http.StatusInternalServerError,
			}
		}

		if alloc.TaskStates[defaultTaskName] == nil {
			allocErr = fmt.Errorf("task state '%s' is nil", defaultTaskName)

			telemetry.ReportCriticalError(childCtx, allocErr)

			return nil, &api.APIError{
				Err:       allocErr,
				ClientMsg: "Cannot create a environment instance right now",
				Code:      http.StatusInternalServerError,
			}
		}

		task := alloc.TaskStates[defaultTaskName]

		var instanceErr error

		if task.Failed {
			for _, event := range task.Events {
				if event.Type == "Terminated" {
					instanceErr = fmt.Errorf("%s", event.Message)
				}
			}

			if instanceErr == nil {
				allocErr = fmt.Errorf("starting instance failed")
			} else {
				allocErr = fmt.Errorf("starting instance failed %w", instanceErr)
			}

			telemetry.ReportCriticalError(childCtx, allocErr)

			return nil, &api.APIError{
				Err:       allocErr,
				ClientMsg: "Cannot create a environment instance right now",
				Code:      http.StatusInternalServerError,
			}
		}

		// We accept pending state as well because it means that the task is starting (and the env exists because we checked the DB)
		// This usually speeds up the start from client
		if task.State != taskRunningState && task.State != taskPendingState {
			allocErr = fmt.Errorf("task state is not '%s' - it is '%s'", taskRunningState, task.State)

			telemetry.ReportCriticalError(childCtx, allocErr)

			return nil, &api.APIError{
				Err:       allocErr,
				ClientMsg: "Cannot create a environment instance right now",
				Code:      http.StatusInternalServerError,
			}
		}

		return &api.Instance{
			ClientID:   strings.Clone(alloc.NodeID[:shortNodeIDLength]),
			InstanceID: instanceID,
			EnvID:      envID,
		}, nil
	}
}

func (n *NomadClient) DeleteInstance(instanceID string, purge bool) *api.APIError {
	_, _, err := n.client.Jobs().Deregister(instanceJobNameWithSlash+instanceID, purge, nil)
	if err != nil {
		errMsg := fmt.Errorf("cannot delete job '%s%s' job: %w", instanceJobNameWithSlash, instanceID, err)

		return &api.APIError{
			Err:       errMsg,
			ClientMsg: "Cannot delete the environment instance right now",
			Code:      http.StatusInternalServerError,
		}
	}

	return nil
}
