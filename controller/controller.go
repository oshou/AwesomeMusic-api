package controller

import (
	"errors"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oshou/AwesomeMusic-api/service"
)

const (
	OK         = 200
	CREATED    = 201
	BadRequest = 400
	NotFound   = 404
)

type Controller struct{}

// Index: GET /v1/users
func (c Controller) GetUsers(ctx *gin.Context) {
	var us service.UserService
	val, err := us.GetAll()

	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(NotFound)

		return
	}

	ctx.JSON(OK, val)
}

// Create: POST /v1/users
func (c Controller) AddUser(ctx *gin.Context) {
	name := ctx.Query("name")

	var us service.UserService
	val, err := us.Add(name)

	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}

	ctx.JSON(CREATED, val)
}

// Detail: GET /v1/users/:user_id
func (c Controller) GetUserByID(ctx *gin.Context) {
	// user_id
	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}

	var us service.UserService
	val, err := us.GetByID(userID)

	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(NotFound)

		return
	}

	ctx.JSON(OK, val)
}

// Index: GET /v1/posts
func (c Controller) GetPosts(ctx *gin.Context) {
	var ps service.PostService
	val, err := ps.GetAll()

	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(NotFound)

		return
	}

	ctx.JSON(OK, val)
}

// Index: POST /v1/posts
func (c Controller) AddPost(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Query("user_id"))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}

	url := ctx.Query("url")
	title := ctx.Query("title")
	message := ctx.Query("message")

	var ps service.PostService
	val, err := ps.Add(userID, title, url, message)

	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}

	ctx.JSON(CREATED, val)
}

// Detail: GET /v1/posts/:post_id
func (c Controller) GetPostByID(ctx *gin.Context) {
	postID, err := strconv.Atoi(ctx.Param("post_id"))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}

	var ps service.PostService
	val, err := ps.GetByID(postID)

	if err != nil {
		ctx.AbortWithStatus(NotFound)
		log.Println(err)

		return
	}

	ctx.JSON(OK, val)
}

// Detail: GET /v1/tags/:tag_id/posts
func (c Controller) GetPostsByTagID(ctx *gin.Context) {
	tagID, err := strconv.Atoi(ctx.Param("tag_id"))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}

	var ps service.PostService
	val, err := ps.GetByTagID(tagID)

	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(NotFound)

		return
	}

	ctx.JSON(OK, val)
}

// Detail: GET /v1/tags/:tag_id/users
func (c Controller) GetPostsByUserID(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}

	var ps service.PostService
	val, err := ps.GetByUserID(userID)

	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(NotFound)

		return
	}

	ctx.JSON(OK, val)
}

// Delete: DELETE /v1/posts/:post_id
//func (c Controller) DeletePostByID(ctx *gin.Context) {
//
//	id := ctx.Param("post_id")
//	postID, err := strconv.Atoi(id)
//	if err != nil {
//		log.Println(err)
//		ctx.AbortWithStatus(BadRequest)
//		return
//	}
//
//	var ps service.PostService
//	if err := ps.DeleteByID(postID); err != nil {
//		log.Println(err)
//		ctx.AbortWithStatus(BadRequest)
//    return
//	}
//
//	ctx.JSON(204, gin.H{"id #" + id: "deleted"})
//}

// Index: GET /v1/posts/:post_id/comments
func (c Controller) GetComments(ctx *gin.Context) {
	postID, err := strconv.Atoi(ctx.Param("post_id"))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}

	var cs service.CommentService
	val, err := cs.GetAll(postID)

	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(NotFound)

		return
	}

	ctx.JSON(OK, val)
}

// Create: POST /v1/posts/:post_id/comments
func (c Controller) AddComment(ctx *gin.Context) {
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

	comment := ctx.Query("comment")

	var cs service.CommentService
	val, err := cs.Add(postID, userID, comment)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}

	ctx.JSON(CREATED, val)
}

// Detail: GET /v1/posts/:post_id/comments/:comment_id
func (c Controller) GetCommentByID(ctx *gin.Context) {
	commentID, err := strconv.Atoi(ctx.Param("comment_id"))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}

	var cs service.CommentService
	val, err := cs.GetByID(commentID)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(NotFound)

		return
	}

	ctx.JSON(OK, val)
}

// // Index action: DELETE /v1/posts/:post_id/comments/:comment_id
// func (c Controller) DeleteComment(ctx *gin.Context) {
// 	postID := ctx.Param("post_id")
// 	commentID := ctx.Param("comment_id")
//
// 	var s post.Service
// 	val, err := s.DeleteComment(commentID)
// 	if err != nil {
// 		log.Println(err)
// 		ctx.AbortWithStatus(NotFound)
//
// 		return
// 	}
//
// 	ctx.JSON(OK, val)
// }

// Index: GET /v1/tags
func (c Controller) GetTags(ctx *gin.Context) {
	var ts service.TagService
	val, err := ts.GetAll()
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(NotFound)

		return
	}

	ctx.JSON(OK, val)
}

// Create: POST /v1/tags
func (c Controller) AddTag(ctx *gin.Context) {
	name := ctx.Query("name")

	var ts service.TagService
	val, err := ts.Add(name)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}

	ctx.JSON(CREATED, val)
}

// Detail: GET /v1/tags/:tag_id
func (c Controller) GetTagByID(ctx *gin.Context) {
	tagID, err := strconv.Atoi(ctx.Param("tag_id"))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}

	var ts service.TagService
	val, err := ts.GetByID(tagID)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(NotFound)

		return
	}

	ctx.JSON(OK, val)
}

// Index: GET /v1/posts/:post_id/tags
func (c Controller) GetTagsByPostID(ctx *gin.Context) {
	postID, err := strconv.Atoi(ctx.Param("post_id"))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}

	var ts service.TagService
	val, err := ts.GetByPostID(postID)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(NotFound)

		return
	}

	ctx.JSON(OK, val)
}

// Create: POST /v1/posts/:post_id/tags/:tag_id
func (c Controller) AttachTag(ctx *gin.Context) {
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

	var ts service.TagService
	val, err := ts.Attach(postID, tagID)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}

	ctx.JSON(CREATED, val)
}

// Index: GET /v1/search
func (c Controller) SearchByType(ctx *gin.Context) {
	searchType := ctx.Query("type")
	q := ctx.Query("q")

	var ss service.SearchService

	switch searchType {
	case "post_title":
		val, err := ss.GetByPostTitle(q)
		if err != nil {
			log.Println(err)
			ctx.AbortWithStatus(NotFound)

			return
		}

		ctx.JSON(OK, val)
	case "user_name":
		var us service.UserService
		val, err := us.GetByName(q)
		if err != nil {
			log.Println(err)
			ctx.AbortWithStatus(NotFound)

			return
		}

		ctx.JSON(OK, val)
	case "tag_name":
		var ts service.TagService
		val, err := ts.GetByName(q)
		if err != nil {
			log.Println(err)
			ctx.AbortWithStatus(NotFound)

			return
		}

		ctx.JSON(OK, val)
	default:
		err := errors.New("undefined")
		log.Println(err)
		ctx.AbortWithStatus(NotFound)
	}
}
