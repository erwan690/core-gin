package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"core-gin/infrastructure"
	"core-gin/lib"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// TestMetricsMiddleware is a test suite for the MetricsMiddleware struct
type TestMetricsMiddleware struct {
	suite.Suite
}

// TestSetup tests the Setup method of the MetricsMiddleware struct
func (suite *TestMetricsMiddleware) TestSetup() {
	router := infrastructure.Router{
		Engine: gin.New(),
	}
	logger := lib.GetLogger()
	metricsMiddleware := NewMetricsMiddleware(router, logger)

	metricsMiddleware.Setup()

	// Register a Prometheus collector for testing purposes
	c := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "test_counter",
		Help: "Test counter for testing purposes.",
	})
	prometheus.MustRegister(c)

	// Send a request to the /metrics endpoint and check that the test counter is present in the response
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/metrics", nil)
	router.ServeHTTP(w, req)
	assert.Contains(suite.T(), w.Body.String(), "test_counter")
}

// TestMetricsMiddlewareSuite runs the TestMetricsMiddleware test suite
func TestMetricsMiddlewareSuite(t *testing.T) {
	suite.Run(t, new(TestMetricsMiddleware))
}
