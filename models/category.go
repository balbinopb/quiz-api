package models

import "time"

type Category struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:100;unique;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	Questions []Question `gorm:"foreignKey:CategoryID"`
}