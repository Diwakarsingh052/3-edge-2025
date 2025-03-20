package handlers

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
	"rest-api/middleware"
	"rest-api/models"
)

func (h *handler) Signup(w http.ResponseWriter, r *http.Request) {
	// Set response content-type to JSON
	w.Header().Set("Content-Type", "application/json")
	// Extract the trace ID from the request context
	traceId := GetTraceIdOfRequest(r)

	if r.ContentLength > 5*1024 {
		slog.Error("request body limit exceeded",
			slog.String("TraceID", traceId), slog.Int64("Size Received", r.ContentLength))

		err := sendJsonResponse(w, http.StatusBadRequest,
			map[string]string{"error": "Request body too large. Limit is 5KB"})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	// Decode the JSON request into the NewUser model
	var newUser models.NewUser
	// NewDecoder would directly read the data from the request body, and after converting it would put the data in struct
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		slog.Error("JSON decoding error", slog.String("TraceID", traceId), slog.String("Error", err.Error()))
		err := sendJsonResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	// Validate the decoded JSON
	v := validator.New()
	// struct would be validated according to the field tags specified
	err = v.Struct(newUser)
	if err != nil {
		slog.Error("validation error", slog.String("TraceID", traceId), slog.String("Error", err.Error()))

		err := sendJsonResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	u, err := h.conn.CreateUser(newUser)

	if err != nil {
		slog.Error("error creating user", slog.String("TraceID", traceId), slog.String("Error", err.Error()))
		err := sendJsonResponse(w, http.StatusInternalServerError, map[string]string{"error": "signup failed"})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	err = sendJsonResponse(w, http.StatusOK, u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return

}

func GetTraceIdOfRequest(r *http.Request) string {
	// Get the current request context
	ctx := r.Context()

	// Extract the Trace ID from the context
	traceId, ok := ctx.Value(middleware.TraceIdKey).(string)

	// If Trace ID is not available, log an error and return "Unknown"
	if !ok {
		slog.Error("trace id not present in the context")
		return "Unknown"
	}
	return traceId

}

func sendJsonResponse(w http.ResponseWriter, status int, data any) error {
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return err
	}
	return nil
}
