package middlewares_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"core-gin/api/middlewares"
	"core-gin/infrastructure"
	"core-gin/lib"

	"github.com/gin-gonic/gin"
)

var _ = Describe("Module", func() {
	var (
		env                 *lib.Env
		router              infrastructure.Router
		logger              lib.Logger
		corsMiddleware      middlewares.CorsMiddleware
		metricsMiddleware   middlewares.MetricsMiddleware
		rateLimitMiddleware middlewares.RateLimitMiddleware
		inOutMiddleware     middlewares.InOutMiddleware
		swaggerMiddleware   middlewares.SwaggerMiddleware
		middlewaresArray    middlewares.Middlewares
	)

	BeforeEach(func() {
		env = &lib.Env{}
		router = infrastructure.Router{
			Engine: gin.New(),
		}
		logger = lib.GetLogger()
		corsMiddleware = middlewares.NewCorsMiddleware(router, logger)
		metricsMiddleware = middlewares.NewMetricsMiddleware(router, logger)
		rateLimitMiddleware = middlewares.NewRateLimitMiddleware(logger, env, router)
		inOutMiddleware = middlewares.NewInOutMiddlewareMiddleware(logger, router, env)
		swaggerMiddleware = middlewares.NewSwaggerMiddleware(router, logger)
		middlewaresArray = middlewares.NewMiddlewares(corsMiddleware, metricsMiddleware, rateLimitMiddleware, inOutMiddleware, swaggerMiddleware)
	})

	It("creates a slice of middlewares with the correct length and element types", func() {
		Expect(middlewaresArray).To(HaveLen(5))
		Expect(middlewaresArray[0]).To(BeAssignableToTypeOf(corsMiddleware))
		Expect(middlewaresArray[1]).To(BeAssignableToTypeOf(metricsMiddleware))
		Expect(middlewaresArray[2]).To(BeAssignableToTypeOf(rateLimitMiddleware))
		Expect(middlewaresArray[3]).To(BeAssignableToTypeOf(inOutMiddleware))
		Expect(middlewaresArray[4]).To(BeAssignableToTypeOf(swaggerMiddleware))
	})
})
