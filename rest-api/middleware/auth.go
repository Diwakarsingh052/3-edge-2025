package middleware

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"
)

func (m *Mid) Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			err := errors.New("expected authorization header format: Bearer <token>")
			slog.Error("invalid authorization header format",
				//slog.String("Trace ID", traceId),
				slog.String("Error", err.Error()))

			// Respond with HTTP 401 Unauthorized if the header is invalid
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		token := parts[1]
		m.a.ValidateToken(token)
	}
}
