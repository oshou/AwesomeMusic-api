package handler

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oshou/AwesomeMusic-api/usecase"
)

type IUserHandler interface {
	GetUsers(ctx *gin.Context)
	GetUserByID(ctx *gin.Context)
	AddUser(ctx *gin.Context)
}

type userHandler struct {
	usecase usecase.IUserUsecase
}

var _ IUserHandler = (*userHandler)(nil)

func NewUserHandler(u usecase.IUserUsecase) IUserHandler {
	return &userHandler{
		usecase: u,
	}
}

func (uh *userHandler) GetUsers(ctx *gin.Context) {
	users, err := uh.usecase.GetUsers()
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(NotFound)

		return
	}

	ctx.JSON(OK, users)
}

// Create: POST /v1/users
func (uh *userHandler) AddUser(ctx *gin.Context) {
	name := ctx.Query("name")

	user, err := uh.usecase.AddUser(name)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}

	ctx.JSON(Created, user)
}

// Detail: GET /v1/users/:user_id
func (uh *userHandler) GetUserByID(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(BadRequest)

		return
	}

	user, err := uh.usecase.GetUserByID(userID)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(NotFound)

		return
	}

	ctx.JSON(OK, user)
}
