package middlewares_test

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"core-gin/api/middlewares"
	"core-gin/constants"
	"core-gin/lib"

	"github.com/gin-gonic/gin"
)

var _ = Describe("PaginationMiddleware", func() {
	var (
		logger               lib.Logger
		paginationMiddleware middlewares.PaginationMiddleware
		handler              gin.HandlerFunc
		c                    *gin.Context
	)

	BeforeEach(func() {
		logger = lib.GetLogger()
		paginationMiddleware = middlewares.NewPaginationMiddleware(logger)
		handler = paginationMiddleware.Handle()
		w := httptest.NewRecorder()
		c, _ = gin.CreateTestContext(w)
	})

	It("sets the limit and page variables based on the query parameters", func() {
		c.Request, _ = http.NewRequest("GET", "/?limit=20&page=2", nil)
		handler(c)
		Expect(c.GetInt64(constants.Limit)).To(Equal(int64(20)))
		Expect(c.GetInt64(constants.Page)).To(Equal(int64(2)))
	})

	It("sets the limit and page variables to default values when the query parameters are not set", func() {
		handler(c)
		Expect(c.GetInt64(constants.Limit)).To(Equal(int64(10)))
		Expect(c.GetInt64(constants.Page)).To(Equal(int64(0)))
	})

	It("sets the limit and page variables to default values when the query parameters are invalid", func() {
		c.Request, _ = http.NewRequest("GET", "/?limit=abc&page=def", nil)
		handler(c)
		Expect(c.GetInt64(constants.Limit)).To(Equal(int64(10)))
		Expect(c.GetInt64(constants.Page)).To(Equal(int64(0)))
	})
})
