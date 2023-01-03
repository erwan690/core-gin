package utils

import (
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PageInfo struct {
	CurrentPage int `json:"current_page"`
	PageSize    int `json:"page_size"`
	TotalCount  int `json:"total_count"`
	TotalPages  int `json:"total_page"`
}

type BaseResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	TraceID string      `json:"trace_id"`
}

type SuccessListResponse struct {
	BaseResponse
	Meta PageInfo `json:"meta"`
}

type ErrorResponse struct {
	Success bool        `json:"success" default:"false"`
	Message string      `json:"message" example:"failure because of error"`
	Data    interface{} `json:"data"`
	TraceID string      `json:"trace_id"`
}

func GetPageInfo(limit int, page int, total int) *PageInfo {
	return &PageInfo{
		CurrentPage: page,
		PageSize:    limit,
		TotalCount:  total,
		TotalPages:  int(math.Ceil(float64(total) / float64(limit))),
	}
}

// JSONResponse method
func JSONResponse(success bool, message string, object interface{}, traceID string) gin.H {
	return gin.H{"success": success, "message": message, "data": object, "trace_id": traceID}
}

func JSONResponsePage(success bool, message string, object interface{}, meta interface{}, traceID string) gin.H {
	return gin.H{"success": success, "message": message, "data": object, "meta": meta, "trace_id": traceID}
}

func BadRequestResponse(c *gin.Context, msg string) {
	traceID := GetTraceIDFromCtx(c.Request.Context())
	c.JSON(http.StatusBadRequest, JSONResponse(false, msg, nil, traceID))
}

func PreconditionFailedResponse(c *gin.Context, msg string, err map[string][]string) {
	traceID := GetTraceIDFromCtx(c.Request.Context())
	c.JSON(http.StatusPreconditionFailed, JSONResponse(false, msg, err, traceID))
}

func UnprocessableResponse(c *gin.Context, msg string) {
	traceID := GetTraceIDFromCtx(c.Request.Context())
	c.JSON(http.StatusUnprocessableEntity, JSONResponse(false, msg, nil, traceID))
}

func UnauthorizedResponse(c *gin.Context, msg string) {
	traceID := GetTraceIDFromCtx(c.Request.Context())
	c.JSON(http.StatusUnauthorized, JSONResponse(false, msg, nil, traceID))
}

func SuccessResponse(c *gin.Context, msg string, object interface{}) {
	traceID := GetTraceIDFromCtx(c.Request.Context())
	c.JSON(http.StatusOK, JSONResponse(true, msg, object, traceID))
}

func SuccessResponsePage(c *gin.Context, msg string, object interface{}, meta interface{}) {
	traceID := GetTraceIDFromCtx(c.Request.Context())
	c.JSON(http.StatusOK, JSONResponsePage(true, msg, object, meta, traceID))
}

func InternalErrorResponse(c *gin.Context, msg string) {
	traceID := GetTraceIDFromCtx(c.Request.Context())
	c.JSON(http.StatusInternalServerError, JSONResponse(false, msg, nil, traceID))
}

func NotFoundResponse(c *gin.Context) {
	c.JSON(http.StatusNotFound, JSONResponse(false, "what are you looking for ?", nil, ""))
}

func DynamicErrorCodeResponse(c *gin.Context, msg string, code int) {
	traceID := GetTraceIDFromCtx(c.Request.Context())
	c.JSON(code, JSONResponse(false, msg, nil, traceID))
}
