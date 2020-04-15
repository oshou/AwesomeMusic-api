package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/oshou/AwesomeMusic-api/usecase"
)

type ISearchHandler interface {
	SearchByType(ctx *gin.Context)
}

type searchHandler struct {
	usecase usecase.ISearchUsecase
}

var _ ISearchHandler = (*searchHandler)(nil)

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
			ctx.AbortWithStatus(BadRequest)

			return
		}

		ctx.JSON(OK, posts)
	case "user_name":
		posts, err := sh.usecase.GetPostsByUserName(q)
		if err != nil {
			log.Println(err)
			ctx.AbortWithStatus(BadRequest)

			return
		}

		ctx.JSON(OK, posts)
	case "tag_name":
		posts, err := sh.usecase.GetPostsByTagName(q)
		if err != nil {
			log.Println(err)
			ctx.AbortWithStatus(BadRequest)

			return
		}

		ctx.JSON(OK, posts)
	}
}
