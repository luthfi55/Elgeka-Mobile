package controllers

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"elgeka-mobile/initializers"
	"elgeka-mobile/middleware"
	"elgeka-mobile/models"
	"encoding/base64"

	otpresponse "elgeka-mobile/response/OtpResponse"
	userresponse "elgeka-mobile/response/UserResponse"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginController(r *gin.Engine) {
	r.POST("api/user/login", UserLogin)
	r.POST("api/user/login_website", UserLoginWebsite)
	r.POST("api/user/forgot_password", ForgotPassword)
	r.GET("api/user/validate", middleware.RequireAuth, Validate)
	r.GET("api/doctor/validate", middleware.RequireAuth, ValidateDoctor)
	r.POST("api/user/refresh_code/forgot_password/:user_id", RefreshForgotPasswordOtp)
	r.POST("api/user/check_otp/:user_id", CheckOtp)
	r.POST("api/user/change_password/:user_id/:otp_code", ChangePassword)
	r.POST("api/user/logout", UserLogoutWebsite)
}

func UserLogin(c *gin.Context) {
	var body struct {
		EmailOrPhoneNumber string
		Password           string
	}

	var UserId models.UserIdData

	if c.Bind(&body) != nil {
		userresponse.LoginFailedResponse(c, "Failed to read body", UserId, "http://localhost:3000/api/user/login", http.StatusBadRequest)
		return
	}

	var user models.User

	if body.EmailOrPhoneNumber[:2] == "62" {
		initializers.DB.First(&user, "phone_number = ?", body.EmailOrPhoneNumber)
	} else {
		initializers.DB.First(&user, "email = ?", body.EmailOrPhoneNumber)
	}

	if user.ID == uuid.Nil {
		var doctor models.Doctor

		if body.EmailOrPhoneNumber[:2] == "62" {
			initializers.DB.First(&doctor, "phone_number = ?", body.EmailOrPhoneNumber)
		} else {
			initializers.DB.First(&doctor, "email = ?", body.EmailOrPhoneNumber)
		}

		if doctor.ID == uuid.Nil {
			userresponse.LoginFailedResponse(c, "Invalid email or phone number", UserId, "http://localhost:3000/api/user/login", http.StatusBadRequest)
			return
		}
		err := bcrypt.CompareHashAndPassword([]byte(doctor.Password), []byte(body.Password))
		if err != nil {
			userresponse.LoginFailedResponse(c, "Invalid password", UserId, "http://localhost:3000/api/user/login", http.StatusBadRequest)
			return
		}

		//generate a jwt token
		UserId.ID = doctor.ID
		if !doctor.EmailActive {
			userresponse.LoginFailedResponse(c, "Email Account not Active", UserId, "http://localhost:3000/api/doctor/activate_email/"+doctor.ID.String(), http.StatusBadRequest)
			return
		}

		if !doctor.IsActive {
			userresponse.LoginFailedResponse(c, "Account not Active", UserId, "http://localhost:3000/api/doctor/login", http.StatusBadRequest)
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": doctor.ID,
			// "exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
			"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
		})

		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid to create token",
			})
			userresponse.LoginFailedResponse(c, "Invalid to create token", UserId, "http://localhost:3000/api/user/login", http.StatusBadRequest)
			return
		}

		var account models.Login

		account.Email = doctor.Email
		account.Name = doctor.Name

		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

		activationLink := "http://localhost:3000"
		userresponse.LoginSuccessResponse(c, "Login Doctor Success", account, activationLink, http.StatusOK)
		return
	}

	UserId.ID = user.ID

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		userresponse.LoginFailedResponse(c, "Invalid password", UserId, "http://localhost:3000/api/user/login", http.StatusBadRequest)
		return
	}

	if !user.IsActive {
		userresponse.LoginFailedResponse(c, "Account not Active", UserId, "http://localhost:3000/api/user/activate/"+user.ID.String(), http.StatusBadRequest)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		userresponse.LoginFailedResponse(c, "Invalid to create token", UserId, "http://localhost:3000/api/user/login", http.StatusBadRequest)
		return
	}

	var account models.Login

	account.Email = user.Email
	account.Name = user.Name

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	activationLink := "http://localhost:3000"
	userresponse.LoginSuccessResponse(c, "Login Success", account, activationLink, http.StatusOK)
}

