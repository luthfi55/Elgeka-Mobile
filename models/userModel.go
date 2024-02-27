package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Name         string    `validate:"required"`
	Address      string    `validate:"required"`
	PhoneNumber  string    `validate:"required"`
	Email        string    `gorm:"unique" validate:"required,email"`
	Password     string    `validate:"required,min=8"`
	IsActive     bool
	OtpCode      string
	OtpCreatedAt time.Time
	OtpType      string
	gorm.Model
}

type Data struct {
	ID      uuid.UUID `json:"ID"`
	Email   string    `json:"Email"`
	OtpCode string    `json="OtpCode"`
}

type LinkItem struct {
	Name string `json:"Name"`
	Link string `json:"Link"`
}

type RegisterUserSuccessResponse struct {
	Message string     `json:"Message"`
	Data    []Data     `json:"Data"`
	Link    []LinkItem `json:"Link"`
}

type RegisterUserFailledResponse struct {
	ErrorMessage string     `json:"Message"`
	Data         []Data     `json:"Data"`
	Link         []LinkItem `json:"Link"`
}

type LoginUserSuccessResponse struct {
	Message string     `json:"Message"`
	Data    []Data     `json:"Data"`
	Link    []LinkItem `json:"Link"`
}

type LoginUserFailledResponse struct {
	ErrorMessage string     `json:"Message"`
	Data         []Data     `json:"Data"`
	Link         []LinkItem `json:"Link"`
}

type OtpData struct {
	Email string `json:"Email"`
}

type OtpFailledResponse struct {
	ErrorMessage string
	Data         []OtpData  `json:"Data"`
	Link         []LinkItem `json:"Link"`
}

type OtpSuccessResponse struct {
	Message string
	Data    []OtpData  `json:"OtpData"`
	Link    []LinkItem `json:"Link"`
}
