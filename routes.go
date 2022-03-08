package main

import (
	"github.com/TakayaSugiyama/web-service-gin/controller"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine = gin.Default()

func initRoutes() {
	router.GET("/albums", controller.GetAlbums)
	router.GET("/albums/:id", controller.GetAlbumByID)
	router.DELETE("/albums/:id", controller.DeleteAlbumByID)
	router.POST("/albums", controller.PostAlbums)
}
