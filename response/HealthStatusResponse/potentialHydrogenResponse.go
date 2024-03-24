package healthstatusresponse

import (
	"elgeka-mobile/models"

	"github.com/gin-gonic/gin"
)

func PotentialHydrogenFailedResponse(c *gin.Context, message string, data interface{}, link_name string, link string, status int) {
	linkItem := models.LinkItem{
		Name: link_name,
		Link: link,
	}

	response := models.PotentialHydrogenFailed{
		ErrorMessage: message,
		Data:         data,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func PotentialHydrogenSuccessResponse(c *gin.Context, message string, data interface{}, link string, status int) {
	linkItem := models.LinkItem{
		Name: "List Potential Hydrogen",
		Link: link,
	}

	response := models.PotentialHydrogenSuccess{
		Message: message,
		Data:    data,
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}
