package handlers_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"core-gin/infrastructure"
	"core-gin/internal/handlers"
	"core-gin/lib"
	"core-gin/utils"

	"github.com/gin-gonic/gin"
)

type mockHealthService struct {
	err error
}

func (s *mockHealthService) PingDB(ctx context.Context) error {
	return s.err
}

var _ = Describe("Health", func() {
	var (
		service  *mockHealthService
		tracer   infrastructure.ITracer
		handler  handlers.IHealthHandler
		response utils.BaseResponse
		w        *httptest.ResponseRecorder
		c        *gin.Context
		req      *http.Request
	)

	BeforeEach(func() {
		// Set up test data and dependencies
		service = &mockHealthService{}
		tracer = infrastructure.NewTracer(&lib.Env{})
		handler = handlers.NewHealthHandler(service, tracer)
		w = httptest.NewRecorder()
		c, _ = gin.CreateTestContext(w)
		response = utils.BaseResponse{}
	})

	It("should return 'ok' for a successful database ping", func() {
		// Test successful database ping
		service.err = nil
		req, _ = http.NewRequest("GET", "/health", nil)
		c.Request = req
		handler.Health(c)
		Expect(w.Code).To(Equal(http.StatusOK))
		json.Unmarshal(w.Body.Bytes(), &response)
		Expect(response.Message).To(Equal("ok"))
	})

	It("should return 'fail' for a failed database ping", func() {
		// Test failed database ping
		service.err = errors.New("error pinging database")
		req, _ = http.NewRequest("GET", "/health", nil)
		c.Request = req
		handler.Health(c)
		Expect(w.Code).To(Equal(http.StatusOK))
		json.Unmarshal(w.Body.Bytes(), &response)
		Expect(response.Message).To(Equal("fail"))
	})
})
