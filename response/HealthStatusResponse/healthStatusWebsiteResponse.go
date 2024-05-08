package healthstatusresponse

import (
	"elgeka-mobile/models"

	"github.com/gin-gonic/gin"
)

func HealthStatusWebsiteFailedResponse(c *gin.Context, message string, data string, status int) {
	linkItem := models.LinkItem{
		Name: "Health Status Patient",
		Link: "http://localhost:3000/api/user/health_status/list_website/:type",
	}

	response := models.ListAcceptancePatientFailedResponse{
		ErrorMessage: message,
		Data:         data,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func HealthStatusWebsiteSuccessResponse(c *gin.Context, message string, data interface{}, status int) {
	linkItem := models.LinkItem{
		Name: "Health Status Patient",
		Link: "http://localhost:3000/api/user/health_status/list_website/:type",
	}

	response := models.ListAcceptancePatientSuccessResponse{
		Message: message,
		Data:    data,
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}
