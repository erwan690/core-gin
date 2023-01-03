package middlewares_test

import (
	"net/http"
	"net/http/httptest"

	"core-gin/api/middlewares"
	"core-gin/infrastructure"
	"core-gin/lib"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("RateLimitMiddleware", func() {
	var (
		logger     lib.Logger
		env        *lib.Env
		router     infrastructure.Router
		middleware middlewares.RateLimitMiddleware
		req        *http.Request
		recorder   *httptest.ResponseRecorder
	)

	BeforeEach(func() {
		// Create a new gin router
		// Create a new RateLimitMiddleware
		logger = lib.GetLogger()
		env = &lib.Env{
			RateLimitPeriod:   1,
			RateLimitRequests: 5,
		}
		router = infrastructure.Router{
			Engine: gin.New(),
		}
		middleware = middlewares.NewRateLimitMiddleware(logger, env, router)

		// Add the rate limit middleware to the router
		middleware.Setup()

		// Define a test handler
		router.GET("/test", func(c *gin.Context) {
			c.String(http.StatusOK, "test")
		})

		// Create a new HTTP request
		var err error
		req, err = http.NewRequest("GET", "/test", nil)
		Expect(err).NotTo(HaveOccurred())

		// Create a new recorder for the response
		recorder = httptest.NewRecorder()
	})

	It("returns a 200 status code and the correct body", func() {
		// Execute the test handler with the middleware
		router.ServeHTTP(recorder, req)

		// Assert the response status code is 200
		Expect(recorder.Code).To(Equal(http.StatusOK))

		// Assert the response body is "test"
		Expect(recorder.Body.String()).To(Equal("test"))
	})

	It("sets the custom headers", func() {
		// Execute the test handler with the middleware
		router.ServeHTTP(recorder, req)

		// Assert the custom headers are set
		Expect(recorder.Header().Get("X-RateLimit-Limit")).To(Equal("5"))
		// equal to 3 because it has call in it before
		Expect(recorder.Header().Get("X-RateLimit-Remaining")).To(Equal("3"))
		Expect(recorder.Header().Get("X-RateLimit-Reset")).To(Not(BeEmpty()))
	})

	It("sets the reached field on the context", func() {
		// Execute the test handler with the middleware multiple times
		for i := 0; i < 5; i++ {
			router.ServeHTTP(recorder, req)
		}

		// Get the last context used in the middleware
		Expect(recorder.Header().Get("X-RateLimit-Limit")).To(Equal("5"))
		// equal to 3 because it has call in it before
		Expect(recorder.Header().Get("X-RateLimit-Remaining")).To(Equal("0"))
		Expect(recorder.Header().Get("X-RateLimit-Reset")).To(Not(BeEmpty()))
	})
})
