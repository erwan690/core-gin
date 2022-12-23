package infrastructure

import (
	"context"
	"time"

	"core-gin/lib"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktracer "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Otel struct{}

func NewOtel(
	env *lib.Env,
	logger lib.Logger,
) Otel {
	// if disable do nothing
	if !env.OtelEnable {
		return Otel{}
	}

	ctx := context.Background()
	sctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	secureOption := otlptracegrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))
	if env.InsecureMode {
		secureOption = otlptracegrpc.WithInsecure()
	}

	exporter, err := otlptrace.New(
		sctx,
		otlptracegrpc.NewClient(
			secureOption,
			otlptracegrpc.WithEndpoint(env.OtelEndpoint),
			otlptracegrpc.WithDialOption(grpc.WithBlock()),
		),
	)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Infof("Otel connection established : %s", otel.Version())

	resources, err := resource.New(
		sctx,
		resource.WithAttributes(
			attribute.String("service.name", env.ServiceName),
			attribute.String("library.language", "go"),
		),
	)
	if err != nil {
		logger.Infof("Could not set resources: %s", err)
	}

	otel.SetTracerProvider(
		sdktracer.NewTracerProvider(
			sdktracer.WithSampler(sdktracer.AlwaysSample()),
			sdktracer.WithBatcher(exporter),
			sdktracer.WithResource(resources),
		),
	)

	// setup list Of Middleware
	defer exporter.Shutdown(sctx)

	return Otel{}
}
