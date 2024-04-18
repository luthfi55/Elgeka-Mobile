package userresponse

import (
	"elgeka-mobile/models"

	"github.com/gin-gonic/gin"
)

func LogoutFailedResponse(c *gin.Context, message string, link string, status int) {
	linkItem := models.LinkItem{
		Name: "Login",
		Link: link,
	}

	response := models.LogoutUserFailedResponse{
		ErrorMessage: message,
		Data:         "",
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func LogoutSuccessResponse(c *gin.Context, message string, link string, status int) {
	linkItem := models.LinkItem{
		Name: "Login",
		Link: link,
	}

	response := models.LogoutUserFailedResponse{
		ErrorMessage: message,
		Data:         "",
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}
