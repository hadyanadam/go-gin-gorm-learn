package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/hadyanadam/go-gorm-gin/dto"
	"github.com/hadyanadam/go-gorm-gin/helper"
	"github.com/hadyanadam/go-gorm-gin/service"
)


type UserController interface {
	Update(context *gin.Context)
	Profile(context *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService service.JWTService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService: jwtService,
	}
}

func (c *userController) Update(context *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTO
	errDTO := context.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authHeader, errAuth := helper.SplitHeader(context.GetHeader("Authorization"), "Bearer")
	if errAuth != nil {
		panic(errAuth.Error())
	}
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	userUpdateDTO.ID = id
	user := c.userService.Update(userUpdateDTO)
	res := helper.BuildResponse(true, "OK!", user)
	context.JSON(http.StatusOK, res)
}

func (c *userController) Profile(context *gin.Context) {
	authHeader, errAuth := helper.SplitHeader(context.GetHeader("Authorization"), "Bearer")
	if errAuth != nil {
		panic(errAuth.Error())
	}
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	user := c.userService.Profile(fmt.Sprintf("%v", claims["user_id"]))
	res := helper.BuildResponse(true, "OK", user)
	context.JSON(http.StatusOK, res)
}
