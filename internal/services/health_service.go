package services

import (
	"context"

	"core-gin/internal/repositories"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type HealthService struct {
	repository repositories.HealthRepo
}

func NewHealthService(repository repositories.HealthRepo) HealthService {
	return HealthService{repository: repository}
}

func (s *HealthService) PingDB(ctx context.Context) (err error) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("HealthService", "PingDB"))
	db, err := s.repository.GetDB(ctx)
	if err != nil {
		return err
	}
	err = db.PingContext(ctx)
	if err != nil {
		return err
	}
	return nil
}
