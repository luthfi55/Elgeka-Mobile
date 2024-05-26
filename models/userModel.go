package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID                   uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Name                 string    `validate:"required,min=2,max=100"`
	Gender               string    `validate:"required,min=1,max=10"`
	BirthDate            string    `validate:"required,len=10"`
	BloodGroup           string    `validate:"required,min=1,max=3"`
	DiagnosisDate        string    `validate:"required,len=10"`
	Province             string    `validate:"required,min=2,max=100"`
	District             string    `validate:"required,min=2,max=100"`
	SubDistrict          string    `validate:"required,min=2,max=100"`
	Village              string    `validate:"required,min=2,max=100"`
	Address              string    `validate:"required,min=10,max=200"`
	PhoneNumber          string    `validate:"required,min=10,max=14"`
	Email                string    `gorm:"unique" validate:"required,email,max=100"`
	Password             string    `validate:"required,min=8,max=50"`
	PasswordConfirmation string    `validate:"required,eqfield=Password"`
	IsActive             bool
	OtpCode              string
	OtpCreatedAt         time.Time
	OtpType              string
	ForgotPasswordCode   string
	BCR_ABL              []BCR_ABL            `gorm:"foreignKey:UserID"`
	UserPersonalDoctor   []UserPersonalDoctor `gorm:"foreignKey:UserID"`
	Medicine             []Medicine           `gorm:"foreignKey:UserID"`
	MedicineSchedule     []MedicineSchedule   `gorm:"foreignKey:UserID"`
	Symptom              []SymptomAnswer      `gorm:"foreignKey:UserID"`
	UserTreatment        []UserTreatment      `gorm:"foreignKey:UserID"`
	UserInformation      []UserInformation    `gorm:"foreignKey:UserID"`

	gorm.Model
}

type UserInformationData struct {
	ID                uuid.UUID `json:"ID"`
	Name              string    `json:"Name"`
	Email             string    `json:"Email"`
	Address           string    `json:"Address"`
	Province          string    `json:"Province"`
	District          string    `json:"District"`
	SubDistrict       string    `json:"SubDistrict"`
	Village           string    `json:"Village"`
	Gender            string    `json:"Gender"`
	BirthDate         string    `json:"BirthDate"`
	BloodGroup        string    `json:"BloodGroup"`
	DiagnosisDate     string    `json:"DiagnosisDate"`
	PhoneNumber       string    `json:"PhoneNumber"`
	PcrLevel          string    `json:"PcrLevel"`
	TherapyActive     bool      `json:"TherapyActive"`
	TreatmentFree     bool      `json:"TreatmentFree"`
	TreatmentFreeDate string    `json:"TreatmentFreeDate"`
	MonitoringPlace   string    `json:"MonitoringPlace"`
	PcrFrequent       string    `json:"PcrFrequent"`
}

type UserData struct {
	ID            uuid.UUID `json:"ID"`
	Name          string    `json:"Name"`
	Email         string    `json:"Email"`
	Address       string    `json:"Address"`
	Province      string    `json:"Province"`
	District      string    `json:"District"`
	SubDistrict   string    `json:"SubDistrict"`
	Village       string    `json:"Village"`
	Gender        string    `json:"Gender"`
	BirthDate     string    `json:"BirthDate"`
	BloodGroup    string    `json:"BloodGroup"`
	DiagnosisDate string    `json:"DiagnosisDate"`
	PhoneNumber   string    `json:"PhoneNumber"`
}

type UserDataAge struct {
	ID            uuid.UUID `json:"ID"`
	Name          string    `json:"Name"`
	Email         string    `json:"Email"`
	Address       string    `json:"Address"`
	Province      string    `json:"Province"`
	District      string    `json:"District"`
	SubDistrict   string    `json:"SubDistrict"`
	Village       string    `json:"Village"`
	Gender        string    `json:"Gender"`
	BirthDate     string    `json:"BirthDate"`
	Age           int       `json:"Age"`
	BloodGroup    string    `json:"BloodGroup"`
	DiagnosisDate string    `json:"DiagnosisDate"`
	PhoneNumber   string    `json:"PhoneNumber"`
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
	Message string                `json:"Message"`
	Data    []UserInformationData `json:"Data"`
	Link    []LinkItem            `json:"Link"`
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
