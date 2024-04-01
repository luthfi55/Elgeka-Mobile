package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Doctor struct {
	ID                 uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Name               string    `validate:"required"`
	PhoneNumber        string    `validate:"required,max=13"`
	Gender             string    `validate:"required"`
	PolyName           string    `validate:"required"`
	HospitalName       string    `validate:"required"`
	Email              string    `gorm:"unique" validate:"required,email"`
	Password           string    `validate:"required,min=8"`
	EmailActive        bool
	IsActive           bool
	OtpCode            string
	OtpCreatedAt       time.Time
	OtpType            string
	ForgotPasswordCode string
	UserPersonalDoctor []UserPersonalDoctor `gorm:"foreignKey:DoctorID"`
	gorm.Model
}

type DoctorData struct {
	ID   uuid.UUID `json:"ID"`
	Name string    `json:"Name"`
}

// type LinkItem struct {
// 	Name string `json:"Name"`
// 	Link string `json:"Link"`
// }

type GetListDoctorSuccessResponse struct {
	Message string       `json:"Message"`
	Data    []DoctorData `json:"Data"`
	Link    []LinkItem   `json:"Link"`
}

type RegisterDoctorSuccessResponse struct {
	Message string         `json:"Message"`
	Data    []RegisterData `json:"Data"`
	Link    []LinkItem     `json:"Link"`
}

type RegisterDoctorFailedResponse struct {
	ErrorMessage string     `json:"ErrorMessage"`
	Data         []Data     `json:"Data"`
	Link         []LinkItem `json:"Link"`
}

type ForgotPasswordUserSuccess struct {
	Message string     `json:"Message"`
	Data    []Data     `json:"Data"`
	Link    []LinkItem `json:"Link"`
}
