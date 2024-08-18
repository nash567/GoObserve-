package oteltracer

import (
	"github.com/nash-567/goObserve/pkg/tracing/model"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// Span is a wrapper around trace.Span that implements the model.Span interface.
type Span struct {
	traceSpan trace.Span
}

func newSpan(traceSpan trace.Span) *Span {
	return &Span{
		traceSpan: traceSpan,
	}
}

func (s *Span) End(opts ...model.SpanEndOption) {
	spanConfig := model.NewSpanEndConfig(opts...)
	s.traceSpan.End(toSDKSpanEndConfig(&spanConfig)...)
}

// IsRecording can be used to check if the span is recording or not.
func (s *Span) IsRecording() bool {
	return s.traceSpan.IsRecording()
}

// RecordError records an error in the span trace.
func (s *Span) RecordError(err error) {
	s.traceSpan.RecordError(err)
}

// SetAttributes sets the attributes for the span.
func (s *Span) SetAttributes(attributes ...model.KeyValue) {
	s.traceSpan.SetAttributes(toAttributes(attributes)...)
}

// AddEvent can be used to add an event to the span.
// These can be used to add additional information to the span.
func (s *Span) AddEvent(name string, opts ...model.EventOption) {
	eventConfig := model.NewEventConfig(opts...)
	s.traceSpan.AddEvent(name, toSDKEventConfig(&eventConfig)...)
}

func toAttributes(attributes []model.KeyValue) []attribute.KeyValue {
	attrs := make([]attribute.KeyValue, len(attributes))
	for i, v := range attributes {
		attrs[i] = v.GetAttributeKeyValue()
	}
	return attrs
}

func toSDKEventConfig(cfg *model.EventConfig) []trace.EventOption {
	var opts []trace.EventOption
	if cfg.Attributes() != nil {
		opts = append(opts, trace.WithAttributes(toAttributes(cfg.Attributes())...))
	}
	if !cfg.Timestamp().IsZero() {
		opts = append(opts, trace.WithTimestamp(cfg.Timestamp()))
	}
	if cfg.StackTrace() {
		opts = append(opts, trace.WithStackTrace(cfg.StackTrace()))
	}
	return opts
}
