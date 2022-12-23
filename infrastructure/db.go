package infrastructure

import (
	"fmt"

	"core-gin/lib"

	"github.com/uptrace/opentelemetry-go-extra/otelgorm"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database modal
type Database struct {
	*gorm.DB
}

// NewDatabase creates a new database instance
func NewDatabase(logger lib.Logger, env *lib.Env) Database {
	username := env.DBUsername
	password := env.DBPassword
	host := env.DBHost
	port := env.DBPort
	dbname := env.DBName

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, dbname)

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: logger.GetGormLogger(),
	})

	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		logger.Error(err)
	}
	if err != nil {
		logger.Info("Url: ", url)
		logger.Error(err)
	}

	logger.Info("Database connection established")

	return Database{
		DB: db,
	}
}
