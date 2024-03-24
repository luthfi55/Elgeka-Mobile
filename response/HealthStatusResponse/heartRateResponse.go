package healthstatusresponse

import (
	"elgeka-mobile/models"

	"github.com/gin-gonic/gin"
)

func HeartRateFailedResponse(c *gin.Context, message string, data interface{}, link_name string, link string, status int) {
	linkItem := models.LinkItem{
		Name: link_name,
		Link: link,
	}

	response := models.HeartRateFailed{
		ErrorMessage: message,
		Data:         data,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func HeartRateSuccessResponse(c *gin.Context, message string, data interface{}, link string, status int) {
	linkItem := models.LinkItem{
		Name: "List Heart Rate",
		Link: link,
	}

	response := models.HeartRateSuccess{
		Message: message,
		Data:    data,
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}
