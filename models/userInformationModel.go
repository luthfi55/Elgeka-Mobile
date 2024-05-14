package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserInformation struct {
	ID                uuid.UUID `gorm:"type:uuid;primaryKey;"`
	UserID            uuid.UUID `gorm:"type:uuid;foreignKey;"`
	PcrLevel          string    `validate:"required"`
	TherapyActive     bool      `validate:"required"`
	TreatmentFree     bool      `validate:"required"`
	TreatmentFreeDate string    `validate:"required"`
	MonitoringPlace   string    `validate:"required"`
	PcrFrequent       string    `validate:"required"`

	gorm.Model
}
