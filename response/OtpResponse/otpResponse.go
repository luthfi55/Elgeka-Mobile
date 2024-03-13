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

	response := models.OtpFailedResponse{
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

func FailedCheckOtpResponse(c *gin.Context, message string, data string, link string, status int) {
	otpData := models.CheckOtpData{
		OtpCode: data,
	}

	linkItem := models.LinkItem{
		Name: "Forgot Password",
		Link: link,
	}

	response := models.CheckOtpFailedResponse{
		ErrorMessage: message,
		Data:         []models.CheckOtpData{otpData},
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func SuccessCheckOtpResponse(c *gin.Context, message string, data models.CheckSuccessOtpData, link string, status int) {
	Data := models.CheckSuccessOtpData{
		ID:      data.ID,
		Email:   data.Email,
		OtpCode: data.OtpCode,
	}

	linkItem := models.LinkItem{
		Name: "Login",
		Link: link,
	}

	response := models.CheckOtpSuccessResponse{
		Message: message,
		Data:    []models.CheckSuccessOtpData{Data},
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}
