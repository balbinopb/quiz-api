package controllers

import (
	"math/rand"
	"net/http"
	"quiz-api/database"
	"quiz-api/models"
	"time"

	"github.com/gin-gonic/gin"
)

// StartQuiz godoc
// @Summary Start a quiz
// @Description Get randomized questions by category
// @Tags quiz
// @Produce json
// @Param category query int true "Category ID"
// @Success 200 {array} models.Question
// @Router /quiz/start [get]
func StartQuiz(c *gin.Context) {
	categoryID := c.Query("category")
	var questions []models.Question
	database.DB.Preload("Options").Where("category_id = ?", categoryID).Find(&questions)

	// Shuffle questions
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(questions), func(i, j int) { questions[i], questions[j] = questions[j], questions[i] })

	c.JSON(http.StatusOK, questions)
}


// SubmitQuiz godoc
// @Summary Submit quiz answers
// @Description Submit answers and calculate score
// @Tags quiz
// @Accept json
// @Produce json
// @Param answers body map[uint]string true "Answers in {question_id: answer} format"
// @Success 200 {object} map[string]int
// @Failure 400 {object} map[string]string
// @Router /quiz/submit [post]
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


// GetUserResults godoc
// @Summary Get user quiz results
// @Description Get all past quiz results for logged-in user
// @Tags quiz
// @Produce json
// @Success 200 {array} models.QuizResult
// @Router /quiz/results [get]
func GetUserResults(c *gin.Context) {
	userID := c.GetInt("user_id")
	var results []models.QuizResult
	database.DB.Where("user_id = ?", userID).Find(&results)
	c.JSON(http.StatusOK, results)
}
