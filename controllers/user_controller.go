package controllers

import (
	"net/http"
	"quiz-api/database"
	"quiz-api/models"

	"github.com/gin-gonic/gin"
)

// GetProfile godoc
// @Summary      Get user profile
// @Description  Retrieve the profile of the currently authenticated user
// @Tags         users
// @Produce      json
// @Success      200 {object} models.User
// @Failure      401 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Router       /profile [get]
// @Security     BearerAuth
func GetProfile(c *gin.Context) {
	userIDInterface, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found in context"})
		return
	}

	userID, ok := userIDInterface.(uint)
	if !ok || userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user ID"})
		return
	}

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
