package handler

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oshou/AwesomeMusic-api/usecase"
)

const (
	OK         = 200
	Created    = 201
	NoContent  = 204
	BadRequest = 400
	NotFound   = 404
)

type ICommentHandler interface {
	GetComments(ctx *gin.Context)
	GetCommentByID(ctx *gin.Context)
	AddComment(ctx *gin.Context)
}

type commentHandler struct {
	usecase usecase.ICommentUsecase
}

var _ ICommentHandler = (*commentHandler)(nil)

func NewCommentHandler(u usecase.ICommentUsecase) ICommentHandler {
	return &commentHandler{
		usecase: u,
	}
}

// Index: GET /v1/posts/:post_id/comments
func (ch *commentHandler) GetComments(ctx *gin.Context) {
	postID, err := strconv.Atoi(ctx.Param("post_id"))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}

	comments, err := ch.usecase.GetComments(postID)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(NotFound)

		return
	}

	ctx.JSON(OK, comments)
}

func (ch *commentHandler) AddComment(ctx *gin.Context) {
	postID, err := strconv.Atoi(ctx.Param("post_id"))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}

	userID, err := strconv.Atoi(ctx.Query("user_id"))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}

	commentText := ctx.Query("comment")
	comment, err := ch.usecase.AddComment(postID, userID, commentText)

	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}

	ctx.JSON(Created, comment)
}

func (ch *commentHandler) GetCommentByID(ctx *gin.Context) {
	commentID, err := strconv.Atoi(ctx.Param("comment_id"))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}

	comment, err := ch.usecase.GetCommentByID(commentID)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(NotFound)

		return
	}

	ctx.JSON(OK, comment)
}
