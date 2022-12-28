package infrastructure

import (
	"core-gin/lib"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type ITracer interface {
	trace.Tracer
}

func NewTracer(
	env *lib.Env,
) ITracer {
	tracer := otel.Tracer(env.ServiceName)

	return tracer
}
