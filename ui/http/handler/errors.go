package handler

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"github.com/oshou/AwesomeMusic-api/log"
)

type errorResponse struct {
	Message string `json:"message"`
}

func writeErrorResponse(w http.ResponseWriter, statusCode int) {
	e := &errorResponse{
		Message: http.StatusText(statusCode),
	}

	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(e)
	if err != nil {
		log.Logger.Error("failed to json encode", zap.Error(err))
	}
}

func unauthorizedError(w http.ResponseWriter) {
	writeErrorResponse(w, http.StatusUnauthorized)
}

func internalServerError(w http.ResponseWriter) {
	writeErrorResponse(w, http.StatusInternalServerError)
}

func badRequestError(w http.ResponseWriter) {
	writeErrorResponse(w, http.StatusBadRequest)
}

func notFoundError(w http.ResponseWriter) {
	writeErrorResponse(w, http.StatusNotFound)
}
