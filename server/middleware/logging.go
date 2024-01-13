package middleware

import (
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"
	"log/slog"
	"net/http"
	"time"
)

func RequestLogger(logger *slog.Logger) func(handler http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				start := time.Now()
				traceID := trace.SpanFromContext(r.Context()).SpanContext().TraceID()
				var reqID string
				if id, err := uuid.Parse(r.Header.Get("x-request-id")); err == nil {
					reqID = id.String()
				} else {
					reqID = uuid.NewString()
				}

				l := logger.With("reqId", reqID, "traceId", traceID)
				handler.ServeHTTP(w, r)
				l.Info(
					"Request handled",
					slog.String("method", r.Method),
					slog.String("method", r.Method),
					slog.String("path", r.URL.Path),
					slog.String("ua", r.UserAgent()),
					slog.String("ip", r.RemoteAddr),
					slog.Duration("duration", time.Since(start)),
				)
			},
		)
	}
}
