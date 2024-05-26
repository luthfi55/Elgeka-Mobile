package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserInformation struct {
	ID                uuid.UUID `gorm:"type:uuid;primaryKey;"`
	UserID            uuid.UUID `gorm:"type:uuid;foreignKey;"`
	PcrLevel          string    `validate:"required,min=1,max=100"`
	TherapyActive     bool
	TreatmentFree     bool
	TreatmentFreeDate string `validate:"required,len=10"`
	MonitoringPlace   string `validate:"required,min=2,max=100"`
	PcrFrequent       string `validate:"required,min=2,max=100"`

	gorm.Model
}
