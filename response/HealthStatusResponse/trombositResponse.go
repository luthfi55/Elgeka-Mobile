package healthstatusresponse

import (
	"elgeka-mobile/models"

	"github.com/gin-gonic/gin"
)

func TrombositFailedResponse(c *gin.Context, message string, data interface{}, link_name string, link string, status int) {
	linkItem := models.LinkItem{
		Name: link_name,
		Link: link,
	}

	response := models.TrombositFailed{
		ErrorMessage: message,
		Data:         data,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func TrombositSuccessResponse(c *gin.Context, message string, data interface{}, link string, status int) {
	linkItem := models.LinkItem{
		Name: "List Trombosit",
		Link: link,
	}

	response := models.TrombositSuccess{
		Message: message,
		Data:    data,
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}
