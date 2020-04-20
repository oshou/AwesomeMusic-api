// Package handler is ui layer http-handler package
package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oshou/AwesomeMusic-api/service"
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
	svc service.ITagService
}

var _ ITagHandler = (*tagHandler)(nil)

// NewTagHandler is constructor for tagHandler
func NewTagHandler(svc service.ITagService) ITagHandler {
	return &tagHandler{
		svc: svc,
	}
}

func (th *tagHandler) GetTags(ctx *gin.Context) {
	tags, err := th.svc.GetTags()
	if err != nil {
		fmt.Printf("%+v\n", err)
		ctx.AbortWithStatus(http.StatusBadRequest)

		return
	}

	ctx.JSON(http.StatusOK, tags)
}

func (th *tagHandler) GetTagByID(ctx *gin.Context) {
	tagID, err := strconv.Atoi(ctx.Param("tag_id"))
	if err != nil {
		fmt.Printf("%+v\n", err)
		ctx.AbortWithStatus(http.StatusBadRequest)

		return
	}

	tag, err := th.svc.GetTagByID(tagID)
	if err != nil {
		fmt.Printf("%+v\n", err)
		ctx.AbortWithStatus(http.StatusBadRequest)

		return
	}

	ctx.JSON(http.StatusOK, tag)
}

func (th *tagHandler) GetTagsByPostID(ctx *gin.Context) {
	postID, err := strconv.Atoi(ctx.Param("post_id"))
	if err != nil {
		fmt.Printf("%+v\n", err)
		ctx.AbortWithStatus(http.StatusBadRequest)

		return
	}

	tags, err := th.svc.GetTagsByPostID(postID)
	if err != nil {
		fmt.Printf("%+v\n", err)
		ctx.AbortWithStatus(http.StatusBadRequest)

		return
	}

	ctx.JSON(http.StatusOK, tags)
}

func (th *tagHandler) AddTag(ctx *gin.Context) {
	tagName := ctx.Query("name")
	tag, err := th.svc.AddTag(tagName)

	if err != nil {
		fmt.Printf("%+v\n", err)
		ctx.AbortWithStatus(http.StatusBadRequest)

		return
	}

	ctx.JSON(http.StatusCreated, tag)
}

func (th *tagHandler) AttachTag(ctx *gin.Context) {
	postID, err := strconv.Atoi(ctx.Param("post_id"))
	if err != nil {
		fmt.Printf("%+v\n", err)
		ctx.AbortWithStatus(http.StatusBadRequest)

		return
	}

	tagID, err := strconv.Atoi(ctx.Param("tag_id"))
	if err != nil {
		fmt.Printf("%+v\n", err)
		ctx.AbortWithStatus(http.StatusBadRequest)

		return
	}

	postTag, err := th.svc.AttachTag(postID, tagID)
	if err != nil {
		fmt.Printf("%+v\n", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
	}

	ctx.JSON(http.StatusCreated, postTag)
}
