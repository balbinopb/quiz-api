package controllers

import (
	"net/http"
	"quiz-api/database"
	"quiz-api/models"

	"github.com/gin-gonic/gin"
)

func GetTopScores(c *gin.Context) {
	var results []models.QuizResult
	database.DB.Order("score desc").Limit(10).Find(&results)
	c.JSON(http.StatusOK, results)
}

func GetUsersReport(c *gin.Context) {
	var users []models.User
	database.DB.Preload("QuizResults").Find(&users)
	c.JSON(http.StatusOK, users)
}
