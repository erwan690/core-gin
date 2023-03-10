package services

import (
	"context"

	"core-gin/infrastructure"
	"core-gin/internal/repositories"
)

type IHealthService interface {
	PingDB(ctx context.Context) (err error)
}

type HealthService struct {
	repository repositories.IHealthRepo
	tracer     infrastructure.ITracer
}

func NewHealthService(repository repositories.IHealthRepo, tracer infrastructure.ITracer) IHealthService {
	return &HealthService{repository: repository, tracer: tracer}
}

func (s *HealthService) PingDB(ctx context.Context) (err error) {
	ctx, span := s.tracer.Start(ctx, "HealthService.PingDB")
	defer span.End()
	db, err := s.repository.GetDB(ctx)
	if err != nil {
		return err
	}
	return db.PingContext(ctx)
}
