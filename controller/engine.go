package controller

import "github.com/gin-gonic/gin"

func Run(port string) {

	en := gin.Default()

	cControllerInterface := NewCuorseControllerInterface()

	en.POST("/cuorse/create", cControllerInterface.CreateCuorse)
	en.GET("/cuorse/get/:cuorse_id", cControllerInterface.GetCuorseData)

	en.Run(port)
}
