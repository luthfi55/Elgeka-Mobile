package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserPersonalDoctor struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;"`
	UserID    uuid.UUID `gorm:"type:uuid;foreignKey;"`
	DoctorID  uuid.UUID `gorm:"type:uuid;foreignKey;"`
	StartDate string    `validate:"required"`
	EndDate   string    `validate:"required"`
	gorm.Model
}

type UserPersonalDoctorID struct {
	ID string `json:"Personal Doctor ID"`
}

type AddPersonalDoctorFailledResponse struct {
	ErrorMessage string     `json:"Message"`
	Data         string     `json:"Data"`
	Link         []LinkItem `json:"Link"`
}

type AddPersonalDoctorSuccessResponse struct {
	Message string                 `json:"Message"`
	Data    []UserPersonalDoctorID `json:"Data"`
	Link    []LinkItem             `json:"Link"`
}

type GetPersonalDoctorFailledResponse struct {
	ErrorMessage string     `json:"Message"`
	Data         string     `json:"Data"`
	Link         []LinkItem `json:"Link"`
}

type GetPersonalDoctorSuccessResponse struct {
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
	Link    []LinkItem  `json:"Link"`
}
