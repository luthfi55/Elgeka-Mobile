package doctorresponse

import (
	"elgeka-mobile/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckDoctorAccountFailedResponse(c *gin.Context) {
	linkItem := models.LinkItem{
		Name: "Login",
		Link: "http://localhost:3000/api/user/login",
	}
	var UserId []models.UserIdData
	response := models.LoginUserFailedResponse{
		ErrorMessage: "Need Doctor Login",
		Data:         UserId,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(http.StatusUnauthorized, response)
}

func ListAcceptancePatientFailedResponse(c *gin.Context, message string, data string, status int) {
	linkItem := models.LinkItem{
		Name: "Login",
		Link: "http://localhost:3000/api/user/login",
	}

	response := models.ListAcceptancePatientFailedResponse{
		ErrorMessage: message,
		Data:         data,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func ListAcceptancePatientSuccessResponse(c *gin.Context, message string, data interface{}, status int) {
	linkItem := models.LinkItem{
		Name: "Accept Patient",
		Link: "http://localhost:3000/api/doctor/patient_request/accept/:acceptance_id",
	}

	response := models.ListAcceptancePatientSuccessResponse{
		Message: message,
		Data:    data,
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func DoctorPatientAcceptFailedResponse(c *gin.Context, message string, data string, link_name string, link string, status int) {
	linkItem := models.LinkItem{
		Name: link_name,
		Link: link,
	}

	response := models.DoctorPatientAcceptFailledResponse{
		ErrorMessage: message,
		Data:         data,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func DoctorPatientAcceptSuccessResponse(c *gin.Context, message string, data interface{}, status int) {
	linkItem := models.LinkItem{
		Name: "List Patient Acceptance",
		Link: "http://localhost:3000/api/doctor/patient_request",
	}

	response := models.DoctorPatientAcceptSuccessResponse{
		Message: message,
		Data:    data,
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func DoctorPatientRejectFailedResponse(c *gin.Context, message string, data string, link_name string, link string, status int) {
	linkItem := models.LinkItem{
		Name: link_name,
		Link: link,
	}

	response := models.DoctorPatientRejectFailledResponse{
		ErrorMessage: message,
		Data:         data,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func DoctorPatientRejectSuccessResponse(c *gin.Context, message string, data interface{}, status int) {
	linkItem := models.LinkItem{
		Name: "List Patient Acceptance",
		Link: "http://localhost:3000/api/doctor/patient_request",
	}

	response := models.DoctorPatientRejectSuccessResponse{
		Message: message,
		Data:    data,
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}
