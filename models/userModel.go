package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID                 uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Name               string    `validate:"required"`
	Gender             string    `validate:"required"`
	BirthDate          string    `validate:"required"`
	BloodGroup         string    `validate:"required"`
	Province           string    `validate:"required"`
	District           string    `validate:"required"`
	SubDistrict        string    `validate:"required"`
	Village            string    `validate:"required"`
	Address            string    `validate:"required"`
	PhoneNumber        string    `validate:"required,max=14"`
	Email              string    `gorm:"unique" validate:"required,email"`
	Password           string    `validate:"required,min=8"`
	IsActive           bool
	OtpCode            string
	OtpCreatedAt       time.Time
	OtpType            string
	ForgotPasswordCode string
	BCR_ABL            []BCR_ABL            `gorm:"foreignKey:UserID"`
	UserPersonalDoctor []UserPersonalDoctor `gorm:"foreignKey:UserID"`
	Medicine           []Medicine           `gorm:"foreignKey:UserID"`
	MedicineSchedule   []MedicineSchedule   `gorm:"foreignKey:UserID"`
	Symptom            []SymptomAnswer      `gorm:"foreignKey:UserID"`

	gorm.Model
}

type UserData struct {
	ID          uuid.UUID `json:"ID"`
	Name        string    `json:"Name"`
	Email       string    `json:"Email"`
	Address     string    `json:"Address"`
	Province    string    `json:"Province"`
	District    string    `json:"District"`
	SubDistrict string    `json:"SubDistrict"`
	Village     string    `json:"Village"`
	Gender      string    `json:"Gender"`
	BirthDate   string    `json:"BirthDate"`
	BloodGroup  string    `json:"BloodGroup"`
	PhoneNumber string    `json:"PhoneNumber"`
}

type Data struct {
	ID      uuid.UUID `json:"ID"`
	Email   string    `json:"Email"`
	OtpCode string    `json:"OtpCode"`
}

type RegisterData struct {
	ID    uuid.UUID `json:"ID"`
	Email string    `json:"Email"`
}

type UserIdData struct {
	ID uuid.UUID `json:"ID"`
}

type Login struct {
	Name  string `json:"Name"`
	Email string `json:"Email"`
}

type LinkItem struct {
	Name string `json:"Name"`
	Link string `json:"Link"`
}

type RegisterUserSuccessResponse struct {
	Message string         `json:"Message"`
	Data    []RegisterData `json:"Data"`
	Link    []LinkItem     `json:"Link"`
}

type RegisterUserFailedResponse struct {
	ErrorMessage string     `json:"ErrorMessage"`
	Data         []Data     `json:"Data"`
	Link         []LinkItem `json:"Link"`
}

type LoginUserSuccessResponse struct {
	Message string     `json:"Message"`
	Data    []Login    `json:"Data"`
	Link    []LinkItem `json:"Link"`
}

type LoginUserWebsiteSuccessResponse struct {
	Message string     `json:"Message"`
	Data    []Login    `json:"Data"`
	Link    []LinkItem `json:"Link"`
	Token   string
}

type LoginUserFailedResponse struct {
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

type OtpFailedResponse struct {
	ErrorMessage string
	Data         []OtpData  `json:"Data"`
	Link         []LinkItem `json:"Link"`
}

type OtpSuccessResponse struct {
	Message string
	Data    []OtpData  `json:"OtpData"`
	Link    []LinkItem `json:"Link"`
}

type CheckOtpFailedResponse struct {
	ErrorMessage string
	Data         []CheckOtpData `json:"Data"`
	Link         []LinkItem     `json:"Link"`
}

type CheckOtpSuccessResponse struct {
	Message string
	Data    []CheckSuccessOtpData `json:"OtpData"`
	Link    []LinkItem            `json:"Link"`
}

type UpdateUserProfileSuccessResponse struct {
	Message string       `json:"Message"`
	Data    []UserIdData `json:"Data"`
	Link    []LinkItem   `json:"Link"`
}

type UpdateUserProfileFailedResponse struct {
	ErrorMessage string       `json:"ErrorMessage"`
	Data         []UserIdData `json:"Data"`
	Link         []LinkItem   `json:"Link"`
}

type GetProfileSuccessResponse struct {
	Message string     `json:"Message"`
	Data    []UserData `json:"Data"`
	Link    []LinkItem `json:"Link"`
}

type GetProfileFailedResponse struct {
	ErrorMessage string `json:"ErrorMessage"`
	Data         string `json:"Data"`
	Link         []LinkItem
}

type LogoutUserSuccessResponse struct {
	Message string     `json:"Message"`
	Data    string     `json:"Data"`
	Link    []LinkItem `json:"Link"`
}

type LogoutUserFailedResponse struct {
	ErrorMessage string     `json:"ErrorMessage"`
	Data         string     `json:"Data"`
	Link         []LinkItem `json:"Link"`
}

type ListUserWebsiteSuccessResponse struct {
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
	Link    []LinkItem  `json:"Link"`
}

type ListUserWebsiteFailledResponse struct {
	ErrorMessage string      `json:"ErrorMessage"`
	Data         interface{} `json:"Data"`
	Link         []LinkItem  `json:"Link"`
}
