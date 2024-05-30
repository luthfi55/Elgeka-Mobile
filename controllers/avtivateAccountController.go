package controllers

import (
	"elgeka-mobile/initializers"
	"elgeka-mobile/models"
	doctorresponse "elgeka-mobile/response/DoctorResponse"
	otpresponse "elgeka-mobile/response/OtpResponse"
	userresponse "elgeka-mobile/response/UserResponse"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func ActivateAccountController(r *gin.Engine) {
	//user
	r.POST("api/user/email_otp/:user_id", SendEmailOtp)
	r.POST("api/user/whatsapp_otp/:user_id", SendWhatsappOtp)
	r.POST("api/user/activate/:user_id", Activate)
	r.POST("api/user/email_refresh_code/:user_id", RefreshOtpCode)
	r.POST("api/user/whatsapp_refresh_code/:user_id", RefreshWhatsappOtpCode)

	//doctor activate otp
	r.POST("api/doctor/email_otp/:doctor_id", SendDoctorEmailOtp)
	r.POST("api/doctor/whatsapp_otp/:doctor_id", SendDoctorWhatsappOtp)
	r.POST("api/doctor/activate_otp/:doctor_id", ActivateOtpDoctor)
	r.POST("api/doctor/email_refresh_code/:doctor_id", RefreshDoctorEmailOtpCode)
	r.POST("api/doctor/whatsapp_refresh_code/:doctor_id", RefreshDoctorWhatsappOtpCode)

	//doctor activate account admin website
	r.POST("api/doctor/activate_account/:doctor_id", ActivateDoctor)
	r.POST("api/doctor/reject_activation/:doctor_id", RejectDoctor)
	r.POST("api/doctor/refresh_code/:doctor_id", RefreshDoctorOtpCode)
	r.GET("api/doctor/list_inactive", ListInactiveDoctor)
}

func SendEmailOtp(c *gin.Context) {
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
	user.OtpCreatedAt = time.Now().Add(3 * time.Minute)
	user.OtpType = "Activation"

	if err := initializers.DB.Save(&user).Error; err != nil {
		activationLink := "http://localhost:3000/api/user/register"
		otpresponse.FailedResponse(c, "Failed to Update Otp Code", data, activationLink, http.StatusInternalServerError)
		return
	}

	SendEmailWithGmail(user.Email, otpCode)

	activationLink := "http://localhost:3000/api/user/activate/" + userID
	otpresponse.SuccessResponse(c, "Send Email OTP Successfully", user.Email, activationLink, http.StatusOK)
}

func SendWhatsappOtp(c *gin.Context) {
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
	user.OtpCreatedAt = time.Now().Add(3 * time.Minute)
	user.OtpType = "Activation"

	if err := initializers.DB.Save(&user).Error; err != nil {
		activationLink := "http://localhost:3000/api/user/register"
		otpresponse.FailedResponse(c, "Failed to Update Otp Code", data, activationLink, http.StatusInternalServerError)
		return
	}

	initializers.SendMessageToUser(user.PhoneNumber, otpCode)

	activationLink := "http://localhost:3000/api/user/activate/" + userID
	otpresponse.SuccessResponse(c, "Send Whatsapp OTP Successfully", user.Email, activationLink, http.StatusOK)
}

func RefreshWhatsappOtpCode(c *gin.Context) {
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
	user.OtpCreatedAt = time.Now().Add(3 * time.Minute)
	user.OtpType = "Activation"

	if err := initializers.DB.Save(&user).Error; err != nil {
		activationLink := "http://localhost:3000/api/user/register"
		otpresponse.FailedResponse(c, "Failed to Update Otp Code", data, activationLink, http.StatusInternalServerError)
		return
	}

	initializers.SendMessageToUser(user.PhoneNumber, otpCode)

	activationLink := "http://localhost:3000/api/user/activate/" + userID
	otpresponse.SuccessResponse(c, "Refresh Whatsapp OTP Successfully", data, activationLink, http.StatusOK)
}

func Activate(c *gin.Context) {
	userID := c.Param("user_id")

	var body struct {
		OtpCode string `json:"OtpCode"`
	}

	if c.Bind(&body) != nil {
		activationLink := "http://localhost:3000/api/user/register"
		otpresponse.FailedResponse(c, "Failed To Read Body", "", activationLink, http.StatusBadRequest)

		return
	}

	var user models.User

	result := initializers.DB.First(&user, "id = ?", userID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			activationLink := "http://localhost:3000/api/user/register"
			otpresponse.FailedResponse(c, "User Not Found", "", activationLink, http.StatusNotFound)

			return
		} else {
			activationLink := "http://localhost:3000/api/user/register"
			otpresponse.FailedResponse(c, "Database Error", "", activationLink, http.StatusInternalServerError)

			return
		}
	}

	if user.IsActive {
		activationLink := "http://localhost:3000/api/user/register"
		otpresponse.FailedResponse(c, "User Account Already Active", user.Email, activationLink, http.StatusUnauthorized)

		return
	}

	if user.OtpCode != body.OtpCode {
		activationLink := "http://localhost:3000/api/user/register"
		otpresponse.FailedResponse(c, "Incorrect OTP code", user.Email, activationLink, http.StatusUnauthorized)
		return
	} else {
		if time.Now().After(user.OtpCreatedAt) || time.Now().Equal(user.OtpCreatedAt) {
			activationLink := "http://localhost:3000/api/doctor/register"
			otpresponse.FailedResponse(c, "OTP Code Expired", user.Email, activationLink, http.StatusUnauthorized)
			return
		}

		if user.OtpType != "Activation" {
			activationLink := "http://localhost:3000/api/user/register"
			otpresponse.FailedResponse(c, "Incorrect OTP code", user.Email, activationLink, http.StatusUnauthorized)
			return
		}

		user.IsActive = true
		// Save the updated user
		if err := initializers.DB.Save(&user).Error; err != nil {
			activationLink := "http://localhost:3000/api/user/register"
			otpresponse.FailedResponse(c, "Failed to Activate", user.Email, activationLink, http.StatusInternalServerError)
			return
		}

		var treatment models.UserTreatment
		newUUID := uuid.New()
		treatment.ID = newUUID
		treatment.UserID = user.ID
		if err := initializers.DB.Create(&treatment).Error; err != nil {
			userresponse.GetTreatmentDataFailedResponse(c, "Failed To Create Treatment Data", "", http.StatusInternalServerError)
			return
		}

		var user_information models.UserInformation
		user_newUUID := uuid.New()
		user_information.ID = user_newUUID
		user_information.UserID = user.ID
		if err := initializers.DB.Create(&user_information).Error; err != nil {
			userresponse.GetTreatmentDataFailedResponse(c, "Failed To Create User Information Data", "", http.StatusInternalServerError)
			return
		}

		activationLink := "http://localhost:3000/api/user/login"
		otpresponse.SuccessResponse(c, "User Activated Successfully", user.Email, activationLink, http.StatusOK)
		return
	}
}

