package main

import (
	server "url-shortener/server"
)

func main() {
	r := server.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
