package doctorresponse

import (
	"elgeka-mobile/models"

	"github.com/gin-gonic/gin"
)

func GetDoctorProfileFailedResponse(c *gin.Context, message string, data string, status int) {
	linkItem := models.LinkItem{
		Name: "Get Doctor Profile",
		Link: "http://localhost:3000/api/doctor/profile",
	}

	response := models.ListAcceptancePatientFailedResponse{
		ErrorMessage: message,
		Data:         data,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func GetDoctorProfileSuccessResponse(c *gin.Context, message string, data interface{}, status int) {
	linkItem := models.LinkItem{
		Name: "Get Doctor Profile",
		Link: "http://localhost:3000/api/doctor/profile",
	}

	response := models.ListAcceptancePatientSuccessResponse{
		Message: message,
		Data:    data,
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func UpdateDoctorProfileFailedResponse(c *gin.Context, message string, data string, status int) {
	linkItem := models.LinkItem{
		Name: "Update Doctor Profile",
		Link: "http://localhost:3000/api/doctor/profile/edit",
	}

	response := models.ListAcceptancePatientFailedResponse{
		ErrorMessage: message,
		Data:         data,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func UpdateDoctorProfileSuccessResponse(c *gin.Context, message string, data interface{}, status int) {
	linkItem := models.LinkItem{
		Name: "Update Doctor Profile",
		Link: "http://localhost:3000/api/doctor/profile/edit",
	}

	response := models.ListAcceptancePatientSuccessResponse{
		Message: message,
		Data:    data,
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func UpdatePasswordDoctorFailedResponse(c *gin.Context, message string, data string, status int) {
	linkItem := models.LinkItem{
		Name: "Update Doctor Password",
		Link: "http://localhost:3000/api/doctor/profile/password/edit",
	}

	response := models.GetTreatmentDataFailedResponse{
		ErrorMessage: message,
		Data:         data,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func UpdatePasswordDoctorSuccessResponse(c *gin.Context, message string, data interface{}, status int) {
	linkItem := models.LinkItem{
		Name: "Update Doctor Password",
		Link: "http://localhost:3000/api/doctor/profile/password/edit",
	}

	response := models.GetTreatmentDataSuccessResponse{
		Message: message,
		Data:    data,
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}
