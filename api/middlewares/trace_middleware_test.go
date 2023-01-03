package middlewares_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"core-gin/api/middlewares"
	"core-gin/infrastructure"
	"core-gin/lib"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("InOutMiddleware", func() {
	var (
		env             *lib.Env
		router          infrastructure.Router
		logger          lib.Logger
		inOutMiddleware middlewares.InOutMiddleware
		w               *httptest.ResponseRecorder
		req             *http.Request
	)

	BeforeEach(func() {
		env = &lib.Env{
			ServiceName: "test",
		}
		router = infrastructure.Router{
			Engine: gin.New(),
		}
		logger = lib.GetLogger()
		inOutMiddleware = middlewares.NewInOutMiddlewareMiddleware(logger, router, env)
		inOutMiddleware.Setup()
		w = httptest.NewRecorder()
	})

	It("handles a POST request", func() {
		postBody, _ := json.Marshal(map[string]string{
			"name":  "Toby",
			"email": "Toby@example.com",
		})
		req, _ = http.NewRequest("POST", "/", bytes.NewBuffer(postBody))
		req.Header = http.Header{
			"X-Client-Id":      {"Test"},
			"X-Client-Version": {"1.0"},
		}

		router.ServeHTTP(w, req)
		Expect(w.Body.String()).To(ContainSubstring("404 page not found"))
	})
})
