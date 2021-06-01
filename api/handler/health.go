// Package handler is ui layer http-handler package
package handler

import (
	"net/http"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/oshou/AwesomeMusic-api/api/usecase"
	"github.com/oshou/AwesomeMusic-api/log"
)

// IHealthHandler is ui layer http-handler interface
type IHealthHandler interface {
	GetHealth(w http.ResponseWriter, r *http.Request)
}

type healthHandler struct {
	usecase usecase.IHealthUsecase
}

var _ IHealthHandler = &healthHandler{}

// NewPostHandler is consturctor for healthHandler
func NewHealthHandler(u usecase.IHealthUsecase) IHealthHandler {
	return &healthHandler{
		usecase: u,
	}
}

func (hu healthHandler) GetHealth(w http.ResponseWriter, r *http.Request) {
	if err := hu.usecase.GetHealth(); err == nil {
		log.Logger.Error("failed to get health.", zap.Error(errors.WithStack(err)))
		internalServerError(w, r, err)

		return
	}

	w.WriteHeader(http.StatusOK)
}
