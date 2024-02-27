package controllers

import (
	"elgeka-mobile/initializers"
	"elgeka-mobile/models"
	doctorresponse "elgeka-mobile/response/DoctorResponse"
	userresponse "elgeka-mobile/response/UserResponse"
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func RegisterController(r *gin.Engine) {
	//user
	r.POST("api/user/register", UserRegister)
	//doctor
	r.POST("api/doctor/register", DoctorRegister)
}

func isPasswordValid(password string) bool {
	if len(password) < 8 {
		return false
	}
	// Use separate checks for uppercase letter, digit, and special character
	hasUppercase, _ := regexp.MatchString(`[A-Z]`, password)
	hasDigit, _ := regexp.MatchString(`\d`, password)
	hasSpecialChar, _ := regexp.MatchString(`[^\w\d\s]`, password)

	// Ensure that the password contains at least 1 uppercase letter, 1 digit, and 1 special character
	return hasUppercase && hasDigit && hasSpecialChar
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

	if !isEmailUnique(body.Email) {
		data := body
		activationLink := "http://localhost:3000/api/user/register"
		userresponse.RegisterFailedResponse(c, "Email Already Use", data, activationLink, http.StatusBadRequest)
		return
	}

	if !isPasswordValid(body.Password) {
		errorMessage := "Password must contain at least 1 uppercase letter, 1 digit, and 1 symbol."
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
		activationLink := "http://localhost:3000/api/user/register"
		userresponse.RegisterFailedResponse(c, "Email Already Use", user, activationLink, http.StatusConflict)
		return
	}

	rand.Seed(time.Now().UnixNano())
	otpCode := fmt.Sprintf("%04d", rand.Intn(10000))

	user.OtpCode = otpCode
	//2 minute otp code expired
	user.OtpCreatedAt = time.Now().Add(time.Minute)
	user.OtpType = "Activation"

	if err := initializers.DB.Save(&user).Error; err != nil {
		activationLink := "http://localhost:3000/api/user/refresh_code/" + newUUID.String()
		userresponse.RegisterFailedResponse(c, "Failed To Update Otp Code", user, activationLink, http.StatusInternalServerError)
		return
	}

	SendEmailWithGmail(body.Email, otpCode)

	//respond
	activationLink := "http://localhost:3000/api/user/activate/" + newUUID.String()
	userresponse.RegisterSuccessResponse(c, "Register Success", user, activationLink)
}

func DoctorRegister(c *gin.Context) {
	var body models.Doctor

	if c.Bind(&body) != nil {
		data := body
		activationLink := "http://localhost:3000/api/doctor/register"
		doctorresponse.RegisterFailedResponse(c, "Failed To Read Body", data, activationLink, http.StatusBadRequest)

		return
	}

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("%s %s", err.Field(), getValidationErrorTagMessage(err.Tag())))
		}
		errorMessage := strings.Join(validationErrors, ", ")
		data := body
		activationLink := "http://localhost:3000/api/doctor/register"
		doctorresponse.RegisterFailedResponse(c, errorMessage, data, activationLink, http.StatusBadRequest)
		return
	}

	if !isEmailUnique(body.Email) {
		data := body
		activationLink := "http://localhost:3000/api/doctor/register"
		doctorresponse.RegisterFailedResponse(c, "Email Already Use", data, activationLink, http.StatusBadRequest)
		return
	}

	if !isPasswordValid(body.Password) {
		errorMessage := "Password must contain at least 1 uppercase letter, 1 digit, and 1 symbol."
		data := body
		activationLink := "http://localhost:3000/api/user/register"
		doctorresponse.RegisterFailedResponse(c, errorMessage, data, activationLink, http.StatusBadRequest)
		return
	}

	//hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		data := body
		activationLink := "http://localhost:3000/api/doctor/register"
		doctorresponse.RegisterFailedResponse(c, "Failed To Hash Password", data, activationLink, http.StatusBadRequest)
		return
	}

	//create the doctor
	newUUID := uuid.New()
	doctor := models.Doctor{ID: newUUID, Name: body.Name, PolyName: body.PolyName, HospitalName: body.HospitalName, Email: body.Email, Password: string(hash)}

	if err := initializers.DB.Create(&doctor).Error; err != nil {
		activationLink := "http://localhost:3000/api/doctor/register"
		doctorresponse.RegisterFailedResponse(c, err.Error(), doctor, activationLink, http.StatusConflict)
		return
	}
	rand.Seed(time.Now().UnixNano())
	otpCode := fmt.Sprintf("%04d", rand.Intn(10000))
	doctor.OtpCode = otpCode
	doctor.OtpCreatedAt = time.Now().Add(time.Minute)
	doctor.OtpType = "Activation"

	if err := initializers.DB.Save(&doctor).Error; err != nil {
		activationLink := "http://localhost:3000/api/doctor/refresh_code/" + newUUID.String()
		doctorresponse.RegisterFailedResponse(c, "Failed To Update Otp Code", doctor, activationLink, http.StatusInternalServerError)
		return
	}

	SendEmailWithGmail(doctor.Email, otpCode)

	//respond
	activationLink := "http://localhost:3000/api/doctor/activate/" + newUUID.String()
	doctorresponse.RegisterSuccessResponse(c, "Register Success", doctor, activationLink)
}
