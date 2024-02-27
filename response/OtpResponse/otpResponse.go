package otpresponse

import (
	"elgeka-mobile/models"

	"github.com/gin-gonic/gin"
)

func FailedResponse(c *gin.Context, message string, data string, link string, status int) {
	otpData := models.OtpData{
		Email: data,
	}

	linkItem := models.LinkItem{
		Name: "Register",
		Link: link,
	}

	response := models.OtpFailledResponse{
		ErrorMessage: message,
		Data:         []models.OtpData{otpData},
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func SuccessResponse(c *gin.Context, message string, data string, link string, status int) {
	Data := models.OtpData{
		Email: data,
	}

	linkItem := models.LinkItem{
		Name: "Login",
		Link: link,
	}

	response := models.OtpSuccessResponse{
		Message: message,
		Data:    []models.OtpData{Data},
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}
