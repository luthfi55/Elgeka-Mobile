package doctorresponse

import (
	"elgeka-mobile/models"

	"github.com/gin-gonic/gin"
)

func GenderChartFailedResponse(c *gin.Context, message string, data string, status int) {
	linkItem := models.LinkItem{
		Name: "Patient Data by Gender",
		Link: "http://localhost:3000/api/doctor/patient/data/gender",
	}

	response := models.ListAcceptancePatientFailedResponse{
		ErrorMessage: message,
		Data:         data,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func GenderChartSuccessResponse(c *gin.Context, message string, data interface{}, status int) {
	linkItem := models.LinkItem{
		Name: "Patient Data by Gender",
		Link: "http://localhost:3000/api/doctor/patient/data/gender",
	}

	response := models.ListAcceptancePatientSuccessResponse{
		Message: message,
		Data:    data,
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func AgeChartFailedResponse(c *gin.Context, message string, data string, status int) {
	linkItem := models.LinkItem{
		Name: "Patient Data by Gender",
		Link: "http://localhost:3000/api/doctor/patient/data/age",
	}

	response := models.ListAcceptancePatientFailedResponse{
		ErrorMessage: message,
		Data:         data,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func AgeChartSuccessResponse(c *gin.Context, message string, data interface{}, status int) {
	linkItem := models.LinkItem{
		Name: "Patient Data by Gender",
		Link: "http://localhost:3000/api/doctor/patient/data/age",
	}

	response := models.ListAcceptancePatientSuccessResponse{
		Message: message,
		Data:    data,
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}
