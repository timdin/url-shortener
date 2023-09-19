package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"url-shortener/route"
	"url-shortener/service"
)

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// setup service
	urlService := service.NewMockURLHandler()
	route.Route(r, urlService)

	return r
}
