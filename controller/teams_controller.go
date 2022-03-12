package controller

import (
	"database/sql"
	"net/http"

	"github.com/TakayaSugiyama/web-service-gin/model"
	"github.com/gin-gonic/gin"
)

func TeamIndex(c *gin.Context) {
	teams, err := model.GetTeams()
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, teams)
}

func TeamShow(c *gin.Context) {
	var id = c.Param("id")
	result, err := model.GetTeam(id)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	if err == nil {
		c.IndentedJSON(http.StatusOK, result)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status":        http.StatusNotFound,
			"error_message": "team is not found",
		})
	}
}

func TeamMembers(c *gin.Context) {
	var id = c.Param("id")
	result, err := model.GetTeamMembers(id)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	if err == nil && result != nil {
		c.IndentedJSON(http.StatusOK, result)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status":        http.StatusNotFound,
			"error_message": "team or team members is not found",
		})
	}
}
