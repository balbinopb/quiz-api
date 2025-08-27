package main

import (
	"net/http"
	"quiz-api/config/initializers"
	"quiz-api/database"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	database.ConnectDB()
}

func main() {
	database.DBMigrate()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
