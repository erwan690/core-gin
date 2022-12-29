package middlewares

import (
	"testing"

	"core-gin/infrastructure"
	"core-gin/lib"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestModule(t *testing.T) {
	env := &lib.Env{}
	router := infrastructure.Router{
		Engine: gin.New(),
	}
	logger := lib.GetLogger()
	corsMiddleware := NewCorsMiddleware(router, logger)
	metricsMiddleware := NewMetricsMiddleware(router, logger)
	rateLimitMiddleware := NewRateLimitMiddleware(logger, env, router)
	inOutMiddleware := NewInOutMiddlewareMiddleware(logger, router, env)
	swaggerMiddleware := NewSwaggerMiddleware(router, logger)

	middlewares := NewMiddlewares(corsMiddleware, metricsMiddleware, rateLimitMiddleware, inOutMiddleware, swaggerMiddleware)

	assert.Equal(t, 5, len(middlewares))
	assert.IsType(t, corsMiddleware, middlewares[0])
	assert.IsType(t, metricsMiddleware, middlewares[1])
	assert.IsType(t, rateLimitMiddleware, middlewares[2])
	assert.IsType(t, inOutMiddleware, middlewares[3])
	assert.IsType(t, swaggerMiddleware, middlewares[4])
}
