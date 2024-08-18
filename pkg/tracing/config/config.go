package config

import (
	"github.com/nash-567/goObserve/pkg/tracing/model"
	"time"
)

type TracingConfig struct {
	Enabled                bool                         `koanf:"Enabled"`
	InstrumentationLibrary InstrumentationLibraryConfig `koanf:"InstrumentationLibrary"`
	ExporterConfig         TraceExporterConfig          `koanf:"ExporterConfig"`
}

type InstrumentationLibraryConfig struct {
	Name      string `koanf:"Name"`
	SchemaURL string `koanf:"SchemaURL"`
	Version   string `koanf:"Version"`
}

// TraceExporterConfig is the configuration for the trace exporter.
type TraceExporterConfig struct {
	Type model.TraceExporterType `koanf:"Type"`
	// in case of stdout exporter, rest of the config is ignored
	EndpointURL  string                   `koanf:"EndpointURL"`
	Timeout      time.Duration            `koanf:"Timeout"`
	RetryConfig  TraceExporterRetryConfig `koanf:"RetryConfig"`
	BatchTimeout time.Duration            `koanf:"BatchTimeout"`
}

// TraceExporterRetryConfig is the configuration for the trace exporter retry.
type TraceExporterRetryConfig struct {
	Enabled         bool          `koanf:"Enabled"`
	InitialInterval time.Duration `koanf:"InitialInterval"`
	MaxInterval     time.Duration `koanf:"MaxInterval"`
	MaxElapsedTime  time.Duration `koanf:"MaxElapsedTime"`
}
