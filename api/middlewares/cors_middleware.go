package middlewares

import (
	"time"

	"core-gin/infrastructure"
	"core-gin/lib"

	"github.com/gin-contrib/cors"
)

// CorsMiddleware middleware for cors
type CorsMiddleware struct {
	router infrastructure.Router
	logger lib.Logger
}

// NewCorsMiddleware creates new cors middleware
func NewCorsMiddleware(router infrastructure.Router, logger lib.Logger) CorsMiddleware {
	return CorsMiddleware{
		router: router,
		logger: logger,
	}
}

// Setup sets up cors middleware
func (m CorsMiddleware) Setup() {
	m.logger.Info("Setting up cors middleware")

	m.router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Accept", "Access-Control-Allow-Headers", "Authorization", "X-App-Token", "Webauthn-Session", "X-Client-Id", "X-Client-Version"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		// default allow all origins
		AllowAllOrigins: true,
	}))
}
