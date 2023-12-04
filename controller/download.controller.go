package controller

import (
	"github.com/eco-challenge/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type Body struct {
	Id string `json:"id" binding:"required"`
}

func GetFile(c *gin.Context) {
	var body Body
	err := c.BindJSON(&body)
	if err != nil || body.Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "You must provide an id",
		})
		return
	}
	path, err := service.NewFileManager().GetBasePath(body.Id)
	if err != nil {
		if os.Getenv("ENV") == "dev" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "ID is invalid or the the service is temporally unavailable",
				"error":   err,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "ID is invalid or the the service is temporally unavailable",
			})
		}

		return
	}
	c.File(path)
}
