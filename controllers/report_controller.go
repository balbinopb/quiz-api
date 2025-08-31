package controllers

import (
	"net/http"
	"quiz-api/database"
	"quiz-api/models"

	"github.com/gin-gonic/gin"
)


// GetTopScores godoc
// @Summary Get Top 10 Scores
// @Description Retrieve the top 10 quiz results ordered by score.
// @Tags Results
// @Produce json
// @Success 200 {array} models.QuizResult
// @Failure 500 {object} map[string]string
// @Router /results/top [get]
func GetTopScores(c *gin.Context) {
	var results []models.QuizResult
	database.DB.Order("score desc").Limit(10).Find(&results)
	c.JSON(http.StatusOK, results)
}



// GetUsersReport godoc
// @Summary Get Users Report
// @Description Retrieve all users with their quiz results.
// @Tags Reports
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} map[string]string
// @Router /reports/users [get]
func GetUsersReport(c *gin.Context) {
	var users []models.User
	database.DB.Preload("QuizResults").Find(&users)
	c.JSON(http.StatusOK, users)
}