func ActivateDoctor(c *gin.Context) {
	if !ParseWebToken(c) {
		return
	}
	doctorID := c.Param("doctor_id")

	var doctor models.Doctor

	// Find the doctor by ID
	result := initializers.DB.First(&doctor, "id = ?", doctorID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			activationLink := "http://localhost:3000/api/user/register"
			otpresponse.FailedResponse(c, "Doctor Not Found", "", activationLink, http.StatusNotFound)

			return
		} else {
			activationLink := "http://localhost:3000/api/user/register"
			otpresponse.FailedResponse(c, "Database Error", "", activationLink, http.StatusInternalServerError)

			return
		}
	}

	if doctor.IsActive {
		activationLink := "http://localhost:3000/api/user/register"
		otpresponse.FailedResponse(c, "Doctor Account Already Active", doctor.Email, activationLink, http.StatusUnauthorized)

		return
	}

	if !doctor.EmailActive {
		activationLink := "http://localhost:3000/api/user/register"
		otpresponse.FailedResponse(c, "Doctor Email Account Must Active", doctor.Email, activationLink, http.StatusUnauthorized)

		return
	}

	doctor.IsActive = true
	// Save the updated user
	if err := initializers.DB.Save(&doctor).Error; err != nil {
		activationLink := "http://localhost:3000/api/user/register"
		otpresponse.FailedResponse(c, "Failed to Activate", doctor.Email, activationLink, http.StatusInternalServerError)
		return
	}

	activationLink := "http://localhost:3000/api/user/login"
	otpresponse.SuccessResponse(c, "Doctor Activated Successfully", doctor.Email, activationLink, http.StatusOK)
}

