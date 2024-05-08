package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type HealthStatusData struct {
	ID          uuid.UUID `json:"ID"`
	UserID      uuid.UUID `json:"UserID"`
	Name        string    `json:"Name"`
	Email       string    `json:"Email"`
	PhoneNumber string    `json:"PhoneNumber"`
	Data        float32   `json:"Data"`
	Notes       string    `json:"Notes"`
	Date        string    `json:"Date"`
}

type HealthStatusDataBloodPressure struct {
	ID          uuid.UUID `json:"ID"`
	UserID      uuid.UUID `json:"UserID"`
	Name        string    `json:"Name"`
	Email       string    `json:"Email"`
	PhoneNumber string    `json:"PhoneNumber"`
	DataSys     float32   `json:"DataSys"`
	DataDia     float32   `json:"DataDia"`
	Notes       string    `json:"Notes"`
	Date        string    `json:"Date"`
}

type BCR_ABL struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey;"`
	UserID uuid.UUID `gorm:"type:uuid;foreignKey;"`
	Data   float32   `validate:"required" json:"data"`
	Notes  string    `validate:"required" json:"notes"`
	Date   string    `validate:"required" json:"date"`
	gorm.Model
}

type BcrAblFailed struct {
	ErrorMessage string      `json:"errormessage"`
	Data         interface{} `json:"data"`
	Link         []LinkItem  `json:"link"`
}

type BcrAblSuccess struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Link    []LinkItem  `json:"link"`
}

type Leukocytes struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey;"`
	UserID uuid.UUID `gorm:"type:uuid;foreignKey;"`
	Data   float32   `validate:"required" json:"data"`
	Notes  string    `validate:"required" json:"notes"`
	Date   string    `validate:"required" json:"date"`
	gorm.Model
}

type LeukocytesFailed struct {
	ErrorMessage string      `json:"errormessage"`
	Data         interface{} `json:"data"`
	Link         []LinkItem  `json:"link"`
}

type LeukocytesSuccess struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Link    []LinkItem  `json:"link"`
}

type PotentialHydrogen struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey;"`
	UserID uuid.UUID `gorm:"type:uuid;foreignKey;"`
	Data   float32   `validate:"required" json:"data"`
	Notes  string    `validate:"required" json:"notes"`
	Date   string    `validate:"required" json:"date"`
	gorm.Model
}

type PotentialHydrogenFailed struct {
	ErrorMessage string      `json:"errormessage"`
	Data         interface{} `json:"data"`
	Link         []LinkItem  `json:"link"`
}

type PotentialHydrogenSuccess struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Link    []LinkItem  `json:"link"`
}

type Hemoglobin struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey;"`
	UserID uuid.UUID `gorm:"type:uuid;foreignKey;"`
	Data   float32   `validate:"required" json:"data"`
	Notes  string    `validate:"required" json:"notes"`
	Date   string    `validate:"required" json:"date"`
	gorm.Model
}

type HemoglobinFailed struct {
	ErrorMessage string      `json:"errormessage"`
	Data         interface{} `json:"data"`
	Link         []LinkItem  `json:"link"`
}

type HemoglobinSuccess struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Link    []LinkItem  `json:"link"`
}

type HeartRate struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey;"`
	UserID uuid.UUID `gorm:"type:uuid;foreignKey;"`
	Data   float32   `validate:"required" json:"data"`
	Notes  string    `validate:"required" json:"notes"`
	Date   string    `validate:"required" json:"date"`
	gorm.Model
}

type HeartRateFailed struct {
	ErrorMessage string      `json:"errormessage"`
	Data         interface{} `json:"data"`
	Link         []LinkItem  `json:"link"`
}

type HeartRateSuccess struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Link    []LinkItem  `json:"link"`
}

type BloodPressure struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey;"`
	UserID  uuid.UUID `gorm:"type:uuid;foreignKey;"`
	DataSys float32   `validate:"required" json:"datasys"`
	DataDia float32   `validate:"required" json:"datadia"`
	Notes   string    `validate:"required" json:"notes"`
	Date    string    `validate:"required" json:"date"`
	gorm.Model
}

type BloodPressureFailed struct {
	ErrorMessage string      `json:"errormessage"`
	Data         interface{} `json:"data"`
	Link         []LinkItem  `json:"link"`
}

type BloodPressureSuccess struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Link    []LinkItem  `json:"link"`
}
