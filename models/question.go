package models

import "time"

// type Question struct {
// 	ID            uint      `gorm:"primaryKey"`
// 	CategoryID    uint      `gorm:"not null"`
// 	Category      Category  `gorm:"foreignKey:CategoryID;constraint:OnDelete:CASCADE"`
// 	QuestionText  string    `gorm:"type:text;not null"`
// 	CorrectAnswer string    `gorm:"size:255;not null"`
// 	CreatedAt     time.Time `gorm:"autoCreateTime"`

// 	Options []Option `gorm:"foreignKey:QuestionID"`
// }
type Question struct {
    ID           uint      `gorm:"primaryKey" json:"id"`
    CategoryID   uint      `json:"category_id"`
    QuestionText string    `json:"question_text"`
    CorrectAnswer string   `json:"correct_answer"`
    Options      []Option  `gorm:"foreignKey:QuestionID" json:"options"`
    CreatedAt    time.Time `json:"created_at"`
}

