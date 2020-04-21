// Package handler is ui layer http-handler package
package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oshou/AwesomeMusic-api/service"
)

// ICommentHandler is ui layer http-handler interface
type ICommentHandler interface {
	GetComments(ctx *gin.Context)
	GetCommentByID(ctx *gin.Context)
	AddComment(ctx *gin.Context)
}

type commentHandler struct {
	svc service.ICommentService
}

var _ ICommentHandler = (*commentHandler)(nil)

// NewCommentHandler is constructor for commentHandler
func NewCommentHandler(svc service.ICommentService) ICommentHandler {
	return &commentHandler{
		svc: svc,
	}
}

func (ch *commentHandler) GetComments(ctx *gin.Context) {
	postID, err := strconv.Atoi(ctx.Param("post_id"))
	if err != nil {
		fmt.Printf("%+v\n", err)
		ctx.AbortWithStatus(http.StatusBadRequest)

		return
	}

	comments, err := ch.svc.GetComments(postID)
	if err != nil {
		fmt.Printf("%+v\n", err)
		ctx.AbortWithStatus(http.StatusNotFound)

		return
	}

	ctx.JSON(http.StatusOK, comments)
}

func (ch *commentHandler) AddComment(ctx *gin.Context) {
	postID, err := strconv.Atoi(ctx.Param("post_id"))
	if err != nil {
		fmt.Printf("%+v\n", err)
		ctx.AbortWithStatus(http.StatusBadRequest)

		return
	}

	userID, err := strconv.Atoi(ctx.Query("user_id"))
	if err != nil {
		fmt.Printf("%+v\n", err)
		ctx.AbortWithStatus(http.StatusBadRequest)

		return
	}

	commentText := ctx.Query("comment")
	comment, err := ch.svc.AddComment(postID, userID, commentText)

	if err != nil {
		fmt.Printf("%+v\n", err)
		ctx.AbortWithStatus(http.StatusBadRequest)

		return
	}

	ctx.JSON(http.StatusCreated, comment)
}

func (ch *commentHandler) GetCommentByID(ctx *gin.Context) {
	commentID, err := strconv.Atoi(ctx.Param("comment_id"))
	if err != nil {
		fmt.Printf("%+v\n", err)
		ctx.AbortWithStatus(http.StatusBadRequest)

		return
	}

	comment, err := ch.svc.GetCommentByID(commentID)
	if err != nil {
		fmt.Printf("%+v\n", err)
		ctx.AbortWithStatus(http.StatusNotFound)

		return
	}

	ctx.JSON(http.StatusOK, comment)
}