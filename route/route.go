package route

import (
	"url-shortener/service"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine, service *service.URLHandler) {
	r.POST("/shortern", service.Shortern)
	g := r.Group("/")
	{
		g.GET("/:id", service.Redirect)
	}
}
