package controller

import (
	"database/sql"
	"net/http"

	"github.com/TakayaSugiyama/web-service-gin/model"
	"github.com/gin-gonic/gin"
)

func MemberInfoIndex(c *gin.Context) {
	members, err := model.GetMemberInfos()
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, members)
}

func MemberInfoShow(c *gin.Context) {
	var id = c.Param("id")
	result, err := model.GetMemberInfo(id)
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
