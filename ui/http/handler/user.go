// Package handler is ui layer http-handler package
package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/oshou/AwesomeMusic-api/usecase"
)

// IUserHandler is ui layer http-handler interface
type IUserHandler interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
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

func (uh *userHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users, err := uh.usecase.GetUsers()
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}

// Create: POST /v1/users
func (uh *userHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	user, err := uh.usecase.AddUser(name)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}

// Detail: GET /v1/users/:user_id
func (uh *userHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	userIDString := chi.URLParam(r, "user_id")
	userID, err := strconv.Atoi(userIDString)

	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	user, err := uh.usecase.GetUserByID(userID)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusNotFound)

		return
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}
