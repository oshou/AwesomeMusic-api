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

// ITagHandler is ui layer http-handler interface
type ITagHandler interface {
	GetTags(w http.ResponseWriter, r *http.Request)
	GetTagByID(w http.ResponseWriter, r *http.Request)
	GetTagsByPostID(w http.ResponseWriter, r *http.Request)
	AddTag(w http.ResponseWriter, r *http.Request)
	AttachTag(w http.ResponseWriter, r *http.Request)
}

type tagHandler struct {
	svc service.ITagService
}

var _ ITagHandler = &tagHandler{}

// NewTagHandler is constructor for tagHandler
func NewTagHandler(svc service.ITagService) ITagHandler {
	return &tagHandler{
		svc: svc,
	}
}

func (th *tagHandler) GetTags(w http.ResponseWriter, r *http.Request) {
	tags, err := th.svc.GetTags()
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	if err := json.NewEncoder(w).Encode(tags); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (th *tagHandler) GetTagByID(w http.ResponseWriter, r *http.Request) {
	tagIDString := chi.URLParam(r, "tag_id")
	tagID, err := strconv.Atoi(tagIDString)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	tag, err := th.svc.GetTagByID(tagID)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	if err := json.NewEncoder(w).Encode(tag); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (th *tagHandler) GetTagsByPostID(w http.ResponseWriter, r *http.Request) {
	postIDString := chi.URLParam(r, "post_id")
	postID, err := strconv.Atoi(postIDString)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	tags, err := th.svc.GetTagsByPostID(postID)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	if err := json.NewEncoder(w).Encode(tags); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (th *tagHandler) AddTag(w http.ResponseWriter, r *http.Request) {
	tagName := r.URL.Query().Get("name")
	tag, err := th.svc.AddTag(tagName)

	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	if err := json.NewEncoder(w).Encode(tag); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (th *tagHandler) AttachTag(w http.ResponseWriter, r *http.Request) {
	postIDString := chi.URLParam(r, "post_id")
	postID, err := strconv.Atoi(postIDString)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	tagIDString := chi.URLParam(r, "tag_id")
	tagID, err := strconv.Atoi(tagIDString)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	postTag, err := th.svc.AttachTag(postID, tagID)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	if err := json.NewEncoder(w).Encode(postTag); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}
