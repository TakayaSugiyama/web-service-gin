package main

import (
	"github.com/TakayaSugiyama/web-service-gin/controller"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine = gin.Default()

func initRoutes() {
	//チーム情報
	router.GET("/teams", controller.TeamIndex)
	router.GET("/teams/:id", controller.TeamShow)
	//メンバー情報
	router.GET("/members", controller.MemberIndex)
	router.GET("/members/:id", controller.MemberShow)
	//メンバー詳細情報
	router.GET("/memberdetails", controller.MemberInfoIndex)
	router.GET("/memberdetails/:id", controller.MemberInfoShow)
}
