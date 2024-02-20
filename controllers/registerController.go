package controllers

import (
	"elgeka-mobile/initializers"
	"elgeka-mobile/models"
	otpresponse "elgeka-mobile/response/OtpResponse"
	userresponse "elgeka-mobile/response/UserResponse"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterController(r *gin.Engine) {
	r.POST("api/user/register", UserRegister)
	r.POST("api/user/activate/:user_id", Activate)
	r.POST("api/user/refresh_code/:user_id", RefreshOtpCode)
}

func UserRegister(c *gin.Context) {
	var body models.User

	if c.Bind(&body) != nil {
		data := body
		activationLink := "http://localhost:3000/api/user/register"
		userresponse.RegisterFailedResponse(c, "Failed To Read Body", data, activationLink, http.StatusBadRequest)

		return
	}

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("%s %s", err.Field(), getValidationErrorTagMessage(err.Tag())))
		}
		errorMessage := strings.Join(validationErrors, ", ") // Join errors into a single string
		data := body
		activationLink := "http://localhost:3000/api/user/register"
		userresponse.RegisterFailedResponse(c, errorMessage, data, activationLink, http.StatusBadRequest)
		return
	}

	//hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		data := body
		activationLink := "http://localhost:3000/api/user/register"
		userresponse.RegisterFailedResponse(c, "Failed To Hash Password", data, activationLink, http.StatusBadRequest)
		return
	}

	//create the user
	newUUID := uuid.New()
	user := models.User{ID: newUUID, Name: body.Name, Address: body.Address, PhoneNumber: body.PhoneNumber, Email: body.Email, Password: string(hash)}

	if err := initializers.DB.Create(&user).Error; err != nil {
		data := body
		activationLink := "http://localhost:3000/api/user/register"
		userresponse.RegisterFailedResponse(c, "Email Already Use", data, activationLink, http.StatusConflict)
		return
	}

	rand.Seed(time.Now().UnixNano())
	otpCode := fmt.Sprintf("%04d", rand.Intn(10000))

	user.OtpCode = otpCode
	//2 minute otp code expired
	user.OtpCreatedAt = time.Now().Add(2 * time.Minute)

	if err := initializers.DB.Save(&user).Error; err != nil {
		data := body
		activationLink := "http://localhost:3000/api/user/refresh_code/" + newUUID.String()
		userresponse.RegisterFailedResponse(c, "Failed To Update Otp Code", data, activationLink, http.StatusInternalServerError)
		return
	}

	SendEmailWithGmail(body.Email, otpCode)

	//respond
	data := body
	activationLink := "http://localhost:3000/api/user/activate/" + newUUID.String()
	userresponse.RegisterSuccessResponse(c, "Register Success", data, activationLink)
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

	if err := initializers.DB.Save(&user).Error; err != nil {
		activationLink := "http://localhost:3000/api/user/register"
		otpresponse.FailedResponse(c, "Failed to Update Otp Code", data, activationLink, http.StatusInternalServerError)
		return
	}

	SendEmailWithGmail(user.Email, otpCode)

	activationLink := "http://localhost:3000/api/user/activate/" + userID
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
		return "Password Must Be At Least 6 Letters"
	default:
		return fmt.Sprintf("validation Failed for Tag: %s", tag)
	}
}
