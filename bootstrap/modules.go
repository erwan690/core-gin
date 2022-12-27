package bootstrap

import (
	"core-gin/infrastructure"
	"core-gin/internal/handlers"
	"core-gin/internal/repositories"
	"core-gin/lib"

	"core-gin/api/middlewares"
	"core-gin/api/routes"

	"core-gin/internal/services"

	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	lib.Module,
	infrastructure.Module,
	middlewares.Module,
	repositories.Module,
	services.Module,
	handlers.Module,
	routes.Module,
)
