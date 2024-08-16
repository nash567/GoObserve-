package config_test

import (
	"github.com/nash-567/goObserve/pkg/logger/config"
	"log/slog"
	"testing"

	logModel "github.com/nash-567/goObserve/pkg/logger/model"
)

func TestConfig_GetLevel(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		level string
		want  logModel.Level
	}{
		{name: "Debug", level: logModel.DebugLevel.String(), want: logModel.DebugLevel},
		{name: "Info", level: logModel.InfoLevel.String(), want: logModel.InfoLevel},
		{name: "Warn", level: logModel.WarnLevel.String(), want: logModel.WarnLevel},
		{name: "Error", level: logModel.ErrorLevel.String(), want: logModel.ErrorLevel},
		{name: "invalid", level: "invalid", want: logModel.InfoLevel},
	}
	for _, tC := range tests {
		tt := tC
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := &config.Config{
				Level: tt.level,
			}
			if got := c.GetLevel(); got != tt.want {
				t.Errorf("GetLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_GetSlogLevel(t *testing.T) {
	t.Parallel()
	type fields struct {
		Level string
	}
	tests := []struct {
		name   string
		fields fields
		want   slog.Level
	}{
		{name: "Debug", fields: fields{Level: logModel.DebugLevel.String()}, want: slog.LevelDebug},
		{name: "Info", fields: fields{Level: logModel.InfoLevel.String()}, want: slog.LevelInfo},
		{name: "Warn", fields: fields{Level: logModel.WarnLevel.String()}, want: slog.LevelWarn},
		{name: "Error", fields: fields{Level: logModel.ErrorLevel.String()}, want: slog.LevelError},
	}
	for _, tC := range tests {
		tt := tC
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := &config.Config{
				Level: tt.fields.Level,
			}
			if got := c.GetSlogLevel(); got != tt.want {
				t.Errorf("GetSlogLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
