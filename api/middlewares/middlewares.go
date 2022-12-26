package middlewares

import "go.uber.org/fx"

// Module Middleware exported
var Module = fx.Options(
	fx.Provide(NewCorsMiddleware),
	fx.Provide(NewMetricsMiddleware),
	fx.Provide(NewDBTransactionMiddleware),
	fx.Provide(NewPaginationMiddleware),
	fx.Provide(NewRateLimitMiddleware),
	fx.Provide(NewInOutMiddlewareMiddleware),
	fx.Provide(NewSwaggerMiddleware),
	fx.Provide(NewMiddlewares),
)

// IMiddleware middleware interface
type IMiddleware interface {
	Setup()
}

// Middlewares contains multiple middleware
type Middlewares []IMiddleware

// NewMiddlewares creates new middlewares
// Register the middleware that should be applied directly (globally)
func NewMiddlewares(
	corsMiddleware CorsMiddleware,
	metricsMiddleware MetricsMiddleware,
	rateLimitMiddleware RateLimitMiddleware,
	inOutMiddleware InOutMiddleware,
	swaggerMiddleware SwaggerMiddleware,
) Middlewares {
	return Middlewares{
		corsMiddleware,
		metricsMiddleware,
		rateLimitMiddleware,
		inOutMiddleware,
		swaggerMiddleware,
	}
}

// Setup sets up middlewares
func (m Middlewares) Setup() {
	for _, middleware := range m {
		middleware.Setup()
	}
}
