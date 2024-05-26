package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SymptomAnswer struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;"`
	UserID     uuid.UUID `gorm:"type:uuid;foreignKey;"`
	Type       string    `validate:"required,min=3,max=20"`
	Answer     string    `validate:"required,min=2,max=500"`
	WordAnswer string    `validate:"required,min=2,max=500"`
	Date       string    `validate:"required,len=10"`
	gorm.Model
}

type SymptomAnswerData struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Type       string    `validate:"required"`
	Date       string    `validate:"required"`
	WordAnswer string    `validate:"required"`
}

type SymptomAnswerType struct {
	Type   string `validate:"required"`
	Answer string `validate:"required"`
	Date   string `validate:"required"`
}

type SubmitSymptomAnswerSuccess struct {
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
	Link    []LinkItem  `json:"Link"`
}

type SubmitSymptomAnswerFailed struct {
	ErrorMessage string     `json:"ErrorMessage"`
	Data         []string   `json:"Data"`
	Link         []LinkItem `json:"Link"`
}

type SymptomTypeNotFound struct {
	ErrorMessage string            `json:"ErrorMessage"`
	Data         SymptomAnswerType `json:"Data"`
	Link         []LinkItem        `json:"Link"`
}
