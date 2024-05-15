package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Doctor struct {
	ID                   uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Name                 string    `validate:"required"`
	PhoneNumber          string    `validate:"required,max=13"`
	Gender               string    `validate:"required"`
	PolyName             string    `validate:"required"`
	HospitalName         string    `validate:"required"`
	Email                string    `gorm:"unique" validate:"required,email"`
	Password             string    `validate:"required,min=8"`
	PasswordConfirmation string    `validate:"required,eqfield=Password"`
	EmailActive          bool
	IsActive             bool
	OtpCode              string
	OtpCreatedAt         time.Time
	OtpType              string
	ForgotPasswordCode   string
	UserPersonalDoctor   []UserPersonalDoctor `gorm:"foreignKey:DoctorID"`
	gorm.Model
}

type DoctorProfile struct {
	ID           uuid.UUID `json:"ID"`
	Name         string    `json:"Name"`
	PhoneNumber  string    `json:"PhoneNumber"`
	Email        string    `json:"Email"`
	Gender       string    `json:"Gender"`
	PolyName     string    `json:"PolyName"`
	HospitalName string    `json:"HospitalName"`
}

type DoctorData struct {
	ID   uuid.UUID `json:"ID"`
	Name string    `json:"Name"`
}

type DoctorPatientData struct {
	ID          uuid.UUID   `json:"ID"`
	DoctorName  string      `json:"DoctorName"`
	PatientData interface{} `json:"PatientData"`
}

type GetListDoctorSuccessResponse struct {
	Message string          `json:"Message"`
	Data    []DoctorProfile `json:"Data"`
	Link    []LinkItem      `json:"Link"`
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

type ListDoctorWebsiteSuccessResponse struct {
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
	Link    []LinkItem  `json:"Link"`
}

type ListDoctorWebsiteFailledResponse struct {
	ErrorMessage string      `json:"ErrorMessage"`
	Data         interface{} `json:"Data"`
	Link         []LinkItem  `json:"Link"`
}

type ListPatientDoctorWebsiteSuccessResponse struct {
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
	Link    []LinkItem  `json:"Link"`
}

type ListPatientDoctorWebsiteFailledResponse struct {
	ErrorMessage string      `json:"ErrorMessage"`
	Data         interface{} `json:"Data"`
	Link         []LinkItem  `json:"Link"`
}
