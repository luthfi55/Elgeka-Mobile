package userresponse

import (
	"elgeka-mobile/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UpdateUserProfileFailedResponse(c *gin.Context, message string, data models.User, link_name string, link string, status int) {
	datas := models.UserIdData{
		ID: data.ID,
	}

	linkItem := models.LinkItem{
		Name: link_name,
		Link: link,
	}

	response := models.LoginUserFailedResponse{
		ErrorMessage: message,
		Data:         []models.UserIdData{datas},
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func UpdateUserProfileSuccessResponse(c *gin.Context, message string, data uuid.UUID, link string, status int) {
	datas := models.UserIdData{
		ID: data,
	}

	linkItem := models.LinkItem{
		Name: "Edit Profile",
		Link: link,
	}

	response := models.UpdateUserProfileSuccessResponse{
		Message: message,
		Data:    []models.UserIdData{datas},
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}
