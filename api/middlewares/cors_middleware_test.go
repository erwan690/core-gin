package middlewares_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"core-gin/api/middlewares"
	"core-gin/infrastructure"
	"core-gin/lib"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var _ = Describe("CorsMiddleware", func() {
	var (
		router         infrastructure.Router
		logger         lib.Logger
		corsMiddleware middlewares.CorsMiddleware
	)

	BeforeEach(func() {
		router = infrastructure.Router{
			Engine: gin.New(),
		}
		logger = lib.GetLogger()
		corsMiddleware = middlewares.NewCorsMiddleware(router, logger)
	})

	It("adds the correct middleware to the router", func() {
		corsMiddleware.Setup()
		Expect(router.Handlers).To(HaveLen(1))
		Expect(router.Handlers[0]).To(BeAssignableToTypeOf(cors.Default()))
	})
})
