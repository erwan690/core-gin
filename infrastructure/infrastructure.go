package infrastructure

import "go.uber.org/fx"

// Module exports dependency
var Module = fx.Options(
	fx.Provide(NewRouter),
	fx.Provide(NewDatabase),
	fx.Provide(NewOtel),
	fx.Provide(NewTracer),
	fx.Provide(NewSlack),
)
