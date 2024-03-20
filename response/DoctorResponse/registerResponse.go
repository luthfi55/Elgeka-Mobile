package doctorresponse

import (
	"net/http"

	"elgeka-mobile/models"

	"github.com/gin-gonic/gin"
)

func RegisterFailedResponse(c *gin.Context, message string, data models.Doctor, link string, status int) {
	linkItem := models.LinkItem{
		Name: "Register",
		Link: link,
	}

	response := models.RegisterDoctorFailedResponse{
		ErrorMessage: message,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func RegisterSuccessResponse(c *gin.Context, message string, data models.Doctor, link string) {
	datas := models.RegisterData{
		ID:    data.ID,
		Email: data.Email,
	}

	linkItem := models.LinkItem{
		Name: "Activate Account",
		Link: link,
	}

	response := models.RegisterDoctorSuccessResponse{
		Message: message,
		Data:    []models.RegisterData{datas},
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(http.StatusCreated, response)
}
