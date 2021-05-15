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

type AddCommentRequest struct {
	UserID  string `json:"user_id" validate:"required"`
	Comment string `json:"comment" validate:"required"`
}

// ICommentHandler is ui layer http-handler interface
type ICommentHandler interface {
	ListComments(w http.ResponseWriter, r *http.Request)
	GetCommentByID(w http.ResponseWriter, r *http.Request)
	AddComment(w http.ResponseWriter, r *http.Request)
}

type commentHandler struct {
	usecase usecase.ICommentUsecase
}

var _ ICommentHandler = &commentHandler{}

// NewCommentHandler is constructor for commentHandler
func NewCommentHandler(usecase usecase.ICommentUsecase) ICommentHandler {
	return &commentHandler{
		usecase: usecase,
	}
}

func (ch *commentHandler) ListComments(w http.ResponseWriter, r *http.Request) {
	postIDString := chi.URLParam(r, "post_id")
	postID, err := strconv.Atoi(postIDString)
	if err != nil {
		log.Logger.Error("failed to convert string", zap.Error(errors.WithStack(err)))
		badRequestError(w)

		return
	}

	comments, err := ch.usecase.ListComments(postID)
	if err != nil {
		log.Logger.Error("failed to list comments", zap.Error(errors.WithStack(err)))
		httpError(w, err)

		return
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(comments); err != nil {
		internalServerError(w)

		return
	}
}

func (ch *commentHandler) AddComment(w http.ResponseWriter, r *http.Request) {
	postIDString := chi.URLParam(r, "post_id")
	postID, err := strconv.Atoi(postIDString)
	if err != nil {
		log.Logger.Error("failed to convert string", zap.Error(errors.WithStack(err)))
		badRequestError(w)

		return
	}

	req := AddCommentRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Logger.Error("failed to add comment", zap.Error(errors.WithStack(err)))
		badRequestError(w)

		return
	}

	userID, err := strconv.Atoi(req.UserID)
	if err != nil {
		log.Logger.Error("failed to add comment", zap.Error(errors.WithStack(err)))
		internalServerError(w)

		return
	}

	comment, err := ch.usecase.AddComment(postID, userID, req.Comment)
	if err != nil {
		log.Logger.Error("failed to add comment", zap.Error(errors.WithStack(err)))
		httpError(w, err)

		return
	}

	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		log.Logger.Error("failed to add comment", zap.Error(errors.WithStack(err)))
		internalServerError(w)

		return
	}
}

func (ch *commentHandler) GetCommentByID(w http.ResponseWriter, r *http.Request) {
	commentIDString := chi.URLParam(r, "comment_id")
	commentID, err := strconv.Atoi(commentIDString)
	if err != nil {
		log.Logger.Error("failed to convert string", zap.Error(errors.WithStack(err)))
		badRequestError(w)

		return
	}

	comment, err := ch.usecase.GetCommentByID(commentID)
	if err != nil {
		log.Logger.Error("failed to get comment by commentID", zap.Error(errors.WithStack(err)))
		httpError(w, err)

		return
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		internalServerError(w)

		return
	}
}
