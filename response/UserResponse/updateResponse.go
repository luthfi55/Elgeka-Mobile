package userresponse

import (
	"elgeka-mobile/models"

	"github.com/gin-gonic/gin"
)

func UpdateUserProfileFailedResponse(c *gin.Context, message string, data models.User, link_name string, link string, status int) {
	datas := models.UserIdData{
		ID: data.ID,
	}

	linkItem := models.LinkItem{
		Name: link_name,
		Link: link,
	}

	response := models.LoginUserFailledResponse{
		ErrorMessage: message,
		Data:         []models.UserIdData{datas},
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func UpdateUserProfileSuccessResponse(c *gin.Context, message string, data string, link string, status int) {
	datas := models.Login{
		Email: data,
	}

	linkItem := models.LinkItem{
		Name: "Edit Profile",
		Link: link,
	}

	response := models.LoginUserSuccessResponse{
		Message: message,
		Data:    []models.Login{datas},
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}
