package handler

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

type UserPresentor interface{
}

// Index: GET /v1/users
func (c Controller) GetUsers(ctx *gin.Context) {
	var us service.UserService
	users, err := us.GetAll()

	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(NotFound)

		return
	}

	ctx.JSON(OK, users)
}

// Create: POST /v1/users
func (c Controller) AddUser(ctx *gin.Context) {
	name := ctx.Query("name")

	var us service.UserService
	user, err := us.Add(name)

	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}

	ctx.JSON(CREATED, user)
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
	user, err := us.GetByID(userID)

	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(NotFound)

		return
	}

	ctx.JSON(OK, user)
}
}
