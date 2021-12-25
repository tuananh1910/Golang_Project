package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"Assignment_3/src/models/dto"
	"Assignment_3/src/models/entity"
	"Assignment_3/src/response"
	"Assignment_3/src/services"
)

type AuthController interface {
	LogIn(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService services.AuthService
	jwtService  services.JWTService
}

func (a *authController) LogIn(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBindJSON(&loginDTO)

	if errDTO != nil {
		res := response.BuildErrResponse("Failed to process request", errDTO.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authResult := a.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(entity.Account); ok {
		generator := a.jwtService.GenerateToken(strconv.FormatUint(uint64(v.ID), 10))
		v.Token = generator

		res := response.BuildResponse(true, "OKELA", v)
		ctx.JSON(http.StatusOK, res)
		return
	}
	res := response.BuildErrResponse("Please check your credential", "Invalid", response.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
	return
}

func (a *authController) Register(ctx *gin.Context) {

	var regiserAccountDTO dto.RegisterDTO

	errDTO := ctx.ShouldBindJSON(&regiserAccountDTO)

	if errDTO != nil {

		fmt.Println("error DTO ?")

		res := response.BuildErrResponse("Failed to process request", errDTO.Error(), response.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	if !a.authService.IsDuplicateEmail(regiserAccountDTO.Email) {
		res := response.BuildErrResponse("Please enter email", "Email already exists", nil)
		ctx.AbortWithStatusJSON(http.StatusConflict, res)
		return
	} else {
		registerAccount := a.authService.Register(regiserAccountDTO)
		token := a.jwtService.GenerateToken(strconv.FormatUint(uint64(registerAccount.ID), 10))
		registerAccount.Token = token

		res := response.BuildResponse(
			true, "Successfully", registerAccount)
		ctx.JSON(http.StatusOK, res)

		return
	}

}

func NewAuthController(authService services.AuthService, jwtService services.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}
