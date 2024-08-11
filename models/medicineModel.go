package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Medicine struct {
	ID               uuid.UUID          `gorm:"type:uuid;primaryKey;"`
	UserID           uuid.UUID          `gorm:"type:uuid;foreignKey;"`
	Name             string             `validate:"required,min=2,max=100" json:"Name"`
	Dosage           string             `validate:"required,min=1,max=10" json:"Dosage"`
	Category         string             `validate:"required,min=1,max=10" json:"Category"`
	Stock            int                `json:"Stock"`
	GetMedicineDate  time.Time          `json:"Time"`
	MedicineSchedule []MedicineSchedule `gorm:"foreignKey:MedicineID"`
	gorm.Model
}

type MedicineSchedule struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;"`
	UserID       uuid.UUID `gorm:"type:uuid;foreignKey;"`
	MedicineID   uuid.UUID `gorm:"type:uuid;foreignKey;"`
	MedicineName string    `validate:"required,min=2,max=100" json:"MedicineName"`
	Dosage       string    `validate:"required,min=1,max=10" json:"Dosage"`
	Day          string    `validate:"required,min=1,max=10" json:"Day"`
	Hour         string    `validate:"required,min=1,max=10" json:"Hour"`
	Status       bool      `json:"Status" gorm:"default:true"`
	gorm.Model
}

type MedicineData struct {
	ID       uuid.UUID `json:"ID"`
	Name     string    `json:"Name"`
	Dosage   string    `json:"Dosage"`
	Category string    `json:"Category"`
	Stock    int       `json:"Stock"`
}

type MedicineDataDate struct {
	ID       uuid.UUID `json:"ID"`
	Name     string    `json:"Name"`
	Dosage   string    `json:"Dosage"`
	Category string    `json:"Category"`
	Stock    int       `json:"Stock"`
	Date     string    `json:"Date"`
}

type MedicineDataResponse struct {
	UserID   uuid.UUID   `json:"UserID"`
	Name     string      `json:"Name"`
	Medicine interface{} `json:"Medicine"`
}

type MedicineDataWebsite struct {
	ID           uuid.UUID   `json:"ID"`
	Name         string      `json:"Name"`
	Email        string      `json:"Email"`
	PhoneNumber  string      `json:"PhoneNumber"`
	ListMedicine interface{} `json:"ListMedicine"`
}

type MedicineScheduleData struct {
	ID           uuid.UUID `json:"ID"`
	MedicineID   uuid.UUID `json:"MedicineID"`
	MedicineName string    `json:"MedicineName"`
	Dosage       string    `json:"Dosage"`
	Day          string    `json:"Day"`
	Hour         string    `json:"Hour"`
	Status       bool      `json:"Status"`
}

type AddMedicineFailedResponse struct {
	ErrorMessage string         `json:"ErrorMessage"`
	Data         []MedicineData `json:"Data"`
	Link         []LinkItem     `json:"Link"`
}

type AddMedicineScheduleFailedResponse struct {
	ErrorMessage string                 `json:"ErrorMessage"`
	Data         []MedicineScheduleData `json:"Data"`
	Link         []LinkItem             `json:"Link"`
}

type AddMedicineSuccessResponse struct {
	Message string         `json:"Message"`
	Data    []MedicineData `json:"Data"`
	Link    []LinkItem     `json:"Link"`
}

type AddMedicineScheduleSuccessResponse struct {
	Message string                 `json:"Message"`
	Data    []MedicineScheduleData `json:"Data"`
	Link    []LinkItem             `json:"Link"`
}

type GetMedicineFailedResponse struct {
	ErrorMessage string         `json:"ErrorMessage"`
	Data         []MedicineData `json:"Data"`
	Link         []LinkItem     `json:"Link"`
}

type GetMedicineScheduleFailedResponse struct {
	ErrorMessage string                 `json:"ErrorMessage"`
	Data         []MedicineScheduleData `json:"Data"`
	Link         []LinkItem             `json:"Link"`
}

type GetMedicineSuccessResponse struct {
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
	Link    []LinkItem  `json:"Link"`
}

type GetMedicineScheduleSuccessResponse struct {
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
	Link    []LinkItem  `json:"Link"`
}

type UpdateMedicineFailedResponse struct {
	ErrorMessage string         `json:"ErrorMessage"`
	Data         []MedicineData `json:"Data"`
	Link         []LinkItem     `json:"Link"`
}

type UpdateMedicineScheduleFailedResponse struct {
	ErrorMessage string                 `json:"ErrorMessage"`
	Data         []MedicineScheduleData `json:"Data"`
	Link         []LinkItem             `json:"Link"`
}

type UpdateMedicineSuccessResponse struct {
	Message string         `json:"Message"`
	Data    []MedicineData `json:"Data"`
	Link    []LinkItem     `json:"Link"`
}

type UpdateMedicineScheduleSuccessResponse struct {
	Message string                 `json:"Message"`
	Data    []MedicineScheduleData `json:"Data"`
	Link    []LinkItem             `json:"Link"`
}

type DeleteMedicineFailedResponse struct {
	ErrorMessage string         `json:"ErrorMessage"`
	Data         []MedicineData `json:"Data"`
	Link         []LinkItem     `json:"Link"`
}

type DeleteMedicineScheduleFailedResponse struct {
	ErrorMessage string                 `json:"ErrorMessage"`
	Data         []MedicineScheduleData `json:"Data"`
	Link         []LinkItem             `json:"Link"`
}

type DeleteMedicineSuccessResponse struct {
	Message string         `json:"Message"`
	Data    []MedicineData `json:"Data"`
	Link    []LinkItem     `json:"Link"`
}

type DeleteMedicineScheduleSuccessResponse struct {
	Message string                 `json:"Message"`
	Data    []MedicineScheduleData `json:"Data"`
	Link    []LinkItem             `json:"Link"`
}

type GetMedicineWebsiteFailedResponse struct {
	ErrorMessage string      `json:"ErrorMessage"`
	Data         interface{} `json:"Data"`
	Link         []LinkItem  `json:"Link"`
}

type GetMedicineWebsiteSuccessResponse struct {
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
	Link    []LinkItem  `json:"Link"`
}

type GetPatientMedicineWebsiteFailedResponse struct {
	ErrorMessage string      `json:"ErrorMessage"`
	Data         interface{} `json:"Data"`
	Link         []LinkItem  `json:"Link"`
}

type GetPatientMedicineWebsiteSuccessResponse struct {
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
	Link    []LinkItem  `json:"Link"`
}
