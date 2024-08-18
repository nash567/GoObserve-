package model

import "context"

type Tracer interface {
	// StartSpan creates a span and a context.Context containing the newly-created span.
	//
	// If the context.Context provided in `ctx` contains a Span then the newly-created
	// Span will be a child of that span, otherwise it will be a root span. This behavior
	// can be overridden by providing `WithNewRoot()` as a SpanOption, causing the
	// newly-created Span to be a root span even if `ctx` contains a Span.
	//
	// When creating a Span it is recommended to provide all known span attributes using
	// the `WithAttributes()` SpanOption as samplers will only have access to the
	// attributes provided when a Span is created.
	//
	// Any Span that is created MUST also be ended. This is the responsibility of the user.
	StartSpan(ctx context.Context, spanName string, opts ...SpanStartOption) (context.Context, Span)
}

type Span interface {
	// End completes the Span. The Span is considered complete and ready to be
	// delivered through the rest of the telemetry pipeline after this method
	// is called. Therefore, updates to the Span are not allowed after this
	// method has been called.
	End(opts ...SpanEndOption)
	// IsRecording returns the recording state of the Span. It will return
	// true if the Span is active and events can be recorded.
	IsRecording() bool
	// RecordError will record err as an exception span event for this span.
	RecordError(err error)
	// SetAttributes sets kv as attributes of the Span. If a key from kv
	// already exists for an attribute of the Span it will be overwritten with
	// the value contained in kv.
	SetAttributes(attributes ...KeyValue)
	// AddEvent adds an event with the provided name and options.
	AddEvent(name string, opts ...EventOption)
}
