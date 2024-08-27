package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserInformation struct {
	ID                uuid.UUID `gorm:"type:uuid;primaryKey;"`
	UserID            uuid.UUID `gorm:"type:uuid;foreignKey;"`
	PcrLevel          string    `validate:"min=1,max=100"`
	TherapyActive     bool
	TreatmentFree     bool
	TreatmentFreeDate string `validate:"len=10"`
	MonitoringPlace   string `validate:"min=2,max=100"`
	PcrFrequent       string `validate:"min=2,max=100"`
	PotentialHydrogen string `validate:"min=1, max=5"`

	gorm.Model
}
