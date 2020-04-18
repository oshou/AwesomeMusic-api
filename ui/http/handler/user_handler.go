// Package handler is ui layer http-handler package
package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oshou/AwesomeMusic-api/service"
)

// IUserHandler is ui layer http-handler interface
type IUserHandler interface {
	GetUsers(ctx *gin.Context)
	GetUserByID(ctx *gin.Context)
	AddUser(ctx *gin.Context)
}

type userHandler struct {
	service service.IUserService
}

var _ IUserHandler = (*userHandler)(nil)

// NewUserHandler is constructor for userHandler
func NewUserHandler(s service.IUserService) IUserHandler {
	return &userHandler{
		service: s,
	}
}

func (uh *userHandler) GetUsers(ctx *gin.Context) {
	users, err := uh.service.GetUsers()
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusNotFound)

		return
	}

	ctx.JSON(http.StatusOK, users)
}

// Create: POST /v1/users
func (uh *userHandler) AddUser(ctx *gin.Context) {
	name := ctx.Query("name")

	user, err := uh.service.AddUser(name)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)

		return
	}

	ctx.JSON(http.StatusCreated, user)
}

// Detail: GET /v1/users/:user_id
func (uh *userHandler) GetUserByID(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)

		return
	}

	user, err := uh.service.GetUserByID(userID)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusNotFound)

		return
	}

	ctx.JSON(http.StatusOK, user)
}
