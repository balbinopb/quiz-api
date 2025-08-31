package controllers

import (
	"net/http"
	"quiz-api/database"
	"quiz-api/models"

	"github.com/gin-gonic/gin"
)

// GET METHOD
func GetCategories(c *gin.Context) {
	var categories []models.Category
	database.DB.Preload("Questions.Options").Find(&categories)
	c.JSON(http.StatusOK, categories)
}

// POST METHOD
func CreateCategory(c *gin.Context) {
	var input models.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&input)
	c.JSON(http.StatusCreated, input)
}

// PUT METHOD
func UpdateCategory(c *gin.Context) {
	id := c.Param("id")

	var category models.Category
	if err := database.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	}

	// Create a struct for update input
	var input struct {
		Name string `json:"name"`
		// add other updatable fields
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update only fields provided
	category.Name = input.Name

	database.DB.Save(&category)
	c.JSON(http.StatusOK, category)
}

// DELETE METHOD
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Category{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "category deleted"})
}
