package userresponse

import (
	"elgeka-mobile/models"

	"github.com/gin-gonic/gin"
)

func GetProfileSuccessResponse(c *gin.Context, message string, data models.UserData, link string, status int) {
	datas := models.UserData{
		ID:          data.ID,
		Name:        data.Name,
		Email:       data.Email,
		Address:     data.Address,
		Gender:      data.Gender,
		BirthDate:   data.BirthDate,
		BloodGroup:  data.BloodGroup,
		PhoneNumber: data.PhoneNumber,
	}

	linkItem := models.LinkItem{
		Name: "Get Profile",
		Link: link,
	}

	response := models.GetProfileSuccessResponse{
		Message: message,
		Data:    []models.UserData{datas},
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func GetProfileFailedResponse(c *gin.Context, message string, data string, link_name string, link string, status int) {
	linkItem := models.LinkItem{
		Name: link_name,
		Link: link,
	}

	response := models.GetProfileFailedResponse{
		ErrorMessage: message,
		Data:         data,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}
