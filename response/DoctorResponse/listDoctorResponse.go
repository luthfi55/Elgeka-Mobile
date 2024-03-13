package doctorresponse

import (
	"elgeka-mobile/models"

	"github.com/gin-gonic/gin"
)

func GetInactiveDoctorFailedResponse(c *gin.Context, message string, data []models.DoctorData, link string, status int) {
	linkItem := models.LinkItem{
		Name: "Activate Doctor",
		Link: link,
	}

	response := models.RegisterDoctorFailedResponse{
		ErrorMessage: message,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func GetInactiveDoctorSuccessResponse(c *gin.Context, message string, data []models.DoctorData, link string, status int) {
	linkItem := models.LinkItem{
		Name: "Activate Doctor",
		Link: link,
	}

	response := models.GetListDoctorSuccessResponse{
		Message: message,
		Data:    data,
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}
