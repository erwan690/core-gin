package infrastructure

import (
	"core-gin/lib"

	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	"github.com/gin-gonic/gin"
)

// Router -> Gin Router
type Router struct {
	*gin.Engine
}

// NewRouter : all the routes are defined here
func NewRouter(
	env *lib.Env,
	logger lib.Logger,
) Router {
	gin.DefaultWriter = logger.GetGinLogger()
	appEnv := env.Environment
	if appEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	httpRouter := gin.Default()

	httpRouter.SetTrustedProxies(nil)

	httpRouter.Use(otelgin.Middleware(env.ServiceName))

	httpRouter.MaxMultipartMemory = env.MaxMultipartMemory

	return Router{
		httpRouter,
	}
}
