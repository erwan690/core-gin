package bootstrap

import (
	"core-gin/infrastructure"
	"core-gin/lib"

	"core-gin/api/middlewares"

	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	lib.Module,
	infrastructure.Module,
	middlewares.Module,
)
