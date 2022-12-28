package routes

import (
	"testing"

	"core-gin/infrastructure"
	"core-gin/lib"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockHealthHandler struct{}

func (m *MockHealthHandler) Health(c *gin.Context) {
	// Do nothing, this is a mock implementation
}

func TestHealthRoutes_Setup(t *testing.T) {
	// Create a router and a mock handler
	router := infrastructure.NewRouter(&lib.Env{})
	mockHandler := new(MockHealthHandler)

	// Create a new HealthRoutes instance
	routes := NewHealthRoutes(router, mockHandler)

	// Call the Setup method
	routes.Setup()

	// Assert that the router has the correct route registered
	assert.Len(t, router.Routes(), 1)
	route := router.Routes()[0]
	assert.Equal(t, "GET", route.Method)
	assert.Equal(t, "/health", route.Path)
}
