package server

import (
	"github.com/gin-gonic/gin"

	"url-shortener/config"
	"url-shortener/route"
	"url-shortener/service"
)

func SetupServer(config *config.Config) *gin.Engine {
	r := gin.Default()

	// TODO: make the following init in the new url handler
	// db, err := dao.InitDB(config.DB.Conn)
	// if err != nil {
	// 	panic(err)
	// }
	// cache := redis.NewClient(&redis.Options{
	// 	Network: "tcp",
	// 	Addr:    config.Cache.Conn,
	// })
	// urlWrapper := internal.NewURLWrapper(config.Server)
	// validator := validator.NewUrlValidator(config.Server.AcceptExpired, config.Server.AcceptNoExpire)

	// setup service
	urlService := service.NewURLHandler(config)
	route.Route(r, urlService)

	return r
}
