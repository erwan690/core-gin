package middlewares

import (
	"core-gin/infrastructure"
	"core-gin/lib"

	"github.com/gin-contrib/gzip"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "core-gin/docs"
)

// SwaggerMiddleware middleware for documentation
type SwaggerMiddleware struct {
	router infrastructure.Router
	logger lib.Logger
}

// NewSwaggerMiddleware creates new documentation middleware
func NewSwaggerMiddleware(router infrastructure.Router, logger lib.Logger) SwaggerMiddleware {
	return SwaggerMiddleware{
		router: router,
		logger: logger,
	}
}

// Setup sets up documentation middleware
func (m SwaggerMiddleware) Setup() {
	m.logger.Info("Setting up documentation middleware")

	m.router.Use(gzip.Gzip(gzip.BestSpeed))

	m.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
