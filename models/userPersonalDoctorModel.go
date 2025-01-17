package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserPersonalDoctor struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;"`
	UserID    uuid.UUID `gorm:"type:uuid;foreignKey;"`
	DoctorID  uuid.UUID `gorm:"type:uuid;foreignKey;"`
	StartDate string    `validate:"required,len=10"`
	EndDate   string    `validate:"required,len=10"`
	Request   string    `validate:"required,min=2,max=8"`
	gorm.Model
}

type UserPersonalDoctorData struct {
	UserID        uuid.UUID `json:"UserID"`
	Name          string    `json:"Name"`
	Email         string    `json:"Email"`
	PhoneNumber   string    `json:"PhoneNumber"`
	Gender        string    `json:"Gender"`
	BloodGroup    string    `json:"BloodGroup"`
	DiagnosisDate string    `json:"DiagnosisDate"`
}

type UserPersonalDoctorID struct {
	ID string `json:"Personal Doctor ID"`
}

type AddPersonalDoctorFailledResponse struct {
	ErrorMessage string     `json:"ErrorMessage"`
	Data         string     `json:"Data"`
	Link         []LinkItem `json:"Link"`
}

type AddPersonalDoctorSuccessResponse struct {
	Message string                 `json:"Message"`
	Data    []UserPersonalDoctorID `json:"Data"`
	Link    []LinkItem             `json:"Link"`
}

type GetPersonalDoctorFailledResponse struct {
	ErrorMessage string     `json:"ErrorMessage"`
	Data         string     `json:"Data"`
	Link         []LinkItem `json:"Link"`
}

type GetPersonalDoctorSuccessResponse struct {
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
	Link    []LinkItem  `json:"Link"`
}

type ListAcceptancePatientFailedResponse struct {
	ErrorMessage string      `json:"ErrorMessage"`
	Data         interface{} `json:"Data"`
	Link         []LinkItem  `json:"Link"`
}

type ListAcceptancePatientSuccessResponse struct {
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
	Link    []LinkItem  `json:"Link"`
}

type DoctorPatientAcceptFailledResponse struct {
	ErrorMessage string      `json:"ErrorMessage"`
	Data         interface{} `json:"Data"`
	Link         []LinkItem  `json:"Link"`
}

type DoctorPatientAcceptSuccessResponse struct {
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
	Link    []LinkItem  `json:"Link"`
}

type DoctorPatientRejectFailledResponse struct {
	ErrorMessage string      `json:"ErrorMessage"`
	Data         interface{} `json:"Data"`
	Link         []LinkItem  `json:"Link"`
}

type DoctorPatientRejectSuccessResponse struct {
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
	Link    []LinkItem  `json:"Link"`
}

type ListDoctorPatientFailledResponse struct {
	ErrorMessage string      `json:"ErrorMessage"`
	Data         interface{} `json:"Data"`
	Link         []LinkItem  `json:"Link"`
}

type ListDoctorPatientSuccessResponse struct {
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
	Link    []LinkItem  `json:"Link"`
}

type DoctorPatientProfileFailledResponse struct {
	ErrorMessage string      `json:"ErrorMessage"`
	Data         interface{} `json:"Data"`
	Link         []LinkItem  `json:"Link"`
}

type DoctorPatientProfileSuccessResponse struct {
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
	Link    []LinkItem  `json:"Link"`
}

type DoctorPatientHealthStatusFailledResponse struct {
	ErrorMessage string      `json:"ErrorMessage"`
	Data         interface{} `json:"Data"`
	Link         []LinkItem  `json:"Link"`
}

type DoctorPatientHealthStatusSuccessResponse struct {
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
	Link    []LinkItem  `json:"Link"`
}
