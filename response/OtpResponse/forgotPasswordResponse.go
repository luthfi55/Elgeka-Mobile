package otpresponse

import (
	"elgeka-mobile/models"

	"github.com/gin-gonic/gin"
)

func ForgotPasswordUserSuccessResponse(c *gin.Context, message string, data models.User, link string, status int) {
	datas := models.Data{
		ID:      data.ID,
		Email:   data.Email,
		OtpCode: data.OtpCode,
	}

	linkItem := models.LinkItem{
		Name: "Check Otp",
		Link: link,
	}

	response := models.ForgotPasswordUserSuccess{
		Message: message,
		Data:    []models.Data{datas},
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func ForgotPasswordDoctorSuccessResponse(c *gin.Context, message string, data models.Doctor, link string, status int) {
	datas := models.Data{
		ID:      data.ID,
		Email:   data.Email,
		OtpCode: data.OtpCode,
	}

	linkItem := models.LinkItem{
		Name: "Activate Account",
		Link: link,
	}

	response := models.ForgotPasswordUserSuccess{
		Message: message,
		Data:    []models.Data{datas},
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}
