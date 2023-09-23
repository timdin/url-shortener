package route

import (
	"url-shortener/service"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine, service *service.URLHandler) {
	// shortern routing
	r.POST("/shortern", service.Shortern)
	// redirect routing
	r.GET("/:id", service.Redirect)
}
