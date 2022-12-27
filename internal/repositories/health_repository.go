package repositories

import (
	"context"
	"database/sql"

	"core-gin/infrastructure"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type HealthRepo struct {
	infrastructure.Database
}

func NewHealthRepo(db infrastructure.Database) HealthRepo {
	return HealthRepo{Database: db}
}

func (r *HealthRepo) GetDB(ctx context.Context) (*sql.DB, error) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("HealthRepo", "GetDB"))
	return r.WithContext(ctx).DB()
}
