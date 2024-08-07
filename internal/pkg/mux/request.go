package mux

import (
	"context"
	clogger "diving-log-book-service/internal/pkg/logger"
	"net/http"
)

type contextKey string

const LoggerKey = contextKey("logger")

type Request struct {
	*http.Request
	Logger *clogger.Logger
}

func GetLoggerFromContext(ctx context.Context) *clogger.Logger {
	if logger, ok := ctx.Value(LoggerKey).(*clogger.Logger); ok {
		return logger
	}
	return nil
}
