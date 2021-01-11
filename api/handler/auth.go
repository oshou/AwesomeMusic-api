package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/sessions"
	"go.uber.org/zap"

	"github.com/oshou/AwesomeMusic-api/api/usecase"
	"github.com/oshou/AwesomeMusic-api/log"
)

const (
	sessionKey = "aw-session-id"
)

type loginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type IAuthHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	usecase usecase.IUserUsecase
	store   sessions.Store
}

var _ IAuthHandler = &authHandler{}

func NewAuthHandler(u usecase.IUserUsecase, s sessions.Store) IAuthHandler {
	return &authHandler{
		usecase: u,
		store:   s,
	}
}

func (h *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	req := loginRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Logger.Error("failed to decode request.", zap.Error(err))
		badRequestError(w)

		return
	}

	authenticatedUser, err := h.usecase.Authenticate(req.Name, req.Password)
	if err != nil {
		log.Logger.Error("failed to auth request.", zap.Error(err))
		unauthorizedError(w)

		return
	}

	session, err := h.store.Get(r, sessionKey)
	if err != nil || session.ID == "" {
		session, err := h.store.New(r, sessionKey)
		if err != nil {
			log.Logger.Error("failed to create new session.", zap.Error(err))
			internalServerError(w)

			return
		}

		if err := h.store.Save(r, w, session); err != nil {
			log.Logger.Error("failed to save new session.", zap.Error(err))
			internalServerError(w)

			return
		}
	}

	if err := json.NewEncoder(w).Encode(authenticatedUser); err != nil {
		log.Logger.Error("failed to encode json.", zap.Error(err))
		internalServerError(w)

		return
	}
}

func (h *authHandler) Logout(w http.ResponseWriter, r *http.Request) {
	session, err := h.store.Get(r, sessionKey)
	if err != nil || session.ID == "" {
		log.Logger.Error("failed to get session.", zap.Error(err))
		badRequestError(w)

		return
	}

	session.Options.MaxAge = -1
	if err := h.store.Save(r, w, session); err != nil {
		log.Logger.Error("failed to expire session.", zap.Error(err))
		internalServerError(w)

		return
	}

	w.WriteHeader(http.StatusOK)
}
