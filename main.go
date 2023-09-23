package main

import (
	"time"
	"url-shortener/config"
	"url-shortener/internal"
	"url-shortener/logging"
	server "url-shortener/server"
)

func main() {
	configs := config.NewConfig()
	logging.InitLogger()
	logging.SugarLogger.Infow("init service",
		"configs", internal.DumpStruct(configs),
		"timestamp", time.Now().Format(time.RFC3339),
	)

	r := server.SetupServer(configs)

	// Listen and Server in 0.0.0.0:8080
	r.Run(":" + configs.Server.Port)
}
