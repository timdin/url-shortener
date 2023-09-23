package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"

	"url-shortener/config"
	"url-shortener/dao"
	"url-shortener/internal"
	"url-shortener/route"
	"url-shortener/service"
	"url-shortener/validator"
)

func SetupServer(config *config.Config) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	db, err := dao.InitDB(config.DB.Conn)
	if err != nil {
		panic(err)
	}
	cache := redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    config.Cache.Conn,
	})
	urlWrapper := internal.NewURLWrapper(config.Server)
	validator := validator.NewUrlValidator(config.Server.AcceptExpired, config.Server.AcceptNoExpire)

	// setup service
	urlService := service.NewURLHandler(db, cache, urlWrapper, validator)
	route.Route(r, urlService)

	return r
}
