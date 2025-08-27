package database

import (
	"log"
	"quiz-api/models"
)



func DBMigrate() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Question{},
		&models.Option{},
		&models.QuizResult{},
	)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
}