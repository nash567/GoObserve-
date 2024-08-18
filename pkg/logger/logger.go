package logger

import (
	"context"
	"github.com/nash-567/goObserve/pkg/logger/config"
	"github.com/nash-567/goObserve/pkg/logger/model"
	"log/slog"
	"os"
	"sync"
)

//nolint:gochecknoglobals // to be used to return default logger when not set in context
var (
	mux        = &sync.Mutex{}
	defaultLog *SlogLogger
	levelInfo  = "INFO"
)

func defaultLogger() *SlogLogger {
	mux.Lock()
	defer mux.Unlock()

	if defaultLog == nil {
		defaultLog = NewSlogLogger(&config.Config{
			Level:  levelInfo,
			Output: os.Stdout,
		})
	}
	return defaultLog
}

// SlogLogger is the default implementation of Logger. It is backed by the slog logging package.
type SlogLogger struct {
	entry *slog.Logger
	cfg   *config.Config
	level *slog.LevelVar
}

func NewSlogLogger(config *config.Config) *SlogLogger {
	loggingLevel := new(slog.LevelVar)
	loggingLevel.Set(config.GetSlogLevel())
	s := &SlogLogger{
		entry: buildLogger(config, loggingLevel),
		cfg:   config,
		level: loggingLevel,
	}

	return s
}

func buildLogger(config *config.Config, level slog.Leveler) *slog.Logger {
	handler := slog.NewJSONHandler(config.Output, &slog.HandlerOptions{
		AddSource:   config.IncludeSource,
		Level:       level,
		ReplaceAttr: replaceAttribute,
	})

	l := slog.New(handler)

	// output from the log package's default Logger (as with log.Print, etc.) will be logged using slog Handler
	slog.SetDefault(l)
	return l
}

// func to map the custom log levels to their respective labels
// e.g.-> slog doesn't have FATAL level.
func replaceAttribute(_ []string, a slog.Attr) slog.Attr {
	if a.Key == slog.LevelKey {
		level, ok := a.Value.Any().(slog.Level)
		if !ok {
			return a
		}
		levelLabel, exists := getCustomLevelMap()[level]
		if !exists {
			levelLabel = level.String()
		}
		a.Value = slog.StringValue(levelLabel)
	}
	return a
}

// map contains custom slog levels.
func getCustomLevelMap() map[slog.Leveler]string {
	return map[slog.Leveler]string{
		model.LevelFatal: "FATAL",
	}
}

func (log *SlogLogger) Debug(msg string) {
	log.entry.Debug(msg)
}

func (log *SlogLogger) Info(msg string) {
	log.entry.Info(msg)
}

func (log *SlogLogger) Warn(msg string) {
	log.entry.Warn(msg)
}

func (log *SlogLogger) Fatal(msg string) {
	log.entry.Log(context.Background(), model.LevelFatal, msg)
	os.Exit(1)
}

func (log *SlogLogger) Error(msg string) {
	log.entry.Error(msg)
}

func (log *SlogLogger) WithField(key string, value interface{}) model.Logger {
	return &SlogLogger{
		entry: log.entry.With(key, value),
		level: log.level,
	}
}

//nolint:ireturn // implements model.Logger interface
func (log *SlogLogger) WithFields(fields model.Fields) model.Logger {
	sFields := make([]any, 0)
	for key, value := range fields {
		sFields = append(sFields, key, value)
	}
	return &SlogLogger{
		entry: log.entry.With(sFields...),
		level: log.level,
	}
}

//nolint:ireturn // implements model.Logger interface
func (log *SlogLogger) WithError(err error) model.Logger {
	return &SlogLogger{
		entry: log.entry.With("error", err),
		level: log.level,
	}
}

func (log *SlogLogger) SetLevel(lvl model.Level) {
	log.level.Set(lvl.SlogLevel())
}

func (log *SlogLogger) GetLevel() model.Level {
	return model.ParseLevel(log.level.Level().String())
}
