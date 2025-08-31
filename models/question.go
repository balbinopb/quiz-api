package models

import "time"


type Question struct {
    ID           uint      `gorm:"primaryKey" json:"id"`
    CategoryID   uint      `json:"category_id"`
    QuestionText string    `json:"question_text"`
    CorrectAnswer string   `json:"correct_answer"`
    Options      []Option  `gorm:"foreignKey:QuestionID" json:"options"`//do we need to remove this?
    CreatedAt    time.Time `json:"created_at"`
}

