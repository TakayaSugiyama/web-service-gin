package main

import (
	"github.com/TakayaSugiyama/web-service-gin/controller"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine = gin.Default()

func initRoutes() {
	router.GET("/teams", controller.TeamIndex)
	//router.GET("/albums/:id", controller.TeamShow)
	//router.DELETE("/albums/:id", controller.TeamDestroy)
	//router.POST("/albums", controller.TeamCreate)
}
