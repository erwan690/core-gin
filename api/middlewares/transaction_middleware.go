package middlewares

import (
	"net/http"

	"core-gin/constants"
	"core-gin/infrastructure"
	"core-gin/lib"
	"core-gin/utils"

	"github.com/gin-gonic/gin"
)

// DBTransactionMiddleware -> struct for transaction
type DBTransactionMiddleware struct {
	logger lib.Logger
	db     infrastructure.Database
}

// NewDBTransactionMiddleware -> new instance of transaction
func NewDBTransactionMiddleware(
	logger lib.Logger,
	db infrastructure.Database,
) DBTransactionMiddleware {
	return DBTransactionMiddleware{
		logger: logger,
		db:     db,
	}
}

// Setup sets up database transaction middleware
func (m DBTransactionMiddleware) Handle() gin.HandlerFunc {
	m.logger.Info("setting up database transaction middleware")

	return func(c *gin.Context) {
		txHandle := m.db.DB.Begin()
		ctx := c.Request.Context()
		m.logger.InfofContext(ctx, "beginning database transaction")

		defer func() {
			if r := recover(); r != nil {
				txHandle.Rollback()
			}
		}()

		c.Set(constants.DBTransaction, txHandle)
		c.Next()

		// commit transaction on success status
		if utils.StatusInList(c.Writer.Status(), []int{http.StatusOK, http.StatusCreated, http.StatusNoContent}) {
			m.logger.InfofContext(ctx, "committing transactions")
			if err := txHandle.Commit().Error; err != nil {
				m.logger.Error("trx commit error: ", err)
			}
		} else {
			m.logger.InfofContext(ctx, "rolling back transaction due to status code: 500")
			txHandle.Rollback()
		}
	}
}
