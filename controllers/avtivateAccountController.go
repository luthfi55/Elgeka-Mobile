package controllers

import (
	"elgeka-mobile/initializers"
	"elgeka-mobile/models"
	doctorresponse "elgeka-mobile/response/DoctorResponse"
	otpresponse "elgeka-mobile/response/OtpResponse"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ActivateAccountController(r *gin.Engine) {
	//user
	r.POST("api/user/activate/:user_id", Activate)
	r.POST("api/user/refresh_code/:user_id", RefreshOtpCode)

	//doctor
	r.POST("api/doctor/activate_account/:doctor_id", ActivateDoctor)
	r.POST("api/doctor/activate_email/:doctor_id", ActivateEmailDoctor)
	r.GET("api/doctor/list_activate", ListActivateDoctor)
	r.POST("api/doctor/refresh_code/:doctor_id", RefreshDoctorOtpCode)
	r.POST("api/doctor/reject_activation/:doctor_id", RejectDoctor)
}

func Activate(c *gin.Context) {
	userID := c.Param("user_id")
	data := userID

	var body struct {
		OtpCode string `json:"OtpCode"`
	}

	if c.Bind(&body) != nil {
		activationLink := "http://localhost:3000/api/user/register"
		otpresponse.FailedResponse(c, "Failed To Read Body", data, activationLink, http.StatusBadRequest)

		return
	}

	var user models.User

	// Find the user by ID
	result := initializers.DB.First(&user, "id = ?", userID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			activationLink := "http://localhost:3000/api/user/register"
			otpresponse.FailedResponse(c, "User Not Found", data, activationLink, http.StatusNotFound)

			return
		} else {
			activationLink := "http://localhost:3000/api/user/register"
			otpresponse.FailedResponse(c, "Database Error", data, activationLink, http.StatusInternalServerError)

			return
		}
	}

	if user.IsActive {
		activationLink := "http://localhost:3000/api/user/register"
		otpresponse.FailedResponse(c, "User Account Already Active", data, activationLink, http.StatusUnauthorized)

		return
	}

	if user.OtpCode != body.OtpCode {
		activationLink := "http://localhost:3000/api/user/register"
		otpresponse.FailedResponse(c, "Incorrect OTP code", data, activationLink, http.StatusUnauthorized)
		return
	} else {
		// 1 minute expired
		if time.Since(user.OtpCreatedAt) > time.Minute {
			activationLink := "http://localhost:3000/api/user/register"
			otpresponse.FailedResponse(c, "OTP Code Expired", data, activationLink, http.StatusUnauthorized)
			return
		}

		if user.OtpType != "Activation" {
			activationLink := "http://localhost:3000/api/user/register"
			otpresponse.FailedResponse(c, "Incorrect OTP code", data, activationLink, http.StatusUnauthorized)
			return
		}

		user.IsActive = true
		// Save the updated user
		if err := initializers.DB.Save(&user).Error; err != nil {
			activationLink := "http://localhost:3000/api/user/register"
			otpresponse.FailedResponse(c, "Failed to Activate", data, activationLink, http.StatusInternalServerError)
			return
		}

		activationLink := "http://localhost:3000/api/user/login"
		otpresponse.SuccessResponse(c, "User Activated Successfully", data, activationLink, http.StatusOK)
		return
	}
}

func ActivateDoctor(c *gin.Context) {
	doctorID := c.Param("doctor_id")
	data := doctorID

	var doctor models.Doctor

	// Find the doctor by ID
	result := initializers.DB.First(&doctor, "id = ?", doctorID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			activationLink := "http://localhost:3000/api/user/register"
			otpresponse.FailedResponse(c, "Doctor Not Found", data, activationLink, http.StatusNotFound)

			return
		} else {
			activationLink := "http://localhost:3000/api/user/register"
			otpresponse.FailedResponse(c, "Database Error", data, activationLink, http.StatusInternalServerError)

			return
		}
	}

	if doctor.IsActive {
		activationLink := "http://localhost:3000/api/user/register"
		otpresponse.FailedResponse(c, "Doctor Account Already Active", data, activationLink, http.StatusUnauthorized)

		return
	}

	if !doctor.EmailActive {
		activationLink := "http://localhost:3000/api/user/register"
		otpresponse.FailedResponse(c, "Doctor Email Account Must Active", data, activationLink, http.StatusUnauthorized)

		return
	}

	doctor.IsActive = true
	// Save the updated user
	if err := initializers.DB.Save(&doctor).Error; err != nil {
		activationLink := "http://localhost:3000/api/user/register"
		otpresponse.FailedResponse(c, "Failed to Activate", data, activationLink, http.StatusInternalServerError)
		return
	}

	activationLink := "http://localhost:3000/api/user/login"
	otpresponse.SuccessResponse(c, "Doctor Activated Successfully", data, activationLink, http.StatusOK)
}

