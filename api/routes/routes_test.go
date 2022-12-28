package routes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockRoute struct {
	SetupCalled bool
}

func (m *MockRoute) Setup() {
	m.SetupCalled = true
}

func TestRoutes(t *testing.T) {
	// Create a mock health route
	mockHealthRoute := new(MockRoute)

	// Call the NewRoutes function
	routes := NewRoutes(mockHealthRoute)

	// Assert that the returned value is a Routes instance with the correct route
	assert.Len(t, routes, 1)
	assert.Equal(t, mockHealthRoute, routes[0])

	// Call the Setup method
	routes.Setup()

	// Assert that the Setup method was called on the mock health route
	assert.True(t, mockHealthRoute.SetupCalled)
}
