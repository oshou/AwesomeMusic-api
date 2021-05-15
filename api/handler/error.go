package handler

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"github.com/oshou/AwesomeMusic-api/api/usecase"
	"github.com/oshou/AwesomeMusic-api/log"
)

type errorMessage struct {
	Message string `json:"message"`
}

func errorResponse(w http.ResponseWriter, statusCode int) {
	e := &errorMessage{
		Message: http.StatusText(statusCode),
	}

	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(e)
	if err != nil {
		log.Logger.Error("failed to json encode", zap.Error(err))
	}
}

func httpError(w http.ResponseWriter, err error) {
	switch err.(type) {
	case usecase.InvalidParamError:
		badRequestError(w)
	case usecase.UnauthorizedError:
		unauthorizedError(w)
	case usecase.ForbiddenError:
		forbiddenError(w)
	case usecase.NotFoundError:
		notFoundError(w)
	case usecase.ConflictError:
		confictError(w)
	default:
		internalServerError(w)
	}
}

func badRequestError(w http.ResponseWriter) {
	errorResponse(w, http.StatusBadRequest)
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

func confictError(w http.ResponseWriter) {
	errorResponse(w, http.StatusConflict)
}

func internalServerError(w http.ResponseWriter) {
	errorResponse(w, http.StatusInternalServerError)
}
