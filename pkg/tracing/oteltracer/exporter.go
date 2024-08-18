package oteltracer

import (
	"context"
	"fmt"
	"github.com/nash-567/goObserve/pkg/tracing/config"
	"github.com/nash-567/goObserve/pkg/tracing/model"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
)

var ErrUnknownTraceExporterType = fmt.Errorf("unknown trace exporter type")

// NewTraceExporter creates a new trace exporter based on the provided configuration.
// stdout exporter exports the spans to the stdout at regular intervals, the interval is configurable ExporterConfig.BatchTimeout.
// http exporter exports the spans to the specified endpoint.
func NewTraceExporter(ctx context.Context, cfg *config.TracingConfig) (sdkTrace.SpanExporter, error) {
	var (
		exporter sdkTrace.SpanExporter
		err      error
	)

	switch cfg.ExporterConfig.Type {
	case model.TraceExporterTypeStdout:
		exporter, err = newStdOutExporter()
	case model.TraceExporterTypeHTTP:
		exporter, err = newOTLPTraceHTTPExporter(ctx, &cfg.ExporterConfig)
	default:
		err = ErrUnknownTraceExporterType
	}
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	return exporter, nil
}

func newStdOutExporter() (*stdouttrace.Exporter, error) {
	exporter, err := stdouttrace.New(
		stdouttrace.WithPrettyPrint(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create stdout exporter: %w", err)
	}
	return exporter, nil
}

func newOTLPTraceHTTPExporter(ctx context.Context, cfg *config.TraceExporterConfig) (*otlptrace.Exporter, error) {
	exporter, err := otlptracehttp.New(ctx,
		otlptracehttp.WithRetry(otlptracehttp.RetryConfig{
			Enabled:         cfg.RetryConfig.Enabled,
			InitialInterval: cfg.RetryConfig.InitialInterval,
			MaxInterval:     cfg.RetryConfig.MaxInterval,
			MaxElapsedTime:  cfg.RetryConfig.MaxElapsedTime,
		}),
		otlptracehttp.WithTimeout(cfg.Timeout),
		otlptracehttp.WithEndpointURL(cfg.EndpointURL),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create otlptracehttp exporter: %w", err)
	}
	return exporter, nil
}
