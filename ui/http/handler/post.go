// Package handler is ui layer http-handler package
package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"go.uber.org/zap"

	"github.com/oshou/AwesomeMusic-api/log"
	"github.com/oshou/AwesomeMusic-api/usecase"
)

// IPostHandler is ui layer http-handler interface
type IPostHandler interface {
	GetPosts(w http.ResponseWriter, r *http.Request)
	GetPostByID(w http.ResponseWriter, r *http.Request)
	GetPostsByTagID(w http.ResponseWriter, r *http.Request)
	GetPostsByUserID(w http.ResponseWriter, r *http.Request)
	AddPost(w http.ResponseWriter, r *http.Request)
	DeletePostByID(w http.ResponseWriter, r *http.Request)
}

type postHandler struct {
	usecase usecase.IPostUsecase
}

var _ IPostHandler = &postHandler{}

// NewPostHandler is constructor for postHandler
func NewPostHandler(usecase usecase.IPostUsecase) IPostHandler {
	return &postHandler{
		usecase: usecase,
	}
}

func (ph *postHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := ph.usecase.GetPosts()
	if err != nil {
		log.Logger.Error("failed to get posts", zap.Error(err))
		w.WriteHeader(http.StatusNotFound)

		return
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(posts); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (ph *postHandler) GetPostByID(w http.ResponseWriter, r *http.Request) {
	postIDString := chi.URLParam(r, "post_id")
	postID, err := strconv.Atoi(postIDString)

	if err != nil {
		log.Logger.Error("failed to convert string", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	post, err := ph.usecase.GetPostByID(postID)
	if err != nil {
		log.Logger.Error("failed to get posts by postID", zap.Error(err))
		w.WriteHeader(http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(post); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (ph *postHandler) GetPostsByTagID(w http.ResponseWriter, r *http.Request) {
	tagIDString := chi.URLParam(r, "tag_id")
	tagID, err := strconv.Atoi(tagIDString)

	if err != nil {
		log.Logger.Error("failed to convert string", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	posts, err := ph.usecase.GetPostsByTagID(tagID)
	if err != nil {
		log.Logger.Error("failed to get posts by tagID", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(posts); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (ph *postHandler) GetPostsByUserID(w http.ResponseWriter, r *http.Request) {
	userIDString := chi.URLParam(r, "user_Id")
	userID, err := strconv.Atoi(userIDString)

	if err != nil {
		log.Logger.Error("failed to convert string", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	posts, err := ph.usecase.GetPostsByUserID(userID)
	if err != nil {
		log.Logger.Error("failed to get posts by userID", zap.Error(err))
		w.WriteHeader(http.StatusNotFound)

		return
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(posts); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (ph *postHandler) AddPost(w http.ResponseWriter, r *http.Request) {
	userIDString := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(userIDString)

	if err != nil {
		log.Logger.Error("failed to convert string", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	title := r.URL.Query().Get("title")
	url := r.URL.Query().Get("url")
	message := r.URL.Query().Get("message")

	post, err := ph.usecase.AddPost(userID, title, url, message)
	if err != nil {
		log.Logger.Error("failed to add post", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(post); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (ph *postHandler) DeletePostByID(w http.ResponseWriter, r *http.Request) {
	postIDString := chi.URLParam(r, "post_id")
	postID, err := strconv.Atoi(postIDString)

	if err != nil {
		log.Logger.Error("failed to convert string", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusNoContent)

	if err := ph.usecase.DeletePostByID(postID); err != nil {
		log.Logger.Error("failed to delete post by postID", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)

		return
	}
}
