package services_test

import (
	"context"
	"database/sql"
	"fmt"

	"core-gin/infrastructure"
	"core-gin/internal/services"
	"core-gin/lib"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.opentelemetry.io/otel/trace"
)

type mockHealthRepo struct {
	err error
}

func (r *mockHealthRepo) GetDB(ctx context.Context) (*sql.DB, error) {
	// return a mock database connection and an error
	return nil, fmt.Errorf("error getting database connection")
}

var _ = Describe("HealthService", func() {
	var (
		s      services.IHealthService
		repo   *mockHealthRepo
		tracer trace.Tracer
		ctx    context.Context
	)

	BeforeEach(func() {
		// create a mock repository and tracer
		repo = &mockHealthRepo{}
		tracer = infrastructure.NewTracer(&lib.Env{})

		// create a new HealthService using the mock repository and tracer
		s = services.NewHealthService(repo, tracer)

		// create a context
		ctx = context.Background()
	})

	It("should return an error when pinging the database", func() {
		// call the PingDB function
		err := s.PingDB(ctx)

		// assert that the correct error is returned
		Expect(err).To(HaveOccurred())
	})
})
