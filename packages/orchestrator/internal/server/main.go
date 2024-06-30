package server

import (
	"context"
	"fmt"
	"log"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	consulapi "github.com/hashicorp/consul/api"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc/filters"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/e2b-dev/infra/packages/orchestrator/internal/constants"
	"github.com/e2b-dev/infra/packages/orchestrator/internal/consul"
	"github.com/e2b-dev/infra/packages/orchestrator/internal/dns"
	"github.com/e2b-dev/infra/packages/orchestrator/internal/pool"
	"github.com/e2b-dev/infra/packages/orchestrator/internal/sandbox"
	"github.com/e2b-dev/infra/packages/shared/pkg/grpc/orchestrator"
	"github.com/e2b-dev/infra/packages/shared/pkg/logging"
	"github.com/e2b-dev/infra/packages/shared/pkg/smap"
	"github.com/e2b-dev/infra/packages/shared/pkg/telemetry"
)

const (
	ipSlotConcurrency = 4
	ipSlotPoolSize    = 25
)

type server struct {
	orchestrator.UnimplementedSandboxServer
	sandboxes   *smap.Map[*sandbox.Sandbox]
	dns         *dns.DNS
	tracer      trace.Tracer
	consul      *consulapi.Client
	networkPool *pool.Pool[*sandbox.IPSlot]
}

func New(logger *zap.Logger) *grpc.Server {
	opts := []grpc_zap.Option{logging.WithoutHealthCheck()}

	s := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler(otelgrpc.WithInterceptorFilter(filters.Not(filters.HealthCheck())))),
		grpc.ChainUnaryInterceptor(
			grpc_zap.UnaryServerInterceptor(logger, opts...),
			recovery.UnaryServerInterceptor(),
		),
	)

	log.Println("Initializing orchestrator server")

	ctx := context.Background()

	dns := dns.New()
	go dns.Start("127.0.0.1:53")

	tracer := otel.Tracer(constants.ServiceName)

	consulClient, err := consul.New(ctx)
	if err != nil {
		panic(err)
	}

	createNetwork := func() (*sandbox.IPSlot, error) {
		ips, err := sandbox.NewSlot(ctx, tracer, consulClient)

		err = ips.CreateNetwork(ctx, tracer)
		if err != nil {
			errMsg := fmt.Errorf("failed to create namespaces: %w", err)
			telemetry.ReportCriticalError(ctx, errMsg)

			return nil, errMsg
		}

		return ips, err
	}

	networkPool := pool.New[*sandbox.IPSlot](ipSlotPoolSize)

	go func() {
		err := networkPool.Populate(
			ctx,
			ipSlotConcurrency,
			createNetwork,
		)
		if err != nil {
			logger.Fatal("failed to populate network pool", zap.Error(err))
			panic(err)
		}
	}()

	orchestrator.RegisterSandboxServer(s, &server{
		tracer:      tracer,
		consul:      consulClient,
		dns:         dns,
		sandboxes:   smap.New[*sandbox.Sandbox](),
		networkPool: networkPool,
	})

	grpc_health_v1.RegisterHealthServer(s, health.NewServer())

	return s
}