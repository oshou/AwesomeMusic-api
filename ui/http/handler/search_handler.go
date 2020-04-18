// Package handler is ui layer http-handler package
package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oshou/AwesomeMusic-api/usecase"
)

// ISearchHandler is ui layer http-handler interface
type ISearchHandler interface {
	SearchByType(ctx *gin.Context)
}

type searchHandler struct {
	usecase usecase.ISearchUsecase
}

var _ ISearchHandler = (*searchHandler)(nil)

// NewSearchHandler is constructor for searchHandler
func NewSearchHandler(u usecase.ISearchUsecase) ISearchHandler {
	return &searchHandler{
		usecase: u,
	}
}

func (sh *searchHandler) SearchByType(ctx *gin.Context) {
	searchType := ctx.Query("type")
	q := ctx.Query("q")

	switch searchType {
	case "post_title":
		posts, err := sh.usecase.GetPostsByTitle(q)
		if err != nil {
			log.Println(err)
			ctx.AbortWithStatus(http.StatusBadRequest)

			return
		}

		ctx.JSON(http.StatusOK, posts)
	case "user_name":
		posts, err := sh.usecase.GetPostsByUserName(q)
		if err != nil {
			log.Println(err)
			ctx.AbortWithStatus(http.StatusBadRequest)

			return
		}

		ctx.JSON(http.StatusOK, posts)
	case "tag_name":
		posts, err := sh.usecase.GetPostsByTagName(q)
		if err != nil {
			log.Println(err)
			ctx.AbortWithStatus(http.StatusBadRequest)

			return
		}

		ctx.JSON(http.StatusOK, posts)
	}
}
