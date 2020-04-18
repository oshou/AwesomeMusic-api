// Package handler is ui layer http-handler package
package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oshou/AwesomeMusic-api/usecase"
)

// ITagHandler is ui layer http-handler interface
type ITagHandler interface {
	GetTags(ctx *gin.Context)
	GetTagByID(ctx *gin.Context)
	GetTagsByPostID(ctx *gin.Context)
	AddTag(ctx *gin.Context)
	AttachTag(ctx *gin.Context)
}

type tagHandler struct {
	usecase usecase.ITagUsecase
}

var _ ITagHandler = (*tagHandler)(nil)

// NewTagHandler is constructor for tagHandler
func NewTagHandler(u usecase.ITagUsecase) ITagHandler {
	return &tagHandler{
		usecase: u,
	}
}

func (th *tagHandler) GetTags(ctx *gin.Context) {
	tags, err := th.usecase.GetTags()
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)

		return
	}

	ctx.JSON(http.StatusOK, tags)
}

func (th *tagHandler) GetTagByID(ctx *gin.Context) {
	tagID, err := strconv.Atoi(ctx.Param("tag_id"))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)

		return
	}

	tag, err := th.usecase.GetTagByID(tagID)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)

		return
	}

	ctx.JSON(http.StatusOK, tag)
}

func (th *tagHandler) GetTagsByPostID(ctx *gin.Context) {
	postID, err := strconv.Atoi(ctx.Param("post_id"))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)

		return
	}

	tags, err := th.usecase.GetTagsByPostID(postID)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)

		return
	}

	ctx.JSON(http.StatusOK, tags)
}

func (th *tagHandler) AddTag(ctx *gin.Context) {
	tagName := ctx.Query("name")
	tag, err := th.usecase.AddTag(tagName)

	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)

		return
	}

	ctx.JSON(http.StatusCreated, tag)
}

func (th *tagHandler) AttachTag(ctx *gin.Context) {
	postID, err := strconv.Atoi(ctx.Param("post_id"))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)

		return
	}

	tagID, err := strconv.Atoi(ctx.Param("tag_id"))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)

		return
	}

	postTag, err := th.usecase.AttachTag(postID, tagID)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
	}

	ctx.JSON(http.StatusCreated, postTag)
}
