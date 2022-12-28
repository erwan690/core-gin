package handlers

import (
	"net/http"

	"core-gin/infrastructure"
	"core-gin/internal/services"

	"github.com/gin-gonic/gin"
)

type IHealthHandler interface {
	Health(c *gin.Context)
}

type HealthHandler struct {
	service services.IHealthService
	tracer  infrastructure.Tracer
}

func NewHealthHandler(service services.IHealthService, tracer infrastructure.Tracer) IHealthHandler {
	return &HealthHandler{service: service, tracer: tracer}
}

func (h *HealthHandler) Health(c *gin.Context) {
	ctx, span := h.tracer.Start(c.Request.Context(), "HealthHandler.Health")
	defer span.End()

	err := h.service.PingDB(ctx)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"db": "fail"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"db": "ok"})
}
