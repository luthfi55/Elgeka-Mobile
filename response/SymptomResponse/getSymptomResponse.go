package symptomresponse

import (
	"elgeka-mobile/models"

	"github.com/gin-gonic/gin"
)

func GetSymptomSuccessResponse(c *gin.Context, message string, data interface{}, link string, status int) {
	linkItem := models.LinkItem{
		Name: "Get Symptom",
		Link: link,
	}

	response := models.SubmitSymptomAnswerSuccess{
		Message: message,
		Data:    data,
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func GetSymptomFailedResponse(c *gin.Context, message string, data []string, link string, status int) {
	linkItem := models.LinkItem{
		Name: "Get Symptom",
		Link: link,
	}

	response := models.SubmitSymptomAnswerFailed{
		ErrorMessage: message,
		Data:         data,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

// func SymptomTypeNotFoundResponse(c *gin.Context, message string, data models.SymptomAnswerType, link string, status int) {
// 	linkItem := models.LinkItem{
// 		Name: "Get Symptom",
// 		Link: link,
// 	}

// 	response := models.SymptomTypeNotFound{
// 		ErrorMessage: message,
// 		Data:         data,
// 		Link:         []models.LinkItem{linkItem},
// 	}

// 	c.JSON(status, response)
// }
