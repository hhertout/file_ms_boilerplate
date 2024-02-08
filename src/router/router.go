package router

import (
	"github.com/eco-challenge/src/controller"
	"github.com/eco-challenge/src/middlewares"
	"github.com/gin-gonic/gin"
)

func Provider() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.CORSMiddleware())

	r.GET("/ping", controller.Ping)
	r.GET("/download", controller.GetFile)

	r.POST("/upload", controller.UploadFile)

	r.POST("/export/json", controller.ExportJson)
	r.POST("/export/xml", controller.ExportXml)

	return r
}
