
// @title Quiz API
// @version 1.0
// @description A REST API for managing users, quizzes, and results.
// @termsOfService http://swagger.io/terms/

// @host localhost:8080
// @BasePath /api/v1
package main

import (
	"os"
	"quiz-api/config/initializers"
	"quiz-api/database"
	"quiz-api/routes"
	"quiz-api/seed"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "quiz-api/docs"
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

	if os.Getenv("APP_ENV") != "production" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	r.Run()
}
