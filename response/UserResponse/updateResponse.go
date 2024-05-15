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

func UpdateUserInformationProfileFailedResponse(c *gin.Context, message string, data models.UserInformation, link_name string, link string, status int) {
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

func UpdatePasswordUserFailedResponse(c *gin.Context, message string, data string, status int) {
	linkItem := models.LinkItem{
		Name: "Update Password",
		Link: "http://localhost:3000/api/user/profile/password/edit",
	}

	response := models.GetTreatmentDataFailedResponse{
		ErrorMessage: message,
		Data:         data,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func UpdatePasswordUserSuccessResponse(c *gin.Context, message string, data interface{}, status int) {
	linkItem := models.LinkItem{
		Name: "Update Password",
		Link: "http://localhost:3000/api/user/profile/password/edit",
	}

	response := models.GetTreatmentDataSuccessResponse{
		Message: message,
		Data:    data,
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}
