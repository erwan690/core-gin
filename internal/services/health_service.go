package services

import (
	"context"

	"core-gin/infrastructure"
	"core-gin/internal/repositories"
)

type HealthService struct {
	repository repositories.HealthRepo
	tracer     infrastructure.Tracer
}

func NewHealthService(repository repositories.HealthRepo, tracer infrastructure.Tracer) HealthService {
	return HealthService{repository: repository, tracer: tracer}
}

func (s *HealthService) PingDB(ctx context.Context) (err error) {
	ctx, span := s.tracer.Start(ctx, "HealthService.PingDB")
	defer span.End()
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
