package oteltracer

import (
	"context"
	"github.com/nash-567/goObserve/pkg/tracing/config"
	"github.com/nash-567/goObserve/pkg/tracing/model"
	"go.opentelemetry.io/otel/trace"
)

// Tracer is a wrapper around the OpenTelemetry tracer.
type Tracer struct {
	tracerProvider trace.TracerProvider
	sdkTracer      trace.Tracer
}

// NewTracer creates a new tracer instance which is used across the application.
func NewTracer(cfg *config.TracingConfig, tp trace.TracerProvider) *Tracer {
	return &Tracer{
		tracerProvider: tp,
		sdkTracer: tp.Tracer(
			cfg.InstrumentationLibrary.Name,
			trace.WithInstrumentationVersion(cfg.InstrumentationLibrary.Version),
			trace.WithSchemaURL(cfg.InstrumentationLibrary.SchemaURL),
		),
	}
}

// StartSpan starts a new span with the given name and options.
func (t *Tracer) StartSpan(ctx context.Context, spanName string, opts ...model.SpanStartOption) (context.Context, model.Span) {
	spanConfig := model.NewSpanStartConfig(opts...)
	ctx, s := t.sdkTracer.Start(ctx, spanName, toSDKSpanStartConfig(&spanConfig)...)
	return ctx, newSpan(s)
}

// SpanFromContext returns the active span from the context.
func (t *Tracer) SpanFromContext(ctx context.Context) model.Span {
	return newSpan(trace.SpanFromContext(ctx))
}

// TracerProvider returns the tracer provider.
func (t *Tracer) TracerProvider() trace.TracerProvider {
	return t.tracerProvider
}

// ParentTracer returns the parent tracer.
func (t *Tracer) ParentTracer() interface{} {
	return t.sdkTracer
}

func toSDKSpanStartConfig(cfg *model.SpanConfig) []trace.SpanStartOption {
	var opts []trace.SpanStartOption
	if cfg.Attributes() != nil {
		opts = append(opts, trace.WithAttributes(toAttributes(cfg.Attributes())...))
	}
	if cfg.NewRoot() {
		opts = append(opts, trace.WithNewRoot())
	}
	if !cfg.Timestamp().IsZero() {
		opts = append(opts, trace.WithTimestamp(cfg.Timestamp()))
	}
	return opts
}

func toSDKSpanEndConfig(cfg *model.SpanConfig) []trace.SpanEndOption {
	var opts []trace.SpanEndOption
	if cfg.StackTrace() {
		opts = append(opts, trace.WithStackTrace(cfg.StackTrace()))
	}
	if !cfg.Timestamp().IsZero() {
		opts = append(opts, trace.WithTimestamp(cfg.Timestamp()))
	}
	return opts
}
