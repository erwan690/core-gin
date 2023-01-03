package routes_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"core-gin/api/routes"
	"core-gin/infrastructure"

	"github.com/gin-gonic/gin"
)

type MockHealthHandler struct{}

func (m *MockHealthHandler) Health(c *gin.Context) {
	// Do nothing, this is a mock implementation
}

var _ = Describe("HealthRoutes", func() {
	var (
		router       infrastructure.Router
		mockHandler  *MockHealthHandler
		healthRoutes routes.IHealthRoutes
	)

	BeforeEach(func() {
		router = infrastructure.Router{
			Engine: gin.New(),
		}
		mockHandler = new(MockHealthHandler)
		healthRoutes = routes.NewHealthRoutes(router, mockHandler)
	})

	It("registers the correct route in the router", func() {
		healthRoutes.Setup()
		Expect(router.Routes()).To(HaveLen(1))
		route := router.Routes()[0]
		Expect(route.Method).To(Equal("GET"))
		Expect(route.Path).To(Equal("/health"))
	})
})
