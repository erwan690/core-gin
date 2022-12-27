package infrastructure

import (
	"core-gin/lib"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type Tracer struct {
	trace.Tracer
}

func NewTracer(
	env *lib.Env,
) Tracer {
	tracer := otel.Tracer(env.ServiceName)

	return Tracer{
		tracer,
	}
}
