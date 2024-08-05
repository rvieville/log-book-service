package dashlogger

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

type DashLogger struct {
	traceID uuid.UUID
	logger  zerolog.Logger
}

func New(traceID uuid.UUID) *DashLogger {
	logger := zerolog.New(os.Stdout).With().Str("trace_id", traceID.String()).Timestamp().Logger()
	return &DashLogger{
		traceID,
		logger,
	}
}

func (l DashLogger) Info(log LogMessage) {
	l.logger.Info().Interface("log", log).Send()
}

func (l DashLogger) Warn(log LogMessage) {
	l.logger.Warn().Interface("log", log).Send()
}

func (l DashLogger) Error(log LogMessage) {
	l.logger.Error().Interface("log", log).Send()
}
