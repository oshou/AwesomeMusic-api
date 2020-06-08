// Package handler is ui layer http-handler package
package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/oshou/AwesomeMusic-api/service"
)

// ISearchHandler is ui layer http-handler interface
type ISearchHandler interface {
	SearchByType(w http.ResponseWriter, r *http.Request)
}

type searchHandler struct {
	svc service.ISearchService
}

var _ ISearchHandler = &searchHandler{}

// NewSearchHandler is constructor for searchHandler
func NewSearchHandler(svc service.ISearchService) ISearchHandler {
	return &searchHandler{
		svc: svc,
	}
}

func (sh *searchHandler) SearchByType(w http.ResponseWriter, r *http.Request) {
	searchType := r.URL.Query().Get("type")
	q := r.URL.Query().Get("q")

	switch searchType {
	case "post_title":
		posts, err := sh.svc.GetPostsByTitle(q)
		if err != nil {
			fmt.Printf("%+v\n", err)
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		if err := json.NewEncoder(w).Encode(posts); err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}
		w.WriteHeader(http.StatusOK)
	case "user_name":
		posts, err := sh.svc.GetPostsByUserName(q)
		if err != nil {
			fmt.Printf("%+v\n", err)
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		if err := json.NewEncoder(w).Encode(posts); err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		w.WriteHeader(http.StatusOK)
	case "tag_name":
		posts, err := sh.svc.GetPostsByTagName(q)
		if err != nil {
			fmt.Printf("%+v\n", err)
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		if err := json.NewEncoder(w).Encode(posts); err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
