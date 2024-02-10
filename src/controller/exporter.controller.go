package controller

import (
	"encoding/json"
	"fmt"
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

func ExportXml(c *gin.Context) {
	var body interface{}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to parse Body",
			"error":   err.Error(),
		})
		return
	}

	j, err := service.NewExporter().JsonToXml(body)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed create xml content",
			"error":   err.Error(),
		})
		return
	}

	c.Header("Content-Disposition", "attachment; filename=export.xml")
	c.Data(http.StatusOK, "application/xml", []byte(j))
}

func ExportCsv(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	var parsedBody []map[string]interface{}
	fmt.Printf("%v", parsedBody)
	if err := json.Unmarshal(body, &parsedBody); err != nil {
		c.JSON(400, gin.H{
			"message": "Error: Bad Request",
		})
		return
	}
	fmt.Printf("%v", parsedBody)

	csv, err := service.NewExporter().Csv(parsedBody)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed create csv content",
			"error":   err.Error(),
		})
		return
	}

	c.Header("Content-Disposition", "attachment; filename=export.csv")
	c.Data(http.StatusOK, "text/csv", csv.Bytes())
}
