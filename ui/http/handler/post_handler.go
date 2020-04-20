// Package handler is ui layer http-handler package
package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oshou/AwesomeMusic-api/service"
)

// IPostHandler is ui layer http-handler interface
type IPostHandler interface {
	GetPosts(ctx *gin.Context)
	GetPostByID(ctx *gin.Context)
	GetPostsByTagID(ctx *gin.Context)
	GetPostsByUserID(ctx *gin.Context)
	AddPost(ctx *gin.Context)
	DeletePostByID(ctx *gin.Context)
}

type postHandler struct {
	svc service.IPostService
}

var _ IPostHandler = (*postHandler)(nil)

// NewPostHandler is constructor for postHandler
func NewPostHandler(svc service.IPostService) IPostHandler {
	return &postHandler{
		svc: svc,
	}
}

func (ph *postHandler) GetPosts(ctx *gin.Context) {
	posts, err := ph.svc.GetPosts()
	if err != nil {
		fmt.Printf("%+v\n", err)
		ctx.AbortWithStatus(http.StatusNotFound)

		return
	}

	ctx.JSON(http.StatusOK, posts)
}

func (ph *postHandler) GetPostByID(ctx *gin.Context) {
	postID, err := strconv.Atoi(ctx.Param("post_id"))
	if err != nil {
		fmt.Printf("%+v\n", err)
		ctx.AbortWithStatus(http.StatusBadRequest)

		return
	}

	post, err := ph.svc.GetPostByID(postID)
	if err != nil {
		fmt.Printf("%+v\n", err)
		ctx.AbortWithStatus(http.StatusNotFound)
	}

	ctx.JSON(http.StatusOK, post)
}

func (ph *postHandler) GetPostsByTagID(ctx *gin.Context) {
	tagID, err := strconv.Atoi(ctx.Param("tag_id"))
	if err != nil {
		fmt.Printf("%+v\n", err)
		ctx.AbortWithStatus(http.StatusBadRequest)

		return
	}

	posts, err := ph.svc.GetPostsByTagID(tagID)
	if err != nil {
		fmt.Printf("%+v\n", err)
		ctx.AbortWithStatus(http.StatusBadRequest)

		return
	}

	ctx.JSON(http.StatusOK, posts)
}

func (ph *postHandler) GetPostsByUserID(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		fmt.Printf("%+v\n", err)
		ctx.AbortWithStatus(http.StatusBadRequest)

		return
	}

	posts, err := ph.svc.GetPostsByUserID(userID)
	if err != nil {
		fmt.Printf("%+v\n", err)
		ctx.AbortWithStatus(http.StatusNotFound)

		return
	}

	ctx.JSON(http.StatusOK, posts)
}

func (ph *postHandler) AddPost(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Query("user_id"))
	if err != nil {
		fmt.Printf("%+v\n", err)
		ctx.AbortWithStatus(http.StatusBadRequest)

		return
	}

	title := ctx.Query("title")
	url := ctx.Query("url")
	message := ctx.Query("message")

	post, err := ph.svc.AddPost(userID, title, url, message)
	if err != nil {
		fmt.Printf("%+v\n", err)
		ctx.AbortWithStatus(http.StatusBadRequest)

		return
	}

	ctx.JSON(http.StatusCreated, post)
}

func (ph *postHandler) DeletePostByID(ctx *gin.Context) {
	id := ctx.Param("post_id")
	postID, err := strconv.Atoi(id)

	if err != nil {
		fmt.Printf("%+v\n", err)
		ctx.AbortWithStatus(http.StatusBadRequest)

		return
	}

	if err := ph.svc.DeletePostByID(postID); err != nil {
		fmt.Printf("%+v\n", err)
		ctx.AbortWithStatus(http.StatusBadRequest)

		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{"id #" + id: "deleted"})
}
