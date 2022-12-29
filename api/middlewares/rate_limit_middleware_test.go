package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"core-gin/infrastructure"
	"core-gin/lib"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRateLimitMiddleware(t *testing.T) {
	// Create a new gin router
	// Create a new RateLimitMiddleware
	logger := lib.GetLogger()
	env := &lib.Env{
		RateLimitPeriod:   1,
		RateLimitRequests: 5,
	}
	router := infrastructure.Router{
		Engine: gin.New(),
	}
	middleware := NewRateLimitMiddleware(logger, env, router)

	// Add the rate limit middleware to the router
	middleware.Setup()

	// Define a test handler
	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "test")
	})

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/test", nil)
	assert.NoError(t, err)

	// Create a new recorder for the response
	recorder := httptest.NewRecorder()

	// Execute the test handler with the middleware
	router.ServeHTTP(recorder, req)

	// Assert the response status code is 200
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Assert the response body is "test"
	assert.Equal(t, "test", recorder.Body.String())

	// Assert the custom headers are set
	assert.Equal(t, "5", recorder.Header().Get("X-RateLimit-Limit"))
	assert.Equal(t, "4", recorder.Header().Get("X-RateLimit-Remaining"))
	assert.NotEmpty(t, recorder.Header().Get("X-RateLimit-Reset"))
}
