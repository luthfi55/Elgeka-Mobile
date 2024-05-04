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

func ListDoctorPatientFailedResponse(c *gin.Context, message string, data string, status int) {
	linkItem := models.LinkItem{
		Name: "List Patient Acceptance",
		Link: "http://localhost:3000/api/doctor/patient_request",
	}

	response := models.ListDoctorPatientFailledResponse{
		ErrorMessage: message,
		Data:         data,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func ListDoctorPatientSuccessResponse(c *gin.Context, message string, data interface{}, status int) {
	linkItem := models.LinkItem{
		Name: "List Patient Acceptance",
		Link: "http://localhost:3000/api/doctor/patient_request",
	}

	response := models.ListDoctorPatientSuccessResponse{
		Message: message,
		Data:    data,
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func DoctorPatientProfileFailedResponse(c *gin.Context, message string, data string, status int) {
	linkItem := models.LinkItem{
		Name: "List Patient",
		Link: "http://localhost:3000/api/doctor/patient/list",
	}

	response := models.DoctorPatientProfileFailledResponse{
		ErrorMessage: message,
		Data:         data,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func DoctorPatientProfileSuccessResponse(c *gin.Context, message string, data interface{}, status int) {
	linkItem := models.LinkItem{
		Name: "List Patient",
		Link: "http://localhost:3000/api/doctor/patient/list",
	}

	response := models.DoctorPatientProfileSuccessResponse{
		Message: message,
		Data:    data,
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func DoctorPatientHealthStatusFailedResponse(c *gin.Context, message string, data string, status int) {
	linkItem := models.LinkItem{
		Name: "List Patient",
		Link: "http://localhost:3000/api/doctor/patient/list",
	}

	response := models.DoctorPatientHealthStatusFailledResponse{
		ErrorMessage: message,
		Data:         data,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func DoctorPatientHealthStatusSuccessResponse(c *gin.Context, message string, data interface{}, status int) {
	linkItem := models.LinkItem{
		Name: "List Patient",
		Link: "http://localhost:3000/api/doctor/patient/list",
	}

	response := models.DoctorPatientHealthStatusSuccessResponse{
		Message: message,
		Data:    data,
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func ListDoctorWebsiteFailedResponse(c *gin.Context, message string, data string, status int) {
	linkItem := models.LinkItem{
		Name: "List Doctor Web",
		Link: "http://localhost:3000/api/doctor/list/website",
	}

	response := models.ListDoctorWebsiteFailledResponse{
		ErrorMessage: message,
		Data:         data,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func ListDoctorWebsiteSuccessResponse(c *gin.Context, message string, data interface{}, status int) {
	linkItem := models.LinkItem{
		Name: "List Doctor Web",
		Link: "http://localhost:3000/api/doctor/list/website",
	}

	response := models.ListDoctorWebsiteSuccessResponse{
		Message: message,
		Data:    data,
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func ListPatientDoctorWebsiteFailedResponse(c *gin.Context, message string, data string, status int) {
	linkItem := models.LinkItem{
		Name: "List Patient Doctor Web",
		Link: "http://localhost:3000/api/doctor/list_patient/website",
	}

	response := models.ListPatientDoctorWebsiteFailledResponse{
		ErrorMessage: message,
		Data:         data,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func ListPatientDoctorWebsiteSuccessResponse(c *gin.Context, message string, data interface{}, status int) {
	linkItem := models.LinkItem{
		Name: "List Patient Doctor Web",
		Link: "http://localhost:3000/api/doctor/list_patient/website",
	}

	response := models.ListPatientDoctorWebsiteSuccessResponse{
		Message: message,
		Data:    data,
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}
