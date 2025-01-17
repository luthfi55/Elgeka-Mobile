package controllers

import (
	"elgeka-mobile/initializers"
	"elgeka-mobile/middleware"
	"elgeka-mobile/models"
	medicineresponse "elgeka-mobile/response/MedicineResponse"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func MedicineController(r *gin.Engine) {
	r.GET("api/user/medicine", middleware.RequireAuth, ListMedicine)
	r.GET("api/user/medicine/:medicine_id", middleware.RequireAuth, GetMedicine)
	r.POST("api/user/medicine", middleware.RequireAuth, AddMedicine)
	r.PUT("api/user/medicine/:medicine_id", middleware.RequireAuth, UpdateMedicine)
	r.DELETE("api/user/medicine/:medicine_id", middleware.RequireAuth, DeleteMedicine)

	r.POST("api/user/medicine/schedule/:medicine_id", middleware.RequireAuth, AddMedicineSchedule)
	r.GET("api/user/medicine/schedule", middleware.RequireAuth, ListMedicineSchedule)
	r.PUT("api/user/medicine/schedule/:schedule_id", middleware.RequireAuth, UpdateMedicineSchedule)
	r.DELETE("api/user/medicine/schedule/:schedule_id", middleware.RequireAuth, DeleteMedicineSchedule)

	r.GET("api/user/medicine/list/website", ListMedicineWebsite)
	r.GET("api/user/medicine/list_patient/website", ListPatientMedicineWebsite)
}

func ListMedicine(c *gin.Context) {
	user, _ := c.Get("user")

	var medicine []models.Medicine
	var medicine_data []struct {
		ID       uuid.UUID
		Name     string
		Category string
		Dosage   string
		Stock    int
	}

	if err := initializers.DB.Where("user_id = ?", user).Find(&medicine).Error; err != nil {
		medicineresponse.GetMedicineFailedResponse(c, "Failed to Get Medicine List", []models.MedicineData{}, "List Medicine", "http://localhost:3000/api/medicine", http.StatusBadRequest)
		return
	}

	for _, item := range medicine {
		medicine_data = append(medicine_data, struct {
			ID       uuid.UUID
			Name     string
			Category string
			Dosage   string
			Stock    int
		}{
			ID:       item.ID,
			Name:     item.Name,
			Category: item.Category,
			Dosage:   item.Dosage,
			Stock:    item.Stock,
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

	if body.Dosage == "" {
		medicineresponse.AddMedicineFailedResponse(c, "Dosage Can't be Empty", body, "ADd Medicine", "http://localhost:3000/api/medicine/add", http.StatusBadRequest)
		return
	}

	if body.Category == "" {
		medicineresponse.AddMedicineFailedResponse(c, "Category Can't be Empty", body, "ADd Medicine", "http://localhost:3000/api/medicine/add", http.StatusBadRequest)
		return
	}

	var medicinedata models.Medicine

	result := initializers.DB.Where("user_id = ? AND Name = ?", user, body.Name).First(&medicinedata)

	if result.RowsAffected > 0 {
		medicineresponse.AddMedicineFailedResponse(c, "Can't add Medicine, Medicine has registered", body, "Add Medicine", "http://localhost:3000/api/medicine/add", http.StatusBadRequest)
		return
	}

	body.ID = uuid.New()
	body.UserID = user.(uuid.UUID)
	body.GetMedicineDate = time.Now()

	if err := initializers.DB.Create(&body).Error; err != nil {
		medicineresponse.AddMedicineFailedResponse(c, "Failed to Add Medicine Data", body, "Add Medicine", "http://localhost:3000/api/medicine/add", http.StatusBadRequest)
		return
	}

	medicineresponse.AddMedicineSuccessResponse(c, "Success to Add Medicine Data", body, "http://localhost:3000/api/medicine/add", http.StatusCreated)

}

func GetMedicine(c *gin.Context) {
	user, _ := c.Get("user")

	medicine_id := c.Param("medicine_id")

	var medicine models.Medicine
	if err := initializers.DB.First(&medicine, "id = ? AND user_id = ?", medicine_id, user).Error; err != nil {
		medicineresponse.UpdateMedicineFailedResponse(c, "Medicine Not Found", medicine, "Update Medicine", "http://localhost:3000/api/medicine/update/"+medicine_id, http.StatusNotFound)
		return
	}

	var medicine_data models.MedicineData
	medicine_data.ID = medicine.ID
	medicine_data.Name = medicine.Name
	medicine_data.Category = medicine.Category
	medicine_data.Dosage = medicine.Dosage
	medicine_data.Stock = medicine.Stock

	medicineresponse.GetMedicineSuccessResponse(c, "Success to Get Medicine Data", medicine_data, "http://localhost:3000/api/medicine", http.StatusOK)
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

	if body.Category == "" {
		medicineresponse.UpdateMedicineFailedResponse(c, "Category Can't be Empty", body, "Update Medicine", "http://localhost:3000/api/medicine/update/"+medicine_id, http.StatusBadRequest)
		return
	}

	if body.Dosage == "" {
		medicineresponse.UpdateMedicineFailedResponse(c, "Dosage Can't be Empty", body, "Update Medicine", "http://localhost:3000/api/medicine/update/"+medicine_id, http.StatusBadRequest)
		return
	}

	var medicine models.Medicine
	if err := initializers.DB.First(&medicine, "id = ? AND user_id = ?", medicine_id, user).Error; err != nil {
		medicineresponse.UpdateMedicineFailedResponse(c, "Medicine Not Found", body, "Update Medicine", "http://localhost:3000/api/medicine/update/"+medicine_id, http.StatusNotFound)
		return
	}

	medicine.Name = body.Name
	medicine.Category = body.Category
	medicine.Dosage = body.Dosage

	var dateMedicine = medicine.GetMedicineDate
	fmt.Println("Data 1 : ", dateMedicine)

	if body.Stock > medicine.Stock {
		medicine.Stock = body.Stock
		medicine.GetMedicineDate = time.Now()
		println("tes 1")
	} else {
		medicine.Stock = body.Stock
		medicine.GetMedicineDate = dateMedicine
		println("tes 2")
	}

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

func AddMedicineSchedule(c *gin.Context) {
	user, _ := c.Get("user")
	medicine_id := c.Param("medicine_id")

	var body models.MedicineSchedule

	if c.Bind(&body) != nil {
		medicineresponse.AddMedicineScheduleFailedResponse(c, "Failed to read body", body, "Add Medicine Schedule", "http://localhost:3000/api/medicine/add", http.StatusBadRequest)
		return
	}

	if body.MedicineName == "" {
		medicineresponse.AddMedicineScheduleFailedResponse(c, "Medicine Name Can't be Empty", body, "Add Medicine", "http://localhost:3000/api/medicine/add", http.StatusBadRequest)
		return
	}

	if body.Dosage == "" {
		medicineresponse.AddMedicineScheduleFailedResponse(c, "Dosage Can't be Empty", body, "Add Medicine", "http://localhost:3000/api/medicine/add", http.StatusBadRequest)
		return
	}

	if body.Day == "" {
		medicineresponse.AddMedicineScheduleFailedResponse(c, "Day Can't be Empty", body, "Add Medicine", "http://localhost:3000/api/medicine/add", http.StatusBadRequest)
		return
	}

	if body.Hour == "" {
		medicineresponse.AddMedicineScheduleFailedResponse(c, "Hour Can't be Empty", body, "Add Medicine", "http://localhost:3000/api/medicine/add", http.StatusBadRequest)
		return
	}

	body.ID = uuid.New()
	body.UserID = user.(uuid.UUID)

	parse_medicine_id, err := uuid.Parse(medicine_id)
	if err != nil {
		medicineresponse.AddMedicineScheduleFailedResponse(c, "Invalid Medicine ID", body, "Add Medicine", "http://localhost:3000/api/medicine/add", http.StatusBadRequest)
		return
	}

	body.MedicineID = parse_medicine_id

	if err := initializers.DB.Create(&body).Error; err != nil {
		medicineresponse.AddMedicineScheduleFailedResponse(c, "Failed to Add Medicine Schedule", body, "Add Medicine", "http://localhost:3000/api/medicine/add", http.StatusBadRequest)
		return
	}

	medicineresponse.AddMedicineScheduleSuccessResponse(c, "Success to Add Medicine Schedule", body, "http://localhost:3000/api/medicine/add", http.StatusCreated)

}

func ListMedicineSchedule(c *gin.Context) {
	user, _ := c.Get("user")

	var medicine_schedule []models.MedicineSchedule
	var medicine_schedule_data []struct {
		ID            uuid.UUID
		MedicineID    uuid.UUID
		Medicine_Name string
		Dosage        string
		Day           string
		Hour          string
		Status        bool
	}

	if err := initializers.DB.Where("user_id = ?", user).Find(&medicine_schedule).Error; err != nil {
		medicineresponse.GetMedicineScheduleFailedResponse(c, "Failed to Get Medicine Schedule List", []models.MedicineScheduleData{}, "List Medicine Schedule", "http://localhost:3000/api/medicine/schedule", http.StatusBadRequest)
		return
	}

	for _, item := range medicine_schedule {
		medicine_schedule_data = append(medicine_schedule_data, struct {
			ID            uuid.UUID
			MedicineID    uuid.UUID
			Medicine_Name string
			Dosage        string
			Day           string
			Hour          string
			Status        bool
		}{
			ID:            item.ID,
			MedicineID:    item.MedicineID,
			Medicine_Name: item.MedicineName,
			Dosage:        item.Dosage,
			Day:           item.Day,
			Hour:          item.Hour,
			Status:        item.Status,
		})
	}

	medicineresponse.GetMedicineScheduleSuccessResponse(c, "Success to Get Medicine Schedule List", medicine_schedule_data, "http://localhost:3000/api/medicine/schedule", http.StatusOK)
}

func UpdateMedicineSchedule(c *gin.Context) {
	var body models.MedicineSchedule
	var medicine_schedule models.MedicineSchedule
	schedule_id := c.Param("schedule_id")

	if c.Bind(&body) != nil {
		medicineresponse.AddMedicineScheduleFailedResponse(c, "Failed to read body", body, "Add Medicine Schedule", "http://localhost:3000/api/medicine/add", http.StatusBadRequest)
		return
	}

	if err := initializers.DB.First(&medicine_schedule, "id = ?", schedule_id).Error; err != nil {
		medicineresponse.UpdateMedicineScheduleFailedResponse(c, "Failed to Find Medicine Schedule", []models.MedicineScheduleData{}, "Get Medicine Schedule", "http://localhost:3000/api/medicine/schedule", http.StatusBadRequest)
		return
	}

	if body.MedicineID.String() == "" {
		medicineresponse.AddMedicineScheduleFailedResponse(c, "Medicine ID Can't be Empty", body, "Add Medicine", "http://localhost:3000/api/medicine/add", http.StatusBadRequest)
		return
	}

	if body.MedicineName == "" {
		medicineresponse.AddMedicineScheduleFailedResponse(c, "Medicine Name Can't be Empty", body, "Add Medicine", "http://localhost:3000/api/medicine/add", http.StatusBadRequest)
		return
	}

	if body.Dosage == "" {
		medicineresponse.AddMedicineScheduleFailedResponse(c, "Dosage Can't be Empty", body, "Add Medicine", "http://localhost:3000/api/medicine/add", http.StatusBadRequest)
		return
	}

	if body.Day == "" {
		medicineresponse.AddMedicineScheduleFailedResponse(c, "Day Can't be Empty", body, "Add Medicine", "http://localhost:3000/api/medicine/add", http.StatusBadRequest)
		return
	}

	if body.Hour == "" {
		medicineresponse.AddMedicineScheduleFailedResponse(c, "Hour Can't be Empty", body, "Add Medicine", "http://localhost:3000/api/medicine/add", http.StatusBadRequest)
		return
	}

	medicine_schedule.MedicineID = body.MedicineID
	medicine_schedule.MedicineName = body.MedicineName
	medicine_schedule.Dosage = body.Dosage
	medicine_schedule.Day = body.Day
	medicine_schedule.Hour = body.Hour
	medicine_schedule.Status = body.Status

	if err := initializers.DB.Save(&medicine_schedule).Error; err != nil {
		medicineresponse.UpdateMedicineScheduleFailedResponse(c, "Failed to Update Medicine Schedule", []models.MedicineScheduleData{}, "Update Medicine Schedule", "http://localhost:3000/api/user/medicine/schedule/:schedule_id", http.StatusBadRequest)
		return
	}

	var medicine_schedule_data models.MedicineScheduleData
	medicine_schedule_data.ID = medicine_schedule.ID
	medicine_schedule_data.MedicineID = medicine_schedule.MedicineID
	medicine_schedule_data.MedicineName = body.MedicineName
	medicine_schedule_data.Dosage = body.Dosage
	medicine_schedule_data.Day = body.Day
	medicine_schedule_data.Hour = body.Hour
	medicine_schedule_data.Status = body.Status

	medicineresponse.UpdateMedicineScheduleSuccessResponse(c, "Success to Update Medicine Schedule", medicine_schedule_data, "http://localhost:3000/api/user/medicine/schedule/:schedule_id", http.StatusOK)
}

func DeleteMedicineSchedule(c *gin.Context) {
	var medicine_schedule models.MedicineSchedule
	schedule_id := c.Param("schedule_id")

	if err := initializers.DB.First(&medicine_schedule, "id = ?", schedule_id).Error; err != nil {
		medicineresponse.DeleteMedicineScheduleFailedResponse(c, "Failed to Find Medicine Schedule", []models.MedicineScheduleData{}, "Get Medicine Schedule", "http://localhost:3000/api/medicine/schedule", http.StatusBadRequest)
		return
	}

	var medicine_schedule_data models.MedicineScheduleData
	medicine_schedule_data.ID = medicine_schedule.ID
	medicine_schedule_data.MedicineName = medicine_schedule.MedicineName
	medicine_schedule_data.Status = medicine_schedule.Status

	if err := initializers.DB.Delete(&medicine_schedule).Error; err != nil {
		medicineresponse.DeleteMedicineScheduleFailedResponse(c, "Failed to Delete Medicine Schedule", []models.MedicineScheduleData{}, "Delete Medicine Schedule", "http://localhost:3000/api/user/medicine/schedule/:schedule_id", http.StatusBadRequest)
		return
	}

	medicineresponse.DeleteMedicineScheduleSuccessResponse(c, "Success to Delete Medicine Schedule", medicine_schedule_data, "http://localhost:3000/api/user/medicine/schedule/:schedule_id", http.StatusOK)
}

func ListMedicineWebsite(c *gin.Context) {
	if !ParseWebToken(c) {
		return
	}

	var totalPatientHaveMedicine []models.Medicine
	initializers.DB.Distinct("user_id").Find(&totalPatientHaveMedicine)

	numTotalPatientHaveMedicine := len(totalPatientHaveMedicine)

	var medicines []models.Medicine
	result := initializers.DB.Find(&medicines)
	if result.Error != nil {
		medicineresponse.GetMedicineWebsiteFailedResponse(c, "Failed to Get Medicine List Website", "", http.StatusInternalServerError)
	}

	var medicineData []struct {
		Medicine_Name string
		Category      string
		Total_Patient int
	}
	medicineCount := make(map[string]int)
	medicineCategories := make(map[string]string)

	for _, medicine := range medicines {
		medicineCount[medicine.Name]++
		medicineCategories[medicine.Name] = medicine.Category
	}

	for name, count := range medicineCount {
		medicineData = append(medicineData, struct {
			Medicine_Name string
			Category      string
			Total_Patient int
		}{
			Medicine_Name: name,
			Category:      medicineCategories[name],
			Total_Patient: count,
		})
	}

	var Data struct {
		Total_Patient_Have_Medicine int
		Medicine                    interface{}
	}

	Data.Total_Patient_Have_Medicine = numTotalPatientHaveMedicine
	Data.Medicine = medicineData

	medicineresponse.GetMedicineWebsiteSuccessResponse(c, "Success to Get Medicine List Website", Data, http.StatusOK)
}

func ListPatientMedicineWebsite(c *gin.Context) {
	if !ParseWebToken(c) {
		return
	}

	var user []models.User
	var userData []models.MedicineDataWebsite

	result := initializers.DB.Where("is_active = ?", true).Find(&user)

	if result.Error != nil {
		medicineresponse.GetPatientMedicineWebsiteFailedResponse(c, "Failed to Get Patient Medicine List Website", "", http.StatusInternalServerError)
		return
	}

	for _, item := range user {
		var medicine []models.Medicine
		var medicineDataDate []models.MedicineDataDate
		initializers.DB.Where("user_id = ?", item.ID).Find(&medicine)
		for _, item := range medicine {
			// Contoh input datetime
			inputDatetime := item.GetMedicineDate

			wib, err := time.LoadLocation("Asia/Jakarta")
			if err != nil {
				fmt.Println("Error loading location:", err)
				return
			}

			// Mengonversi waktu ke zona waktu WIB
			inputDatetime = inputDatetime.In(wib)

			// Mengonversi objek time.Time menjadi string dengan format yang diinginkan
			layoutOutput := "2006-01-02 15:04:05"
			formattedDatetime := inputDatetime.Format(layoutOutput)
			medicineDataDate = append(medicineDataDate, models.MedicineDataDate{
				ID:       item.ID,
				Name:     item.Name,
				Category: item.Category,
				Dosage:   item.Dosage,
				Stock:    item.Stock,
				Date:     formattedDatetime,
			})
		}
		if len(medicine) > 0 {
			userData = append(userData, models.MedicineDataWebsite{
				ID:           item.ID,
				Name:         item.Name,
				Email:        item.Email,
				PhoneNumber:  item.PhoneNumber,
				ListMedicine: medicineDataDate,
			})
		}
	}

	medicineresponse.GetPatientMedicineWebsiteSuccessResponse(c, "Success to Get Patient Medicine List Website", userData, http.StatusOK)
}
