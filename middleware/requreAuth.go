package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"elgeka-mobile/initializers"
	"elgeka-mobile/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func RequireAuth(c *gin.Context) {
	//get the cookie of req
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Need Login",
		})
	}

	//decode/validate it

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

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
		sub, subIsString := claims["sub"].(string)

		if !subIsString {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		initializers.DB.Where("id = ?", sub).First(&user)

		if user.ID == uuid.Nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//attach to req
		c.Set("user", user.ID)

		//continue
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
