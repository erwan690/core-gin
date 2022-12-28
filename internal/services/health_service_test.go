package services

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"core-gin/infrastructure"
	"core-gin/lib"
)

type mockHealthRepo struct {
	err error
}

func (r *mockHealthRepo) GetDB(ctx context.Context) (*sql.DB, error) {
	// return a mock database connection and an error
	return nil, fmt.Errorf("error getting database connection")
}

func TestHealthService_PingDB(t *testing.T) {
	// create a mock repository and tracer
	repo := &mockHealthRepo{}
	tracer := infrastructure.NewTracer(&lib.Env{})

	// create a new HealthService using the mock repository and tracer
	s := NewHealthService(repo, tracer)

	// create a context
	ctx := context.Background()

	// call the PingDB function
	err := s.PingDB(ctx)

	// assert that the correct error is returned
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}
