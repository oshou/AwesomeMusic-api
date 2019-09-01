package controller

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oshou/AwesomeMusic-api/service"
)

type Controller struct{}

// Index action: GET /v1/users
func (c Controller) GetUsers(ctx *gin.Context) {

	var us service.UserService
	val, err := us.GetAll()

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(404)
	} else {
		ctx.JSON(200, val)
	}
}

// Index action: GET /v1/users/:user_id
func (c Controller) GetUserById(ctx *gin.Context) {

	user_id, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(400)
	}

	var us service.UserService
	val, err := us.GetById(user_id)

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(404)
	} else {
		ctx.JSON(200, val)
	}
}

// Index action: POST /v1/users
func (c Controller) AddUser(ctx *gin.Context) {

	name := ctx.Query("name")

	var us service.UserService
	val, err := us.Add(name)

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(400)
		return
	} else {
		ctx.JSON(201, val)
	}
}

// Index action: GET /v1/posts
func (c Controller) GetPosts(ctx *gin.Context) {

	var ps service.PostService
	val, err := ps.GetAll()

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(404)
	} else {
		ctx.JSON(200, val)
	}
}

// Index action: GET /v1/posts/:post_id
func (c Controller) GetPostById(ctx *gin.Context) {

	post_id, err := strconv.Atoi(ctx.Param("post_id"))
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(400)
	}

	var ps service.PostService
	val, err := ps.GetById(post_id)

	if err != nil {
		ctx.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		ctx.JSON(200, val)
	}
}

// Index action: POST /v1/posts
func (c Controller) AddPost(ctx *gin.Context) {

	user_id, err := strconv.Atoi(ctx.Query("user_id"))
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(400)
		return
	}
	url := ctx.Query("url")
	message := ctx.Query("message")

	var ps service.PostService
	val, err := ps.Add(user_id, url, message)

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(400)
	} else {
		ctx.JSON(201, val)
	}
}

// Index action: DELETE /v1/posts/:post_id
func (c Controller) DeletePostById(ctx *gin.Context) {

	id := ctx.Param("post_id")
	post_id, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(400)
		return
	}

	var ps service.PostService
	if err := ps.DeleteById(post_id); err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(400)
	} else {
		ctx.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}

// Index action: GET /v1/posts/:post_id/comments
func (c Controller) GetComments(ctx *gin.Context) {

	post_id, err := strconv.Atoi(ctx.Param("post_id"))
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(400)
	}

	var cs service.CommentService
	val, err := cs.GetAll(post_id)

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(404)
	} else {
		ctx.JSON(200, val)
	}
}

// Index action: POST /v1/posts/:post_id/comments
func (c Controller) AddComment(ctx *gin.Context) {

	post_id, err := strconv.Atoi(ctx.Param("post_id"))
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(400)
		return
	}
	user_id, err := strconv.Atoi(ctx.Query("user_id"))
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(400)
		return
	}
	comment := ctx.Query("comment")

	var cs service.CommentService
	val, err := cs.Add(post_id, user_id, comment)

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(400)
	} else {
		ctx.JSON(201, val)
	}
}

// Index action: GET /v1/posts/:post_id/comments/:comment_id
func (c Controller) GetCommentById(ctx *gin.Context) {

	comment_id, err := strconv.Atoi(ctx.Param("comment_id"))
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(400)
	}

	var cs service.CommentService
	val, err := cs.GetById(comment_id)

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(404)
	} else {
		ctx.JSON(200, val)
	}
}

//
// // Index action: DELETE /v1/posts/:post_id/comments/:comment_id
// func (c Controller) DeleteComment(ctx *gin.Context) {
//
// 	post_id := ctx.Param("post_id")
// 	comment_id := ctx.Param("comment_id")
//
// 	var s post.Service
// 	val, err := s.DeleteComment(comment_id)
//
// 	if err != nil {
// 		fmt.Println(err)
// 		ctx.AbortWithStatus(404)
// 	} else {
// 		ctx.JSON(200, val)
// 	}
// }

// Index action: GET /v1/tags
func (c Controller) GetTags(ctx *gin.Context) {

	var ts service.TagService
	val, err := ts.GetAll()

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(404)
	} else {
		ctx.JSON(200, val)
	}
}

// Index action: POST /v1/tags
func (c Controller) AddTag(ctx *gin.Context) {

	name := ctx.Query("name")

	var ts service.TagService
	val, err := ts.Add(name)

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(400)
	} else {
		ctx.JSON(201, val)
	}
}

// Index action: GET /v1/tags/:tag_id
func (c Controller) GetTagById(ctx *gin.Context) {

	tag_id, err := strconv.Atoi(ctx.Param("tag_id"))
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(400)
	}

	var ts service.TagService
	val, err := ts.GetById(tag_id)

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(404)
	} else {
		ctx.JSON(200, val)
	}
}

// Index action: GET /v1/search
func (c Controller) SearchByType(ctx *gin.Context) {

	search_type := ctx.Query("type")
	q := ctx.Query("q")

	var ss service.SearchService

	switch search_type {
	case "user_id":
		val, err := ss.GetByUserId(q)
		if err != nil {
			fmt.Println(err)
			ctx.AbortWithStatus(404)
		} else {
			ctx.JSON(200, val)
		}
	case "user_name":
		val, err := ss.GetByUserName(q)
		if err != nil {
			fmt.Println(err)
			ctx.AbortWithStatus(404)
		} else {
			ctx.JSON(200, val)
		}
	case "tag_id":
		val, err := ss.GetByTagId(q)
		if err != nil {
			fmt.Println(err)
			ctx.AbortWithStatus(404)
		} else {
			ctx.JSON(200, val)
		}
	case "tag_name":
		val, err := ss.GetByTagName(q)
		if err != nil {
			fmt.Println(err)
			ctx.AbortWithStatus(404)
		} else {
			ctx.JSON(200, val)
		}
	default:
		err := errors.New("undefined")
		fmt.Println(err)
		ctx.AbortWithStatus(404)
	}
}
