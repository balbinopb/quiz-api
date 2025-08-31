package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"size:100;unique;not null"`
	Email     string    `gorm:"size:150;unique;not null"`
	Password  string    `gorm:"size:255;not null"`
	Role      string    `gorm:"type:VARCHAR(20);default:player;check:role IN ('admin','player')"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}


//untuk response
type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}