func UserLoginWebsite(c *gin.Context) {
	var body struct {
		EmailOrPhoneNumber string
		Password           string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	var UserId models.UserIdData

	var user models.User
	if body.EmailOrPhoneNumber[:2] == "62" {
		initializers.DB.First(&user, "phone_number = ?", body.EmailOrPhoneNumber)
	} else {
		initializers.DB.First(&user, "email = ?", body.EmailOrPhoneNumber)
	}

	if user.ID == uuid.Nil {
		userresponse.LoginFailedResponse(c, "Invalid email or phone number", UserId, "http://localhost:3000/api/user/login", http.StatusBadRequest)
		return
	}

	UserId.ID = user.ID

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		userresponse.LoginFailedResponse(c, "Invalid password", UserId, "http://localhost:3000/api/user/login", http.StatusBadRequest)
		return
	}

	if !user.IsActive {
		userresponse.LoginFailedResponse(c, "Account not Active", UserId, "http://localhost:3000/api/user/activate/"+user.ID.String(), http.StatusBadRequest)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		userresponse.LoginFailedResponse(c, "Invalid to create token", UserId, "http://localhost:3000/api/user/login", http.StatusBadRequest)
		return
	}

	var account models.Login

	account.Email = user.Email
	account.Name = user.Name

	cookie := http.Cookie{
		Name:     "Authorization",
		Value:    tokenString,
		MaxAge:   3600 * 24 * 30,
		Path:     "/",
		Domain:   "",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}
	http.SetCookie(c.Writer, &cookie)

	activationLink := "http://localhost:3000"
	userresponse.LoginWebsiteSuccessResponse(c, "Login Success", account, activationLink, tokenString, http.StatusOK)
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func ValidateDoctor(c *gin.Context) {
	doctor, _ := c.Get("doctor")

	c.JSON(http.StatusOK, gin.H{
		"message": doctor,
	})
}

func ForgotPassword(c *gin.Context) {
	var body struct {
		Email string
	}

	var data struct {
		ID      string
		Email   string
		OtpCode string
	}

	if c.Bind(&body) != nil {

		activationLink := "http://localhost:3000/api/user/register"
		otpresponse.FailedResponse(c, "Failed to read body", body.Email, activationLink, http.StatusBadRequest)

		return
	}

	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == uuid.Nil {
		var doctor models.Doctor
		initializers.DB.First(&doctor, "email = ?", body.Email)

		if doctor.ID == uuid.Nil {
			activationLink := "http://localhost:3000/api/doctor/register"
			otpresponse.FailedResponse(c, "Email Not Found", body.Email, activationLink, http.StatusInternalServerError)
			return
		}

		rand.Seed(time.Now().UnixNano())
		otpCode := fmt.Sprintf("%04d", rand.Intn(10000))

		doctor.OtpCode = otpCode
		doctor.OtpCreatedAt = time.Now().Add(3 * time.Minute)
		doctor.OtpType = "ForgotPassword"

		if err := initializers.DB.Save(&doctor).Error; err != nil {
			activationLink := "http://localhost:3000/api/doctor/register"
			otpresponse.FailedResponse(c, "Failed to Send Otp Code", body.Email, activationLink, http.StatusInternalServerError)
			return
		}

		SendEmailForgotPassword(body.Email, otpCode)

		data.ID = doctor.ID.String()
		data.Email = doctor.Email
		data.OtpCode = otpCode

		activationLink := "http://localhost:3000"
		otpresponse.ForgotPasswordDoctorSuccessResponse(c, "Success to Send Otp Code", doctor, activationLink, http.StatusOK)
		return
	}

	rand.Seed(time.Now().UnixNano())
	otpCode := fmt.Sprintf("%04d", rand.Intn(10000))

	user.OtpCode = otpCode
	user.OtpCreatedAt = time.Now().Add(3 * time.Minute)
	user.OtpType = "ForgotPassword"

	if err := initializers.DB.Save(&user).Error; err != nil {
		activationLink := "http://localhost:3000/api/user/register"
		otpresponse.FailedResponse(c, "Failed to Send Otp Code", body.Email, activationLink, http.StatusInternalServerError)
		return
	}

	SendEmailForgotPassword(body.Email, otpCode)

	data.ID = user.ID.String()
	data.Email = user.Email
	data.OtpCode = otpCode

	activationLink := "http://localhost:3000/api/user/check_otp/" + data.ID
	otpresponse.ForgotPasswordUserSuccessResponse(c, "Success to Send Otp Code", user, activationLink, http.StatusOK)
}

func CheckOtp(c *gin.Context) {
	userID := c.Param("user_id")

	var data struct {
		ID      string
		Email   string
		OtpCode string
	}

	var body struct {
		OtpCode string
	}

	if c.Bind(&body) != nil {
		activationLink := "http://localhost:3000"
		otpresponse.FailedCheckOtpResponse(c, "Failed to read body", body.OtpCode, activationLink, http.StatusBadRequest)

		return
	}

	var user models.User
	var doctor models.Doctor

	result := initializers.DB.First(&user, "id = ?", userID)

	if result.Error != nil {
		result := initializers.DB.First(&doctor, "id = ?", userID)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				activationLink := "http://localhost:3000/api/user/forgot_password"
				otpresponse.FailedResponse(c, "User Not Found", user.Email, activationLink, http.StatusNotFound)
				return
			} else {
				activationLink := "http://localhost:3000/api/user/forgot_password"
				otpresponse.FailedResponse(c, "Database Error", user.Email, activationLink, http.StatusInternalServerError)
				return
			}
		}
		if !doctor.IsActive {
			activationLink := "http://localhost:3000/api/doctor/activate_account/" + userID
			otpresponse.FailedResponse(c, "Doctor Account Must Active", user.Email, activationLink, http.StatusUnauthorized)
			return
		}

		if doctor.OtpCode != body.OtpCode {
			activationLink := "http://localhost:3000/api/user/register"
			otpresponse.FailedResponse(c, "Incorrect OTP code", user.Email, activationLink, http.StatusUnauthorized)
			return
		} else {
			if time.Since(doctor.OtpCreatedAt) > time.Minute {
				activationLink := "http://localhost:3000/api/user/register"
				otpresponse.FailedResponse(c, "OTP Code Expired", user.Email, activationLink, http.StatusUnauthorized)
				return
			}

			if doctor.OtpType != "ForgotPassword" {
				activationLink := "http://localhost:3000/api/user/register"
				otpresponse.FailedResponse(c, "Incorrect OTP code", user.Email, activationLink, http.StatusUnauthorized)
				return
			}

			hashOtp, err := bcrypt.GenerateFromPassword([]byte(body.OtpCode), 10)
			if err != nil {
				activationLink := "http://localhost:3000/api/user/register"
				otpresponse.FailedResponse(c, "Failed To Hash Password", user.Email, activationLink, http.StatusBadRequest)
				return
			}
			encodedHash := base64.URLEncoding.EncodeToString(hashOtp)

			doctor.ForgotPasswordCode = encodedHash
			if err := initializers.DB.Save(&doctor).Error; err != nil {
				activationLink := "http://localhost:3000/api/user/register"
				otpresponse.FailedResponse(c, "Failed To Update Forgot Password Code", user.Email, activationLink, http.StatusBadRequest)
				return
			}

			data.Email = doctor.Email
			data.ID = userID
			data.OtpCode = encodedHash

			activationLink := "http://localhost:3000/api/user/change_password/" + userID + "/" + encodedHash
			otpresponse.SuccessCheckOtpResponse(c, "Check Otp Successfully", data, activationLink, http.StatusOK)
			return
		}
	}

	if !user.IsActive {
		activationLink := "http://localhost:3000/api/user/activate/" + userID
		otpresponse.FailedResponse(c, "User Email Account Must Active", user.Email, activationLink, http.StatusUnauthorized)
		return
	}

	if user.OtpCode != body.OtpCode {
		activationLink := "http://localhost:3000/api/user/register"
		otpresponse.FailedResponse(c, "Incorrect OTP code", user.Email, activationLink, http.StatusUnauthorized)
		return
	} else {
		if time.Since(user.OtpCreatedAt) > time.Minute {
			activationLink := "http://localhost:3000/api/user/register"
			otpresponse.FailedResponse(c, "OTP Code Expired", user.Email, activationLink, http.StatusUnauthorized)
			return
		}

		if user.OtpType != "ForgotPassword" {
			activationLink := "http://localhost:3000/api/user/register"
			otpresponse.FailedResponse(c, "Incorrect OTP code", user.Email, activationLink, http.StatusUnauthorized)
			return
		}

		hashOtp, err := bcrypt.GenerateFromPassword([]byte(body.OtpCode), 10)
		if err != nil {
			activationLink := "http://localhost:3000/api/user/register"
			otpresponse.FailedResponse(c, "Failed To Hash Password", user.Email, activationLink, http.StatusBadRequest)
			return
		}
		encodedHash := base64.URLEncoding.EncodeToString(hashOtp)

		user.ForgotPasswordCode = encodedHash
		if err := initializers.DB.Save(&user).Error; err != nil {
			activationLink := "http://localhost:3000/api/user/register"
			otpresponse.FailedResponse(c, "Failed To Update Forgot Password Code", user.Email, activationLink, http.StatusBadRequest)
			return
		}

		data.Email = user.Email
		data.ID = userID
		data.OtpCode = encodedHash

		activationLink := "http://localhost:3000/api/user/change_password/" + userID + "/" + encodedHash
		otpresponse.SuccessCheckOtpResponse(c, "Check Otp Successfully", data, activationLink, http.StatusOK)
		return
	}

}

