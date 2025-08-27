package models

import "time"

type Question struct {
	ID            uint      `gorm:"primaryKey"`
	CategoryID    uint      `gorm:"not null"`
	Category      Category  `gorm:"foreignKey:CategoryID;constraint:OnDelete:CASCADE"`
	QuestionText  string    `gorm:"type:text;not null"`
	CorrectAnswer string    `gorm:"size:255;not null"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`

	Options []Option `gorm:"foreignKey:QuestionID"`
}