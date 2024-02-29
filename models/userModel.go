package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID                 uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Name               string    `validate:"required"`
	Address            string    `validate:"required"`
	PhoneNumber        string    `validate:"required"`
	Email              string    `gorm:"unique" validate:"required,email"`
	Password           string    `validate:"required,min=8"`
	IsActive           bool
	OtpCode            string
	OtpCreatedAt       time.Time
	OtpType            string
	ForgotPasswordCode string
	BCR_ABL            []BCR_ABL `gorm:"foreignKey:UserID"`

	gorm.Model
}

type Data struct {
	ID      uuid.UUID `json:"ID"`
	Email   string    `json:"Email"`
	OtpCode string    `json="OtpCode"`
}
type UserIdData struct {
	ID uuid.UUID `json:"ID"`
}

type Login struct {
	Email string `json:"Email"`
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
	ErrorMessage string     `json:"ErrorMessage"`
	Data         []Data     `json:"Data"`
	Link         []LinkItem `json:"Link"`
}

type LoginUserSuccessResponse struct {
	Message string     `json:"Message"`
	Data    []Login    `json:"Data"`
	Link    []LinkItem `json:"Link"`
}

type LoginUserFailledResponse struct {
	ErrorMessage string       `json:"ErrorMessage"`
	Data         []UserIdData `json:"Data"`
	Link         []LinkItem   `json:"Link"`
}

type OtpData struct {
	Email string `json:"Email"`
}

type CheckOtpData struct {
	OtpCode string `json:"OtpCode"`
}

type CheckSuccessOtpData struct {
	ID      string
	Email   string
	OtpCode string
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

type CheckOtpFailledResponse struct {
	ErrorMessage string
	Data         []CheckOtpData `json:"Data"`
	Link         []LinkItem     `json:"Link"`
}

type CheckOtpSuccessResponse struct {
	Message string
	Data    []CheckSuccessOtpData `json:"OtpData"`
	Link    []LinkItem            `json:"Link"`
}
