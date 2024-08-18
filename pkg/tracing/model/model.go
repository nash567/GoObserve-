package model

import (
	"fmt"
	"go.opentelemetry.io/otel/attribute"
)

// TraceExporterType is an enum for the type of trace exporter.
type TraceExporterType int8

const (
	TraceExporterTypeStdout TraceExporterType = iota
	TraceExporterTypeHTTP
)

//go:generate enumer -type=TraceExporterType -json -text -yaml -trimprefix=TraceExporterType -transform=snake -output=enum_traceexportertype_gen.go
type Value interface {
	int | int64 | float64 | bool | string | []int64 | []float64 | []bool | []string
}

type KeyValue struct {
	keyValue attribute.KeyValue
}

func NewKeyValue[T Value](key string, value T) KeyValue {
	return KeyValue{
		keyValue: toAttributeKeyValue(key, value),
	}
}

func (kv KeyValue) GetAttributeKeyValue() attribute.KeyValue {
	return kv.keyValue
}

func toAttributeKeyValue[T Value](key string, value T) attribute.KeyValue {
	switch v := any(value).(type) {
	case int64:
		return attribute.Int64(key, v)
	case float64:
		return attribute.Float64(key, v)
	case bool:
		return attribute.Bool(key, v)
	case string:
		return attribute.String(key, v)
	case []int64:
		return attribute.Int64Slice(key, v)
	case []float64:
		return attribute.Float64Slice(key, v)
	case []bool:
		return attribute.BoolSlice(key, v)
	case []string:
		return attribute.StringSlice(key, v)
	default:
		return attribute.String(key, fmt.Sprintf("%v", v))
	}
}
