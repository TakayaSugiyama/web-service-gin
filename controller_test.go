package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TakayaSugiyama/web-service-gin/routes"
	"github.com/stretchr/testify/assert"
)

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
