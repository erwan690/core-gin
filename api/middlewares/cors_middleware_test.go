package middlewares

import (
	"testing"

	"core-gin/infrastructure"
	"core-gin/lib"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// TestCorsMiddleware is a test suite for the CorsMiddleware struct
type TestCorsMiddleware struct {
	suite.Suite
}

// TestSetup tests the Setup method of the CorsMiddleware struct
func (suite *TestCorsMiddleware) TestSetup() {
	router := infrastructure.Router{
		Engine: gin.New(),
	}
	logger := lib.GetLogger()
	corsMiddleware := NewCorsMiddleware(router, logger)

	corsMiddleware.Setup()

	// Verify that the correct middleware was added to the router
	middlewares := router.Handlers
	suite.Equal(1, len(middlewares))
	assert.IsType(suite.T(), cors.Default(), middlewares[0])
}

// TestCorsMiddlewareSuite runs the TestCorsMiddleware test suite
func TestCorsMiddlewareSuite(t *testing.T) {
	suite.Run(t, new(TestCorsMiddleware))
}
