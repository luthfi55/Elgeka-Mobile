package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"elgeka-mobile/initializers"
	"elgeka-mobile/models"
	userresponse "elgeka-mobile/response/UserResponse"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		userresponse.CheckAccountFailedResponse(c)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		userresponse.CheckAccountFailedResponse(c)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//check te exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token EXP",
			})
			// c.AbortWithStatus(http.StatusUnauthorized)
		}
		//find the user with token sub
		var user models.User
		var doctor models.Doctor
		sub, subIsString := claims["sub"].(string)

		if !subIsString {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		initializers.DB.Where("id = ?", sub).First(&user)

		if user.ID == uuid.Nil {
			initializers.DB.Where("id = ?", sub).First(&doctor)
			//attach to req
			c.Set("doctor", doctor.ID)

			//continue
			c.Next()
		}

		//attach to req
		c.Set("user", user.ID)

		//continue
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
