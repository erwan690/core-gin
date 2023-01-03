package middlewares_test

import (
	"net/http"
	"net/http/httptest"

	"core-gin/api/middlewares"
	"core-gin/infrastructure"
	"core-gin/lib"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DBTransactionMiddleware", func() {
	It("should Commit the transaction on success status codes", func() {
		logger := lib.GetLogger()
		mockDB, mock, err := sqlmock.New()
		Expect(err).ToNot(HaveOccurred())
		logger.Info(mockDB.Ping())
		defer mockDB.Close()
		dialector := postgres.New(postgres.Config{
			DSN:                  "sqlmock_db_1",
			DriverName:           "postgres",
			Conn:                 mockDB,
			PreferSimpleProtocol: true,
		})
		mockGorm, err := gorm.Open(dialector, &gorm.Config{})
		Expect(err).ToNot(HaveOccurred())
		db := infrastructure.Database{DB: mockGorm}
		middleware := middlewares.NewDBTransactionMiddleware(logger, db)
		router := gin.New()
		recorder := httptest.NewRecorder()
		testMiddleware := middleware.Handle()
		router.Use(testMiddleware)

		router.GET("/test", func(c *gin.Context) {
			c.String(http.StatusOK, "test")
		})
		mock.ExpectBegin()
		mock.ExpectCommit()
		req, _ := http.NewRequest("GET", "/test", nil)
		router.ServeHTTP(recorder, req)
		Expect(recorder.Code).To(Equal(http.StatusOK))
		Expect(recorder.Body.String()).To(Equal("test"))
		Expect(mock.ExpectationsWereMet()).ToNot(HaveOccurred())
	})

	It("should rollback the transaction on panic", func() {
		logger := lib.GetLogger()
		mockDB, mock, err := sqlmock.New()
		Expect(err).ToNot(HaveOccurred())
		logger.Info(mockDB.Ping())
		defer mockDB.Close()
		dialector := postgres.New(postgres.Config{
			DSN:                  "sqlmock_db_1",
			DriverName:           "postgres",
			Conn:                 mockDB,
			PreferSimpleProtocol: true,
		})
		mockGorm, err := gorm.Open(dialector, &gorm.Config{})
		Expect(err).ToNot(HaveOccurred())
		db := infrastructure.Database{DB: mockGorm}
		middleware := middlewares.NewDBTransactionMiddleware(logger, db)
		router := gin.New()
		recorder := httptest.NewRecorder()
		testMiddleware := middleware.Handle()
		router.Use(testMiddleware)

		router.GET("/test", func(c *gin.Context) {
			panic("testpanic")
		})
		mock.ExpectBegin()
		mock.ExpectRollback()
		req, _ := http.NewRequest("GET", "/test", nil)
		router.ServeHTTP(recorder, req)
		Expect(recorder.Code).To(Equal(http.StatusOK))
		Expect(mock.ExpectationsWereMet()).ToNot(HaveOccurred())
	})

	It("should rollback the transaction on internal server error", func() {
		logger := lib.GetLogger()
		mockDB, mock, err := sqlmock.New()
		Expect(err).ToNot(HaveOccurred())
		logger.Info(mockDB.Ping())
		defer mockDB.Close()
		dialector := postgres.New(postgres.Config{
			DSN:                  "sqlmock_db_1",
			DriverName:           "postgres",
			Conn:                 mockDB,
			PreferSimpleProtocol: true,
		})
		mockGorm, err := gorm.Open(dialector, &gorm.Config{})
		Expect(err).ToNot(HaveOccurred())
		db := infrastructure.Database{DB: mockGorm}
		middleware := middlewares.NewDBTransactionMiddleware(logger, db)
		router := gin.New()
		recorder := httptest.NewRecorder()
		testMiddleware := middleware.Handle()
		router.Use(testMiddleware)

		router.GET("/test", func(c *gin.Context) {
			c.String(http.StatusInternalServerError, "test")
		})
		mock.ExpectBegin()
		mock.ExpectRollback()
		req, _ := http.NewRequest("GET", "/test", nil)
		router.ServeHTTP(recorder, req)
		Expect(recorder.Code).To(Equal(http.StatusInternalServerError))
		Expect(recorder.Body.String()).To(Equal("test"))
		Expect(mock.ExpectationsWereMet()).ToNot(HaveOccurred())
	})
})
