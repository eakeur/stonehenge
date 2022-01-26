package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"stonehenge/app/config"
	contract "stonehenge/app/core/types/logger"
	"time"
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

func (m logger) log(event *zerolog.Event, operation string, message string){
	event.
		Time("time", time.Now()).
		Str("operation", operation).
		Msg(message)
}

func (m logger) Trace(operation string, message string) {
	m.log(log.Trace(), operation, message)
}

func (m logger) Debug(operation string, message string) {
	m.log(log.Debug(), operation, message)
}

func (m logger) Info(operation string, message string) {
	m.log(log.Info(), operation, message)
}

func (m logger) Warn(operation string, message string) {
	m.log(log.Warn(), operation, message)
}

func (m logger) Error(operation string, message string) {
	m.log(log.Error(), operation, message)
}

func (m logger) Fatal(operation string, message string) {
	m.log(log.Fatal(), operation, message)
}

func (m logger) Panic(operation string, message string) {
	m.log(log.Panic(), operation, message)
}
