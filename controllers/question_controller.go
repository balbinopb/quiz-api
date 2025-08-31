package controllers

import (
	"net/http"
	"quiz-api/database"
	"quiz-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetQuestions(c *gin.Context) {
	var questions []models.Question
	database.DB.Preload("Options").Find(&questions)
	c.JSON(http.StatusOK, questions)
}

func GetQuestionByID(c *gin.Context) {
	id := c.Param("id")
	var question models.Question
	if err := database.DB.Preload("Options").First(&question, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "question not found"})
		return
	}
	c.JSON(http.StatusOK, question)
}

func CreateQuestion(c *gin.Context) {
	var input models.Question
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&input)
	c.JSON(http.StatusCreated, input)
}

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

func DeleteQuestion(c *gin.Context) {
    var question models.Question

    // Convert id from string to uint
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }

    // First check if it exists
    if err := database.DB.First(&question, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "question not found"})
        return
    }

    // Delete (hard delete if needed)
    if err := database.DB.Delete(&question).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "question deleted"})
}

