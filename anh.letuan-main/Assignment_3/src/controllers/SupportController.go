package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"Assignment_3/src/models/dto"
	"Assignment_3/src/response"
	"Assignment_3/src/services"
)

type SupportController interface {
	Save(ctx *gin.Context)
}

type supportController struct {
	supportService services.SupportService
}

func (s *supportController) Save(ctx *gin.Context) {
	var supportDTO dto.Support
	errDTO := ctx.ShouldBindJSON(&supportDTO)

	if errDTO != nil {
		res := response.BuildErrResponse("Failed to process request", errDTO.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	support := s.supportService.Save(supportDTO)

	res := response.BuildResponse(
		true, "Successfully", support)
	ctx.JSON(http.StatusOK, res)
}

func NewSupportController(service services.SupportService) SupportController {
	return &supportController{supportService: service}
}
