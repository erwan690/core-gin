package routes

import (
	"core-gin/infrastructure"
	"core-gin/internal/handlers"
)

// HealthRoutes struct
type HealthRoutes struct {
	route   infrastructure.Router
	handler handlers.HealthHandler
}

// NewHealthRoutes creates new user controller
func NewHealthRoutes(
	route infrastructure.Router,
	handler handlers.HealthHandler,
) HealthRoutes {
	return HealthRoutes{
		handler: handler,
		route:   route,
	}
}

// Setup user routes
func (s HealthRoutes) Setup() {
	s.route.GET("/health", s.handler.Health)
}
