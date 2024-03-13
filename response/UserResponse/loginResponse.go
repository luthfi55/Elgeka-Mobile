package userresponse

import (
	"elgeka-mobile/models"

	"github.com/gin-gonic/gin"
)

func LoginFailedResponse(c *gin.Context, message string, data models.UserIdData, link string, status int) {
	datas := models.UserIdData{
		ID: data.ID,
	}

	linkItem := models.LinkItem{
		Name: "Login",
		Link: link,
	}

	response := models.LoginUserFailedResponse{
		ErrorMessage: message,
		Data:         []models.UserIdData{datas},
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func LoginSuccessResponse(c *gin.Context, message string, data string, link string, status int) {
	datas := models.Login{
		Email: data,
	}

	linkItem := models.LinkItem{
		Name: "Dashboard",
		Link: link,
	}

	response := models.LoginUserSuccessResponse{
		Message: message,
		Data:    []models.Login{datas},
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}
