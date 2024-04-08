package medicineresponse

import (
	"elgeka-mobile/models"

	"github.com/gin-gonic/gin"
)

func AddMedicineScheduleFailedResponse(c *gin.Context, message string, data models.MedicineSchedule, link_name string, link string, status int) {
	datas := models.MedicineScheduleData{
		ID:     data.ID,
		Date:   data.Date,
		Status: data.Status,
	}

	linkItem := models.LinkItem{
		Name: link_name,
		Link: link,
	}

	response := models.AddMedicineScheduleFailedResponse{
		ErrorMessage: message,
		Data:         []models.MedicineScheduleData{datas},
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func AddMedicineScheduleSuccessResponse(c *gin.Context, message string, data models.MedicineSchedule, link string, status int) {
	datas := models.MedicineScheduleData{
		ID:     data.ID,
		Date:   data.Date,
		Status: data.Status,
	}

	linkItem := models.LinkItem{
		Name: "Add Medicine Schedule",
		Link: link,
	}

	response := models.AddMedicineScheduleSuccessResponse{
		Message: message,
		Data:    []models.MedicineScheduleData{datas},
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func GetMedicineScheduleFailedResponse(c *gin.Context, message string, data []models.MedicineScheduleData, link_name string, link string, status int) {
	linkItem := models.LinkItem{
		Name: link_name,
		Link: link,
	}

	response := models.GetMedicineScheduleFailedResponse{
		ErrorMessage: message,
		Data:         data,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func GetMedicineScheduleSuccessResponse(c *gin.Context, message string, data interface{}, link string, status int) {
	linkItem := models.LinkItem{
		Name: "List Medicine Schedule",
		Link: link,
	}

	response := models.GetMedicineScheduleSuccessResponse{
		Message: message,
		Data:    data,
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

// func UpdateMedicineScheduleFailedResponse(c *gin.Context, message string, data models.MedicineSchedule, link_name string, link string, status int) {
// 	linkItem := models.LinkItem{
// 		Name: link_name,
// 		Link: link,
// 	}

// 	response := models.GetMedicineScheduleFailedResponse{
// 		ErrorMessage: message,
// 		Data:         data,
// 		Link:         []models.LinkItem{linkItem},
// 	}

// 	c.JSON(status, response)
// }

func UpdateMedicineScheduleSuccessResponse(c *gin.Context, message string, data models.MedicineScheduleData, link string, status int) {
	linkItem := models.LinkItem{
		Name: "Update Medicine Schedule",
		Link: link,
	}

	response := models.GetMedicineScheduleSuccessResponse{
		Message: message,
		Data:    data,
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}
