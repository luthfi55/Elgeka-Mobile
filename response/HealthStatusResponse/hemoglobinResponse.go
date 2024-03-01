package healthstatusresponse

import (
	"elgeka-mobile/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HemoglobinFailedResponse(c *gin.Context, message string, data interface{}, link_name string, link string, status int) {
	linkItem := models.LinkItem{
		Name: link_name,
		Link: link,
	}

	response := models.HemoglobinFailed{
		ErrorMessage: message,
		Data:         data,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func HemoglobinSuccessResponse(c *gin.Context, message string, data interface{}, link string) {
	linkItem := models.LinkItem{
		Name: "List Hemoglobin",
		Link: link,
	}

	response := models.HemoglobinSuccess{
		Message: message,
		Data:    data,
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(http.StatusCreated, response)
}
