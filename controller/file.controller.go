package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strings"
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

	filename := strings.Replace(file.Filename, " ", "_", -1)
	filePath := filepath.Base(".") + "/upload/" + filename

	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "upload successfully",
	})
}
