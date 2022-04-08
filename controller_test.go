package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/TakayaSugiyama/web-service-gin/routes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	code := m.Run()
	os.Exit(code)
}

func TestMemberIndex(t *testing.T) {
	router := routes.InitRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/members", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestMemberShow(t *testing.T) {
	router := routes.InitRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/members/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "秋吉 優花")
}

func TestMemberShowNotFound(t *testing.T) {
	router := routes.InitRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/members/99", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	assert.Equal(t, w.Body.String(), "{\"error_message\":\"member is not found\",\"status\":404}")
}
