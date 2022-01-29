package logger

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"stonehenge/app/config"
	contract "stonehenge/app/core/types/logger"
)

func NewLogger(cfg config.LoggerConfigurations) contract.Logger {
	if cfg.Environment == "development" {
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	}
	if cfg.Environment == "production" {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	return logger{}
}

type logger struct {
}

func (m logger) log(ctx context.Context, event *zerolog.Event, operation string, message string) {
	res := ctx.Value(contract.TraceKey)
	var traceID string
	if res != nil {
		traceID = res.(string)
	}
	event.
		Str("req_id", traceID).
		Str("operation", operation).
		Msg(message)
}

func (m logger) Trace(ctx context.Context, operation string, message string) {
	m.log(ctx, log.Trace(), operation, message)
}

func (m logger) Debug(ctx context.Context, operation string, message string) {
	m.log(ctx, log.Debug(), operation, message)
}

func (m logger) Info(ctx context.Context, operation string, message string) {
	m.log(ctx, log.Info(), operation, message)
}

func (m logger) Warn(ctx context.Context, operation string, message string) {
	m.log(ctx, log.Warn(), operation, message)
}

func (m logger) Error(ctx context.Context, operation string, message string) {
	m.log(ctx, log.Error(), operation, message)
}

func (m logger) Fatal(ctx context.Context, operation string, message string) {
	m.log(ctx, log.Fatal(), operation, message)
}

func (m logger) Panic(ctx context.Context, operation string, message string) {
	m.log(ctx, log.Panic(), operation, message)
}
