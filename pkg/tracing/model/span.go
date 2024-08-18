package model

import "time"

// SpanStartOption applies an option to a SpanConfig. These options are applicable
// only when the span is created.
type SpanStartOption interface {
	applySpanStart(SpanConfig) SpanConfig
}

// SpanEndOption applies an option to a SpanConfig. These options are
// applicable only when the span is ended.
type SpanEndOption interface {
	applySpanEnd(SpanConfig) SpanConfig
}

// SpanConfig is a group of options for a Span.
type SpanConfig struct {
	attributes []KeyValue
	timestamp  time.Time
	newRoot    bool
	stackTrace bool
}

// Attributes describe the associated qualities of a Span.
func (cfg *SpanConfig) Attributes() []KeyValue {
	return cfg.attributes
}

// Timestamp is a time in a Span life-cycle.
func (cfg *SpanConfig) Timestamp() time.Time {
	return cfg.timestamp
}

// NewRoot identifies a Span as the root Span for a new trace. This is
// commonly used when an existing trace crosses trust boundaries and the
// remote parent span context should be ignored for security.
func (cfg *SpanConfig) NewRoot() bool {
	return cfg.newRoot
}

// StackTrace checks whether stack trace capturing is enabled.
func (cfg *SpanConfig) StackTrace() bool {
	return cfg.stackTrace
}

type spanStartOptionFunc func(SpanConfig) SpanConfig

func (fn spanStartOptionFunc) applySpanStart(cfg SpanConfig) SpanConfig {
	return fn(cfg)
}

// NewSpanStartConfig applies all the options to a returned SpanConfig.
// No validation is performed on the returned SpanConfig (e.g. no uniqueness
// checking or bounding of data), it is left to the SDK to perform this
// action.
func NewSpanStartConfig(options ...SpanStartOption) SpanConfig {
	var c SpanConfig
	for _, option := range options {
		c = option.applySpanStart(c)
	}
	return c
}

// NewSpanEndConfig applies all the options to a returned SpanConfig.
// No validation is performed on the returned SpanConfig (e.g. no uniqueness
// checking or bounding of data), it is left to the SDK to perform this
// action.
func NewSpanEndConfig(options ...SpanEndOption) SpanConfig {
	var c SpanConfig
	for _, option := range options {
		c = option.applySpanEnd(c)
	}
	return c
}