func ActivateEmailDoctor(c *gin.Context) {
	doctorID := c.Param("doctor_id")
	data := doctorID

	var body struct {
		OtpCode string `json:"OtpCode"`
	}

	if c.Bind(&body) != nil {
		activationLink := "http://localhost:3000/api/doctor/register"
		otpresponse.FailedResponse(c, "Failed To Read Body", data, activationLink, http.StatusBadRequest)

		return
	}

	var doctor models.Doctor

	// Find the doctor by ID
	result := initializers.DB.First(&doctor, "id = ?", doctorID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			activationLink := "http://localhost:3000/api/doctor/register"
			otpresponse.FailedResponse(c, "Doctor Not Found", data, activationLink, http.StatusNotFound)

			return
		} else {
			activationLink := "http://localhost:3000/api/doctor/register"
			otpresponse.FailedResponse(c, "Database Error", data, activationLink, http.StatusInternalServerError)

			return
		}
	}

	if doctor.EmailActive {
		activationLink := "http://localhost:3000/api/doctor/register"
		otpresponse.FailedResponse(c, "Doctor Account Already Active", data, activationLink, http.StatusUnauthorized)

		return
	}

	if doctor.OtpCode != body.OtpCode {
		activationLink := "http://localhost:3000/api/doctor/register"
		otpresponse.FailedResponse(c, "Incorrect OTP code", data, activationLink, http.StatusUnauthorized)
		return
	} else {
		// 1 minute expired
		if time.Since(doctor.OtpCreatedAt) > time.Minute {
			activationLink := "http://localhost:3000/api/doctor/register"
			otpresponse.FailedResponse(c, "OTP Code Expired", data, activationLink, http.StatusUnauthorized)
			return
		}

		if doctor.OtpType != "Activation" {
			activationLink := "http://localhost:3000/api/doctor/register"
			otpresponse.FailedResponse(c, "Incorrect OTP code", data, activationLink, http.StatusUnauthorized)
			return
		}

		doctor.EmailActive = true
		// Save the updated doctor
		if err := initializers.DB.Save(&doctor).Error; err != nil {
			activationLink := "http://localhost:3000/api/doctor/register"
			otpresponse.FailedResponse(c, "Failed to Activate", data, activationLink, http.StatusInternalServerError)
			return
		}

		activationLink := "http://localhost:3000/api/doctor/login"
		otpresponse.SuccessResponse(c, "Doctor Email Activated Successfully", data, activationLink, http.StatusOK)
		return
	}
}

func RefreshOtpCode(c *gin.Context) {
	userID := c.Param("user_id")
	data := userID

	var user models.User
	initializers.DB.First(&user, "id = ?", userID)
	result := initializers.DB.First(&user, "id = ?", userID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			activationLink := "http://localhost:3000/api/user/register"
			otpresponse.FailedResponse(c, "User Not Found", data, activationLink, http.StatusNotFound)

			return
		} else {
			activationLink := "http://localhost:3000/api/user/register"
			otpresponse.FailedResponse(c, "Database Error", data, activationLink, http.StatusInternalServerError)

			return
		}
	}

	rand.Seed(time.Now().UnixNano())
	otpCode := fmt.Sprintf("%04d", rand.Intn(10000))

	user.OtpCode = otpCode
	user.OtpCreatedAt = time.Now().Add(time.Minute)
	user.OtpType = "Activation"

	if err := initializers.DB.Save(&user).Error; err != nil {
		activationLink := "http://localhost:3000/api/user/register"
		otpresponse.FailedResponse(c, "Failed to Update Otp Code", data, activationLink, http.StatusInternalServerError)
		return
	}

	SendEmailWithGmail(user.Email, otpCode)

	activationLink := "http://localhost:3000/api/user/activate/" + userID
	otpresponse.SuccessResponse(c, "Refresh OTP Successfully", data, activationLink, http.StatusOK)
}

