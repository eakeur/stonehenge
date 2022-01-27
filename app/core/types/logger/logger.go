package logger

import "context"

type TraceContextKey string

const TraceKey TraceContextKey = "logger-trace-key"

type Logger interface {
	Trace(ctx context.Context, operation string, message string)
	Debug(ctx context.Context, operation string, message string)
	Info(ctx context.Context, operation string, message string)
	Warn(ctx context.Context, operation string, message string)
	Error(ctx context.Context, operation string, message string)
	Fatal(ctx context.Context, operation string, message string)
	Panic(ctx context.Context, operation string, message string)
}
