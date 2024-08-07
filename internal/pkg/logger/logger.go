package clogger

import (
	"os"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type Datatype map[string]interface{}

type LogMessage struct {
	Event string `json:"event"`
	Msg   string `json:"msg"`
	Data  any    `json:"data,omitempty"`
	Error any    `json:"error,omitempty"`
}

type Logger struct {
	traceID uuid.UUID
	logger  zerolog.Logger
}

func New(traceID uuid.UUID) *Logger {
	logger := zerolog.New(os.Stdout).With().Str("trace_id", traceID.String()).Timestamp().Logger()
	return &Logger{
		traceID,
		logger,
	}
}

func (l Logger) Info(log LogMessage) {
	l.logger.Info().Interface("log", log).Send()
}

func (l Logger) Warn(log LogMessage) {
	l.logger.Warn().Interface("log", log).Send()
}

func (l Logger) Error(log LogMessage) {
	l.logger.Error().Interface("log", log).Send()
}
