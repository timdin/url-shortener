package route

import (
	"time"
	"url-shortener/service"

	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine, service *service.URLHandler) {
	redisStore := persist.NewRedisStore(service.Redis)
	r.POST("/shortern", service.Shortern)
	g := r.Group("/")
	{
		g.GET("/:id", service.Redirect)
		g.GET("/cache/:id", service.Redirect, cache.CacheByRequestURI(redisStore, 1*time.Minute))
	}
}
