package models

import (
	"gorm.io/gorm"
)

type SymptomQuestion struct {
	ID       string `gorm:"primaryKey;"`
	Type     string `validate:"required"`
	Question string `validate:"required"`
	Option   string `validate:"required"`
	gorm.Model
}
