package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"elgeka-mobile/initializers"
	"elgeka-mobile/middleware"
	"elgeka-mobile/models"

	otpresponse "elgeka-mobile/response/OtpResponse"
	loginresponse "elgeka-mobile/response/UserResponse"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func LoginController(r *gin.Engine) {
	r.POST("api/user/login", UserLogin)
	r.POST("api/user/forgot_password", ForgotPassword)
	r.GET("api/user/validate", middleware.RequireAuth, Validate)
	r.GET("api/doctor/validate", middleware.RequireAuth, ValidateDoctor)
}

func UserLogin(c *gin.Context) {
	//get the email and pass of req body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	//look up requested user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == uuid.Nil {
		var doctor models.Doctor
		initializers.DB.First(&doctor, "email = ?", body.Email)

		if doctor.ID == uuid.Nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid email or password",
			})
			return
		}
		err := bcrypt.CompareHashAndPassword([]byte(doctor.Password), []byte(body.Password))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid email or password",
			})
			return
		}
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"error": "Login Doctor Success",
		// })
		//generate a jwt token
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

			return
		}

		//send it back
		c.SetSameSite(http.SameSiteLaxMode)
		//expire set with second
		c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

		activationLink := "http://localhost:3000"
		loginresponse.LoginSuccessResponse(c, "Login Doctor Success", body.Email, activationLink, http.StatusOK)
		return
	}

	//compare sent in pass with saved user pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	if !user.IsActive {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Account not Active",
		})

		return
	}

	//generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		// "exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid to create token",
		})

		return
	}

	//send it back
	c.SetSameSite(http.SameSiteLaxMode)
	//expire set with second
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})

	activationLink := "http://localhost:3000"
	loginresponse.LoginSuccessResponse(c, "Login Success", body.Email, activationLink, http.StatusOK)
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
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

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
		doctor.OtpCreatedAt = time.Now().Add(time.Minute)
		doctor.OtpType = "ForgotPassword"

		if err := initializers.DB.Save(&doctor).Error; err != nil {
			activationLink := "http://localhost:3000/api/doctor/register"
			otpresponse.FailedResponse(c, "Failed to Send Otp Code", doctor.ID.String(), activationLink, http.StatusInternalServerError)
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
	user.OtpCreatedAt = time.Now().Add(time.Minute)
	user.OtpType = "ForgotPassword"

	if err := initializers.DB.Save(&user).Error; err != nil {
		activationLink := "http://localhost:3000/api/user/register"
		otpresponse.FailedResponse(c, "Failed to Send Otp Code", user.ID.String(), activationLink, http.StatusInternalServerError)
		return
	}

	SendEmailForgotPassword(body.Email, otpCode)

	data.ID = user.ID.String()
	data.Email = user.Email
	data.OtpCode = otpCode

	activationLink := "http://localhost:3000"
	otpresponse.ForgotPasswordUserSuccessResponse(c, "Success to Send Otp Code", user, activationLink, http.StatusOK)
	return

}
