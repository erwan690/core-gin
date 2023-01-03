package middlewares_test

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"core-gin/api/middlewares"
	"core-gin/infrastructure"
	"core-gin/lib"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var _ = Describe("MetricsMiddleware", func() {
	var (
		router            infrastructure.Router
		logger            lib.Logger
		metricsMiddleware middlewares.MetricsMiddleware
	)

	BeforeEach(func() {
		router = infrastructure.Router{
			Engine: gin.New(),
		}
		logger = lib.GetLogger()
		metricsMiddleware = middlewares.NewMetricsMiddleware(router, logger)
	})

	It("exposes the Prometheus metrics through the /metrics endpoint", func() {
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
		Expect(w.Body.String()).To(ContainSubstring("test_counter"))
	})
})
