package middleware

import (
	"context"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
	"time"
)

type key string

const TraceIdKey key = "1"

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		traceId := uuid.NewString()
		requestStartTime := time.Now()

		// taking the context out from the request object
		ctx := r.Context()
		// putting the traceId in the context
		ctx = context.WithValue(ctx, TraceIdKey, traceId)

		slog.Info("started", slog.String("Trace ID", traceId),
			slog.String("Method", r.Method), slog.String("URL Path", r.URL.Path),
		)
		// r.WithContext would update the request to use the updated context
		next(w, r.WithContext(ctx)) // call the next thing in the chain

		slog.Info("completed", slog.String("Trace ID", traceId),
			slog.String("Method", r.Method), slog.String("URL Path", r.URL.Path),
			slog.Int64("duration μs",
				time.Since(requestStartTime).Microseconds()),
		)
	}
}
