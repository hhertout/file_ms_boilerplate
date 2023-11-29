package router

import (
	"github.com/eco-challenge/controller"
	"github.com/gin-gonic/gin"
)

func Provider() *gin.Engine {
	r := gin.New()

	r.GET("/ping", controller.Ping)

	return r
}
