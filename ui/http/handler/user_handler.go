// Package handler is ui layer http-handler package
package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oshou/AwesomeMusic-api/usecase"
)

// IUserHandler is ui layer http-handler interface
type IUserHandler interface {
	GetUsers(ctx *gin.Context)
	GetUserByID(ctx *gin.Context)
	AddUser(ctx *gin.Context)
}

type userHandler struct {
	usecase usecase.IUserUsecase
}

var _ IUserHandler = (*userHandler)(nil)

// NewUserHandler is constructor for userHandler
func NewUserHandler(u usecase.IUserUsecase) IUserHandler {
	return &userHandler{
		usecase: u,
	}
}

func (uh *userHandler) GetUsers(ctx *gin.Context) {
	users, err := uh.usecase.GetUsers()
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

	user, err := uh.usecase.AddUser(name)
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

	user, err := uh.usecase.GetUserByID(userID)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusNotFound)

		return
	}

	ctx.JSON(http.StatusOK, user)
}
