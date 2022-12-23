package middlewares

import (
	"strconv"

	"core-gin/constants"
	"core-gin/lib"

	"github.com/gin-gonic/gin"
)

type PaginationMiddleware struct {
	logger lib.Logger
}

func NewPaginationMiddleware(logger lib.Logger) PaginationMiddleware {
	return PaginationMiddleware{logger: logger}
}

func (p PaginationMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		p.logger.Info("setting up pagination middleware")

		perPage, err := strconv.ParseInt(c.Query("limit"), 10, 0)
		if err != nil {
			perPage = 10
		}

		page, err := strconv.ParseInt(c.Query("page"), 10, 0)
		if err != nil {
			page = 0
		}

		c.Set(constants.Limit, perPage)
		c.Set(constants.Page, page)

		c.Next()
	}
}
