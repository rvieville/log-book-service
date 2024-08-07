package middlewares

import (
	"context"
	clogger "diving-log-book-service/internal/pkg/logger"
	"diving-log-book-service/internal/pkg/mux"
	"net/http"

	"github.com/google/uuid"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create an instance of CustomResponseWriter
		uuid := uuid.New()
		logger := clogger.New(uuid)
		ctx := context.WithValue(r.Context(), mux.LoggerKey, logger)

		// Call the next handler with the custom ResponseWriter
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
