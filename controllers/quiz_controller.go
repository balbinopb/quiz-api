package controllers

import (
	"math/rand"
	"net/http"
	"quiz-api/database"
	"quiz-api/models"
	"time"

	"github.com/gin-gonic/gin"
)

func StartQuiz(c *gin.Context) {
	categoryID := c.Query("category")
	var questions []models.Question
	database.DB.Preload("Options").Where("category_id = ?", categoryID).Find(&questions)

	// Shuffle questions
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(questions), func(i, j int) { questions[i], questions[j] = questions[j], questions[i] })

	c.JSON(http.StatusOK, questions)
}

func SubmitQuiz(c *gin.Context) {
	userID := c.GetUint("user_id")

	var input struct {
		Answers map[uint]string `json:"answers"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	score := 0
	for qID, ans := range input.Answers {
		var question models.Question
		database.DB.First(&question, qID)
		if question.CorrectAnswer == ans {
			score++
		}
	}

	// Save result
	result := models.QuizResult{
		UserID: userID,
		Score:  score,
	}
	database.DB.Create(&result)

	c.JSON(http.StatusOK, gin.H{
		"score": score,
	})
}

func GetUserResults(c *gin.Context) {
	userID := c.GetInt("user_id")
	var results []models.QuizResult
	database.DB.Where("user_id = ?", userID).Find(&results)
	c.JSON(http.StatusOK, results)
}
