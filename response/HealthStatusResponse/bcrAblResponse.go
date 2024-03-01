package healthstatusresponse

import (
	"elgeka-mobile/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BcrAblFailedResponse(c *gin.Context, message string, data interface{}, link_name string, link string, status int) {
	linkItem := models.LinkItem{
		Name: link_name,
		Link: link,
	}

	response := models.BcrAblFailed{
		ErrorMessage: message,
		Data:         data,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func BcrAblSuccessResponse(c *gin.Context, message string, data interface{}, link string) {
	linkItem := models.LinkItem{
		Name: "List BCR-ABL",
		Link: link,
	}

	response := models.BcrAblSuccess{
		Message: message,
		Data:    data,
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(http.StatusCreated, response)
}
