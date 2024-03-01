package healthstatusresponse

import (
	"elgeka-mobile/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LeukocytesFailedResponse(c *gin.Context, message string, data interface{}, link_name string, link string, status int) {
	linkItem := models.LinkItem{
		Name: link_name,
		Link: link,
	}

	response := models.LeukocytesFailed{
		ErrorMessage: message,
		Data:         data,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func LeukocytesSuccessResponse(c *gin.Context, message string, data interface{}, link string) {
	linkItem := models.LinkItem{
		Name: "List Leukocytes",
		Link: link,
	}

	response := models.LeukocytesSuccess{
		Message: message,
		Data:    data,
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(http.StatusCreated, response)
}
