// Package handler is ui layer http-handler package
package handler

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/oshou/AwesomeMusic-api/api/usecase"
	"github.com/oshou/AwesomeMusic-api/log"
)

// ISearchHandler is ui layer http-handler interface
type ISearchHandler interface {
	SearchByType(w http.ResponseWriter, r *http.Request)
}

type searchHandler struct {
	usecase usecase.ISearchUsecase
}

var _ ISearchHandler = &searchHandler{}

// NewSearchHandler is constructor for searchHandler
func NewSearchHandler(usecase usecase.ISearchUsecase) ISearchHandler {
	return &searchHandler{
		usecase: usecase,
	}
}

func (sh *searchHandler) SearchByType(w http.ResponseWriter, r *http.Request) {
	searchType := r.URL.Query().Get("type")
	q := r.URL.Query().Get("q")

	switch searchType {
	case "post_title":
		posts, err := sh.usecase.ListPostsByTitle(q)
		if err != nil {
			log.Logger.Error("failed to get posts by title", zap.Error(errors.WithStack(err)))
			httpError(w, r, err)

			return
		}

		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(posts); err != nil {
			internalServerError(w, r, err)

			return
		}
	case "user_name":
		posts, err := sh.usecase.ListPostsByUserName(q)
		if err != nil {
			log.Logger.Error("failed to get posts by user name", zap.Error(errors.WithStack(err)))
			httpError(w, r, err)

			return
		}

		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(posts); err != nil {
			internalServerError(w, r, err)

			return
		}
	case "tag_name":
		posts, err := sh.usecase.ListPostsByTagName(q)
		if err != nil {
			log.Logger.Error("failed to get posts by tag name", zap.Error(errors.WithStack(err)))
			httpError(w, r, err)

			return
		}

		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(posts); err != nil {
			internalServerError(w, r, err)

			return
		}
	}
}
