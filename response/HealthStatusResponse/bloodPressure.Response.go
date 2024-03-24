package healthstatusresponse

import (
	"elgeka-mobile/models"

	"github.com/gin-gonic/gin"
)

func BloodPressureFailedResponse(c *gin.Context, message string, data interface{}, link_name string, link string, status int) {
	linkItem := models.LinkItem{
		Name: link_name,
		Link: link,
	}

	response := models.BloodPressureFailed{
		ErrorMessage: message,
		Data:         data,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func BloodPressureSuccessResponse(c *gin.Context, message string, data interface{}, link string, status int) {
	linkItem := models.LinkItem{
		Name: "List Blood Pressure",
		Link: link,
	}

	response := models.BloodPressureSuccess{
		Message: message,
		Data:    data,
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}
