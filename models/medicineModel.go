package models

import "github.com/google/uuid"

type Medicine struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey;"`
	UserID uuid.UUID `gorm:"type:uuid;foreignKey;"`
	Name   string    `json:"Name"`
	Stock  int       `json:"Stock"`
}

type MedicineData struct {
	ID    uuid.UUID `json:"ID"`
	Name  string    `json:"Name"`
	Stock int       `json:"Stock"`
}

type AddMedicineFailedResponse struct {
	ErrorMessage string         `json:"ErrorMessage"`
	Data         []MedicineData `json:"Data"`
	Link         []LinkItem     `json:"Link"`
}

type AddMedicineSuccessResponse struct {
	Message string         `json:"Message"`
	Data    []MedicineData `json:"Data"`
	Link    []LinkItem     `json:"Link"`
}

type GetMedicineFailedResponse struct {
	ErrorMessage string         `json:"ErrorMessage"`
	Data         []MedicineData `json:"Data"`
	Link         []LinkItem     `json:"Link"`
}

type GetMedicineSuccessResponse struct {
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
	Link    []LinkItem  `json:"Link"`
}

type UpdateMedicineFailedResponse struct {
	ErrorMessage string         `json:"ErrorMessage"`
	Data         []MedicineData `json:"Data"`
	Link         []LinkItem     `json:"Link"`
}

type UpdateMedicineSuccessResponse struct {
	Message string         `json:"Message"`
	Data    []MedicineData `json:"Data"`
	Link    []LinkItem     `json:"Link"`
}

type DeleteMedicineFailedResponse struct {
	ErrorMessage string         `json:"ErrorMessage"`
	Data         []MedicineData `json:"Data"`
	Link         []LinkItem     `json:"Link"`
}

type DeleteMedicineSuccessResponse struct {
	Message string         `json:"Message"`
	Data    []MedicineData `json:"Data"`
	Link    []LinkItem     `json:"Link"`
}
