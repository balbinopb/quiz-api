package main

import (
	"os"
	"quiz-api/config/initializers"
	"quiz-api/database"
	"quiz-api/routes"
	"quiz-api/seed"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	database.ConnectDB()
}

func main() {
	database.DBMigrate()

	if os.Getenv("APP_ENV") != "production" {
		seed.RunSeed()
	}

	r := gin.Default()
	routes.Routes(r)
	r.Run()
}
