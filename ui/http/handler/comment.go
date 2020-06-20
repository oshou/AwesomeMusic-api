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

// ICommentHandler is ui layer http-handler interface
type ICommentHandler interface {
	GetComments(w http.ResponseWriter, r *http.Request)
	GetCommentByID(w http.ResponseWriter, r *http.Request)
	AddComment(w http.ResponseWriter, r *http.Request)
}

type commentHandler struct {
	svc usecase.ICommentUsecase
}

var _ ICommentHandler = &commentHandler{}

// NewCommentHandler is constructor for commentHandler
func NewCommentHandler(svc usecase.ICommentUsecase) ICommentHandler {
	return &commentHandler{
		svc: svc,
	}
}

func (ch *commentHandler) GetComments(w http.ResponseWriter, r *http.Request) {
	postIDString := chi.URLParam(r, "post_id")
	postID, err := strconv.Atoi(postIDString)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	comments, err := ch.svc.GetComments(postID)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusNotFound)

		return
	}

	if err := json.NewEncoder(w).Encode(comments); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (ch *commentHandler) AddComment(w http.ResponseWriter, r *http.Request) {
	postIDString := chi.URLParam(r, "post_id")
	postID, err := strconv.Atoi(postIDString)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	userIDString := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	commentString := r.URL.Query().Get("comment")
	comment, err := ch.svc.AddComment(postID, userID, commentString)

	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (ch *commentHandler) GetCommentByID(w http.ResponseWriter, r *http.Request) {
	commentIDString := chi.URLParam(r, "comment_id")
	commentID, err := strconv.Atoi(commentIDString)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	comment, err := ch.svc.GetCommentByID(commentID)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusNotFound)

		return
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}
