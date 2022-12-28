package utils

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"go.opentelemetry.io/otel/trace"
)

func GetAppSource(r *http.Request) string {
	if r == nil {
		return ""
	}
	return r.Header.Get("X-Client-Id")
}

func GetRequestID(r *http.Request) string {
	if r == nil {
		return ""
	}
	return r.Header.Get("X-Request-Id")
}

func GetAppVersion(r *http.Request) string {
	if r == nil {
		return ""
	}
	return r.Header.Get("X-Client-Version")
}

func GetTraceIDFromCtx(ctx context.Context) string {
	if ctx == nil {
		return "00000000000000000000000000000000"
	}
	span := trace.SpanFromContext(ctx)

	return span.SpanContext().TraceID().String()
}

func GetSpanIDFromCtx(ctx context.Context) string {
	if ctx == nil {
		return "0000000000000000"
	}
	span := trace.SpanFromContext(ctx)

	return span.SpanContext().SpanID().String()
}

func GetBodyTrace(r *http.Request) string {
	if r == nil || r.Body == nil {
		return ""
	}
	var body []byte
	var buf bytes.Buffer
	tee := io.TeeReader(r.Body, &buf)
	body, _ = io.ReadAll(tee)
	r.Body = io.NopCloser(&buf)
	return string(body)
}
