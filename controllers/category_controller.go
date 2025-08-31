package controllers

import (
	"net/http"
	"quiz-api/database"
	"quiz-api/models"

	"github.com/gin-gonic/gin"
)


// GetCategories godoc
// @Summary Get all categories
// @Description Fetch all categories with related questions and options
// @Tags categories
// @Produce json
// @Success 200 {array} models.Category
// @Router /categories [get]
func GetCategories(c *gin.Context) {
    var categories []models.Category
    database.DB.Preload("Questions.Options").Find(&categories)
    c.JSON(http.StatusOK, categories)
}


// CreateCategory godoc
// @Summary Create new category
// @Description Add a new category
// @Tags categories
// @Accept json
// @Produce json
// @Param category body models.Category true "Category Data"
// @Success 201 {object} models.Category
// @Failure 400 {object} map[string]string
// @Router /categories [post]
func CreateCategory(c *gin.Context) {
	var input models.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&input)
	c.JSON(http.StatusCreated, input)
}

// UpdateCategory godoc
// @Summary Update a category
// @Description Update category by ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param category body models.Category true "Category Data"
// @Success 200 {object} models.Category
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /categories/{id} [put]
func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	if err := database.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	}

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&category)
	c.JSON(http.StatusOK, category)
}


// DeleteCategory godoc
// @Summary Delete a category
// @Description Delete category by ID
// @Tags categories
// @Param id path int true "Category ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /categories/{id} [delete]
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Category{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "category deleted"})
}
