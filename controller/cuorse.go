package controller

import (
	"net/http"
	"uni/dataModels"
	"uni/utils"

	"github.com/gin-gonic/gin"
)

type CuorseController interface {
	CreateCuorse(ctx *gin.Context)
	GetCuorseData(ctx *gin.Context)
}

type cuorse struct {
	cuorseUtil utils.CuorseInterfaceUtils
}

func NewCuorseControllerInterface() CuorseController {
	var Ccontroller CuorseController

	c := cuorse{
		cuorseUtil: utils.NewCuorseInterfaceUtils(),
	}

	Ccontroller = &c

	return Ccontroller
}

func (c *cuorse) CreateCuorse(ctx *gin.Context) {

	var req dataModels.CreateCuorseRequest

	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "error on json binding: "+err.Error())
		return
	}

	id, err := c.cuorseUtil.CreateCuorse(req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, *id)

}

func (c *cuorse) GetCuorseData(ctx *gin.Context) {
	cId := ctx.Param("cuorse_id")

	cData, err := c.cuorseUtil.GetCuorseData(cId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, *cData)
}
