package main

import (
	"quiz-api/config/initializers"
	"quiz-api/database"
	"quiz-api/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	database.ConnectDB()
}

func main() {
	database.DBMigrate()

	r := gin.Default()
	routes.Routes(r)
	r.Run()
}
