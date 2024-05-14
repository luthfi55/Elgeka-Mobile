package userresponse

import (
	"elgeka-mobile/models"

	"github.com/gin-gonic/gin"
)

func GetProfileSuccessResponse(c *gin.Context, message string, data models.UserInformationData, link string, status int) {
	linkItem := models.LinkItem{
		Name: "Get Profile",
		Link: link,
	}

	response := models.GetProfileSuccessResponse{
		Message: message,
		Data:    []models.UserInformationData{data},
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func GetProfileFailedResponse(c *gin.Context, message string, data string, link_name string, link string, status int) {
	linkItem := models.LinkItem{
		Name: link_name,
		Link: link,
	}

	response := models.GetProfileFailedResponse{
		ErrorMessage: message,
		Data:         data,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func ListUserWebsiteFailedResponse(c *gin.Context, message string, data string, link_name string, link string, status int) {
	linkItem := models.LinkItem{
		Name: link_name,
		Link: link,
	}

	response := models.ListUserWebsiteFailledResponse{
		ErrorMessage: message,
		Data:         data,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func ListUserWebsiteSuccessResponse(c *gin.Context, message string, data interface{}, status int) {
	linkItem := models.LinkItem{
		Name: "List Patient Acceptance",
		Link: "http://localhost:3000/api/doctor/patient_request",
	}

	response := models.ListUserWebsiteSuccessResponse{
		Message: message,
		Data:    data,
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}
