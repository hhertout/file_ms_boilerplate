package router

import (
	"github.com/eco-challenge/controller"
	"github.com/gin-gonic/gin"
)

func Provider() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", controller.Ping)
	r.POST("/upload", controller.UploadFile)

	return r
}
