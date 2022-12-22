package commands

import (
	"context"
	"time"

	"core-gin/api/middlewares"
	"core-gin/infrastructure"
	"core-gin/lib"

	"github.com/spf13/cobra"
)

// ServeCommand test command
type ServeCommand struct{}

func (s *ServeCommand) Short() string {
	return "serve application"
}

func (s *ServeCommand) Setup(cmd *cobra.Command) {}

func (s *ServeCommand) Run() lib.CommandRunner {
	return func(
		env *lib.Env,
		middleware middlewares.Middlewares,
		logger lib.Logger,
		router infrastructure.Router,
		database infrastructure.Database,
		otel infrastructure.Otel,
	) {
		logger.Info("Running server")
		// Using time zone as specified in env file
		loc, _ := time.LoadLocation(env.TimeZone)
		time.Local = loc

		// setup list Of Middleware
		defer otel.Shutdown(context.Background())
		middleware.Setup()
		// Set Default Port
		const defaultServerPort = "8080"

		serverPort := env.ServerPort
		if serverPort == "" {
			serverPort = defaultServerPort
		}

		if err := router.Run(":" + serverPort); err != nil {
			logger.Fatal(err)
		}
	}
}

func NewServeCommand() *ServeCommand {
	return &ServeCommand{}
}
