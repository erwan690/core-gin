package utils_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"core-gin/utils"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUtils(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Utils Suite")
}

var _ = Describe("GetPageInfo", func() {
	It("should return the correct page info", func() {
		// Test with limit = 10, page = 1, and total = 100
		pageInfo := utils.GetPageInfo(10, 1, 100)
		Expect(pageInfo.CurrentPage).To(Equal(1))
		Expect(pageInfo.PageSize).To(Equal(10))
		Expect(pageInfo.TotalCount).To(Equal(100))
		Expect(pageInfo.TotalPages).To(Equal(10))
	})
})

var _ = Describe("JSONResponse", func() {
	It("should return the correct JSON response", func() {
		expectedResponse := gin.H{"success": true, "message": "success", "data": "some data", "trace_id": "123"}
		response := utils.JSONResponse(true, "success", "some data", "123")
		Expect(response).To(Equal(expectedResponse))
	})
})

var _ = Describe("JSONResponsePage", func() {
	It("should return the correct JSON response with pagination", func() {
		pageInfo := utils.GetPageInfo(10, 1, 100)
		expectedResponse := gin.H{"success": true, "message": "success", "data": "some data", "meta": pageInfo, "trace_id": "123"}
		response := utils.JSONResponsePage(true, "success", "some data", pageInfo, "123")
		Expect(response).To(Equal(expectedResponse))
	})
})

var _ = Describe("BadRequestResponse", func() {
	It("should return a Bad Request response with the correct JSON", func() {
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("GET", "/", nil)
		utils.BadRequestResponse(c, "invalid request")
		Expect(w.Result().StatusCode).To(Equal(http.StatusBadRequest))
		Expect(w.Result().Header.Get("Content-Type")).To(Equal("application/json; charset=utf-8"))
		Expect(w.Body.String()).To(MatchJSON(`{"success": false, "message": "invalid request", "data": null, "trace_id": "00000000000000000000000000000000"}`))
	})
})

var _ = Describe("PreconditionFailedResponse", func() {
	It("should return a Precondition Failed response with the correct JSON", func() {
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("GET", "/", nil)
		utils.PreconditionFailedResponse(c, "precondition failed", map[string][]string{"field1": {"error1"}})
		Expect(w.Result().StatusCode).To(Equal(http.StatusPreconditionFailed))
		Expect(w.Result().Header.Get("Content-Type")).To(Equal("application/json; charset=utf-8"))
		Expect(w.Body.String()).To(MatchJSON(`{"success": false, "message": "precondition failed", "data": {"field1": ["error1"]}, "trace_id": "00000000000000000000000000000000"}`))
	})
})

var _ = Describe("UnprocessableResponse", func() {
	It("should return an Unprocessable Entity response with the correct JSON", func() {
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("GET", "/", nil)
		utils.UnprocessableResponse(c, "unprocessable entity")
		Expect(w.Result().StatusCode).To(Equal(http.StatusUnprocessableEntity))
		Expect(w.Result().Header.Get("Content-Type")).To(Equal("application/json; charset=utf-8"))
		Expect(w.Body.String()).To(MatchJSON(`{"success": false, "message": "unprocessable entity", "data": null, "trace_id": "00000000000000000000000000000000"}`))
	})
})

var _ = Describe("UnauthorizedResponse", func() {
	It("should return an Unauthorized response with the correct JSON", func() {
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("GET", "/", nil)
		utils.UnauthorizedResponse(c, "unauthorized")
		Expect(w.Result().StatusCode).To(Equal(http.StatusUnauthorized))
		Expect(w.Result().Header.Get("Content-Type")).To(Equal("application/json; charset=utf-8"))
		Expect(w.Body.String()).To(MatchJSON(`{"success": false, "message": "unauthorized", "data": null, "trace_id": "00000000000000000000000000000000"}`))
	})
})

var _ = Describe("SuccessResponse", func() {
	It("should return an OK response with the correct JSON", func() {
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("GET", "/", nil)
		utils.SuccessResponse(c, "success", map[string]string{"field1": "value1"})
		Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
		Expect(w.Result().Header.Get("Content-Type")).To(Equal("application/json; charset=utf-8"))
		Expect(w.Body.String()).To(MatchJSON(`{"success": true, "message": "success", "data": {"field1": "value1"}, "trace_id": "00000000000000000000000000000000"}`))
	})
})

var _ = Describe("SuccessResponsePage", func() {
	It("should return an OK response with the correct JSON", func() {
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("GET", "/", nil)
		utils.SuccessResponsePage(c, "success", map[string]string{"field1": "value1"}, map[string]int{"page": 1, "pageSize": 10, "totalCount": 100, "totalPages": 10})
		Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
		Expect(w.Result().Header.Get("Content-Type")).To(Equal("application/json; charset=utf-8"))
		Expect(w.Body.String()).To(MatchJSON(`{"success": true, "message": "success", "data": {"field1": "value1"}, "meta": {"page": 1, "pageSize": 10, "totalCount": 100, "totalPages": 10}, "trace_id": "00000000000000000000000000000000"}`))
	})
})

var _ = Describe("InternalErrorResponse", func() {
	It("should return an Internal Server Error response with the correct JSON", func() {
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("GET", "/", nil)
		utils.InternalErrorResponse(c, "internal error")
		Expect(w.Result().StatusCode).To(Equal(http.StatusInternalServerError))
		Expect(w.Result().Header.Get("Content-Type")).To(Equal("application/json; charset=utf-8"))
		Expect(w.Body.String()).To(MatchJSON(`{"success": false, "message": "internal error", "data": null, "trace_id": "00000000000000000000000000000000"}`))
	})
})

var _ = Describe("NotFoundResponse", func() {
	It("should return a Not Found response with the correct JSON", func() {
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("GET", "/", nil)
		utils.NotFoundResponse(c)
		Expect(w.Result().StatusCode).To(Equal(http.StatusNotFound))
		Expect(w.Result().Header.Get("Content-Type")).To(Equal("application/json; charset=utf-8"))
		Expect(w.Body.String()).To(MatchJSON(`{"success": false, "message": "what are you looking for ?", "data": null, "trace_id": ""}`))
	})
})

var _ = Describe("DynamicErrorCodeResponse", func() {
	It("should return a response with the correct status code and JSON", func() {
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("GET", "/", nil)
		utils.DynamicErrorCodeResponse(c, "error", http.StatusBadRequest)
		Expect(w.Result().StatusCode).To(Equal(http.StatusBadRequest))
		Expect(w.Result().Header.Get("Content-Type")).To(Equal("application/json; charset=utf-8"))
		Expect(w.Body.String()).To(MatchJSON(`{"success": false, "message": "error", "data": null, "trace_id": "00000000000000000000000000000000"}`))
	})
})
