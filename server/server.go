package server

import (
	"github.com/gin-gonic/gin"

	"url-shortener/config"
	"url-shortener/logging"
	"url-shortener/route"
	"url-shortener/service"
)

func SetupServer(config *config.Config) *gin.Engine {
	r := gin.Default()
	r.Use(logging.GinLoggerMiddleware())

	// setup service
	urlService := service.NewURLHandler(config)
	route.Route(r, urlService)

	return r
}
