package logging

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// create a global logger
var SugarLogger *zap.SugaredLogger

func InitLogger() {
	logger, err := zap.NewProduction(zap.AddCaller())
	if err != nil {
		panic(err)
	}
	SugarLogger = logger.Sugar()
}

func GinLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		SugarLogger.Infow("gin request",
			"method", c.Request.Method,
			"timestamp", start.Format(time.RFC3339),
		)
		c.Next()
		elapsed := time.Since(start)

		SugarLogger.Infow("gin response",
			"method", c.Request.Method,
			"status", c.Writer.Status(),
			"elapsed", elapsed,
		)
	}
}
