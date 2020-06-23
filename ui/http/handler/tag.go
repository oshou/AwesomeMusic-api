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

// ITagHandler is ui layer http-handler interface
type ITagHandler interface {
	GetTags(w http.ResponseWriter, r *http.Request)
	GetTagByID(w http.ResponseWriter, r *http.Request)
	GetTagsByPostID(w http.ResponseWriter, r *http.Request)
	AddTag(w http.ResponseWriter, r *http.Request)
	AttachTag(w http.ResponseWriter, r *http.Request)
}

type tagHandler struct {
	usecase usecase.ITagUsecase
}

var _ ITagHandler = &tagHandler{}

// NewTagHandler is constructor for tagHandler
func NewTagHandler(usecase usecase.ITagUsecase) ITagHandler {
	return &tagHandler{
		usecase: usecase,
	}
}

func (th *tagHandler) GetTags(w http.ResponseWriter, r *http.Request) {
	tags, err := th.usecase.GetTags()
	if err != nil {
		log.Logger.Error("failed to get tags", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(tags); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (th *tagHandler) GetTagByID(w http.ResponseWriter, r *http.Request) {
	tagIDString := chi.URLParam(r, "tag_id")
	tagID, err := strconv.Atoi(tagIDString)

	if err != nil {
		log.Logger.Error("failed to convert string", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	tag, err := th.usecase.GetTagByID(tagID)
	if err != nil {
		log.Logger.Error("failed to get tag by tagID", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(tag); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (th *tagHandler) GetTagsByPostID(w http.ResponseWriter, r *http.Request) {
	postIDString := chi.URLParam(r, "post_id")
	postID, err := strconv.Atoi(postIDString)

	if err != nil {
		log.Logger.Error("failed to convert string", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	tags, err := th.usecase.GetTagsByPostID(postID)
	if err != nil {
		log.Logger.Error("failed to get tags by postID", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(tags); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (th *tagHandler) AddTag(w http.ResponseWriter, r *http.Request) {
	tagName := r.URL.Query().Get("name")
	tag, err := th.usecase.AddTag(tagName)

	if err != nil {
		log.Logger.Error("failed to add tag", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(tag); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (th *tagHandler) AttachTag(w http.ResponseWriter, r *http.Request) {
	postIDString := chi.URLParam(r, "post_id")
	postID, err := strconv.Atoi(postIDString)

	if err != nil {
		log.Logger.Error("failed to convert string", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	tagIDString := chi.URLParam(r, "tag_id")
	tagID, err := strconv.Atoi(tagIDString)

	if err != nil {
		log.Logger.Error("failed to convert string", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	postTag, err := th.usecase.AttachTag(postID, tagID)
	if err != nil {
		log.Logger.Error("failed to attach tag", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(postTag); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}
