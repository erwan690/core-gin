package middlewares

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"core-gin/infrastructure"
	"core-gin/lib"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// TestInOutMiddleware is a test suite for the InOutMiddleware struct
type TestInOutMiddleware struct {
	suite.Suite
}

// TestSetup tests the Setup method of the InOutMiddleware struct
func (suite *TestInOutMiddleware) TestSetup() {
	env := &lib.Env{
		ServiceName: "test",
	}
	router := infrastructure.Router{
		Engine: gin.New(),
	}
	logger := lib.GetLogger()
	inOutMiddleware := NewInOutMiddlewareMiddleware(logger, router, env)

	inOutMiddleware.Setup()
	postBody, _ := json.Marshal(map[string]string{
		"name":  "Toby",
		"email": "Toby@example.com",
	})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(postBody))

	req.Header = http.Header{
		"X-Client-Id":      {"Test"},
		"X-Client-Version": {"1.0"},
	}
	router.ServeHTTP(w, req)
	assert.Contains(suite.T(), w.Body.String(), "404 page not found")
}

// TestInOutMiddlewareSuite runs the TestInOutMiddleware test suite
func TestInOutMiddlewareSuite(t *testing.T) {
	suite.Run(t, new(TestInOutMiddleware))
}
