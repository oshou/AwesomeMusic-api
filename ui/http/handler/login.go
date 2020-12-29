package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/sessions"
	"go.uber.org/zap"

	"github.com/oshou/AwesomeMusic-api/log"
	"github.com/oshou/AwesomeMusic-api/usecase"
)

type loginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type ILoginHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

type loginHandler struct {
	usecase usecase.IUserUsecase
	store   *session.IStore
}

type loginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

var _ ILoginHandler = &loginHandler{}

func NewLoginHandler(usecase usecase.IUserUsecase, store session.IStore) ILoginHandler {
	return &loginHandler{
		usecase: usecase,
		store:   &store,
	}
}

func (h *loginHandler) Login(w http.ResponseWriter, r *http.Request) {
	req := loginRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		badRequestError(w)

		return
	}

	user, err := h.usecase.Authenticate(req.Name, req.Password)
	if err != nil {
		unauthorizedError(w)

		return
	}

	if err := h.store.Set(r, w, user.ID); err != nil {
		unauthorizedError(w)

		return
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Logger.Error("failed to encode json")
	}
}

func (h *loginHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// TODO
}
