package route

import (
	"url-shortener/service"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine, service *service.URLHandler) {
	g := r.Group("/redirect")
	{
		g.POST("/", service.Shortern)
		g.GET("/:id", service.Redirect)
	}
}
