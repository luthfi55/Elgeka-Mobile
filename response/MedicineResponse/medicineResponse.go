package medicineresponse

import (
	"elgeka-mobile/models"

	"github.com/gin-gonic/gin"
)

func AddMedicineFailedResponse(c *gin.Context, message string, data models.Medicine, link_name string, link string, status int) {
	datas := models.MedicineData{
		ID:       data.ID,
		Name:     data.Name,
		Category: data.Category,
		Dosage:   data.Dosage,
		Stock:    data.Stock,
	}

	linkItem := models.LinkItem{
		Name: link_name,
		Link: link,
	}

	response := models.AddMedicineFailedResponse{
		ErrorMessage: message,
		Data:         []models.MedicineData{datas},
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func AddMedicineSuccessResponse(c *gin.Context, message string, data models.Medicine, link string, status int) {
	datas := models.MedicineData{
		ID:       data.ID,
		Name:     data.Name,
		Category: data.Category,
		Dosage:   data.Dosage,
		Stock:    data.Stock,
	}

	linkItem := models.LinkItem{
		Name: "Add Medicine Data",
		Link: link,
	}

	response := models.AddMedicineSuccessResponse{
		Message: message,
		Data:    []models.MedicineData{datas},
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func GetMedicineFailedResponse(c *gin.Context, message string, data []models.MedicineData, link_name string, link string, status int) {
	linkItem := models.LinkItem{
		Name: link_name,
		Link: link,
	}

	response := models.GetMedicineFailedResponse{
		ErrorMessage: message,
		Data:         data,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func GetMedicineSuccessResponse(c *gin.Context, message string, data interface{}, link string, status int) {
	linkItem := models.LinkItem{
		Name: "List Medicine Data",
		Link: link,
	}

	response := models.GetMedicineSuccessResponse{
		Message: message,
		Data:    data,
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func UpdateMedicineFailedResponse(c *gin.Context, message string, data models.Medicine, link_name string, link string, status int) {
	datas := models.MedicineData{
		ID:       data.ID,
		Name:     data.Name,
		Category: data.Category,
		Dosage:   data.Dosage,
		Stock:    data.Stock,
	}

	linkItem := models.LinkItem{
		Name: link_name,
		Link: link,
	}

	response := models.UpdateMedicineFailedResponse{
		ErrorMessage: message,
		Data:         []models.MedicineData{datas},
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func UpdateMedicineSuccessResponse(c *gin.Context, message string, data models.Medicine, link string, status int) {
	datas := models.MedicineData{
		ID:       data.ID,
		Name:     data.Name,
		Category: data.Category,
		Dosage:   data.Dosage,
		Stock:    data.Stock,
	}

	linkItem := models.LinkItem{
		Name: "Update Medicine Data",
		Link: link,
	}

	response := models.UpdateMedicineSuccessResponse{
		Message: message,
		Data:    []models.MedicineData{datas},
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func DeleteMedicineFailedResponse(c *gin.Context, message string, data models.Medicine, link_name string, link string, status int) {
	datas := models.MedicineData{
		ID:       data.ID,
		Name:     data.Name,
		Category: data.Category,
		Dosage:   data.Dosage,
		Stock:    data.Stock,
	}

	linkItem := models.LinkItem{
		Name: link_name,
		Link: link,
	}

	response := models.DeleteMedicineFailedResponse{
		ErrorMessage: message,
		Data:         []models.MedicineData{datas},
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func DeleteMedicineSuccessResponse(c *gin.Context, message string, data models.Medicine, link string, status int) {
	datas := models.MedicineData{
		ID:       data.ID,
		Name:     data.Name,
		Category: data.Category,
		Dosage:   data.Dosage,
		Stock:    data.Stock,
	}

	linkItem := models.LinkItem{
		Name: "Delete Medicine Data",
		Link: link,
	}

	response := models.DeleteMedicineSuccessResponse{
		Message: message,
		Data:    []models.MedicineData{datas},
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}
