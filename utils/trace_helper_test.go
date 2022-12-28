package utils

import (
	"bytes"
	"context"
	"net/http"
	"testing"

	"go.opentelemetry.io/otel"
	sdktracer "go.opentelemetry.io/otel/sdk/trace"
)

func TestGetAppSource(t *testing.T) {
	// Test with a non-nil request
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("X-Client-Id", "test-client")
	result := GetAppSource(req)
	if result != "test-client" {
		t.Errorf("Expected 'test-client', got '%s'", result)
	}

	// Test with a nil request
	result = GetAppSource(nil)
	if result != "" {
		t.Errorf("Expected '', got '%s'", result)
	}
}

func TestGetRequestID(t *testing.T) {
	// Test with a non-nil request
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("X-Request-Id", "12345")
	result := GetRequestID(req)
	if result != "12345" {
		t.Errorf("Expected '12345', got '%s'", result)
	}

	// Test with a nil request
	result = GetRequestID(nil)
	if result != "" {
		t.Errorf("Expected '', got '%s'", result)
	}
}

func TestGetAppVersion(t *testing.T) {
	// Test with a non-nil request
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("X-Client-Version", "12345")
	result := GetAppVersion(req)
	if result != "12345" {
		t.Errorf("Expected '12345', got '%s'", result)
	}

	// Test with a nil request
	result = GetAppVersion(nil)
	if result != "" {
		t.Errorf("Expected '', got '%s'", result)
	}
}

func TestGetTraceIDFromCtx(t *testing.T) {
	// Test with a non-nil context
	tracer := otel.Tracer("test-tracer")
	otel.SetTracerProvider(
		sdktracer.NewTracerProvider(
			sdktracer.WithSampler(sdktracer.AlwaysSample()),
		),
	)
	ctx, _ := tracer.Start(context.Background(), "test-span")
	result := GetTraceIDFromCtx(ctx)
	if result == "00000000000000000000000000000000" {
		t.Errorf("Expected a non-zero trace ID, got '%s'", result)
	}

	// Test with a nil context
	result = GetTraceIDFromCtx(nil)
	if result != "00000000000000000000000000000000" {
		t.Errorf("Expected '00000000000000000000000000000000', got '%s'", result)
	}
}

func TestGetSpanIDFromCtx(t *testing.T) {
	// Test with a non-nil context
	tracer := otel.Tracer("test-tracer")
	otel.SetTracerProvider(
		sdktracer.NewTracerProvider(
			sdktracer.WithSampler(sdktracer.AlwaysSample()),
		),
	)
	ctx, _ := tracer.Start(context.Background(), "test-span")
	result := GetSpanIDFromCtx(ctx)
	if result == "0000000000000000" {
		t.Errorf("Expected a non-zero span ID, got '%s'", result)
	}

	// Test with a nil context
	result = GetSpanIDFromCtx(nil)
	if result != "0000000000000000" {
		t.Errorf("Expected '0000000000000000', got '%s'", result)
	}
}

func TestGetBodyTrace(t *testing.T) {
	// Test with a non-nil request
	body := "test body"
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer([]byte(body)))
	result := GetBodyTrace(req)
	if result != body {
		t.Errorf("Expected '%s', got '%s'", body, result)
	}

	// Test with a nil request
	result = GetBodyTrace(nil)
	if result != "" {
		t.Errorf("Expected '', got '%s'", result)
	}
}