func ParseWebToken(c *gin.Context) bool {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		activationLink := "http://localhost:3000/api"
		otpresponse.FailedResponse(c, "Token is required", "", activationLink, http.StatusBadRequest)
		// c.JSON(400, gin.H{"error": "Token is required"})
		return false
	}

	tokenString = tokenString[7:] // Remove "Bearer " prefix

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("kFJ9CPC7av3X7VuddYR3AF7"), nil
	})
	if err != nil {
		// c.JSON(401, gin.H{"error": "Failed to parse token"})
		activationLink := "http://localhost:3000/api"
		otpresponse.FailedResponse(c, "Failed to parse token", "", activationLink, http.StatusUnauthorized)
		return false
	}
	if token.Valid {
		return true
	} else {
		return false
	}
}

func SendDoctorEmailOtp(c *gin.Context) {
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
	doctor.OtpCreatedAt = time.Now().Add(3 * time.Minute)
	doctor.OtpType = "Activation"

	if err := initializers.DB.Save(&doctor).Error; err != nil {
		activationLink := "http://localhost:3000/api/doctor/register"
		otpresponse.FailedResponse(c, "Failed to Update Otp Code", data, activationLink, http.StatusInternalServerError)
		return
	}

	SendEmailWithGmail(doctor.Email, otpCode)

	activationLink := "http://localhost:3000/api/doctor/activate/" + doctorID
	otpresponse.SuccessResponse(c, "Send Email OTP Successfully", doctor.Email, activationLink, http.StatusOK)
}

func SendDoctorWhatsappOtp(c *gin.Context) {
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
	doctor.OtpCreatedAt = time.Now().Add(3 * time.Minute)
	doctor.OtpType = "Activation"

	if err := initializers.DB.Save(&doctor).Error; err != nil {
		activationLink := "http://localhost:3000/api/doctor/register"
		otpresponse.FailedResponse(c, "Failed to Update Otp Code", data, activationLink, http.StatusInternalServerError)
		return
	}

	initializers.SendMessageToUser(doctor.PhoneNumber, otpCode)

	activationLink := "http://localhost:3000/api/doctor/activate/" + doctorID
	otpresponse.SuccessResponse(c, "Send Whatsapp OTP Successfully", doctor.Email, activationLink, http.StatusOK)
}

func ActivateOtpDoctor(c *gin.Context) {
	doctorID := c.Param("doctor_id")

	var body struct {
		OtpCode string `json:"OtpCode"`
	}

	if c.Bind(&body) != nil {
		activationLink := "http://localhost:3000/api/doctor/register"
		otpresponse.FailedResponse(c, "Failed To Read Body", "", activationLink, http.StatusBadRequest)

		return
	}

	var doctor models.Doctor

	// Find the doctor by ID
	result := initializers.DB.First(&doctor, "id = ?", doctorID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			activationLink := "http://localhost:3000/api/doctor/register"
			otpresponse.FailedResponse(c, "Doctor Not Found", "", activationLink, http.StatusNotFound)

			return
		} else {
			activationLink := "http://localhost:3000/api/doctor/register"
			otpresponse.FailedResponse(c, "Database Error", "", activationLink, http.StatusInternalServerError)

			return
		}
	}

	if doctor.EmailActive {
		activationLink := "http://localhost:3000/api/doctor/register"
		otpresponse.FailedResponse(c, "Doctor Account Already Active", doctor.Email, activationLink, http.StatusUnauthorized)

		return
	}

	if doctor.OtpCode != body.OtpCode {
		activationLink := "http://localhost:3000/api/doctor/register"
		otpresponse.FailedResponse(c, "Incorrect OTP code", doctor.Email, activationLink, http.StatusUnauthorized)
		return
	} else {
		// 1 minute expired
		if time.Now().After(doctor.OtpCreatedAt) || time.Now().Equal(doctor.OtpCreatedAt) {
			activationLink := "http://localhost:3000/api/doctor/register"
			otpresponse.FailedResponse(c, "OTP Code Expired", doctor.Email, activationLink, http.StatusUnauthorized)
			return
		}

		if doctor.OtpType != "Activation" {
			activationLink := "http://localhost:3000/api/doctor/register"
			otpresponse.FailedResponse(c, "Incorrect OTP code", doctor.Email, activationLink, http.StatusUnauthorized)
			return
		}

		doctor.EmailActive = true

		if err := initializers.DB.Save(&doctor).Error; err != nil {
			activationLink := "http://localhost:3000/api/doctor/register"
			otpresponse.FailedResponse(c, "Failed to Activate", doctor.Email, activationLink, http.StatusInternalServerError)
			return
		}

		activationLink := "http://localhost:3000/api/doctor/login"
		otpresponse.SuccessResponse(c, "Doctor Otp Activated Successfully", doctor.Email, activationLink, http.StatusOK)
		return
	}
}

