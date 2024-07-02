package healthstatusresponse

import (
	"elgeka-mobile/models"

	"github.com/gin-gonic/gin"
)

func HematokritFailedResponse(c *gin.Context, message string, data interface{}, link_name string, link string, status int) {
	linkItem := models.LinkItem{
		Name: link_name,
		Link: link,
	}

	response := models.HematokritFailed{
		ErrorMessage: message,
		Data:         data,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func HematokritSuccessResponse(c *gin.Context, message string, data interface{}, link string, status int) {
	linkItem := models.LinkItem{
		Name: "List Hematokrit",
		Link: link,
	}

	response := models.HematokritSuccess{
		Message: message,
		Data:    data,
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}
