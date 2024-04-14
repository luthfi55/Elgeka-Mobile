package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Medicine struct {
	ID               uuid.UUID          `gorm:"type:uuid;primaryKey;"`
	UserID           uuid.UUID          `gorm:"type:uuid;foreignKey;"`
	Name             string             `json:"Name"`
	Stock            int                `json:"Stock"`
	MedicineSchedule []MedicineSchedule `gorm:"foreignKey:MedicineID"`
	gorm.Model
}

type MedicineSchedule struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;"`
	UserID     uuid.UUID `gorm:"type:uuid;foreignKey;"`
	MedicineID uuid.UUID `gorm:"type:uuid;foreignKey;"`
	Date       string    `json:"Date"`
	Status     bool      `json:"Status" gorm:"default:false"`
	gorm.Model
}

type MedicineData struct {
	ID    uuid.UUID `json:"ID"`
	Name  string    `json:"Name"`
	Stock int       `json:"Stock"`
}

type MedicineScheduleData struct {
	ID     uuid.UUID `json:"ID"`
	Date   string    `json:"Date"`
	Status bool      `json:"Status"`
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
