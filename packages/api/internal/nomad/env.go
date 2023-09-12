package nomad

import (
	// trunk-ignore(semgrep/go.lang.security.audit.xss.import-text-template.import-text-template)
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"net/http"
	"text/template"

	"github.com/e2b-dev/api/packages/api/internal/api"
	nomadAPI "github.com/hashicorp/nomad/api"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

const (
	buildJobName          = "env-build"
	buildJobNameWithSlash = buildJobName + "/"
	envsDisk              = "/mnt/disks/fc-envs"
)

//go:embed env-build.hcl
var envBuildFile string
var envBuildTemplate = template.Must(template.New(buildJobName).Funcs(template.FuncMap{
	"escapeHCL": escapeHCL,
}).Parse(envBuildFile))

//go:embed provision-env.ubuntu.sh
var provisionEnvScriptFile string

func (n *NomadClient) StartBuildingEnv(
	t trace.Tracer,
	ctx context.Context,
	envID string,
	dockerContextPath string,
) (func() *nomadAPI.Allocation, error) {
	_, childSpan := t.Start(ctx, "build-env",
		trace.WithAttributes(
			attribute.String("env_id", envID),
		),
	)
	defer childSpan.End()

	traceID := childSpan.SpanContext().TraceID().String()
	spanID := childSpan.SpanContext().SpanID().String()

	childSpan.SetAttributes(
		attribute.String("passed_trace_id_hex", traceID),
		attribute.String("passed_span_id_hex", spanID),
	)

	var jobDef bytes.Buffer

	jobVars := struct {
		EnvID             string
		ProvisionScript   string
		DockerContextPath string
		SpanID            string
		TraceID           string
		TaskName          string
		JobName           string
		EnvsDisk          string
	}{
		SpanID:            spanID,
		DockerContextPath: dockerContextPath,
		ProvisionScript:   provisionEnvScriptFile,
		TraceID:           traceID,
		EnvID:             envID,
		TaskName:          defaultTaskName,
		JobName:           buildJobName,
		EnvsDisk:          envsDisk,
	}

	err := envBuildTemplate.Execute(&jobDef, jobVars)
	if err != nil {
		return nil, &api.APIError{
			Msg:       fmt.Sprintf("failed to `envBuildJobTemp.Execute()`: %+v", err),
			ClientMsg: "Cannot build env right now",
			Code:      http.StatusInternalServerError,
		}
	}

	job, err := n.client.Jobs().ParseHCL(jobDef.String(), false)
	if err != nil {
		return nil, &api.APIError{
			Msg:       fmt.Sprintf("failed to parse the HCL job file %+s: %w", jobDef.String(), err),
			ClientMsg: "Cannot create env build job right now",
			Code:      http.StatusInternalServerError,
		}
	}

	res, _, err := n.client.Jobs().Register(job, nil)
	if err != nil {
		fmt.Printf("Failed to register '%s%s' job: %+v", buildJobNameWithSlash, jobVars.EnvID, err)

		return nil, &api.APIError{
			Msg:       err.Error(),
			ClientMsg: "Cannot create env build job right now",
			Code:      http.StatusInternalServerError,
		}
	}

	meta := res.QueryMeta
	evalID := res.EvalID
	index := res.JobModifyIndex

	jobInfo := JobInfo{
		name:   buildJobNameWithSlash + envID,
		evalID: evalID,
		index:  index,
	}

	_, err = n.WaitForJobStart(
		ctx,
		jobInfo,
		meta,
	)
	if err != nil {
		apiErr := n.DeleteInstance(envID, false)
		if apiErr != nil {
			fmt.Printf("error in cleanup after failing to create instance of environment '%s':%+v", envID, apiErr.Msg)
		}

		return nil, &api.APIError{
			Msg:       err.Error(),
			ClientMsg: "Cannot create a environment instance right now",
			Code:      http.StatusInternalServerError,
		}
	}

	waitForFinish := func() *nomadAPI.Allocation {
		alloc, finishErr := n.WaitForJobFinish(
			ctx,
			jobInfo,
			meta,
		)
		if finishErr != nil {
			// TODO: Cleanup
			fmt.Printf("error waiting for env '%s' build: %+w", envID, finishErr)
		}

		return alloc
	}

	return waitForFinish, nil
}
