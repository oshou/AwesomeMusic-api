package handler

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oshou/AwesomeMusic-api/usecase"
)

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

func NewTagHandler(u usecase.ITagUsecase) ITagHandler {
	return &tagHandler{
		usecase: u,
	}
}

func (th *tagHandler) GetTags(ctx *gin.Context) {
	tags, err := th.usecase.GetTags()
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}
	ctx.JSON(OK, tags)
}

func (th *tagHandler) GetTagByID(ctx *gin.Context) {
	tagID, err := strconv.Atoi(ctx.Param("tag_id"))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}

	tag, err := th.usecase.GetTagByID(tagID)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}

	ctx.JSON(OK, tag)
}

func (th *tagHandler) GetTagsByPostID(ctx *gin.Context) {
	postID, err := strconv.Atoi(ctx.Param("post_id"))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}

	tags, err := th.usecase.GetTagsByPostID(postID)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}

	ctx.JSON(OK, tags)
}

func (th *tagHandler) AddTag(ctx *gin.Context) {
	tagName := ctx.Query("name")
	tag, err := th.usecase.AddTag(tagName)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}

	ctx.JSON(Created, tag)
}

func (th *tagHandler) AttachTag(ctx *gin.Context) {
	postID, err := strconv.Atoi(ctx.Param("post_id"))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}

	tagID, err := strconv.Atoi(ctx.Param("tag_id"))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}

	postTag, err := th.usecase.AttachTag(postID, tagID)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)
	}

	ctx.JSON(Created, postTag)
}
