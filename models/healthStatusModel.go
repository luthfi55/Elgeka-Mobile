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
	Notes  string    `validate:"required,min=2,max=200" json:"notes"`
	Date   string    `validate:"required,len=10" json:"date"`
	gorm.Model
}

type Leukocytes struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey;"`
	UserID uuid.UUID `gorm:"type:uuid;foreignKey;"`
	Data   float32   `validate:"required" json:"data"`
	Notes  string    `validate:"required,min=2,max=200" json:"notes"`
	Date   string    `validate:"required,len=10" json:"date"`
	gorm.Model
}

type PotentialHydrogen struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey;"`
	UserID uuid.UUID `gorm:"type:uuid;foreignKey;"`
	Data   float32   `validate:"required" json:"data"`
	Notes  string    `validate:"required,min=2,max=200" json:"notes"`
	Date   string    `validate:"required,len=10" json:"date"`
	gorm.Model
}

type Hemoglobin struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey;"`
	UserID uuid.UUID `gorm:"type:uuid;foreignKey;"`
	Data   float32   `validate:"required" json:"data"`
	Notes  string    `validate:"required,min=2,max=200" json:"notes"`
	Date   string    `validate:"required,len=10" json:"date"`
	gorm.Model
}

type HeartRate struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey;"`
	UserID uuid.UUID `gorm:"type:uuid;foreignKey;"`
	Data   float32   `validate:"required" json:"data"`
	Notes  string    `validate:"required,min=2,max=200" json:"notes"`
	Date   string    `validate:"required,len=10" json:"date"`
	gorm.Model
}

type BloodPressure struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey;"`
	UserID  uuid.UUID `gorm:"type:uuid;foreignKey;"`
	DataSys float32   `validate:"required" json:"datasys"`
	DataDia float32   `validate:"required" json:"datadia"`
	Notes   string    `validate:"required,min=2,max=200" json:"notes"`
	Date    string    `validate:"required,len=10" json:"date"`
	gorm.Model
}

type Hematokrit struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey;"`
	UserID uuid.UUID `gorm:"type:uuid;foreignKey;"`
	Data   float32   `validate:"required" json:"data"`
	Notes  string    `validate:"required,min=2,max=200" json:"notes"`
	Date   string    `validate:"required,len=10" json:"date"`
	gorm.Model
}

type Trombosit struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey;"`
	UserID uuid.UUID `gorm:"type:uuid;foreignKey;"`
	Data   float32   `validate:"required" json:"data"`
	Notes  string    `validate:"required,min=2,max=200" json:"notes"`
	Date   string    `validate:"required,len=10" json:"date"`
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

type HematokritFailed struct {
	ErrorMessage string      `json:"errormessage"`
	Data         interface{} `json:"data"`
	Link         []LinkItem  `json:"link"`
}

type HematokritSuccess struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Link    []LinkItem  `json:"link"`
}

type TrombositFailed struct {
	ErrorMessage string      `json:"errormessage"`
	Data         interface{} `json:"data"`
	Link         []LinkItem  `json:"link"`
}

type TrombositSuccess struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Link    []LinkItem  `json:"link"`
}
