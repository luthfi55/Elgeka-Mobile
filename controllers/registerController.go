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

func getValidationErrorTagMessage(tag string) string {
	switch tag {
	case "required":
		return "Cant Be Empty"
	case "email":
		return "Must Be a Valid Email Address"
	case "min":
		return "Must Be At Least 8 Letters"
	case "max":
		return "Must Be At Most 14 Letters"
	case "eqfield":
		return "Must Match Password"
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

func isValidDateFormat(date string) bool {
	_, err := time.Parse("2006-01-02", date)
	return err == nil
}

func UserRegister(c *gin.Context) {
	var body models.User

	if c.Bind(&body) != nil {
		data := body
		activationLink := "http://localhost:3000/api/user/register"
		userresponse.RegisterFailedResponse(c, "Failed To Read Body", data, activationLink, http.StatusBadRequest)

		return
	}

	if !isEmailUnique(body.Email) {
		data := body
		activationLink := "http://localhost:3000/api/user/register"
		userresponse.RegisterFailedResponse(c, "Email Already Use", data, activationLink, http.StatusBadRequest)
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
		activationLink := "http://localhost:3000/api/user/register"
		userresponse.RegisterFailedResponse(c, errorMessage, data, activationLink, http.StatusBadRequest)
		return
	}

	if !isPasswordValid(body.Password) {
		errorMessage := "Password must contain at least 1 uppercase letter, 1 digit, and 1 symbol."
		data := body
		activationLink := "http://localhost:3000/api/user/register"
		userresponse.RegisterFailedResponse(c, errorMessage, data, activationLink, http.StatusBadRequest)
		return
	}

	if body.Gender != "male" && body.Gender != "female" {
		errorMessage := "Gender must male or female."
		data := body
		activationLink := "http://localhost:3000/api/user/register"
		userresponse.RegisterFailedResponse(c, errorMessage, data, activationLink, http.StatusBadRequest)
		return
	}

	if body.BirthDate == "" || !isValidDateFormat(body.BirthDate) {
		errorMessage := "Birthdate must be in the format 'Year-Month-Day'."
		data := body
		activationLink := "http://localhost:3000/api/user/register"
		userresponse.RegisterFailedResponse(c, errorMessage, data, activationLink, http.StatusBadRequest)
		return
	}

	if body.BloodGroup != "A" && body.BloodGroup != "B" && body.BloodGroup != "AB" && body.BloodGroup != "O" {
		errorMessage := "Blood Group must A, B, AB, or O."
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
	user := models.User{ID: newUUID, Name: body.Name, Address: body.Address, Gender: body.Gender, BirthDate: body.BirthDate, BloodGroup: body.BloodGroup, PhoneNumber: body.PhoneNumber, Email: body.Email, Password: string(hash)}

	if err := initializers.DB.Create(&user).Error; err != nil {
		activationLink := "http://localhost:3000/api/user/register"
		userresponse.RegisterFailedResponse(c, "Email Already Use", user, activationLink, http.StatusConflict)
		return
	}

	// rand.Seed(time.Now().UnixNano())
	// otpCode := fmt.Sprintf("%04d", rand.Intn(10000))

	// user.OtpCode = otpCode
	// //2 minute otp code expired
	// user.OtpCreatedAt = time.Now().Add(3 * time.Minute)
	// user.OtpType = "Activation"

	// if err := initializers.DB.Save(&user).Error; err != nil {
	// 	activationLink := "http://localhost:3000/api/user/refresh_code/" + newUUID.String()
	// 	userresponse.RegisterFailedResponse(c, "Failed To Update Otp Code", user, activationLink, http.StatusInternalServerError)
	// 	return
	// }

	// SendEmailWithGmail(body.Email, otpCode)

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

	if body.Gender != "male" && body.Gender != "female" {
		errorMessage := "Gender must male or female."
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
	doctor := models.Doctor{ID: newUUID, Name: body.Name, PhoneNumber: body.PhoneNumber, Gender: body.Gender, PolyName: body.PolyName, HospitalName: body.HospitalName, Email: body.Email, Password: string(hash)}

	if err := initializers.DB.Create(&doctor).Error; err != nil {
		activationLink := "http://localhost:3000/api/doctor/register"
		doctorresponse.RegisterFailedResponse(c, err.Error(), doctor, activationLink, http.StatusConflict)
		return
	}
	rand.Seed(time.Now().UnixNano())
	otpCode := fmt.Sprintf("%04d", rand.Intn(10000))
	doctor.OtpCode = otpCode
	doctor.OtpCreatedAt = time.Now().Add(3 * time.Minute)
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
