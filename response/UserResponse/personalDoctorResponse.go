package userresponse

import (
	"elgeka-mobile/models"

	"github.com/gin-gonic/gin"
)

func AddPersonalDoctorFailedResponse(c *gin.Context, message string, data string, link_name string, link string, status int) {
	// datas := models.UserIdData{
	// 	ID: data.ID,
	// }

	linkItem := models.LinkItem{
		Name: link_name,
		Link: link,
	}

	response := models.AddPersonalDoctorFailledResponse{
		ErrorMessage: message,
		Data:         data,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func AddPersonalDoctorSuccessResponse(c *gin.Context, message string, data string, link string, status int) {
	datas := models.UserPersonalDoctorID{
		ID: data,
	}

	linkItem := models.LinkItem{
		Name: "Add Personal Doctor",
		Link: link,
	}

	response := models.AddPersonalDoctorSuccessResponse{
		Message: message,
		Data:    []models.UserPersonalDoctorID{datas},
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func GetPersonalDoctorFailedResponse(c *gin.Context, message string, data string, link_name string, link string, status int) {
	// datas := models.UserIdData{
	// 	ID: data.ID,
	// }

	linkItem := models.LinkItem{
		Name: link_name,
		Link: link,
	}

	response := models.AddPersonalDoctorFailledResponse{
		ErrorMessage: message,
		Data:         data,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func GetPersonalDoctorSuccessResponse(c *gin.Context, message string, data interface{}, link string, status int) {
	linkItem := models.LinkItem{
		Name: "List Personal Doctor",
		Link: link,
	}

	response := models.GetPersonalDoctorSuccessResponse{
		Message: message,
		Data:    data,
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}
