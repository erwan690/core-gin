package infrastructure

import (
	"core-gin/lib"

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
	appEnv := env.Environment
	if appEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	httpRouter := gin.New()

	httpRouter.SetTrustedProxies(nil)

	httpRouter.MaxMultipartMemory = env.MaxMultipartMemory

	return Router{
		httpRouter,
	}
}
