package repositories

import (
	"context"
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"

	"core-gin/infrastructure"
	"core-gin/lib"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGetDB(t *testing.T) {
	// Create a mock database

	mockdb, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("Failed to open mock sql db, got error: %v", err)
	}

	if mockdb == nil {
		t.Error("mock db is null")
	}

	if mock == nil {
		t.Error("sqlmock is null")
	}

	// Create a mock tracer
	mockTracer := infrastructure.NewTracer(&lib.Env{})

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 mockdb,
		PreferSimpleProtocol: true,
	})

	mockGorm, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	mockDs := infrastructure.Database{DB: mockGorm}

	// Create a HealthRepo using the mock database and tracer
	repo := NewHealthRepo(mockDs, mockTracer)

	// Call the GetDB method
	db, err := repo.GetDB(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	// Assert that the returned *sql.DB is the same as the mockDB
	if db != mockdb {
		t.Errorf("Expected %v, got %v", mockdb, db)
	}
}
