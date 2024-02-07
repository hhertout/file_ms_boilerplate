package controller

import (
	"github.com/eco-challenge/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ExportJson(c *gin.Context) {
	var body map[string]interface{}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to parse Body",
			"error":   err.Error(),
		})
		return
	}

	j, err := service.NewExporter().Json(body)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed create json content",
			"error":   err.Error(),
		})
		return
	}

	c.Header("Content-Disposition", "attachment; filename=export.json")
	c.Data(http.StatusOK, "application/json", j)
}
