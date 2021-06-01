// Package handler is ui layer http-handler package
package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/oshou/AwesomeMusic-api/api/usecase"
	"github.com/oshou/AwesomeMusic-api/log"
)

type addUserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// IUserHandler is ui layer http-handler interface
type IUserHandler interface {
	ListUsers(w http.ResponseWriter, r *http.Request)
	GetUserByID(w http.ResponseWriter, r *http.Request)
	AddUser(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	usecase usecase.IUserUsecase
}

var _ IUserHandler = &userHandler{}

// NewUserHandler is constructor for userHandler
func NewUserHandler(usecase usecase.IUserUsecase) IUserHandler {
	return &userHandler{
		usecase: usecase,
	}
}

func (uh *userHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uh.usecase.ListUsers()
	if err != nil {
		log.Logger.Error("failed to get users", zap.Error(errors.WithStack(err)))
		httpError(w, r, err)

		return
	}

	if len(users) == 0 {
		log.Logger.Error("failed to get users", zap.Error(errors.WithStack(err)))
		notFoundError(w)

		return
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(users); err != nil {
		log.Logger.Error("failed to get users", zap.Error(errors.WithStack(err)))
		internalServerError(w, r, err)

		return
	}
}

// Create: POST /v1/users
func (uh *userHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	req := addUserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Logger.Error("failed to add user", zap.Error(errors.WithStack(err)))
		badRequestError(w)

		return
	}

	user, err := uh.usecase.AddUser(req.Name, req.Password)
	if err != nil {
		log.Logger.Error("failed to add user", zap.Error(errors.WithStack(err)))
		httpError(w, r, err)

		return
	}

	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(user); err != nil {
		internalServerError(w, r, err)

		return
	}
}

// Detail: GET /v1/users/:user_id
func (uh *userHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	userIDString := chi.URLParam(r, "user_id")
	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		log.Logger.Error("failed to convert string", zap.Error(errors.WithStack(err)))
		badRequestError(w)

		return
	}

	user, err := uh.usecase.GetUserByID(userID)
	if err != nil {
		log.Logger.Error("failed to get user by userID", zap.Error(errors.WithStack(err)))
		httpError(w, r, err)

		return
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Logger.Error("failed to get user by userID", zap.Error(errors.WithStack(err)))
		internalServerError(w, r, err)

		return
	}
}
