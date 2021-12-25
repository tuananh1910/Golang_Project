package controllers

import (
	"net/http"
	"strconv"
	_ "strconv"

	"github.com/gin-gonic/gin"

	"Assignment_3/src/response"
	"Assignment_3/src/services"
)

type ListController interface {
	FindListById(ctx *gin.Context)
}

type listController struct {
	listService services.ListService
}

// watch details weather forecast of a day
func (l *listController) FindListById(ctx *gin.Context) {
	var idList string
	idList = ctx.Query("id_list")

	if idList == "" {
		res := response.BuildErrResponse("Please come back later", "The system is having problems", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	idU64, err := strconv.ParseUint(idList, 2, 32)

	if err != nil {
		res := response.BuildErrResponse("Please come back later", "The system is having problems", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	listResult := l.listService.FindListById(uint32(idU64)) // convert idlist string to uint32 ~~~ 64?

	res := response.BuildResponse(true, "OK", listResult)
	ctx.JSON(http.StatusOK, res)

	return

}

func NewListController(listService services.ListService) ListController {
	return &listController{listService: listService}
}
