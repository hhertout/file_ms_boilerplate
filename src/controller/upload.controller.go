package controller

import (
	"github.com/eco-challenge/src/config"
	"github.com/eco-challenge/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	if file == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "File is missing",
		})
		return
	}
	id, err := service.NewFileManager().Save(file, []string{config.MIME_TYPE.Jpg, config.MIME_TYPE.Png}, "common/")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "upload successfully",
		"id":      id,
	})
}