func RefreshDoctorOtpCode(c *gin.Context) {
	doctorID := c.Param("doctor_id")
	data := doctorID

	var doctor models.Doctor
	initializers.DB.First(&doctor, "id = ?", doctorID)
	result := initializers.DB.First(&doctor, "id = ?", doctorID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			activationLink := "http://localhost:3000/api/doctor/register"
			otpresponse.FailedResponse(c, "Doctor Not Found", data, activationLink, http.StatusNotFound)

			return
		} else {
			activationLink := "http://localhost:3000/api/doctor/register"
			otpresponse.FailedResponse(c, "Database Error", data, activationLink, http.StatusInternalServerError)

			return
		}
	}

	rand.Seed(time.Now().UnixNano())
	otpCode := fmt.Sprintf("%04d", rand.Intn(10000))

	doctor.OtpCode = otpCode
	doctor.OtpCreatedAt = time.Now().Add(time.Minute)
	doctor.OtpType = "Activation"

	if err := initializers.DB.Save(&doctor).Error; err != nil {
		activationLink := "http://localhost:3000/api/doctor/register"
		otpresponse.FailedResponse(c, "Failed to Update Otp Code", data, activationLink, http.StatusInternalServerError)
		return
	}

	SendEmailWithGmail(doctor.Email, otpCode)

	activationLink := "http://localhost:3000/api/doctor/activate_email/" + doctorID
	otpresponse.SuccessResponse(c, "Refresh OTP Successfully", data, activationLink, http.StatusOK)
}

func getValidationErrorTagMessage(tag string) string {
	// Definisi pesan kustom untuk tag validasi tertentu
	switch tag {
	case "required":
		return "Cant Be Empty"
	case "email":
		return "Must Be a Valid Email Address"
	case "min":
		return "Password Must Be At Least 8 Letters"
	default:
		return fmt.Sprintf("validation Failed for Tag: %s", tag)
	}
}

func isEmailUnique(email string) bool {
	var userCount, doctorCount int64

	// Pengecekan email di tabel User
	initializers.DB.Model(&models.User{}).Where("email = ?", email).Count(&userCount)

	// Pengecekan email di tabel Doctor
	initializers.DB.Model(&models.Doctor{}).Where("email = ?", email).Count(&doctorCount)

	// Jika jumlah lebih dari 0, email sudah ada di salah satu tabel
	return (userCount + doctorCount) == 0
}

func ListActivateDoctor(c *gin.Context) {
	var doctors []models.Doctor
	var response []models.DoctorData
	result := initializers.DB.Where("is_active = ?", false).Where("email_active = ?", true).Find(&doctors)

	if result.Error != nil {
		activationLink := "http://localhost:3000/api/user/activate/:user_id"
		doctorresponse.GetInactiveDoctorFailedResponse(c, "Failed To Retrieve Data", response, activationLink, http.StatusInternalServerError)
		return
	}

	for _, doctor := range doctors {
		response = append(response, models.DoctorData{
			ID:   doctor.ID,
			Name: doctor.Name,
		})
	}

	if len(doctors) == 0 {
		activationLink := "http://localhost:3000/api/user/activate/:user_id"
		doctorresponse.GetInactiveDoctorFailedResponse(c, "Data Empty", response, activationLink, http.StatusInternalServerError)
		return
	}

	activationLink := "http://localhost:3000/api/user/activate/:user_id"
	doctorresponse.GetInactiveDoctorSuccessResponse(c, os.Getenv("GMAIL_PASSWORD"), response, activationLink, http.StatusOK)
}

func RejectDoctor(c *gin.Context) {
	doctorID := c.Param("doctor_id")
	data := doctorID

	var doctor models.Doctor
	initializers.DB.First(&doctor, "id = ?", doctorID)
	result := initializers.DB.First(&doctor, "id = ?", doctorID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			activationLink := "http://localhost:3000/api/doctor/register"
			otpresponse.FailedResponse(c, "Doctor Not Found", data, activationLink, http.StatusNotFound)

			return
		} else {
			activationLink := "http://localhost:3000/api/doctor/register"
			otpresponse.FailedResponse(c, "Database Error", data, activationLink, http.StatusInternalServerError)

			return
		}
	}

	result = initializers.DB.Unscoped().Delete(&doctor, "id = ?", doctorID)
	if result.Error != nil {
		activationLink := "http://localhost:3000/api/doctor/register"
		otpresponse.FailedResponse(c, result.Error.Error(), data, activationLink, http.StatusInternalServerError)
		return
	}

	activationLink := "http://localhost:3000/api/doctor/reject/" + doctorID
	otpresponse.SuccessResponse(c, "Reject Doctor Successfully", data, activationLink, http.StatusOK)
}
