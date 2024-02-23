package controllers

import (
	"net/http"
	"os"
	"time"

	"elgeka-mobile/initializers"
	"elgeka-mobile/middleware"
	"elgeka-mobile/models"

	loginresponse "elgeka-mobile/response/UserResponse"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func LoginController(r *gin.Engine) {
	r.POST("api/user/login", UserLogin)
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

		c.JSON(http.StatusOK, gin.H{})
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