func RefreshDoctorEmailOtpCode(c *gin.Context) {
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
	doctor.OtpCreatedAt = time.Now().Add(3 * time.Minute)
	doctor.OtpType = "Activation"

	if err := initializers.DB.Save(&doctor).Error; err != nil {
		activationLink := "http://localhost:3000/api/doctor/register"
		otpresponse.FailedResponse(c, "Failed to Update Otp Code", data, activationLink, http.StatusInternalServerError)
		return
	}

	SendEmailWithGmail(doctor.Email, otpCode)

	activationLink := "http://localhost:3000/api/doctor/activate/" + doctorID
	otpresponse.SuccessResponse(c, "Refresh Email OTP Successfully", doctor.Email, activationLink, http.StatusOK)
}

func RefreshDoctorWhatsappOtpCode(c *gin.Context) {
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
	doctor.OtpCreatedAt = time.Now().Add(3 * time.Minute)
	doctor.OtpType = "Activation"

	if err := initializers.DB.Save(&doctor).Error; err != nil {
		activationLink := "http://localhost:3000/api/doctor/register"
		otpresponse.FailedResponse(c, "Failed to Update Otp Code", data, activationLink, http.StatusInternalServerError)
		return
	}

	initializers.SendMessageToUser(doctor.PhoneNumber, otpCode)

	activationLink := "http://localhost:3000/api/doctor/activate/" + doctorID
	otpresponse.SuccessResponse(c, "Refresh Whatsapp OTP Successfully", doctor.Email, activationLink, http.StatusOK)
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
	user.OtpCreatedAt = time.Now().Add(3 * time.Minute)
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
	doctor.OtpCreatedAt = time.Now().Add(3 * time.Minute)
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

func ListInactiveDoctor(c *gin.Context) {
	if !ParseWebToken(c) {
		return
	}

	var doctors []models.Doctor
	var response []models.DoctorProfile
	result := initializers.DB.Where("is_active = ?", false).Where("email_active = ?", true).Find(&doctors)

	if result.Error != nil {
		activationLink := "http://localhost:3000/api/user/activate/:user_id"
		doctorresponse.GetInactiveDoctorFailedResponse(c, "Failed To Retrieve Data", response, activationLink, http.StatusInternalServerError)
		return
	}

	for _, doctor := range doctors {
		response = append(response, models.DoctorProfile{
			ID:           doctor.ID,
			Name:         doctor.Name,
			PhoneNumber:  doctor.PhoneNumber,
			Email:        doctor.Email,
			Gender:       doctor.Gender,
			PolyName:     doctor.PolyName,
			HospitalName: doctor.HospitalName,
		})
	}

	if len(doctors) == 0 {
		activationLink := "http://localhost:3000/api/user/activate/:user_id"
		doctorresponse.GetInactiveDoctorFailedResponse(c, "Data Empty", response, activationLink, http.StatusInternalServerError)
		return
	}

	activationLink := "http://localhost:3000/api/user/activate/:user_idd"
	doctorresponse.GetInactiveDoctorSuccessResponse(c, "Get Data Successfully", response, activationLink, http.StatusOK)
}

func RejectDoctor(c *gin.Context) {
	if !ParseWebToken(c) {
		return
	}
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
