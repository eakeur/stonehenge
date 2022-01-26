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

func (m logger) log(event *zerolog.Event, operation string, message string, target interface{}){
	event.
		Time("time", time.Now()).
		Str("operation", operation).
		Interface("target", target).
		Msg(message)
}

func (m logger) Trace(operation string, message string, target interface{}) {
	m.log(log.Trace(), operation, message, target)
}

func (m logger) Debug(operation string, message string, target interface{}) {
	m.log(log.Debug(), operation, message, target)
}

func (m logger) Info(operation string, message string, target interface{}) {
	m.log(log.Info(), operation, message, target)
}

func (m logger) Warn(operation string, message string, target interface{}) {
	m.log(log.Warn(), operation, message, target)
}

func (m logger) Error(operation string, message string, target interface{}) {
	m.log(log.Error(), operation, message, target)
}

func (m logger) Fatal(operation string, message string, target interface{}) {
	m.log(log.Fatal(), operation, message, target)
}

func (m logger) Panic(operation string, message string, target interface{}) {
	m.log(log.Panic(), operation, message, target)
}
