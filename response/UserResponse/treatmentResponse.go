package userresponse

import (
	"elgeka-mobile/models"

	"github.com/gin-gonic/gin"
)

func GetTreatmentDataFailedResponse(c *gin.Context, message string, data string, status int) {
	linkItem := models.LinkItem{
		Name: "Get Treatment Data",
		Link: "http://localhost:3000/api/user/profile/treatment",
	}

	response := models.GetTreatmentDataFailedResponse{
		ErrorMessage: message,
		Data:         data,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func GetTreatmentDataSuccessResponse(c *gin.Context, message string, data interface{}, status int) {
	linkItem := models.LinkItem{
		Name: "Get Treatment Data",
		Link: "http://localhost:3000/api/user/profile/treatment",
	}

	response := models.GetTreatmentDataSuccessResponse{
		Message: message,
		Data:    data,
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}
