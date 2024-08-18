package model

import "time"

// EventOption applies span event options to an EventConfig.
type EventOption interface {
	applyEvent(EventConfig) EventConfig
}

// SpanOption are options that can be used at both the beginning and end of a span.
type SpanOption interface {
	SpanStartOption
	SpanEndOption
}

// SpanEventOption are options that can be used with an event or a span.
type SpanEventOption interface {
	SpanOption
	EventOption
}

// SpanStartEventOption are options that can be used at the start of a span, or with an event.
type SpanStartEventOption interface {
	SpanStartOption
	EventOption
}

// SpanEndEventOption are options that can be used at the end of a span, or with an event.
type SpanEndEventOption interface {
	SpanEndOption
	EventOption
}

// EventConfig is a group of options for an Event.
type EventConfig struct {
	attributes []KeyValue
	timestamp  time.Time
	stackTrace bool
}

// NewEventConfig applies all the EventOptions to a returned EventConfig.
// no validation is performed on the returned EventConfig.
func NewEventConfig(options ...EventOption) EventConfig {
	var c EventConfig
	for _, option := range options {
		c = option.applyEvent(c)
	}
	return c
}

// Attributes describe the associated qualities of an Event.
func (cfg *EventConfig) Attributes() []KeyValue {
	return cfg.attributes
}

// Timestamp is a time in an Event life-cycle.
func (cfg *EventConfig) Timestamp() time.Time {
	return cfg.timestamp
}

// StackTrace checks whether stack trace capturing is enabled.
func (cfg *EventConfig) StackTrace() bool {
	return cfg.stackTrace
}

type attributeOption []KeyValue

func (o attributeOption) applySpanStart(c SpanConfig) SpanConfig {
	c.attributes = append(c.attributes, []KeyValue(o)...)
	return c
}

func (o attributeOption) applyEvent(c EventConfig) EventConfig {
	c.attributes = append(c.attributes, []KeyValue(o)...)
	return c
}

// WithAttributes adds the attributes related to a span life-cycle event.
// These attributes are used to describe the work a Span represents when this
// option is provided to a Span's start or end events. Otherwise, these
// attributes provide additional information about the event being recorded
// (e.g., error, state change, processing progress, system event).
func WithAttributes(attributes ...KeyValue) SpanStartEventOption {
	return attributeOption(attributes)
}

type timestampOption time.Time

func (o timestampOption) applySpanStart(c SpanConfig) SpanConfig {
	c.timestamp = time.Time(o)
	return c
}

func (o timestampOption) applySpanEnd(c SpanConfig) SpanConfig { return o.applySpanStart(c) }

func (o timestampOption) applyEvent(c EventConfig) EventConfig {
	c.timestamp = time.Time(o)
	return c
}

var _ SpanStartOption = timestampOption{}

// WithTimestamp sets the time of a Span or Event life-cycle moment (e.g.
// started, stopped, errored).
func WithTimestamp(t time.Time) SpanEventOption {
	return timestampOption(t)
}

var _ SpanEndOption = stackTraceOption(false)

type stackTraceOption bool

func (o stackTraceOption) applyEvent(c EventConfig) EventConfig {
	c.stackTrace = bool(o)
	return c
}

func (o stackTraceOption) applySpanEnd(c SpanConfig) SpanConfig {
	c.stackTrace = bool(o)
	return c
}

// WithStackTrace sets the flag to capture the error with stack trace (e.g. true, false).
func WithStackTrace(b bool) SpanEndEventOption {
	return stackTraceOption(b)
}
