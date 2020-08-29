package middleware

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
}

func writeErrorResponse(w http.ResponseWriter, statusCode int) {
	e := &errorResponse{
		Message: http.StatusText(statusCode),
	}
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(e)
}

func unauthorizedError(w http.ResponseWriter) {
	writeErrorResponse(w, http.StatusUnauthorized)
}

func forbiddenError(w http.ResponseWriter) {
	writeErrorResponse(w, http.StatusForbidden)
}

func notFoundError(w http.ResponseWriter) {
	writeErrorResponse(w, http.StatusNotFound)
}

func internalServerError(w http.ResponseWriter) {
	writeErrorResponse(w, http.StatusInternalServerError)
}
