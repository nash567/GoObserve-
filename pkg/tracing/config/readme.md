# Configuration

This document explains the configuration options for tracing.


## TracingConfig

The main configuration struct for the tracing system.

| Field | Type | Description |
|-------|------|-------------|
| Enabled | bool | Enables or disables tracing. Set to `true` to turn on tracing, `false` to turn it off. |
| InstrumentationLibrary | InstrumentationLibraryConfig | Configuration for the instrumentation library. |
| ExporterConfig | TraceExporterConfig | Configuration for the trace exporter. |

## InstrumentationLibraryConfig

Provides details about the library adding instrumentation to the application.

| Field | Type | Description |
|-------|------|-------------|
| Name | string | The name of the instrumentation library. |
| SchemaURL | string | The URL of the OpenTelemetry schema being used. |
| Version | string | The version of the instrumentation library. |

## TraceExporterConfig

Configuration for the trace exporter.

| Field | Type | Description |
|-------|------|-------------|
| Type | model.TraceExporterType |  The type of exporter. Currently supports two types: "stdout" (writes traces to console) and "http" (exports traces to a specified endpoint).|
| EndpointURL | string | The URL to which traces are exported. Default is "http://localhost:4318" for both HTTP and gRPC, with "/v1/traces" as the path. |
| Timeout | time.Duration | The timeout duration for HTTP calls made by the exporter. |
| BatchTimeout | time.Duration | The maximum delay allowed before the exporter exports any held spans. |
| RetryConfig | TraceExporterRetryConfig | Configuration for the exporter's retry mechanism. |

## TraceExporterRetryConfig

Configuration for the trace exporter's retry mechanism.

| Field | Type | Description |
|-------|------|-------------|
| Enabled | bool | Enables or disables the retry mechanism. |
| InitialInterval | time.Duration | The initial interval between retry attempts. |
| MaxInterval | time.Duration | The maximum interval between retry attempts. |
| MaxElapsedTime | time.Duration | The maximum total time spent on retries. |

