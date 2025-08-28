package models

type Option struct {
    ID         uint   `gorm:"primaryKey" json:"id"`
    QuestionID uint   `json:"question_id"`
    OptionText string `json:"option_text"`
}
