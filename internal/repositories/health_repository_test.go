package repositories_test

import (
	"context"
	"database/sql"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/DATA-DOG/go-sqlmock"

	"core-gin/infrastructure"
	"core-gin/lib"

	"core-gin/internal/repositories"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _ = Describe("HealthRepo", func() {
	var (
		repo       repositories.IHealthRepo
		mockDB     *sql.DB
		mockTracer infrastructure.ITracer
		mockGorm   *gorm.DB
		mockDS     infrastructure.Database
		ctx        context.Context
	)

	BeforeEach(func() {
		// Create a mock database
		var err error
		mockDB, _, err = sqlmock.New()
		Expect(err).ToNot(HaveOccurred())

		// Create a mock tracer
		mockTracer = infrastructure.NewTracer(&lib.Env{})

		dialector := postgres.New(postgres.Config{
			DSN:                  "sqlmock_db_0",
			DriverName:           "postgres",
			Conn:                 mockDB,
			PreferSimpleProtocol: true,
		})

		mockGorm, err = gorm.Open(dialector, &gorm.Config{})
		Expect(err).ToNot(HaveOccurred())
		mockDS = infrastructure.Database{DB: mockGorm}

		// Create a HealthRepo using the mock database and tracer
		repo = repositories.NewHealthRepo(mockDS, mockTracer)

		// Create a context
		ctx = context.Background()
	})

	It("should return the mock database connection", func() {
		// Call the GetDB method
		db, err := repo.GetDB(ctx)
		Expect(err).ToNot(HaveOccurred())

		// Assert that the returned *sql.DB is the same as the mockDB
		Expect(db).To(Equal(mockDB))
	})
})
