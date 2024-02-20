package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Doctor struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Name         string    `validate:"required"`
	PolyName     string    `validate:"required"`
	HospitalName string    `validate:"required"`
	Email        string    `gorm:"unique" validate:"required,email"`
	Password     string    `validate:"required,min=8"`
	IsActive     bool
	gorm.Model
}

// type Data struct {
// 	Email string `json:"Email"`
// }

// type LinkItem struct {
// 	Name string `json:"Name"`
// 	Link string `json:"Link"`
// }

type RegisterDoctorSuccessResponse struct {
	Message string     `json:"Message"`
	Data    []Data     `json:"Data"`
	Link    []LinkItem `json:"Link"`
}

type RegisterDoctorFailledResponse struct {
	Message string     `json:"Message"`
	Data    []Data     `json:"Data"`
	Link    []LinkItem `json:"Link"`
}

// type OtpData struct {
// 	Id string `json:"User ID"`
// }

// type OtpFailledResponse struct {
// 	Message string
// 	Data    []OtpData  `json:"OtpData"`
// 	Link    []LinkItem `json:"Link"`
// }
