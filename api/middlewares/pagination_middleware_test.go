package middlewares

import (
	"core-gin/constants"
	"core-gin/lib"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// TestPaginationMiddleware is a test suite for the PaginationMiddleware struct
type TestPaginationMiddleware struct {
	suite.Suite
}

// TestHandle tests the Handle method of the PaginationMiddleware struct
func (suite *TestPaginationMiddleware) TestHandle() {
	logger := lib.GetLogger()
	paginationMiddleware := NewPaginationMiddleware(logger)
	handler := paginationMiddleware.Handle()

	// Create a mock Context object
	c, _ := gin.CreateTestContext(nil)
	// Set the query parameters of the mock Context object
	c.Request.URL.RawQuery = "limit=20&page=2"
	// Call the handler function
	handler(c)

	// Assert that the limit and page variables were set correctly
	assert.Equal(suite.T(), int64(20), c.GetInt64(constants.Limit))
	assert.Equal(suite.T(), int64(2), c.GetInt64(constants.Page))
}

// TestHandleWithoutParams tests the Handle method of the PaginationMiddleware struct when the query parameters are not set
func (suite *TestPaginationMiddleware) TestHandleWithoutParams() {
	logger := lib.GetLogger()
	paginationMiddleware := NewPaginationMiddleware(logger)
	handler := paginationMiddleware.Handle()

	// Create a mock Context object
	c, _ := gin.CreateTestContext(nil)
	// Call the handler function without setting the query parameters
	handler(c)

	// Assert that the limit and page variables were set to the default values
	assert.Equal(suite.T(), int64(10), c.GetInt64(constants.Limit))
	assert.Equal(suite.T(), int64(0), c.GetInt64(constants.Page))
}

// TestHandleWithInvalidParams tests the Handle method of the PaginationMiddleware struct when the query parameters are invalid
func (suite *TestPaginationMiddleware) TestHandleWithInvalidParams() {
	logger := lib.GetLogger()
	paginationMiddleware := NewPaginationMiddleware(logger)
	handler := paginationMiddleware.Handle()

	// Create a mock Context object
	c, _ := gin.CreateTestContext(nil)
	// Set the query parameters of the mock Context object to invalid values
	c.Request.URL.RawQuery = "limit=abc&page=def"
	// Call the handler function
	handler(c)

	// Assert that the limit and page variables were set to the default values
	assert.Equal(suite.T(), int64(10), c.GetInt64(constants.Limit))
	assert.Equal(suite.T(), int64(0), c.GetInt64(constants.Page))
}
