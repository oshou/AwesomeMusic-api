// Package handler is ui layer http-handler package
package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/oshou/AwesomeMusic-api/service"
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
	svc service.IPostService
}

var _ IPostHandler = &postHandler{}

// NewPostHandler is constructor for postHandler
func NewPostHandler(svc service.IPostService) IPostHandler {
	return &postHandler{
		svc: svc,
	}
}

func (ph *postHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := ph.svc.GetPosts()
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusNotFound)

		return
	}

	if err := json.NewEncoder(w).Encode(posts); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	//ctx.JSON(http.StatusOK, posts)
}

func (ph *postHandler) GetPostByID(w http.ResponseWriter, r *http.Request) {
	postIDString := chi.URLParam(r, "post_id")
	postID, err := strconv.Atoi(postIDString)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	post, err := ph.svc.GetPostByID(postID)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusNotFound)
	}

	if err := json.NewEncoder(w).Encode(post); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	//ctx.JSON(http.StatusOK, post)
}

func (ph *postHandler) GetPostsByTagID(w http.ResponseWriter, r *http.Request) {
	tagIDString := chi.URLParam(r, "tag_id")
	tagID, err := strconv.Atoi(tagIDString)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	posts, err := ph.svc.GetPostsByTagID(tagID)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	if err := json.NewEncoder(w).Encode(posts); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	//ctx.JSON(http.StatusOK, posts)
}

func (ph *postHandler) GetPostsByUserID(w http.ResponseWriter, r *http.Request) {
	userIDString := chi.URLParam(r, "user_Id")
	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	posts, err := ph.svc.GetPostsByUserID(userID)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusNotFound)

		return
	}

	if err := json.NewEncoder(w).Encode(posts); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (ph *postHandler) AddPost(w http.ResponseWriter, r *http.Request) {
	userIDString := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	title := r.URL.Query().Get("title")
	url := r.URL.Query().Get("url")
	message := r.URL.Query().Get("message")

	post, err := ph.svc.AddPost(userID, title, url, message)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	if err := json.NewEncoder(w).Encode(post); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (ph *postHandler) DeletePostByID(w http.ResponseWriter, r *http.Request) {
	postIDString := chi.URLParam(r, "post_id")
	postID, err := strconv.Atoi(postIDString)

	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	if err := ph.svc.DeletePostByID(postID); err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
