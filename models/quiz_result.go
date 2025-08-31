package models

import "time"

type QuizResult struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Score     int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}


type QuizResultResponse struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Score     int       `json:"score"`
}