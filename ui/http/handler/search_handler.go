// Package handler is ui layer http-handler package
package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oshou/AwesomeMusic-api/service"
)

// ISearchHandler is ui layer http-handler interface
type ISearchHandler interface {
	SearchByType(ctx *gin.Context)
}

type searchHandler struct {
	service service.ISearchService
}

var _ ISearchHandler = (*searchHandler)(nil)

// NewSearchHandler is constructor for searchHandler
func NewSearchHandler(s service.ISearchService) ISearchHandler {
	return &searchHandler{
		service: s,
	}
}

func (sh *searchHandler) SearchByType(ctx *gin.Context) {
	searchType := ctx.Query("type")
	q := ctx.Query("q")

	switch searchType {
	case "post_title":
		posts, err := sh.service.GetPostsByTitle(q)
		if err != nil {
			fmt.Printf("%+v\n", err)
			ctx.AbortWithStatus(http.StatusBadRequest)

			return
		}

		ctx.JSON(http.StatusOK, posts)
	case "user_name":
		posts, err := sh.service.GetPostsByUserName(q)
		if err != nil {
			fmt.Printf("%+v\n", err)
			ctx.AbortWithStatus(http.StatusBadRequest)

			return
		}

		ctx.JSON(http.StatusOK, posts)
	case "tag_name":
		posts, err := sh.service.GetPostsByTagName(q)
		if err != nil {
			fmt.Printf("%+v\n", err)
			ctx.AbortWithStatus(http.StatusBadRequest)

			return
		}

		ctx.JSON(http.StatusOK, posts)
	}
}
