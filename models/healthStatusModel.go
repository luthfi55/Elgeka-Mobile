package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BCR_ABL struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey;"`
	UserID uuid.UUID `gorm:"type:uuid;foreignKey;"`
	Data   float32   `validate:"required" json:"data"`
	Notes  string    `validate:"required" json:"notes"`
	Date   string    `validate:"required" json:"date"`
	gorm.Model
}
