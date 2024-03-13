package controllers

import (
	"elgeka-mobile/initializers"
	"elgeka-mobile/middleware"
	"elgeka-mobile/models"
	medicineresponse "elgeka-mobile/response/MedicineResponse"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func MedicineController(r *gin.Engine) {
	r.GET("api/user/medicine", middleware.RequireAuth, ListMedicine)
	r.POST("api/user/medicine", middleware.RequireAuth, AddMedicine)
	r.PUT("api/user/medicine/:medicine_id", middleware.RequireAuth, UpdateMedicine)
	r.DELETE("api/user/medicine/:medicine_id", middleware.RequireAuth, DeleteMedicine)
}

func ListMedicine(c *gin.Context) {
	user, _ := c.Get("user")

	var medicine []models.Medicine
	var medicine_data []struct {
		ID    uuid.UUID
		Name  string
		Stock int
	}

	if err := initializers.DB.Where("user_id = ?", user).Find(&medicine).Error; err != nil {
		medicineresponse.GetMedicineFailedResponse(c, "Failed to Get Medicine List", []models.MedicineData{}, "List Medicine", "http://localhost:3000/api/medicine", http.StatusBadRequest)
		return
	}

	for _, item := range medicine {
		medicine_data = append(medicine_data, struct {
			ID    uuid.UUID
			Name  string
			Stock int
		}{
			ID:    item.ID,
			Name:  item.Name,
			Stock: item.Stock,
		})
	}

	medicineresponse.GetMedicineSuccessResponse(c, "Success to Get Medicine List", medicine_data, "http://localhost:3000/api/medicine", http.StatusOK)
}

func AddMedicine(c *gin.Context) {
	user, _ := c.Get("user")

	var body models.Medicine

	if c.Bind(&body) != nil {
		medicineresponse.AddMedicineFailedResponse(c, "Failed to read body", body, "Add Medicine", "http://localhost:3000/api/medicine/add", http.StatusBadRequest)
		return
	}

	if body.Name == "" || body.Stock == 0 {
		medicineresponse.AddMedicineFailedResponse(c, "Name or Stock Can't be Empty", body, "Add Medicine", "http://localhost:3000/api/medicine/add", http.StatusBadRequest)
		return
	}

	body.ID = uuid.New()
	body.UserID = user.(uuid.UUID)

	if err := initializers.DB.Create(&body).Error; err != nil {
		medicineresponse.AddMedicineFailedResponse(c, "Failed to Add Medicine Data", body, "Add Medicine", "http://localhost:3000/api/medicine/add", http.StatusBadRequest)
		return
	}

	medicineresponse.AddMedicineSuccessResponse(c, "Success to Add Medicine Data", body, "http://localhost:3000/api/medicine/add", http.StatusCreated)

}

func UpdateMedicine(c *gin.Context) {
	user, _ := c.Get("user")

	var body models.Medicine
	medicine_id := c.Param("medicine_id")

	if c.Bind(&body) != nil {
		medicineresponse.UpdateMedicineFailedResponse(c, "Failed to read body", body, "Update Medicine", "http://localhost:3000/api/medicine/update/"+medicine_id, http.StatusBadRequest)
		return
	}

	if body.Name == "" {
		medicineresponse.UpdateMedicineFailedResponse(c, "Name Can't be Empty", body, "Update Medicine", "http://localhost:3000/api/medicine/update/"+medicine_id, http.StatusBadRequest)
		return
	}

	var medicine models.Medicine
	if err := initializers.DB.First(&medicine, "id = ? AND user_id = ?", medicine_id, user).Error; err != nil {
		medicineresponse.UpdateMedicineFailedResponse(c, "Medicine Not Found", body, "Update Medicine", "http://localhost:3000/api/medicine/update/"+medicine_id, http.StatusNotFound)
		return
	}

	medicine.Name = body.Name
	medicine.Stock = body.Stock

	if err := initializers.DB.Save(&medicine).Error; err != nil {
		medicineresponse.UpdateMedicineFailedResponse(c, "Failed to Update Medicine Data", body, "Update Medicine", "http://localhost:3000/api/medicine/update/"+medicine_id, http.StatusBadRequest)
		return
	}

	medicineresponse.UpdateMedicineSuccessResponse(c, "Success to Update Medicine Data", medicine, "http://localhost:3000/api/medicine/update/"+medicine_id, http.StatusOK)
}

func DeleteMedicine(c *gin.Context) {
	user, _ := c.Get("user")

	medicine_id := c.Param("medicine_id")

	var medicine models.Medicine
	if err := initializers.DB.First(&medicine, "id = ? AND user_id = ?", medicine_id, user).Error; err != nil {
		medicineresponse.DeleteMedicineFailedResponse(c, "Medicine Not Found", medicine, "Delete Medicine", "http://localhost:3000/api/medicine/delete/"+medicine_id, http.StatusNotFound)
		return
	}

	if err := initializers.DB.Delete(&medicine).Error; err != nil {
		medicineresponse.DeleteMedicineFailedResponse(c, "Failed to Delete Medicine Data", medicine, "Delete Medicine", "http://localhost:3000/api/medicine/delete/"+medicine_id, http.StatusBadRequest)
		return
	}

	medicineresponse.DeleteMedicineSuccessResponse(c, "Success to Delete Medicine Data", medicine, "http://localhost:3000/api/medicine/delete/"+medicine_id, http.StatusOK)
}