func ChangePassword(c *gin.Context) {
	userID := c.Param("user_id")
	otpCode := c.Param("otp_code")

	var body struct {
		Password             string
		PasswordConfirmation string
	}

	if c.Bind(&body) != nil {
		activationLink := "http://localhost:3000"
		otpresponse.FailedCheckOtpResponse(c, "Failed to read body", otpCode, activationLink, http.StatusBadRequest)

		return
	}

	var user models.User
	var doctor models.Doctor

	result := initializers.DB.First(&user, "id = ?", userID)
	if result.Error != nil {
		result := initializers.DB.First(&doctor, "id = ?", userID)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				activationLink := "http://localhost:3000/api/user/forgot_password"
				otpresponse.FailedResponse(c, "User Not Found", userID, activationLink, http.StatusNotFound)
				return
			} else {
				activationLink := "http://localhost:3000/api/user/forgot_password"
				otpresponse.FailedResponse(c, "Database Error", userID, activationLink, http.StatusInternalServerError)
				return
			}
		}
		if doctor.ForgotPasswordCode != otpCode {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to match",
			})
			return
		}
		c.JSON(http.StatusAccepted, gin.H{
			"message": "Succes",
		})
		return
	}
	if user.ForgotPasswordCode != otpCode {
		activationLink := "http://localhost:3000/api/user/forgot_password"
		otpresponse.FailedResponse(c, "Incorect Otp Code", user.Email, activationLink, http.StatusBadRequest)
		return
	}

	if body.Password != body.PasswordConfirmation {
		activationLink := "http://localhost:3000/api/user/forgot_password"
		otpresponse.FailedResponse(c, "Password Confirmation Must Same as Password", user.Email, activationLink, http.StatusBadRequest)
		return
	}

	if !isPasswordValid(body.Password) {
		errorMessage := "Password must contain at least 8 character, 1 uppercase letter, 1 digit, and 1 symbol."
		activationLink := "http://localhost:3000/api/user/forgot_password"
		otpresponse.FailedResponse(c, errorMessage, user.Email, activationLink, http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		activationLink := "http://localhost:3000/api/user/forgot_password"
		otpresponse.FailedResponse(c, "Failed to Hash Password", user.Email, activationLink, http.StatusBadRequest)
		return
	}

	user.Password = string(hash)

	if err := initializers.DB.Save(&user).Error; err != nil {
		activationLink := "http://localhost:3000/api/user/register"
		otpresponse.FailedResponse(c, "Failed To Update Password", user.Email, activationLink, http.StatusBadRequest)
		return
	}

	activationLink := "http://localhost:3000/api/user/login"
	otpresponse.SuccessResponse(c, "Update Password Successfully", user.Email, activationLink, http.StatusOK)
}

