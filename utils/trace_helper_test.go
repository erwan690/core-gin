package utils_test

import (
	"bytes"
	"context"
	"net/http"

	"core-gin/utils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.opentelemetry.io/otel"
	sdktracer "go.opentelemetry.io/otel/sdk/trace"
)

var _ = Describe("Utils", func() {
	Describe("GetAppSource", func() {
		It("returns the value of the X-Client-Id header when the request is non-nil", func() {
			req, _ := http.NewRequest("GET", "/", nil)
			req.Header.Set("X-Client-Id", "test-client")
			result := utils.GetAppSource(req)
			Expect(result).To(Equal("test-client"))
		})

		It("returns an empty string when the request is nil", func() {
			result := utils.GetAppSource(nil)
			Expect(result).To(Equal(""))
		})
	})

	Describe("GetRequestID", func() {
		It("returns the value of the X-Request-Id header when the request is non-nil", func() {
			req, _ := http.NewRequest("GET", "/", nil)
			req.Header.Set("X-Request-Id", "12345")
			result := utils.GetRequestID(req)
			Expect(result).To(Equal("12345"))
		})

		It("returns an empty string when the request is nil", func() {
			result := utils.GetRequestID(nil)
			Expect(result).To(Equal(""))
		})
	})

	Describe("GetAppVersion", func() {
		It("returns the value of the X-Client-Version header when the request is non-nil", func() {
			req, _ := http.NewRequest("GET", "/", nil)
			req.Header.Set("X-Client-Version", "12345")
			result := utils.GetAppVersion(req)
			Expect(result).To(Equal("12345"))
		})

		It("returns an empty string when the request is nil", func() {
			result := utils.GetAppVersion(nil)
			Expect(result).To(Equal(""))
		})
	})

	Describe("GetTraceIDFromCtx", func() {
		It("returns a non-zero trace ID when the context is non-nil", func() {
			tracer := otel.Tracer("test-tracer")
			otel.SetTracerProvider(
				sdktracer.NewTracerProvider(
					sdktracer.WithSampler(sdktracer.AlwaysSample()),
				),
			)
			ctx, _ := tracer.Start(context.Background(), "test-span")
			result := utils.GetTraceIDFromCtx(ctx)
			Expect(result).ToNot(Equal("00000000000000000000000000000000"))
		})

		It("returns '00000000000000000000000000000000' when the context is nil", func() {
			result := utils.GetTraceIDFromCtx(nil)
			Expect(result).To(Equal("00000000000000000000000000000000"))
		})
	})

	Describe("GetSpanIDFromCtx", func() {
		It("returns a non-zero span ID when the context is non-nil", func() {
			tracer := otel.Tracer("test-tracer")
			otel.SetTracerProvider(
				sdktracer.NewTracerProvider(
					sdktracer.WithSampler(sdktracer.AlwaysSample()),
				),
			)
			ctx, _ := tracer.Start(context.Background(), "test")
			result := utils.GetSpanIDFromCtx(ctx)
			Expect(result).ToNot(Equal("0000000000000000"))
		})

		It("returns '0000000000000000' when the context is nil", func() {
			result := utils.GetSpanIDFromCtx(nil)
			Expect(result).To(Equal("0000000000000000"))
		})
	})

	Describe("GetBodyTrace", func() {
		It("returns the request body as a string when the request is non-nil", func() {
			req, _ := http.NewRequest("POST", "/", bytes.NewBuffer([]byte(`{"key": "value"}`)))
			result := utils.GetBodyTrace(req)
			Expect(result).To(Equal(`{"key": "value"}`))
		})

		It("returns an empty string when the request is nil", func() {
			result := utils.GetBodyTrace(nil)
			Expect(result).To(Equal(""))
		})
	})
})
