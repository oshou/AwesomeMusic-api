package handler

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oshou/AwesomeMusic-api/usecase"
)

type IPostHandler interface {
	GetPosts(ctx *gin.Context)
	GetPostByID(ctx *gin.Context)
	GetPostsByTagID(ctx *gin.Context)
	GetPostsByUserID(ctx *gin.Context)
	AddPost(ctx *gin.Context)
	DeletePostByID(ctx *gin.Context)
}

type postHandler struct {
	usecase usecase.IPostUsecase
}

var _ IPostHandler = (*postHandler)(nil)

func NewPostHandler(u usecase.IPostUsecase) IPostHandler {
	return &postHandler{
		usecase: u,
	}
}

func (ph *postHandler) GetPosts(ctx *gin.Context) {
	posts, err := ph.usecase.GetPosts()
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(NotFound)

		return
	}

	ctx.JSON(OK, posts)
}

func (ph *postHandler) GetPostByID(ctx *gin.Context) {
	postID, err := strconv.Atoi(ctx.Param("post_id"))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}

	post, err := ph.usecase.GetPostByID(postID)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(NotFound)
	}

	ctx.JSON(OK, post)
}

func (ph *postHandler) GetPostsByTagID(ctx *gin.Context) {
	tagID, err := strconv.Atoi(ctx.Param("tag_id"))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}

	posts, err := ph.usecase.GetPostsByTagID(tagID)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}

	ctx.JSON(OK, posts)
}

func (ph *postHandler) GetPostsByUserID(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}

	posts, err := ph.usecase.GetPostsByUserID(userID)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(NotFound)

		return
	}

	ctx.JSON(OK, posts)
}

func (ph *postHandler) AddPost(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Query("user_id"))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}

	title := ctx.Query("title")
	url := ctx.Query("url")
	message := ctx.Query("message")

	post, err := ph.usecase.AddPost(userID, title, url, message)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}

	ctx.JSON(OK, post)
}

func (ph *postHandler) DeletePostByID(ctx *gin.Context) {
	id := ctx.Param("post_id")
	postID, err := strconv.Atoi(id)

	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}

	if err := ph.usecase.DeletePostByID(postID); err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}

	ctx.JSON(NoContent, gin.H{"id #" + id: "deleted"})
}
