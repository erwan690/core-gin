package handlers

import (
	"net/http"

	"core-gin/internal/services"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type HealthHandler struct {
	service services.HealthService
}

func NewHealthHandler(service services.HealthService) HealthHandler {
	return HealthHandler{service: service}
}

func (h *HealthHandler) Health(c *gin.Context) {
	ctx := c.Request.Context()
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("HealthHandler", "Health"))

	err := h.service.PingDB(ctx)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"db": "fail"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"db": "ok"})
}
