package seed

import (
	"quiz-api/database"
	"quiz-api/models"

	"golang.org/x/crypto/bcrypt"
)


// {
// 	"email":"admin@example.com",
// 	"password":"admin123"
// }

func RunSeed() {
	db := database.DB

	// Create admin user
	password, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	admin := models.User{
		Username: "admin",
		Email:    "admin@example.com",
		Password: string(password),
		Role:     "admin",
	}
	db.FirstOrCreate(&admin, models.User{Email: "admin@example.com"})

	// Create sample categories
	categories := []models.Category{
		{Name: "Math"},
		{Name: "Science"},
		{Name: "History"},
	}
	for _, c := range categories {
		db.FirstOrCreate(&c, models.Category{Name: c.Name})
	}

	// Create sample questions
	questions := []models.Question{
		{
			CategoryID:    1,
			QuestionText:  "What is 2 + 2?",
			CorrectAnswer: "4",
			Options: []models.Option{
				{OptionText: "3"},
				{OptionText: "4"},
				{OptionText: "5"},
			},
		},
		{
			CategoryID:    2,
			QuestionText:  "What planet is known as the Red Planet?",
			CorrectAnswer: "Mars",
			Options: []models.Option{
				{OptionText: "Mars"},
				{OptionText: "Venus"},
				{OptionText: "Jupiter"},
			},
		},
	}

	for _, q := range questions {
		var existing models.Question
		if err := db.Where("question_text = ?", q.QuestionText).First(&existing).Error; err != nil {
			// Question does not exist, create it along with options
			db.Create(&q)
		}
	}

}
