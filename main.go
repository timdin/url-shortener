package main

import (
	"os"
	"url-shortener/constants"
	"url-shortener/dao"
	"url-shortener/model"
	server "url-shortener/server"

	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		panic("Error loading .env file")
	}

	port := os.Getenv(constants.PORT)
	dbConfig := os.Getenv(constants.DB_CONFIG)
	db, err := dao.InitDB(dbConfig)
	if err != nil {
		panic("failed to connect database")
	}
	migrateErr := db.AutoMigrate(model.URL{})
	if migrateErr != nil {
		panic("failed to migrate database")
	}

	r := server.SetupRouter()

	// Listen and Server in 0.0.0.0:8080
	r.Run(":" + port)
}
