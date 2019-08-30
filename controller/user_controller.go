package controller

import(
	"fmt"
	"github.com/oshou/AwesomeMusic-api/user"
	"github.com/gin-gonic/gin"
)

type Controller struct{}

// Index action: GET /v1/users
func (c Controller) GetUsers(ctx *gin.Context) {

	var s user.Service
	val, err := s.GetUsers()

	if err != nil {
		ctx.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		ctx.JSON(200, val)
	}
}

// Index action: POST /v1/users
func (c Controller) AddUser(ctx *gin.Context) {

	name := ctx.Query("name")

	var s service.Service
	val, err := s.AddUser(name)

	if err != nil {
		ctx.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		ctx.JSON(200, val)
	}
}

// // Index action: GET /v1/posts
// func (c Controller) GetPosts(ctx *gin.Context) {
// 
// 	var s post.Service
// 	val, err := s.GetPosts()
// 
// 	if err != nil {
// 		ctx.AbortWithStatus(404)
// 		fmt.Println(err)
// 	} else {
// 		ctx.JSON(200, val)
// 	}
// }
// 
// // Index action: POST /v1/posts
// func (c Controller) AddPost(ctx *gin.Context) {
// 	var s post.Service
// 	val, err := s.AddPost()
// 
// 	if err != nil {
// 		ctx.AbortWithStatus(404)
// 		fmt.Println(err)
// 	} else {
// 		ctx.JSON(200, val)
// 	}
// }
// 
// // Index action: GET /v1/posts/:post_id
// func (c Controller) GetPost(ctx *gin.Context) {
// 
// 	post_id := ctx.Param("post_id")
// 
// 	var s post.Service
// 	val, err := s.GetPost(post_id)
// 
// 	if err != nil {
// 		ctx.AbortWithStatus(404)
// 		fmt.Println(err)
// 	} else {
// 		ctx.JSON(200, val)
// 	}
// }
// 
// // Index action: DELETE /v1/posts/:post_id
// func (c Controller) DeletePost(ctx *gin.Context) {
// 
// 	post_id := ctx.Param("post_id")
// 
// 	var s post.Service
// 	val, err := s.DeletePost(post_id)
// 
// 	if err != nil {
// 		ctx.AbortWithStatus(404)
// 		fmt.Println(err)
// 	} else {
// 		ctx.JSON(200, val)
// 	}
// }
// 
// // Index action: GET /v1/posts/:post_id/comments
// func (c Controller) GetComments(ctx *gin.Context) {
// 
// 	post_id := ctx.Param("post_id")
// 
// 	var s post.Service
// 	val, err := s.GetComments(post_id)
// 
// 	if err != nil {
// 		ctx.AbortWithStatus(404)
// 		fmt.Println(err)
// 	} else {
// 		ctx.JSON(200, val)
// 	}
// }
// 
// 
// // Index action: POST /v1/posts/:post_id/comments
// func (c Controller) AddComment(ctx *gin.Context) {
// 
// 	post_id := ctx.Param("post_id")
// 	user_id := ctx.Query("user_id")
// 
// 	var s post.Service
// 	val, err := s.AddComment(post_id,comment_id)
// 
// 	if err != nil {
// 		ctx.AbortWithStatus(404)
// 		fmt.Println(err)
// 	} else {
// 		ctx.JSON(200, val)
// 	}
// }
// 
// // Index action: GET /v1/posts/:post_id/comments/:comment_id
// func (c Controller) GetComment(ctx *gin.Context) {
// 
// 	post_id := ctx.Param("post_id")
// 	comment_id := ctx.Param("comment_id")
// 
// 	var s post.Service
// 	val, err := s.GetComment(post_id,comment_id)
// 
// 	if err != nil {
// 		ctx.AbortWithStatus(404)
// 		fmt.Println(err)
// 	} else {
// 		ctx.JSON(200, val)
// 	}
// }
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
// 		ctx.AbortWithStatus(404)
// 		fmt.Println(err)
// 	} else {
// 		ctx.JSON(200, val)
// 	}
// }