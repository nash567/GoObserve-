package oteltracer

import (
	"fmt"
	"github.com/nash-567/goObserve/pkg/tracing/config"
	"go.opentelemetry.io/otel/sdk/resource"
	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.25.0"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

// NewTraceProvider creates a new trace provider using the exporter and configuration provided.
// This provider is used to initialize the tracer which is then used across the application.
//
//nolint:ireturn
func NewTraceProvider(
	cfg *config.TracingConfig,
	traceExporter sdkTrace.SpanExporter,
	serviceName string,
) (trace.TracerProvider, error) {
	if !cfg.Enabled {
		return noop.NewTracerProvider(), nil
	}
	r, err := resource.Merge(
		resource.Default(),
		resource.NewSchemaless(
			semconv.ServiceName(serviceName),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed creating resource info: %w", err)
	}

	tp := sdkTrace.NewTracerProvider(
		sdkTrace.WithResource(r),
		sdkTrace.WithBatcher(traceExporter,
			sdkTrace.WithBatchTimeout(cfg.ExporterConfig.BatchTimeout)),
	)
	return tp, nil
}
