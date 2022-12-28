package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"core-gin/infrastructure"
	"core-gin/lib"

	"github.com/gin-gonic/gin"
)

type mockHealthService struct {
	err error
}

func (s *mockHealthService) PingDB(ctx context.Context) error {
	return s.err
}

func TestHealth(t *testing.T) {
	// Set up test data and dependencies
	service := &mockHealthService{}
	tracer := infrastructure.NewTracer(&lib.Env{})
	handler := NewHealthHandler(service, tracer)

	// Test successful database ping
	service.err = nil
	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	handler.Health(c)
	if w.Code != http.StatusOK {
		t.Errorf("Expected HTTP status code %d, got %d", http.StatusOK, w.Code)
	}
	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	if response["db"] != "ok" {
		t.Errorf("Expected 'db' to be 'ok', got '%s'", response["db"])
	}

	// Test failed database ping
	service.err = errors.New("error pinging database")
	req, _ = http.NewRequest("GET", "/health", nil)
	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	c.Request = req
	handler.Health(c)
	if w.Code != http.StatusOK {
		t.Errorf("Expected HTTP status code %d, got %d", http.StatusOK, w.Code)
	}
	json.Unmarshal(w.Body.Bytes(), &response)
	if response["db"] != "fail" {
		t.Errorf("Expected 'db' to be 'fail', got '%s'", response["db"])
	}
}
