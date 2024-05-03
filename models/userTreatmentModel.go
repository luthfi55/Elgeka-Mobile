package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserTreatment struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey;"`
	UserID          uuid.UUID `gorm:"type:uuid;foreignKey;"`
	FirstTreatment  string    `validate:"required"`
	SecondTreatment string    `validate:"required"`
	gorm.Model
}

type GetTreatmentDataFailedResponse struct {
	ErrorMessage string      `json:"ErrorMessage"`
	Data         interface{} `json:"Data"`
	Link         []LinkItem  `json:"Link"`
}

type GetTreatmentDataSuccessResponse struct {
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
	Link    []LinkItem  `json:"Link"`
}
