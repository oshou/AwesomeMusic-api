package middleware

import (
	"encoding/json"
	"net/http"
)

type errorMessage struct {
	Message string `json:"message"`
}

func errorResponse(w http.ResponseWriter, statusCode int) {
	e := &errorMessage{
		Message: http.StatusText(statusCode),
	}
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(e)
}

func unauthorizedError(w http.ResponseWriter) {
	errorResponse(w, http.StatusUnauthorized)
}

func forbiddenError(w http.ResponseWriter) {
	errorResponse(w, http.StatusForbidden)
}

func notFoundError(w http.ResponseWriter) {
	errorResponse(w, http.StatusNotFound)
}

func internalServerError(w http.ResponseWriter) {
	errorResponse(w, http.StatusInternalServerError)
}
