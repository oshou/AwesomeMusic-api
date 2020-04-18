// Package handler is ui layer http-handler package
package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oshou/AwesomeMusic-api/usecase"
)

// ICommentHandler is ui layer http-handler interface
type ICommentHandler interface {
	GetComments(ctx *gin.Context)
	GetCommentByID(ctx *gin.Context)
	AddComment(ctx *gin.Context)
}

type commentHandler struct {
	usecase usecase.ICommentUsecase
}

var _ ICommentHandler = (*commentHandler)(nil)

// NewCommentHandler is constructor for commentHandler
func NewCommentHandler(u usecase.ICommentUsecase) ICommentHandler {
	return &commentHandler{
		usecase: u,
	}
}

func (ch *commentHandler) GetComments(ctx *gin.Context) {
	postID, err := strconv.Atoi(ctx.Param("post_id"))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)

		return
	}

	comments, err := ch.usecase.GetComments(postID)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusNotFound)

		return
	}

	ctx.JSON(http.StatusOK, comments)
}

func (ch *commentHandler) AddComment(ctx *gin.Context) {
	postID, err := strconv.Atoi(ctx.Param("post_id"))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)

		return
	}

	userID, err := strconv.Atoi(ctx.Query("user_id"))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)

		return
	}

	commentText := ctx.Query("comment")
	comment, err := ch.usecase.AddComment(postID, userID, commentText)

	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)

		return
	}

	ctx.JSON(http.StatusCreated, comment)
}

func (ch *commentHandler) GetCommentByID(ctx *gin.Context) {
	commentID, err := strconv.Atoi(ctx.Param("comment_id"))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)

		return
	}

	comment, err := ch.usecase.GetCommentByID(commentID)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusNotFound)

		return
	}

	ctx.JSON(http.StatusOK, comment)
}
