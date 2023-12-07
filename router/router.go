package router

import (
	"github.com/eco-challenge/config"
	"github.com/eco-challenge/controller"
	"github.com/gin-gonic/gin"
)

func Provider() *gin.Engine {
	r := gin.Default()

	r.Use(config.CORSMiddleware())

	r.GET("/ping", controller.Ping)
	r.GET("/download", controller.GetFile)

	r.POST("/upload", controller.UploadFile)

	return r
}
