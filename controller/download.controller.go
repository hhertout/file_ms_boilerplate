package controller

import (
	"github.com/eco-challenge/service"
	"github.com/gin-gonic/gin"
)

func GetFile(c *gin.Context) {
	basePath := service.NewUploadManager().GetBasePath()
	c.File(basePath + "123456.jpg")
}
