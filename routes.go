package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine = gin.Default()

func initRoutes() {
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.DELETE("/albums/:id", deleteAlbumByID)
	router.POST("/albums", postAlbums)
}
