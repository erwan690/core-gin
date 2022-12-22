package middlewares

import (
	"context"
	"strings"

	"core-gin/infrastructure"
	"core-gin/lib"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	metrics "github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"
)

type MetricsMiddleware struct {
	router infrastructure.Router
	logger lib.Logger
}

func NewMetricsMiddleware(router infrastructure.Router, logger lib.Logger) MetricsMiddleware {
	return MetricsMiddleware{
		router: router,
		logger: logger,
	}
}

func (m MetricsMiddleware) Setup() {
	m.logger.Info("Setting up metrics middleware")
	// prometheus init
	// Create our middleware.
	mdlw := middleware.New(middleware.Config{
		Recorder: metrics.NewRecorder(metrics.Config{}),
	})
	m.router.Use(m.Handler("", mdlw))
	m.router.GET("/metrics", gin.WrapH(promhttp.Handler()))
}

// Handler returns a Gin measuring middleware.
func (m MetricsMiddleware) Handler(handlerID string, md middleware.Middleware) gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains("/metrics", c.Request.URL.Path) {
			c.Next()
			return
		}

		statusCode := c.Writer.Status()
		if statusCode == 404 {
			c.Next()
			return
		}

		r := &reporter{c: c}
		handlerID := r.c.FullPath()
		md.Measure(handlerID, r, func() {
			c.Next()
		})
	}
}

type reporter struct {
	c *gin.Context
}

func (r *reporter) Method() string { return r.c.Request.Method }

func (r *reporter) Context() context.Context { return r.c.Request.Context() }

func (r *reporter) URLPath() string { return r.c.Request.URL.Path }

func (r *reporter) StatusCode() int { return r.c.Writer.Status() }

func (r *reporter) BytesWritten() int64 { return int64(r.c.Writer.Size()) }
