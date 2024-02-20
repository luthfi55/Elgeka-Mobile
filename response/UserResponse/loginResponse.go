package userresponse

import (
	"elgeka-mobile/models"

	"github.com/gin-gonic/gin"
)

func LoginFailedResponse(c *gin.Context, message string, data models.User, link string, status int) {
	linkItem := models.LinkItem{
		Name: "Login",
		Link: link,
	}

	response := models.SignupFailledResponse{
		Message: message,
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func LoginSuccessResponse(c *gin.Context, message string, data string, link string, status int) {
	datas := models.Data{
		Email: data,
	}

	linkItem := models.LinkItem{
		Name: "Dashboard",
		Link: link,
	}

	response := models.SignupSuccesResponse{
		Message: message,
		Data:    []models.Data{datas},
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}
