package task

import "fmt"

type controller struct{}

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

	var s user.Service
	val, err := s.AddUser(name)

	if err != nil {
		ctx.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		ctx.JSON(200, val)
	}
}

// Index action: GET /v1/posts/:post_id
func (c Controller) GetPost(ctx *gin.Context) {

	post_id := ctx.Param("post_id")

	var s post.Service
	val, err := s.GetPost(post_id)

	if err != nil {
		ctx.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		ctx.JSON(200, val)
	}
}

// Index action: POST /v1/posts/:post_id
func (c Controller) AddPosts(ctx *gin.Context) {
	var s post.Service
	val, err := s.AddPosts()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, val)
	}
}
