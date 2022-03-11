package controller

import (
	"database/sql"
	"net/http"

	"github.com/TakayaSugiyama/web-service-gin/model"
	"github.com/gin-gonic/gin"
)

func MemberIndex(c *gin.Context) {
	members, err := model.GetMembers()
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, members)
}

func MemberShow(c *gin.Context) {
	var id = c.Param("id")
	result, err := model.GetMember(id)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	if err == nil {
		c.IndentedJSON(http.StatusOK, result)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status":        http.StatusNotFound,
			"error_message": "member is not found",
		})
	}
}
