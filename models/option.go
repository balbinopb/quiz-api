package models

type Option struct {
	ID         uint     `gorm:"primaryKey"`
	QuestionID uint     `gorm:"not null"`
	Question   Question `gorm:"foreignKey:QuestionID;constraint:OnDelete:CASCADE"`
	OptionText string   `gorm:"size:255;not null"`
}