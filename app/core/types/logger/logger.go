package logger

import "context"

type TraceContextKey string

const LoggerTraceKey TraceContextKey = "logger-trace-key"

type Logger interface {
	Trace(context.Context, string, string)
	Debug(context.Context, string, string)
	Info(context.Context, string, string)
	Warn(context.Context, string, string)
	Error(context.Context, string, string)
	Fatal(context.Context, string, string)
	Panic(context.Context, string, string)
}
