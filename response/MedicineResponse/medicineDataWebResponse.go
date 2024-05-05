package medicineresponse

import (
	"elgeka-mobile/models"

	"github.com/gin-gonic/gin"
)

func GetMedicineWebsiteFailedResponse(c *gin.Context, message string, data string, status int) {
	linkItem := models.LinkItem{
		Name: "List Medicinte Web",
		Link: "http://localhost:3000/api/user/medicine/list/website",
	}

	response := models.GetMedicineWebsiteFailedResponse{
		ErrorMessage: message,
		Data:         data,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func GetMedicineWebsiteSuccessResponse(c *gin.Context, message string, data interface{}, status int) {
	linkItem := models.LinkItem{
		Name: "List Medicine Web",
		Link: "http://localhost:3000/api/user/medicine/list/website",
	}

	response := models.GetMedicineWebsiteSuccessResponse{
		Message: message,
		Data:    data,
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func GetPatientMedicineWebsiteFailedResponse(c *gin.Context, message string, data string, status int) {
	linkItem := models.LinkItem{
		Name: "List Medicinte Web",
		Link: "http://localhost:3000/api/user/medicine/list/website",
	}

	response := models.GetPatientMedicineWebsiteFailedResponse{
		ErrorMessage: message,
		Data:         data,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func GetPatientMedicineWebsiteSuccessResponse(c *gin.Context, message string, data interface{}, status int) {
	linkItem := models.LinkItem{
		Name: "List Medicine Web",
		Link: "http://localhost:3000/api/user/medicine/list/website",
	}

	response := models.GetPatientMedicineWebsiteSuccessResponse{
		Message: message,
		Data:    data,
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}
