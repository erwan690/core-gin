package repositories

import (
	"context"
	"database/sql"

	"core-gin/infrastructure"
)

type HealthRepo struct {
	infrastructure.Database
	tracer infrastructure.Tracer
}

func NewHealthRepo(db infrastructure.Database, tracer infrastructure.Tracer) HealthRepo {
	return HealthRepo{Database: db, tracer: tracer}
}

func (r *HealthRepo) GetDB(ctx context.Context) (*sql.DB, error) {
	ctx, span := r.tracer.Start(ctx, "HealthRepo.GetDB")
	defer span.End()
	return r.WithContext(ctx).DB()
}