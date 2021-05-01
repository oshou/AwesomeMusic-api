// Package handler is ui layer http-handler package
package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"go.uber.org/zap"

	"github.com/oshou/AwesomeMusic-api/api/usecase"
	"github.com/oshou/AwesomeMusic-api/log"
)

type addPostRequest struct {
	UserID  string `json:"user_id" validate:"required"`
	Title   string `json:"title" validate:"required"`
	URL     string `json:"url" validate:"required"`
	Message string `json:"message" validate:"required"`
}

// IPostHandler is ui layer http-handler interface
type IPostHandler interface {
	ListPosts(w http.ResponseWriter, r *http.Request)
	ListPostsByTagID(w http.ResponseWriter, r *http.Request)
	ListPostsByUserID(w http.ResponseWriter, r *http.Request)
	GetPostByID(w http.ResponseWriter, r *http.Request)
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

func (ph *postHandler) ListPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := ph.usecase.ListPosts()
	if err != nil {
		log.Logger.Error("failed to get posts", zap.Error(err))
		internalServerError(w)

		return
	}

	w.WriteHeader(http.StatusOK)

	if len(posts) == 0 {
		log.Logger.Error("failed to get posts", zap.Error(err))
		notFoundError(w)

		return
	}

	if err := json.NewEncoder(w).Encode(posts); err != nil {
		log.Logger.Error("failed to get posts", zap.Error(err))
		internalServerError(w)
	}
}

func (ph *postHandler) GetPostByID(w http.ResponseWriter, r *http.Request) {
	postIDString := chi.URLParam(r, "post_id")
	postID, err := strconv.Atoi(postIDString)

	if err != nil {
		log.Logger.Error("failed to convert string", zap.Error(err))
		badRequestError(w)

		return
	}

	post, err := ph.usecase.GetPostByID(postID)
	if err != nil {
		log.Logger.Error("failed to get posts by postID", zap.Error(err))
		notFoundError(w)
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(post); err != nil {
		internalServerError(w)

		return
	}
}

func (ph *postHandler) ListPostsByTagID(w http.ResponseWriter, r *http.Request) {
	tagIDString := chi.URLParam(r, "tag_id")
	tagID, err := strconv.Atoi(tagIDString)

	if err != nil {
		log.Logger.Error("failed to convert string", zap.Error(err))
		badRequestError(w)

		return
	}

	posts, err := ph.usecase.ListPostsByTagID(tagID)
	if err != nil {
		log.Logger.Error("failed to get posts by tagID", zap.Error(err))
		badRequestError(w)

		return
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(posts); err != nil {
		internalServerError(w)

		return
	}
}

func (ph *postHandler) ListPostsByUserID(w http.ResponseWriter, r *http.Request) {
	userIDString := chi.URLParam(r, "user_Id")
	userID, err := strconv.Atoi(userIDString)

	if err != nil {
		log.Logger.Error("failed to convert string", zap.Error(err))
		badRequestError(w)

		return
	}

	posts, err := ph.usecase.ListPostsByUserID(userID)
	if err != nil {
		log.Logger.Error("failed to get posts by userID", zap.Error(err))
		notFoundError(w)

		return
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(posts); err != nil {
		internalServerError(w)

		return
	}
}

func (ph *postHandler) AddPost(w http.ResponseWriter, r *http.Request) {
	req := addPostRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Logger.Error("failed to add post", zap.Error(err))
		badRequestError(w)

		return
	}

	userID, err := strconv.Atoi(req.UserID)
	if err != nil {
		internalServerError(w)

		return
	}

	post, err := ph.usecase.AddPost(userID, req.Title, req.URL, req.Message)
	if err != nil {
		log.Logger.Error("failed to add post", zap.Error(err))
		badRequestError(w)

		return
	}

	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(post); err != nil {
		internalServerError(w)

		return
	}
}

func (ph *postHandler) DeletePostByID(w http.ResponseWriter, r *http.Request) {
	postIDString := chi.URLParam(r, "post_id")
	postID, err := strconv.Atoi(postIDString)

	if err != nil {
		log.Logger.Error("failed to convert string", zap.Error(err))
		badRequestError(w)

		return
	}

	w.WriteHeader(http.StatusNoContent)

	if err := ph.usecase.DeletePostByID(postID); err != nil {
		log.Logger.Error("failed to delete post by postID", zap.Error(err))
		badRequestError(w)

		return
	}
}
