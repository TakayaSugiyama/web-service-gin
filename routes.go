package main

import (
	"os"

	"github.com/TakayaSugiyama/web-service-gin/controller"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func setPortandGinMode() string {
	if value, ok := os.LookupEnv("PORT"); ok {
		gin.SetMode(gin.ReleaseMode)
		return value
	}
	return "8080"
}

func initRoutes() {
	port := setPortandGinMode()
	router := gin.Default()
	router.Use(cors.Default())
	//チーム情報
	router.GET("/teams", controller.TeamIndex)
	router.GET("/teams/:id", controller.TeamShow)
	router.GET("/teams/:id/members", controller.TeamMembers)
	//メンバー情報
	router.GET("/members", controller.MemberIndex)
	router.GET("/members/:id", controller.MemberShow)
	router.GET("/randommembers", controller.RandomMembers)
	router.GET("/randommember", controller.RandomMember)
	//メンバー詳細情報
	router.GET("/memberdetails", controller.MemberInfoIndex)
	router.GET("/memberdetails/:id", controller.MemberInfoShow)
	router.Run(":" + port)
}
