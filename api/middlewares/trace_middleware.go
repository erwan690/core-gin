package middlewares

import (
	"time"

	"core-gin/infrastructure"
	"core-gin/lib"
	"core-gin/utils"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type InOutMiddleware struct {
	logger lib.Logger
	router infrastructure.Router
}

// NewDBTransactionMiddleware -> new instance of transaction
func NewInOutMiddlewareMiddleware(
	logger lib.Logger,
	router infrastructure.Router,
) InOutMiddleware {
	return InOutMiddleware{
		logger: logger,
		router: router,
	}
}

func (m InOutMiddleware) Setup() {
	logzp := m.logger.Desugar().Logger
	m.router.Use(ginzap.GinzapWithConfig(logzp, &ginzap.Config{
		TimeFormat: time.RFC3339,
		Context: ginzap.Fn(func(c *gin.Context) []zapcore.Field {
			ctx := c.Request.Context()
			fields := []zapcore.Field{}
			// log request ID

			requestID := utils.GetRequestID(c.Request)
			if requestID == "" {
				requestID = uuid.New().String()
				c.Header("X-Request-Id", requestID)
			}
			fields = append(fields, zap.String("request_id", requestID))

			if clientID := utils.GetAppSource(c.Request); clientID != "" {
				fields = append(fields, zap.String("client_id", clientID))
			}

			if clientVersion := utils.GetAppVersion(c.Request); clientVersion != "" {
				fields = append(fields, zap.String("client_version", clientVersion))
			}

			// log trace and span ID
			if trace.SpanFromContext(ctx).SpanContext().IsValid() {
				fields = append(fields, zap.String("trace_id", utils.GetTraceIDFromCtx(ctx)))
				fields = append(fields, zap.String("span_id", utils.GetSpaneIDFromCtx(ctx)))
			}

			if body := utils.GetBodyTrace(c.Request); body != "" {
				fields = append(fields, zap.String("body", body))
			}

			return fields
		}),
	}))
	m.router.Use(ginzap.RecoveryWithZap(logzp, true))
}
