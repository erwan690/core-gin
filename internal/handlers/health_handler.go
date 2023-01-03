package handlers

import (
	"net/http"

	"core-gin/infrastructure"
	"core-gin/internal/services"
	"core-gin/utils"

	"github.com/gin-gonic/gin"
)

type IHealthHandler interface {
	Health(c *gin.Context)
}

type HealthHandler struct {
	service services.IHealthService
	tracer  infrastructure.ITracer
}

func NewHealthHandler(service services.IHealthService, tracer infrastructure.ITracer) IHealthHandler {
	return &HealthHandler{service: service, tracer: tracer}
}

func (h *HealthHandler) Health(c *gin.Context) {
	ctx, span := h.tracer.Start(c.Request.Context(), "HealthHandler.Health")
	defer span.End()

	err := h.service.PingDB(ctx)
	if err != nil {
		utils.DynamicErrorCodeResponse(c, "fail", http.StatusOK)
		return
	}
	utils.SuccessResponse(c, "ok", nil)
}
