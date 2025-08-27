package routes

import (
	"github.com/gin-gonic/gin"
	"quiz-api/controllers"
	"quiz-api/middleware"
)

func Routes(r *gin.Engine) {
	api := r.Group("/api/v1")

	// Users & Auth
	users := api.Group("/users")
	{
		users.POST("/register", controllers.Register)
		users.POST("/login", controllers.Login)
		users.GET("/profile", middleware.JWTAuth(), controllers.GetProfile)
	}

	// Categories
	categories := api.Group("/categories", middleware.JWTAuth())
	{
		categories.GET("", controllers.GetCategories)

		// Admin only
		categories.POST("", middleware.AdminOnly(), controllers.CreateCategory)
		categories.PUT("/:id", middleware.AdminOnly(), controllers.UpdateCategory)
		categories.DELETE("/:id", middleware.AdminOnly(), controllers.DeleteCategory)
	}

	// Questions (Admin only)
	questions := api.Group("/questions", middleware.JWTAuth(), middleware.AdminOnly())
	{
		questions.GET("", controllers.GetQuestions)
		questions.GET("/:id", controllers.GetQuestionByID)
		questions.POST("", controllers.CreateQuestion)
		questions.PUT("/:id", controllers.UpdateQuestion)
		questions.DELETE("/:id", controllers.DeleteQuestion)
	}

	// Quiz (Player)
	quiz := api.Group("/quiz", middleware.JWTAuth())
	{
		quiz.GET("/start", controllers.StartQuiz)
		quiz.POST("/submit", controllers.SubmitQuiz)
		quiz.GET("/results", controllers.GetUserResults)
	}

	// Reports (Admin only)
	reports := api.Group("/reports", middleware.JWTAuth(), middleware.AdminOnly())
	{
		reports.GET("/top-scores", controllers.GetTopScores)
		reports.GET("/users", controllers.GetUsersReport)
	}
}
