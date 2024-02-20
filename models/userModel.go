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
	Password     string    `validate:"required,min=6"`
	IsActive     bool
	OtpCode      string
	OtpCreatedAt time.Time
	gorm.Model
}

type Data struct {
	Email string `json:"Email"`
}

type LinkItem struct {
	Name string `json:"Name"`
	Link string `json:"Link"`
}

type SignupSuccesResponse struct {
	Message string     `json:"Message"`
	Data    []Data     `json:"Data"`
	Link    []LinkItem `json:"Link"`
}

type SignupFailledResponse struct {
	Message string     `json:"Message"`
	Data    []Data     `json:"Data"`
	Link    []LinkItem `json:"Link"`
}

type OtpData struct {
	Id string `json:"User ID"`
}

type OtpFailledResponse struct {
	Message string
	Data    []OtpData  `json:"OtpData"`
	Link    []LinkItem `json:"Link"`
}