func RefreshForgotPasswordOtp(c *gin.Context) {
	userID := c.Param("user_id")

	var data struct {
		ID      string
		Email   string
		OtpCode string
	}

	var user models.User
	var doctor models.Doctor

	result := initializers.DB.First(&user, "id = ?", userID)
	if result.Error != nil {
		result := initializers.DB.First(&doctor, "id = ?", userID)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				activationLink := "http://localhost:3000/api/user/forgot_password"
				otpresponse.FailedResponse(c, "User Not Found", userID, activationLink, http.StatusNotFound)
				return
			} else {
				activationLink := "http://localhost:3000/api/user/forgot_password"
				otpresponse.FailedResponse(c, "Database Error", userID, activationLink, http.StatusInternalServerError)
				return
			}
		}
		if !doctor.IsActive {
			activationLink := "http://localhost:3000/api/doctor/activate_account/" + userID
			otpresponse.FailedResponse(c, "Doctor Account Must Active", user.Email, activationLink, http.StatusUnauthorized)
			return
		}

		rand.Seed(time.Now().UnixNano())
		otpCode := fmt.Sprintf("%04d", rand.Intn(10000))

		doctor.OtpCode = otpCode
		doctor.OtpCreatedAt = time.Now().Add(3 * time.Minute)
		doctor.OtpType = "ForgotPassword"

		if err := initializers.DB.Save(&doctor).Error; err != nil {
			activationLink := "http://localhost:3000/api/doctor/register"
			otpresponse.FailedResponse(c, "Failed to Send Otp Code", doctor.Email, activationLink, http.StatusInternalServerError)
			return
		}

		SendEmailForgotPassword(doctor.Email, otpCode)

		data.ID = doctor.ID.String()
		data.Email = doctor.Email
		data.OtpCode = otpCode

		activationLink := "http://localhost:3000"
		otpresponse.ForgotPasswordDoctorSuccessResponse(c, "Success to Send Otp Code", doctor, activationLink, http.StatusOK)
		return
	}

	if !user.IsActive {
		activationLink := "http://localhost:3000/api/user/activate/" + userID
		otpresponse.FailedResponse(c, "User Email Account Must Active", user.Email, activationLink, http.StatusUnauthorized)
		return
	}

	rand.Seed(time.Now().UnixNano())
	otpCode := fmt.Sprintf("%04d", rand.Intn(10000))

	user.OtpCode = otpCode
	user.OtpCreatedAt = time.Now().Add(3 * time.Minute)
	user.OtpType = "ForgotPassword"

	if err := initializers.DB.Save(&user).Error; err != nil {
		activationLink := "http://localhost:3000/api/user/register"
		otpresponse.FailedResponse(c, "Failed to Send Otp Code", user.Email, activationLink, http.StatusInternalServerError)
		return
	}

	SendEmailForgotPassword(user.Email, otpCode)

	data.ID = user.ID.String()
	data.Email = user.Email
	data.OtpCode = otpCode

	activationLink := "http://localhost:3000/api/user/check_otp/" + data.ID
	otpresponse.ForgotPasswordUserSuccessResponse(c, "Success to Send Otp Code", user, activationLink, http.StatusOK)
}

func UserLogoutWebsite(c *gin.Context) {
	cookie, err := c.Request.Cookie("Authorization")
	if err != nil {
		userresponse.LogoutSuccessResponse(c, "Failed to Get Cookie", "http://localhost:3000/api/user/login", http.StatusOK)
		return
	}

	cookie.MaxAge = -1
	cookie.Path = "/"
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode

	http.SetCookie(c.Writer, cookie)

	userresponse.LogoutSuccessResponse(c, "Logout Successful", "http://localhost:3000/api/user/login", http.StatusOK)
}
