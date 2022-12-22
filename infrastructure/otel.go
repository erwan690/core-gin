package infrastructure

import (
	"context"

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

type Otel struct {
	*otlptrace.Exporter
}

func NewOtel(
	env *lib.Env,
	logger lib.Logger,
) Otel {
	secureOption := otlptracegrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))
	if env.InsecureMode {
		secureOption = otlptracegrpc.WithInsecure()
	}

	exporter, err := otlptrace.New(
		context.Background(),
		otlptracegrpc.NewClient(
			secureOption,
			otlptracegrpc.WithEndpoint(env.OtelEndpoint),
			otlptracegrpc.WithDialOption(grpc.WithBlock()),
		),
	)
	if err != nil {
		logger.Fatal(err)
	}
	resources, err := resource.New(
		context.Background(),
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

	logger.Infof("Otel connection established : %s", otel.Version())

	return Otel{
		exporter,
	}
}
