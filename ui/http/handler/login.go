package handler

import (
	"encoding/json"
	"net/http"

	"github.com/oshou/AwesomeMusic-api/ui/http/session"
	"github.com/oshou/AwesomeMusic-api/usecase"
)

type ILoginHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

type loginHandler struct {
	usecase usecase.IUserUsecase
	store   *session.Store
}

type loginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

var _ ILoginHandler = &loginHandler{}

func NewLoginHandler(usecase usecase.IUserUsecase, store *session.Store) ILoginHandler {
	return &loginHandler{
		usecase: usecase,
		store:   store,
	}
}

func (h *loginHandler) Login(w http.ResponseWriter, r *http.Request) {
	req := loginRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		badRequestError(w)

		return
	}

	u, err := h.usecase.Authenticate(req.Name, req.Password)
	if err != nil {
		unauthorizedError(w)

		return
	}

	if err := h.store.Save(r, w, u.ID); err != nil {
		unauthorizedError(w)

		return
	}

	type (
		role struct {
			Name string `json:"name`
		}
		user struct {
			ID   int    `json:"id,omitempty" db:"id"`
			Name string `json:"name,omitempty" db:"name"`
		}
		response struct {
			User *user `json:"user"`
		}
	)

	resp := &response{
		User: &user{
			ID:   u.ID,
			Name: u.Name,
		},
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		internalServerError(w)

		return
	}
}

func (h *loginHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// TODO
}
