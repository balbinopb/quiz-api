package controllers

import (
	"net/http"
	"quiz-api/database"
	"quiz-api/models"

	"github.com/gin-gonic/gin"
)

// GetQuestions godoc
// @Summary Get all questions
// @Description Fetch all questions with their options
// @Tags questions
// @Produce json
// @Success 200 {array} models.Question
// @Router /questions [get]
func GetQuestions(c *gin.Context) {
	var questions []models.Question
	database.DB.Preload("Options").Find(&questions)
	c.JSON(http.StatusOK, questions)
}

// GetQuestionByID godoc
// @Summary Get question by ID
// @Description Fetch a question by ID with options
// @Tags questions
// @Produce json
// @Param id path int true "Question ID"
// @Success 200 {object} models.Question
// @Failure 404 {object} map[string]string
// @Router /questions/{id} [get]
func GetQuestionByID(c *gin.Context) {
	id := c.Param("id")
	var question models.Question
	if err := database.DB.Preload("Options").First(&question, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "question not found"})
		return
	}
	c.JSON(http.StatusOK, question)
}

// CreateQuestion godoc
// @Summary Create a question
// @Description Add a new question with options
// @Tags questions
// @Accept json
// @Produce json
// @Param question body models.Question true "Question Data"
// @Success 201 {object} models.Question
// @Failure 400 {object} map[string]string
// @Router /questions [post]
func CreateQuestion(c *gin.Context) {
	var input models.Question
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&input)
	c.JSON(http.StatusCreated, input)
}


// UpdateQuestion godoc
// @Summary Update a question
// @Description Update question by ID
// @Tags questions
// @Accept json
// @Produce json
// @Param id path int true "Question ID"
// @Param question body models.Question true "Question Data"
// @Success 200 {object} models.Question
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /questions/{id} [put]
func UpdateQuestion(c *gin.Context) {
	id := c.Param("id")
	var question models.Question
	if err := database.DB.First(&question, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "question not found"})
		return
	}

	if err := c.ShouldBindJSON(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&question)
	c.JSON(http.StatusOK, question)
}

// DeleteQuestion godoc
// @Summary Delete a question
// @Description Delete question by ID
// @Tags questions
// @Param id path int true "Question ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /questions/{id} [delete]
func DeleteQuestion(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Question{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "question not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "question deleted"})
}
