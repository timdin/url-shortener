package main

import (
	"log"
	"url-shortener/config"
	"url-shortener/internal"
	server "url-shortener/server"
)

func main() {
	configs := config.NewConfig()
	log.Println(internal.DumpStruct(configs))
	// TODO: make a proper logger

	r := server.SetupServer(configs)

	// Listen and Server in 0.0.0.0:8080
	r.Run(":" + configs.Server.Port)
}
