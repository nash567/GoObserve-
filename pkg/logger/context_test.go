package logger_test

import (
	"context"
	"github.com/nash-567/goObserve/pkg/logger"
	"github.com/nash-567/goObserve/pkg/logger/config"
	"github.com/nash-567/goObserve/pkg/logger/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromContext(t *testing.T) {
	t.Parallel()
	type args struct {
		log *logger.SlogLogger
	}
	tests := []struct {
		name       string
		args       args
		wantCtxLog bool
	}{
		{
			name: "logger set in context",
			args: args{
				log: logger.NewSlogLogger(&config.Config{
					Level: "INFO",
				}),
			},

			wantCtxLog: true,
		},
		{
			name: "logger not set in context use default logger",
			args: args{
				log: nil,
			},
			wantCtxLog: false,
		},
	}
	for _, tC := range tests {
		tt := tC
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.WithValue(context.Background(), model.ContextKeyLogger, tt.args.log)
			got := logger.FromContext(ctx)
			if tt.wantCtxLog {
				assert.Same(t, tt.args.log, got)
				return
			}
			assert.Same(t, tt.args.log, got)
		})
	}
}

func TestNewContextWithLogger(t *testing.T) {
	t.Parallel()
	log, _ := makeTestLogger()
	gotCtx := logger.NewContextWithLogger(context.Background(), log)
	assert.NotNil(t, gotCtx)
	got := gotCtx.Value(model.ContextKeyLogger)
	assert.NotNil(t, got)
	assert.Same(t, got, log)
}
