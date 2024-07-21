package controllers

import (
	"elgeka-mobile/initializers"
	"elgeka-mobile/middleware"
	"elgeka-mobile/models"
	healthstatusresponse "elgeka-mobile/response/HealthStatusResponse"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func HealthStatusController(r *gin.Engine) {
	// bcr_abl
	r.POST("api/user/health_status/bcr_abl", middleware.RequireAuth, CreateBcrAbl)
	r.GET("api/user/health_status/bcr_abl", middleware.RequireAuth, GetBcrAbl)
	r.PUT("api/user/health_status/bcr_abl/:bcr_abl_id", middleware.RequireAuth, UpdateBcrAbl)
	r.DELETE("api/user/health_status/bcr_abl/:bcr_abl_id", middleware.RequireAuth, DeleteBcrAbl)

	//leukocytes
	r.POST("api/user/health_status/leukocytes", middleware.RequireAuth, CreateLeukocytes)
	r.GET("api/user/health_status/leukocytes", middleware.RequireAuth, GetLeukocytes)
	r.PUT("api/user/health_status/leukocytes/:leukocytes_id", middleware.RequireAuth, UpdateLeukocytes)
	r.DELETE("api/user/health_status/leukocytes/:leukocytes_id", middleware.RequireAuth, DeleteLeukocytes)

	//potential hydrogen
	r.POST("api/user/health_status/potential_hydrogen", middleware.RequireAuth, CreatePotentialHydrogen)
	r.GET("api/user/health_status/potential_hydrogen", middleware.RequireAuth, GetPotentialHydrogen)
	r.PUT("api/user/health_status/potential_hydrogen/:potential_hydrogen_id", middleware.RequireAuth, UpdatePotentialHydrogen)
	r.DELETE("api/user/health_status/potential_hydrogen/:potential_hydrogen_id", middleware.RequireAuth, DeletePotentialHydrogen)

	//hemoglobin
	r.POST("api/user/health_status/hemoglobin", middleware.RequireAuth, CreateHemoglobin)
	r.GET("api/user/health_status/hemoglobin", middleware.RequireAuth, GetHemoglobin)
	r.PUT("api/user/health_status/hemoglobin/:hemoglobin_id", middleware.RequireAuth, UpdateHemoglobin)
	r.DELETE("api/user/health_status/hemoglobin/:hemoglobin_id", middleware.RequireAuth, DeleteHemoglobin)

	//heart rate
	r.POST("api/user/health_status/heart_rate", middleware.RequireAuth, CreateHeartRate)
	r.GET("api/user/health_status/heart_rate", middleware.RequireAuth, GetHeartRate)
	r.PUT("api/user/health_status/heart_rate/:heart_rate_id", middleware.RequireAuth, UpdateHeartRate)
	r.DELETE("api/user/health_status/heart_rate/:heart_rate_id", middleware.RequireAuth, DeleteHeartRate)

	//blood pressure
	r.POST("api/user/health_status/blood_pressure", middleware.RequireAuth, CreateBloodPressure)
	r.GET("api/user/health_status/blood_pressure", middleware.RequireAuth, GetBloodPressure)
	r.PUT("api/user/health_status/blood_pressure/:blood_pressure_id", middleware.RequireAuth, UpdateBloodPressure)
	r.DELETE("api/user/health_status/blood_pressure/:blood_pressure_id", middleware.RequireAuth, DeleteBloodPressure)

	//hematokrit
	r.POST("api/user/health_status/hematokrit", middleware.RequireAuth, CreateHematokrit)
	r.GET("api/user/health_status/hematokrit", middleware.RequireAuth, GetHematokrit)
	r.PUT("api/user/health_status/hematokrit/:hematokrit_id", middleware.RequireAuth, UpdateHematokrit)
	r.DELETE("api/user/health_status/hematokrit/:hematokrit_id", middleware.RequireAuth, DeleteHematokrit)

	//trombosit
	r.POST("api/user/health_status/trombosit", middleware.RequireAuth, CreateTrombosit)
	r.GET("api/user/health_status/trombosit", middleware.RequireAuth, GetTrombosit)
	r.PUT("api/user/health_status/trombosit/:trombosit_id", middleware.RequireAuth, UpdateTrombosit)
	r.DELETE("api/user/health_status/trombosit/:trombosit_id", middleware.RequireAuth, DeleteTrombosit)

	r.GET("api/user/health_status/list_website/bcr_abl", GetBcrAblPatient)
	r.GET("api/user/health_status/list_website/leukocytes", GetLeukocytesPatient)
	r.GET("api/user/health_status/list_website/potential_hydrogen", GetPotentialHydrogenPatient)
	r.GET("api/user/health_status/list_website/hemoglobin", GetHemoglobinPatient)
	r.GET("api/user/health_status/list_website/heart_rate", GetHeartRatePatient)
	r.GET("api/user/health_status/list_website/blood_pressure", GetBloodPressurePatient)
	r.GET("api/user/health_status/list_website/hematokrit", GetHematokritPatient)
	r.GET("api/user/health_status/list_website/trombosit", GetTrombositPatient)
}

func CreateBcrAbl(c *gin.Context) {
	var body models.BCR_ABL
	user, _ := c.Get("user")

	var bcr_abl_data struct {
		Data  float32
		Notes string
		Date  string
	}

	if c.Bind(&body) != nil {
		healthstatusresponse.BcrAblFailedResponse(c, "Failed to read body", bcr_abl_data, "Create BCR-ABL", "http://localhost:3000/api/user/health_status/bcr_abl", http.StatusBadRequest)
		return
	}

	bcr_abl_data.Data = body.Data
	bcr_abl_data.Notes = body.Notes
	bcr_abl_data.Date = body.Date

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make([]string, len(validationErrors))
		for i, fieldError := range validationErrors {
			errorMessages[i] = getValidationErrorTagMessage(fieldError)
		}

		healthstatusresponse.BcrAblFailedResponse(c, strings.Join(errorMessages, ", "), bcr_abl_data, "Create BCR-ABL", "http://localhost:3000/api/user/health_status/bcr_abl", http.StatusBadRequest)
		return
	}

	newUUID := uuid.New()
	bcr_abl := models.BCR_ABL{
		ID:     newUUID,
		UserID: user.(uuid.UUID),
		Data:   body.Data,
		Notes:  body.Notes,
		Date:   body.Date,
	}

	if err := initializers.DB.Create(&bcr_abl).Error; err != nil {
		healthstatusresponse.BcrAblFailedResponse(c, strings.Title(err.Error()), bcr_abl_data, "Create BCR-ABL", "http://localhost:3000/api/user/health_status/bcr_abl", http.StatusBadRequest)
	}

	healthstatusresponse.BcrAblSuccessResponse(c, "Success Create Data", bcr_abl_data, "http://localhost:3000/api/user/health_status/bcr_abl", http.StatusCreated)
}

func GetBcrAbl(c *gin.Context) {
	var bcr_abl []models.BCR_ABL
	user, _ := c.Get("user")

	var bcr_abl_data []struct {
		Id    uuid.UUID
		Data  float32
		Notes string
		Date  string
	}

	initializers.DB.Where("user_id = ?", user).Order("date asc").Find(&bcr_abl)

	if initializers.DB.Error != nil {
		healthstatusresponse.BcrAblFailedResponse(c, "User Not Found", bcr_abl_data, "Create BCR-ABL", "http://localhost:3000/api/user/health_status/bcr_abl", http.StatusBadRequest)
		return
	}

	for _, item := range bcr_abl {
		bcr_abl_data = append(bcr_abl_data, struct {
			Id    uuid.UUID
			Data  float32
			Notes string
			Date  string
		}{
			Id:    item.ID,
			Data:  item.Data,
			Notes: item.Notes,
			Date:  item.Date,
		})
	}

	if bcr_abl_data == nil {
		healthstatusresponse.BcrAblFailedResponse(c, "Data Not Found", bcr_abl_data, "Create BCR-ABL", "http://localhost:3000/api/user/health_status/bcr_abl", http.StatusNotFound)
		return

	}

	healthstatusresponse.BcrAblSuccessResponse(c, "Success Get Data", bcr_abl_data, "http://localhost:3000/api/user/health_status/bcr_abl", http.StatusOK)
}

func UpdateBcrAbl(c *gin.Context) {
	var body models.BCR_ABL
	bcr_abl_id := c.Param("bcr_abl_id")

	var bcr_abl_data struct {
		Id    uuid.UUID
		Data  float32
		Notes string
		Date  string
	}

	user, _ := c.Get("user")

	if c.Bind(&body) != nil {
		healthstatusresponse.BcrAblFailedResponse(c, "Failed to read body", bcr_abl_data, "Create BCR-ABL", "http://localhost:3000/api/user/health_status/bcr_abl", http.StatusBadRequest)
		return
	}

	body.ID = uuid.MustParse(bcr_abl_id)
	body.UserID = user.(uuid.UUID)

	bcr_abl_data.Id = user.(uuid.UUID)
	bcr_abl_data.Data = body.Data
	bcr_abl_data.Notes = body.Notes
	bcr_abl_data.Date = body.Date

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make([]string, len(validationErrors))
		for i, fieldError := range validationErrors {
			errorMessages[i] = getValidationErrorTagMessage(fieldError)
		}

		healthstatusresponse.BcrAblFailedResponse(c, strings.Join(errorMessages, ", "), bcr_abl_data, "Create BCR-ABL", "http://localhost:3000/api/user/health_status/bcr_abl", http.StatusBadRequest)
		return
	}

	var bcr_abl models.BCR_ABL
	if err := initializers.DB.First(&bcr_abl, "user_id = ? AND ID = ?", user, bcr_abl_id).Error; err != nil {
		healthstatusresponse.BcrAblFailedResponse(c, strings.Title(err.Error()), bcr_abl_data, "Create BCR-ABL", "http://localhost:3000/api/user/health_status/bcr_abl", http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Save(&body).Error; err != nil {
		healthstatusresponse.BcrAblFailedResponse(c, strings.Title(err.Error()), bcr_abl_data, "Create BCR-ABL", "http://localhost:3000/api/user/health_status/bcr_abl", http.StatusBadRequest)
		return
	}

	healthstatusresponse.BcrAblSuccessResponse(c, "Success Update Data", bcr_abl_data, "http://localhost:3000/api/user/health_status/bcr_abl", http.StatusOK)
}

func DeleteBcrAbl(c *gin.Context) {
	bcr_abl_id := c.Param("bcr_abl_id")
	user, _ := c.Get("user")

	var bcr_abl_data struct {
		Id    uuid.UUID
		Data  float32
		Notes string
		Date  string
	}

	var bcr_abl models.BCR_ABL
	if err := initializers.DB.First(&bcr_abl, "user_id = ? AND ID = ?", user, bcr_abl_id).Error; err != nil {
		healthstatusresponse.BcrAblFailedResponse(c, strings.Title(err.Error()), bcr_abl_data, "Create BCR-ABL", "http://localhost:3000/api/user/health_status/bcr_abl", http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Unscoped().Delete(&bcr_abl).Error; err != nil {
		healthstatusresponse.BcrAblFailedResponse(c, strings.Title(err.Error()), bcr_abl_data, "Create BCR-ABL", "http://localhost:3000/api/user/health_status/bcr_abl", http.StatusBadRequest)
	}

	bcr_abl_data.Id = user.(uuid.UUID)
	bcr_abl_data.Data = bcr_abl.Data
	bcr_abl_data.Notes = bcr_abl.Notes
	bcr_abl_data.Date = bcr_abl.Date

	healthstatusresponse.BcrAblSuccessResponse(c, "Success Delete Data", bcr_abl_data, "http://localhost:3000/api/user/health_status/bcr_abl", http.StatusOK)
}

func CreateLeukocytes(c *gin.Context) {
	var body models.Leukocytes
	user, _ := c.Get("user")

	var leukocytesData struct {
		Data  float32
		Notes string
		Date  string
	}

	if c.Bind(&body) != nil {
		healthstatusresponse.LeukocytesFailedResponse(c, "Failed to read body", leukocytesData, "Create Leukocytes", "http://localhost:3000/api/user/health_status/leukocytes", http.StatusBadRequest)
		return
	}

	leukocytesData.Data = body.Data
	leukocytesData.Notes = body.Notes
	leukocytesData.Date = body.Date

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make([]string, len(validationErrors))
		for i, fieldError := range validationErrors {
			errorMessages[i] = getValidationErrorTagMessage(fieldError)
		}
		healthstatusresponse.LeukocytesFailedResponse(c, strings.Join(errorMessages, ", "), leukocytesData, "Create Leukocytes", "http://localhost:3000/api/user/health_status/leukocytes", http.StatusBadRequest)
		return
	}

	newUUID := uuid.New()
	leukocytes := models.Leukocytes{
		ID:     newUUID,
		UserID: user.(uuid.UUID),
		Data:   body.Data,
		Notes:  body.Notes,
		Date:   body.Date,
	}

	if err := initializers.DB.Create(&leukocytes).Error; err != nil {
		healthstatusresponse.LeukocytesFailedResponse(c, strings.Title(err.Error()), leukocytesData, "Create Leukocytes", "http://localhost:3000/api/user/health_status/leukocytes", http.StatusBadRequest)
	}

	healthstatusresponse.LeukocytesSuccessResponse(c, "Success Create Data", leukocytesData, "http://localhost:3000/api/user/health_status/leukocytes", http.StatusCreated)
}

func GetLeukocytes(c *gin.Context) {
	var leukocytes []models.Leukocytes
	user, _ := c.Get("user")

	var leukocytesData []struct {
		Id    uuid.UUID
		Data  float32
		Notes string
		Date  string
	}

	initializers.DB.Where("user_id = ?", user).Order("date asc").Find(&leukocytes)

	if initializers.DB.Error != nil {
		healthstatusresponse.LeukocytesFailedResponse(c, "User Not Found", leukocytesData, "Create Leukocytes", "http://localhost:3000/api/user/health_status/leukocytes", http.StatusBadRequest)
		return
	}

	for _, item := range leukocytes {
		leukocytesData = append(leukocytesData, struct {
			Id    uuid.UUID
			Data  float32
			Notes string
			Date  string
		}{
			Id:    item.ID,
			Data:  item.Data,
			Notes: item.Notes,
			Date:  item.Date,
		})
	}

	if leukocytesData == nil {
		healthstatusresponse.LeukocytesFailedResponse(c, "Data Not Found", leukocytesData, "Create Leukocytes", "http://localhost:3000/api/user/health_status/leukocytes", http.StatusNotFound)
		return

	}

	healthstatusresponse.LeukocytesSuccessResponse(c, "Success Get Data", leukocytesData, "http://localhost:3000/api/user/health_status/leukocytes", http.StatusOK)
}

func UpdateLeukocytes(c *gin.Context) {
	var body models.Leukocytes
	leukocytes_id := c.Param("leukocytes_id")

	var leukocytesData struct {
		Id    uuid.UUID
		Data  float32
		Notes string
		Date  string
	}

	user, _ := c.Get("user")

	if c.Bind(&body) != nil {
		healthstatusresponse.LeukocytesFailedResponse(c, "Failed to read body", leukocytesData, "Create Leukocytes", "http://localhost:3000/api/user/health_status/leukocytes", http.StatusBadRequest)
		return
	}

	body.ID = uuid.MustParse(leukocytes_id)
	body.UserID = user.(uuid.UUID)

	leukocytesData.Id = user.(uuid.UUID)
	leukocytesData.Data = body.Data
	leukocytesData.Notes = body.Notes
	leukocytesData.Date = body.Date

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make([]string, len(validationErrors))
		for i, fieldError := range validationErrors {
			errorMessages[i] = getValidationErrorTagMessage(fieldError)
		}
		healthstatusresponse.LeukocytesFailedResponse(c, strings.Join(errorMessages, ", "), leukocytesData, "Create Leukocytes", "http://localhost:3000/api/user/health_status/leukocytes", http.StatusBadRequest)
		return
	}

	var leukocytes models.Leukocytes
	if err := initializers.DB.First(&leukocytes, "user_id = ? AND ID = ?", user, leukocytes_id).Error; err != nil {
		healthstatusresponse.LeukocytesFailedResponse(c, strings.Title(err.Error()), leukocytesData, "Create Leukocytes", "http://localhost:3000/api/user/health_status/leukocytes", http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Save(&body).Error; err != nil {
		healthstatusresponse.LeukocytesFailedResponse(c, strings.Title(err.Error()), leukocytesData, "Create Leukocytes", "http://localhost:3000/api/user/health_status/leukocytes", http.StatusBadRequest)
		return
	}

	healthstatusresponse.LeukocytesSuccessResponse(c, "Success Update Data", leukocytesData, "http://localhost:3000/api/user/health_status/leukocytes", http.StatusOK)
}

func DeleteLeukocytes(c *gin.Context) {
	leukocytes_id := c.Param("leukocytes_id")
	user, _ := c.Get("user")

	var leukocytesData struct {
		Id    uuid.UUID
		Data  float32
		Notes string
		Date  string
	}

	var leukocytes models.Leukocytes
	if err := initializers.DB.First(&leukocytes, "user_id = ? AND ID = ?", user, leukocytes_id).Error; err != nil {
		healthstatusresponse.LeukocytesFailedResponse(c, strings.Title(err.Error()), leukocytesData, "Create Leukocytes", "http://localhost:3000/api/user/health_status/leukocytes", http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Unscoped().Delete(&leukocytes).Error; err != nil {
		healthstatusresponse.LeukocytesFailedResponse(c, strings.Title(err.Error()), leukocytesData, "Create Leukocytes", "http://localhost:3000/api/user/health_status/leukocytes", http.StatusBadRequest)
	}

	leukocytesData.Id = user.(uuid.UUID)
	leukocytesData.Data = leukocytes.Data
	leukocytesData.Notes = leukocytes.Notes
	leukocytesData.Date = leukocytes.Date

	healthstatusresponse.LeukocytesSuccessResponse(c, "Success Delete Data", leukocytesData, "http://localhost:3000/api/user/health_status/leukocytes", http.StatusOK)
}

func CreatePotentialHydrogen(c *gin.Context) {
	var body models.PotentialHydrogen
	user, _ := c.Get("user")

	var potentialHydrogenData struct {
		Data  float32
		Notes string
		Date  string
	}

	if c.Bind(&body) != nil {
		healthstatusresponse.PotentialHydrogenFailedResponse(c, "Failed to read body", potentialHydrogenData, "Create Potential Hydrogen", "http://localhost:3000/api/user/health_status/potential_hydrogen", http.StatusBadRequest)
		return
	}

	potentialHydrogenData.Data = body.Data
	potentialHydrogenData.Notes = body.Notes
	potentialHydrogenData.Date = body.Date

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make([]string, len(validationErrors))
		for i, fieldError := range validationErrors {
			errorMessages[i] = getValidationErrorTagMessage(fieldError)
		}
		healthstatusresponse.PotentialHydrogenFailedResponse(c, strings.Join(errorMessages, ", "), potentialHydrogenData, "Create Potential Hydrogen", "http://localhost:3000/api/user/health_status/potential_hydrogen", http.StatusBadRequest)
		return
	}

	newUUID := uuid.New()
	potentialHydrogen := models.PotentialHydrogen{
		ID:     newUUID,
		UserID: user.(uuid.UUID),
		Data:   body.Data,
		Notes:  body.Notes,
		Date:   body.Date,
	}

	if err := initializers.DB.Create(&potentialHydrogen).Error; err != nil {
		healthstatusresponse.PotentialHydrogenFailedResponse(c, strings.Title(err.Error()), potentialHydrogenData, "Create Potential Hydrogen", "http://localhost:3000/api/user/health_status/potential_hydrogen", http.StatusBadRequest)
	}

	healthstatusresponse.PotentialHydrogenSuccessResponse(c, "Success Create Data", potentialHydrogenData, "http://localhost:3000/api/user/health_status/potential_hydrogen", http.StatusCreated)
}

func GetPotentialHydrogen(c *gin.Context) {
	var potentialHydrogen []models.PotentialHydrogen
	user, _ := c.Get("user")

	var potentialHydrogenData []struct {
		Id    uuid.UUID
		Data  float32
		Notes string
		Date  string
	}

	initializers.DB.Where("user_id = ?", user).Order("date asc").Find(&potentialHydrogen)

	if initializers.DB.Error != nil {
		healthstatusresponse.PotentialHydrogenFailedResponse(c, "User Not Found", potentialHydrogenData, "Create Potential Hydrogen", "http://localhost:3000/api/user/health_status/potential_hydrogen", http.StatusBadRequest)
		return
	}

	for _, item := range potentialHydrogen {
		potentialHydrogenData = append(potentialHydrogenData, struct {
			Id    uuid.UUID
			Data  float32
			Notes string
			Date  string
		}{
			Id:    item.ID,
			Data:  item.Data,
			Notes: item.Notes,
			Date:  item.Date,
		})
	}

	if potentialHydrogenData == nil {
		healthstatusresponse.PotentialHydrogenFailedResponse(c, "Data Not Found", potentialHydrogenData, "Create Potential Hydrogen", "http://localhost:3000/api/user/health_status/potential_hydrogen", http.StatusNotFound)
		return

	}

	healthstatusresponse.PotentialHydrogenSuccessResponse(c, "Success Get Data", potentialHydrogenData, "http://localhost:3000/api/user/health_status/potential_hydrogen", http.StatusOK)
}

func UpdatePotentialHydrogen(c *gin.Context) {
	var body models.PotentialHydrogen
	potentialHydrogen_id := c.Param("potential_hydrogen_id")

	var potentialHydrogenData struct {
		Id    uuid.UUID
		Data  float32
		Notes string
		Date  string
	}

	user, _ := c.Get("user")

	if c.Bind(&body) != nil {
		healthstatusresponse.PotentialHydrogenFailedResponse(c, "Failed to read body", potentialHydrogenData, "Create Potential Hydrogen", "http://localhost:3000/api/user/health_status/potential_hydrogen", http.StatusBadRequest)
		return
	}

	body.ID = uuid.MustParse(potentialHydrogen_id)
	body.UserID = user.(uuid.UUID)

	potentialHydrogenData.Id = user.(uuid.UUID)
	potentialHydrogenData.Data = body.Data
	potentialHydrogenData.Notes = body.Notes
	potentialHydrogenData.Date = body.Date

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make([]string, len(validationErrors))
		for i, fieldError := range validationErrors {
			errorMessages[i] = getValidationErrorTagMessage(fieldError)
		}
		healthstatusresponse.PotentialHydrogenFailedResponse(c, strings.Join(errorMessages, ", "), potentialHydrogenData, "Create Potential Hydrogen", "http://localhost:3000/api/user/health_status/potential_hydrogen", http.StatusBadRequest)
		return
	}

	var potentialHydrogen models.PotentialHydrogen
	if err := initializers.DB.First(&potentialHydrogen, "user_id = ? AND ID = ?", user, potentialHydrogen_id).Error; err != nil {
		healthstatusresponse.PotentialHydrogenFailedResponse(c, strings.Title(err.Error()), potentialHydrogenData, "Create Potential Hydrogen", "http://localhost:3000/api/user/health_status/potential_hydrogen", http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Save(&body).Error; err != nil {
		healthstatusresponse.PotentialHydrogenFailedResponse(c, strings.Title(err.Error()), potentialHydrogenData, "Create Potential Hydrogen", "http://localhost:3000/api/user/health_status/potential_hydrogen", http.StatusBadRequest)
		return
	}

	healthstatusresponse.PotentialHydrogenSuccessResponse(c, "Success Update Data", potentialHydrogenData, "http://localhost:3000/api/user/health_status/potential_hydrogen", http.StatusOK)
}

func DeletePotentialHydrogen(c *gin.Context) {
	potentialHydrogen_id := c.Param("potential_hydrogen_id")
	user, _ := c.Get("user")

	var potentialHydrogenData struct {
		Id    uuid.UUID
		Data  float32
		Notes string
		Date  string
	}

	var potentialHydrogen models.PotentialHydrogen
	if err := initializers.DB.First(&potentialHydrogen, "user_id = ? AND ID = ?", user, potentialHydrogen_id).Error; err != nil {
		healthstatusresponse.PotentialHydrogenFailedResponse(c, strings.Title(err.Error()), potentialHydrogenData, "Create Potential Hydrogen", "http://localhost:3000/api/user/health_status/potential_hydrogen", http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Unscoped().Delete(&potentialHydrogen).Error; err != nil {
		healthstatusresponse.PotentialHydrogenFailedResponse(c, strings.Title(err.Error()), potentialHydrogenData, "Create Potential Hydrogen", "http://localhost:3000/api/user/health_status/potential_hydrogen", http.StatusBadRequest)
	}

	potentialHydrogenData.Id = user.(uuid.UUID)
	potentialHydrogenData.Data = potentialHydrogen.Data
	potentialHydrogenData.Notes = potentialHydrogen.Notes
	potentialHydrogenData.Date = potentialHydrogen.Date

	healthstatusresponse.PotentialHydrogenSuccessResponse(c, "Success Delete Data", potentialHydrogenData, "http://localhost:3000/api/user/health_status/potential_hydrogen", http.StatusOK)
}

func CreateHemoglobin(c *gin.Context) {
	var body models.Hemoglobin
	user, _ := c.Get("user")

	var hemoglobinData struct {
		Data  float32
		Notes string
		Date  string
	}

	if c.Bind(&body) != nil {
		healthstatusresponse.HemoglobinFailedResponse(c, "Failed to read body", hemoglobinData, "Create Hemoglobin", "http://localhost:3000/api/user/health_status/hemoglobin", http.StatusBadRequest)
		return
	}

	hemoglobinData.Data = body.Data
	hemoglobinData.Notes = body.Notes
	hemoglobinData.Date = body.Date

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make([]string, len(validationErrors))
		for i, fieldError := range validationErrors {
			errorMessages[i] = getValidationErrorTagMessage(fieldError)
		}
		healthstatusresponse.HemoglobinFailedResponse(c, strings.Join(errorMessages, ", "), hemoglobinData, "Create Hemoglobin", "http://localhost:3000/api/user/health_status/hemoglobin", http.StatusBadRequest)
		return
	}

	newUUID := uuid.New()
	hemoglobin := models.Hemoglobin{
		ID:     newUUID,
		UserID: user.(uuid.UUID),
		Data:   body.Data,
		Notes:  body.Notes,
		Date:   body.Date,
	}

	if err := initializers.DB.Create(&hemoglobin).Error; err != nil {
		healthstatusresponse.HemoglobinFailedResponse(c, strings.Title(err.Error()), hemoglobinData, "Create Hemoglobin", "http://localhost:3000/api/user/health_status/hemoglobin", http.StatusBadRequest)
	}

	healthstatusresponse.HemoglobinSuccessResponse(c, "Success Create Data", hemoglobinData, "http://localhost:3000/api/user/health_status/hemoglobin", http.StatusCreated)
}

func GetHemoglobin(c *gin.Context) {
	var hemoglobin []models.Hemoglobin
	user, _ := c.Get("user")

	var hemoglobinData []struct {
		Id    uuid.UUID
		Data  float32
		Notes string
		Date  string
	}

	initializers.DB.Where("user_id = ?", user).Order("date asc").Find(&hemoglobin)

	if initializers.DB.Error != nil {
		healthstatusresponse.HemoglobinFailedResponse(c, "User Not Found", hemoglobinData, "Create Hemoglobin", "http://localhost:3000/api/user/health_status/hemoglobin", http.StatusBadRequest)
		return
	}

	for _, item := range hemoglobin {
		hemoglobinData = append(hemoglobinData, struct {
			Id    uuid.UUID
			Data  float32
			Notes string
			Date  string
		}{
			Id:    item.ID,
			Data:  item.Data,
			Notes: item.Notes,
			Date:  item.Date,
		})
	}

	if hemoglobinData == nil {
		healthstatusresponse.HemoglobinFailedResponse(c, "Data Not Found", hemoglobinData, "Create Hemoglobin", "http://localhost:3000/api/user/health_status/hemoglobin", http.StatusNotFound)
		return

	}

	healthstatusresponse.HemoglobinSuccessResponse(c, "Success Get Data", hemoglobinData, "http://localhost:3000/api/user/health_status/hemoglobin", http.StatusOK)
}

func UpdateHemoglobin(c *gin.Context) {
	var body models.Hemoglobin
	hemoglobin_id := c.Param("hemoglobin_id")

	var hemoglobinData struct {
		Id    uuid.UUID
		Data  float32
		Notes string
		Date  string
	}

	user, _ := c.Get("user")

	if c.Bind(&body) != nil {
		healthstatusresponse.HemoglobinFailedResponse(c, "Failed to read body", hemoglobinData, "Create Hemoglobin", "http://localhost:3000/api/user/health_status/hemoglobin", http.StatusBadRequest)
		return
	}

	body.ID = uuid.MustParse(hemoglobin_id)
	body.UserID = user.(uuid.UUID)

	hemoglobinData.Id = user.(uuid.UUID)
	hemoglobinData.Data = body.Data
	hemoglobinData.Notes = body.Notes
	hemoglobinData.Date = body.Date

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make([]string, len(validationErrors))
		for i, fieldError := range validationErrors {
			errorMessages[i] = getValidationErrorTagMessage(fieldError)
		}
		healthstatusresponse.HemoglobinFailedResponse(c, strings.Join(errorMessages, ", "), hemoglobinData, "Create Hemoglobin", "http://localhost:3000/api/user/health_status/hemoglobin", http.StatusBadRequest)
		return
	}

	var hemoglobin models.Hemoglobin
	if err := initializers.DB.First(&hemoglobin, "user_id = ? AND ID = ?", user, hemoglobin_id).Error; err != nil {
		healthstatusresponse.HemoglobinFailedResponse(c, strings.Title(err.Error()), hemoglobinData, "Create Hemoglobin", "http://localhost:3000/api/user/health_status/hemoglobin", http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Save(&body).Error; err != nil {
		healthstatusresponse.HemoglobinFailedResponse(c, strings.Title(err.Error()), hemoglobinData, "Create Hemoglobin", "http://localhost:3000/api/user/health_status/hemoglobin", http.StatusBadRequest)
		return
	}

	healthstatusresponse.HemoglobinSuccessResponse(c, "Success Update Data", hemoglobinData, "http://localhost:3000/api/user/health_status/hemoglobin", http.StatusOK)
}

func DeleteHemoglobin(c *gin.Context) {
	hemoglobin_id := c.Param("hemoglobin_id")
	user, _ := c.Get("user")

	var hemoglobinData struct {
		Id    uuid.UUID
		Data  float32
		Notes string
		Date  string
	}

	var hemoglobin models.Hemoglobin
	if err := initializers.DB.First(&hemoglobin, "user_id = ? AND ID = ?", user, hemoglobin_id).Error; err != nil {
		healthstatusresponse.HemoglobinFailedResponse(c, strings.Title(err.Error()), hemoglobinData, "Create Hemoglobin", "http://localhost:3000/api/user/health_status/hemoglobin", http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Unscoped().Delete(&hemoglobin).Error; err != nil {
		healthstatusresponse.HemoglobinFailedResponse(c, strings.Title(err.Error()), hemoglobinData, "Create Hemoglobin", "http://localhost:3000/api/user/health_status/hemoglobin", http.StatusBadRequest)
	}

	hemoglobinData.Id = user.(uuid.UUID)
	hemoglobinData.Data = hemoglobin.Data
	hemoglobinData.Notes = hemoglobin.Notes
	hemoglobinData.Date = hemoglobin.Date

	healthstatusresponse.HemoglobinSuccessResponse(c, "Success Delete Data", hemoglobinData, "http://localhost:3000/api/user/health_status/hemoglobin", http.StatusOK)
}

func CreateHeartRate(c *gin.Context) {
	var body models.HeartRate
	user, _ := c.Get("user")

	var heartRateData struct {
		Data  float32
		Notes string
		Date  string
	}

	if c.Bind(&body) != nil {
		healthstatusresponse.HeartRateFailedResponse(c, "Failed to read body", heartRateData, "Create Heart Rate", "http://localhost:3000/api/user/health_status/heart_rate", http.StatusBadRequest)
		return
	}

	heartRateData.Data = body.Data
	heartRateData.Notes = body.Notes
	heartRateData.Date = body.Date

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make([]string, len(validationErrors))
		for i, fieldError := range validationErrors {
			errorMessages[i] = getValidationErrorTagMessage(fieldError)
		}
		healthstatusresponse.HeartRateFailedResponse(c, strings.Join(errorMessages, ", "), heartRateData, "Create Heart Rate", "http://localhost:3000/api/user/health_status/heart_rate", http.StatusBadRequest)
		return
	}

	newUUID := uuid.New()
	heartRate := models.HeartRate{
		ID:     newUUID,
		UserID: user.(uuid.UUID),
		Data:   body.Data,
		Notes:  body.Notes,
		Date:   body.Date,
	}

	if err := initializers.DB.Create(&heartRate).Error; err != nil {
		healthstatusresponse.HeartRateFailedResponse(c, strings.Title(err.Error()), heartRateData, "Create Heart Rate", "http://localhost:3000/api/user/health_status/heart_rate", http.StatusBadRequest)
	}

	healthstatusresponse.HeartRateSuccessResponse(c, "Success Create Data", heartRateData, "http://localhost:3000/api/user/health_status/heart_rate", http.StatusCreated)
}

func GetHeartRate(c *gin.Context) {
	var heartRate []models.HeartRate
	user, _ := c.Get("user")

	var heartRateData []struct {
		Id    uuid.UUID
		Data  float32
		Notes string
		Date  string
	}

	initializers.DB.Where("user_id = ?", user).Order("date asc").Find(&heartRate)

	if initializers.DB.Error != nil {
		healthstatusresponse.HeartRateFailedResponse(c, "User Not Found", heartRateData, "Create Heart Rate", "http://localhost:3000/api/user/health_status/heart_rate", http.StatusBadRequest)
		return
	}

	for _, item := range heartRate {
		heartRateData = append(heartRateData, struct {
			Id    uuid.UUID
			Data  float32
			Notes string
			Date  string
		}{
			Id:    item.ID,
			Data:  item.Data,
			Notes: item.Notes,
			Date:  item.Date,
		})
	}

	if heartRateData == nil {
		healthstatusresponse.HeartRateFailedResponse(c, "Data Not Found", heartRateData, "Create Heart Rate", "http://localhost:3000/api/user/health_status/heart_rate", http.StatusNotFound)
		return

	}

	healthstatusresponse.HeartRateSuccessResponse(c, "Success Get Data", heartRateData, "http://localhost:3000/api/user/health_status/heart_rate", http.StatusOK)
}

func UpdateHeartRate(c *gin.Context) {
	var body models.HeartRate
	heartRate_id := c.Param("heart_rate_id")

	var heartRateData struct {
		Id    uuid.UUID
		Data  float32
		Notes string
		Date  string
	}

	user, _ := c.Get("user")

	if c.Bind(&body) != nil {
		healthstatusresponse.HeartRateFailedResponse(c, "Failed to read body", heartRateData, "Create Heart Rate", "http://localhost:3000/api/user/health_status/heart_rate", http.StatusBadRequest)
		return
	}

	body.ID = uuid.MustParse(heartRate_id)
	body.UserID = user.(uuid.UUID)

	heartRateData.Id = user.(uuid.UUID)
	heartRateData.Data = body.Data
	heartRateData.Notes = body.Notes
	heartRateData.Date = body.Date

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make([]string, len(validationErrors))
		for i, fieldError := range validationErrors {
			errorMessages[i] = getValidationErrorTagMessage(fieldError)
		}
		healthstatusresponse.HeartRateFailedResponse(c, strings.Join(errorMessages, ", "), heartRateData, "Create Heart Rate", "http://localhost:3000/api/user/health_status/heart_rate", http.StatusBadRequest)
		return
	}

	var heartRate models.HeartRate
	if err := initializers.DB.First(&heartRate, "user_id = ? AND ID = ?", user, heartRate_id).Error; err != nil {
		healthstatusresponse.HeartRateFailedResponse(c, strings.Title(err.Error()), heartRateData, "Create Heart Rate", "http://localhost:3000/api/user/health_status/heart_rate", http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Save(&body).Error; err != nil {
		healthstatusresponse.HeartRateFailedResponse(c, strings.Title(err.Error()), heartRateData, "Create Heart Rate", "http://localhost:3000/api/user/health_status/heart_rate", http.StatusBadRequest)
		return
	}

	healthstatusresponse.HeartRateSuccessResponse(c, "Success Update Data", heartRateData, "http://localhost:3000/api/user/health_status/heart_rate", http.StatusOK)
}

func DeleteHeartRate(c *gin.Context) {
	heartRate_id := c.Param("heart_rate_id")
	user, _ := c.Get("user")

	var heartRateData struct {
		Id    uuid.UUID
		Data  float32
		Notes string
		Date  string
	}

	var heartRate models.HeartRate
	if err := initializers.DB.First(&heartRate, "user_id = ? AND ID = ?", user, heartRate_id).Error; err != nil {
		healthstatusresponse.HeartRateFailedResponse(c, strings.Title(err.Error()), heartRateData, "Create Heart Rate", "http://localhost:3000/api/user/health_status/heart_rate", http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Unscoped().Delete(&heartRate).Error; err != nil {
		healthstatusresponse.HeartRateFailedResponse(c, strings.Title(err.Error()), heartRateData, "Create Heart Rate", "http://localhost:3000/api/user/health_status/heart_rate", http.StatusBadRequest)
	}

	heartRateData.Id = user.(uuid.UUID)
	heartRateData.Data = heartRate.Data
	heartRateData.Notes = heartRate.Notes
	heartRateData.Date = heartRate.Date

	healthstatusresponse.HeartRateSuccessResponse(c, "Success Delete Data", heartRateData, "http://localhost:3000/api/user/health_status/heart_rate", http.StatusOK)
}

func CreateBloodPressure(c *gin.Context) {
	var body models.BloodPressure
	user, _ := c.Get("user")

	var bloodPressureData struct {
		DataSys float32
		DataDia float32
		Notes   string
		Date    string
	}

	if c.Bind(&body) != nil {
		healthstatusresponse.BloodPressureFailedResponse(c, "Failed to read body", bloodPressureData, "Create Blood Pressure", "http://localhost:3000/api/user/health_status/blood_pressure", http.StatusBadRequest)
		return
	}

	bloodPressureData.DataSys = body.DataSys
	bloodPressureData.DataDia = body.DataDia
	bloodPressureData.Notes = body.Notes
	bloodPressureData.Date = body.Date

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make([]string, len(validationErrors))
		for i, fieldError := range validationErrors {
			errorMessages[i] = getValidationErrorTagMessage(fieldError)
		}
		healthstatusresponse.BloodPressureFailedResponse(c, strings.Join(errorMessages, ", "), bloodPressureData, "Create Blood Pressure", "http://localhost:3000/api/user/health_status/blood_pressure", http.StatusBadRequest)
		return
	}

	newUUID := uuid.New()
	bloodPressure := models.BloodPressure{
		ID:      newUUID,
		UserID:  user.(uuid.UUID),
		DataSys: body.DataSys,
		DataDia: body.DataDia,
		Notes:   body.Notes,
		Date:    body.Date,
	}

	if err := initializers.DB.Create(&bloodPressure).Error; err != nil {
		healthstatusresponse.BloodPressureFailedResponse(c, strings.Title(err.Error()), bloodPressureData, "Create Blood Pressure", "http://localhost:3000/api/user/health_status/blood_pressure", http.StatusBadRequest)
	}

	healthstatusresponse.BloodPressureSuccessResponse(c, "Success Create Data", bloodPressureData, "http://localhost:3000/api/user/health_status/blood_pressure", http.StatusCreated)
}

func GetBloodPressure(c *gin.Context) {
	var bloodPressure []models.BloodPressure
	user, _ := c.Get("user")

	var bloodPressureData []struct {
		Id      uuid.UUID
		DataSys float32
		DataDia float32
		Notes   string
		Date    string
	}

	initializers.DB.Where("user_id = ?", user).Order("date asc").Find(&bloodPressure)

	if initializers.DB.Error != nil {
		healthstatusresponse.BloodPressureFailedResponse(c, "User Not Found", bloodPressureData, "Create Blood Pressure", "http://localhost:3000/api/user/health_status/blood_pressure", http.StatusBadRequest)
		return
	}

	for _, item := range bloodPressure {
		bloodPressureData = append(bloodPressureData, struct {
			Id      uuid.UUID
			DataSys float32
			DataDia float32
			Notes   string
			Date    string
		}{
			Id:      item.ID,
			DataSys: item.DataSys,
			DataDia: item.DataDia,
			Notes:   item.Notes,
			Date:    item.Date,
		})
	}

	if bloodPressureData == nil {
		healthstatusresponse.BloodPressureFailedResponse(c, "Data Not Found", bloodPressureData, "Create Blood Pressure", "http://localhost:3000/api/user/health_status/blood_pressure", http.StatusNotFound)
		return

	}

	healthstatusresponse.BloodPressureSuccessResponse(c, "Success Get Data", bloodPressureData, "http://localhost:3000/api/user/health_status/blood_pressure", http.StatusOK)
}

func UpdateBloodPressure(c *gin.Context) {
	var body models.BloodPressure
	bloodPressure_id := c.Param("blood_pressure_id")

	var bloodPressureData struct {
		Id      uuid.UUID
		DataSys float32
		DataDia float32
		Notes   string
		Date    string
	}

	user, _ := c.Get("user")

	if c.Bind(&body) != nil {
		healthstatusresponse.BloodPressureFailedResponse(c, "Failed to read body", bloodPressureData, "Create Blood Pressure", "http://localhost:3000/api/user/health_status/blood_pressure", http.StatusBadRequest)
		return
	}

	body.ID = uuid.MustParse(bloodPressure_id)
	body.UserID = user.(uuid.UUID)

	bloodPressureData.Id = user.(uuid.UUID)
	bloodPressureData.DataSys = body.DataSys
	bloodPressureData.DataDia = body.DataDia
	bloodPressureData.Notes = body.Notes
	bloodPressureData.Date = body.Date

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make([]string, len(validationErrors))
		for i, fieldError := range validationErrors {
			errorMessages[i] = getValidationErrorTagMessage(fieldError)
		}
		healthstatusresponse.BloodPressureFailedResponse(c, strings.Join(errorMessages, ", "), bloodPressureData, "Create Blood Pressure", "http://localhost:3000/api/user/health_status/blood_pressure", http.StatusBadRequest)
		return
	}

	var bloodPressure models.BloodPressure
	if err := initializers.DB.First(&bloodPressure, "user_id = ? AND ID = ?", user, bloodPressure_id).Error; err != nil {
		healthstatusresponse.BloodPressureFailedResponse(c, strings.Title(err.Error()), bloodPressureData, "Create Blood Pressure", "http://localhost:3000/api/user/health_status/blood_pressure", http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Save(&body).Error; err != nil {
		healthstatusresponse.BloodPressureFailedResponse(c, strings.Title(err.Error()), bloodPressureData, "Create Blood Pressure", "http://localhost:3000/api/user/health_status/blood_pressure", http.StatusBadRequest)
		return
	}

	healthstatusresponse.BloodPressureSuccessResponse(c, "Success Update Data", bloodPressureData, "http://localhost:3000/api/user/health_status/blood_pressure", http.StatusOK)
}

func DeleteBloodPressure(c *gin.Context) {
	bloodPressure_id := c.Param("blood_pressure_id")
	user, _ := c.Get("user")

	var bloodPressureData struct {
		Id      uuid.UUID
		DataSys float32
		DataDia float32
		Notes   string
		Date    string
	}

	var bloodPressure models.BloodPressure
	if err := initializers.DB.First(&bloodPressure, "user_id = ? AND ID = ?", user, bloodPressure_id).Error; err != nil {
		healthstatusresponse.BloodPressureFailedResponse(c, strings.Title(err.Error()), bloodPressureData, "Create Blood Pressure", "http://localhost:3000/api/user/health_status/blood_pressure", http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Unscoped().Delete(&bloodPressure).Error; err != nil {
		healthstatusresponse.BloodPressureFailedResponse(c, strings.Title(err.Error()), bloodPressureData, "Create Blood Pressure", "http://localhost:3000/api/user/health_status/blood_pressure", http.StatusBadRequest)
	}

	bloodPressureData.Id = user.(uuid.UUID)
	bloodPressureData.DataSys = bloodPressure.DataSys
	bloodPressureData.DataDia = bloodPressure.DataDia
	bloodPressureData.Notes = bloodPressure.Notes
	bloodPressureData.Date = bloodPressure.Date

	healthstatusresponse.BloodPressureSuccessResponse(c, "Success Delete Data", bloodPressureData, "http://localhost:3000/api/user/health_status/blood_pressure", http.StatusOK)
}

func CreateHematokrit(c *gin.Context) {
	var body models.Hematokrit
	user, _ := c.Get("user")

	var hematokrit_data struct {
		Data  float32
		Notes string
		Date  string
	}

	if c.Bind(&body) != nil {
		healthstatusresponse.HematokritFailedResponse(c, "Failed to read body", hematokrit_data, "Create Hematokrit", "http://localhost:3000/api/user/health_status/hematokrit", http.StatusBadRequest)
		return
	}

	hematokrit_data.Data = body.Data
	hematokrit_data.Notes = body.Notes
	hematokrit_data.Date = body.Date

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make([]string, len(validationErrors))
		for i, fieldError := range validationErrors {
			errorMessages[i] = getValidationErrorTagMessage(fieldError)
		}

		healthstatusresponse.HematokritFailedResponse(c, strings.Join(errorMessages, ", "), hematokrit_data, "Create Hematokrit", "http://localhost:3000/api/user/health_status/hematokrit", http.StatusBadRequest)
		return
	}

	newUUID := uuid.New()
	hematokrit := models.Hematokrit{
		ID:     newUUID,
		UserID: user.(uuid.UUID),
		Data:   body.Data,
		Notes:  body.Notes,
		Date:   body.Date,
	}

	if err := initializers.DB.Create(&hematokrit).Error; err != nil {
		healthstatusresponse.HematokritFailedResponse(c, strings.Title(err.Error()), hematokrit_data, "Create Hematokrit", "http://localhost:3000/api/user/health_status/hematokrit", http.StatusBadRequest)
	}

	healthstatusresponse.HematokritSuccessResponse(c, "Success Create Data", hematokrit_data, "http://localhost:3000/api/user/health_status/hematokrit", http.StatusCreated)
}

func GetHematokrit(c *gin.Context) {
	var hematokrit []models.Hematokrit
	user, _ := c.Get("user")

	var hematokrit_data []struct {
		Id    uuid.UUID
		Data  float32
		Notes string
		Date  string
	}

	initializers.DB.Where("user_id = ?", user).Order("date asc").Find(&hematokrit)

	if initializers.DB.Error != nil {
		healthstatusresponse.HematokritFailedResponse(c, "User Not Found", hematokrit_data, "Create Hematokrit", "http://localhost:3000/api/user/health_status/hematokrit", http.StatusBadRequest)
		return
	}

	for _, item := range hematokrit {
		hematokrit_data = append(hematokrit_data, struct {
			Id    uuid.UUID
			Data  float32
			Notes string
			Date  string
		}{
			Id:    item.ID,
			Data:  item.Data,
			Notes: item.Notes,
			Date:  item.Date,
		})
	}

	if hematokrit_data == nil {
		healthstatusresponse.HematokritFailedResponse(c, "Data Not Found", hematokrit_data, "Create Hematokrit", "http://localhost:3000/api/user/health_status/hematokrit", http.StatusNotFound)
		return

	}

	healthstatusresponse.HematokritSuccessResponse(c, "Success Get Data", hematokrit_data, "http://localhost:3000/api/user/health_status/hematokrit", http.StatusOK)
}

func UpdateHematokrit(c *gin.Context) {
	var body models.Hematokrit
	hematokrit_id := c.Param("hematokrit_id")

	var hematokrit_data struct {
		Id    uuid.UUID
		Data  float32
		Notes string
		Date  string
	}

	user, _ := c.Get("user")

	if c.Bind(&body) != nil {
		healthstatusresponse.HematokritFailedResponse(c, "Failed to read body", hematokrit_data, "Create Hematokrit", "http://localhost:3000/api/user/health_status/hematokrit", http.StatusBadRequest)
		return
	}

	body.ID = uuid.MustParse(hematokrit_id)
	body.UserID = user.(uuid.UUID)

	hematokrit_data.Id = user.(uuid.UUID)
	hematokrit_data.Data = body.Data
	hematokrit_data.Notes = body.Notes
	hematokrit_data.Date = body.Date

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make([]string, len(validationErrors))
		for i, fieldError := range validationErrors {
			errorMessages[i] = getValidationErrorTagMessage(fieldError)
		}

		healthstatusresponse.HematokritFailedResponse(c, strings.Join(errorMessages, ", "), hematokrit_data, "Create Hematokrit", "http://localhost:3000/api/user/health_status/hematokrit", http.StatusBadRequest)
		return
	}

	var hematokrit models.Hematokrit
	if err := initializers.DB.First(&hematokrit, "user_id = ? AND ID = ?", user, hematokrit_id).Error; err != nil {
		healthstatusresponse.HematokritFailedResponse(c, strings.Title(err.Error()), hematokrit_data, "Create Hematokrit", "http://localhost:3000/api/user/health_status/hematokrit", http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Save(&body).Error; err != nil {
		healthstatusresponse.HematokritFailedResponse(c, strings.Title(err.Error()), hematokrit_data, "Create Hematokrit", "http://localhost:3000/api/user/health_status/hematokrit", http.StatusBadRequest)
		return
	}

	healthstatusresponse.HematokritSuccessResponse(c, "Success Update Data", hematokrit_data, "http://localhost:3000/api/user/health_status/hematokrit", http.StatusOK)
}

func DeleteHematokrit(c *gin.Context) {
	hematokrit_id := c.Param("hematokrit_id")
	user, _ := c.Get("user")

	var hematokrit_data struct {
		Id    uuid.UUID
		Data  float32
		Notes string
		Date  string
	}

	var hematokrit models.Hematokrit
	if err := initializers.DB.First(&hematokrit, "user_id = ? AND ID = ?", user, hematokrit_id).Error; err != nil {
		healthstatusresponse.HematokritFailedResponse(c, strings.Title(err.Error()), hematokrit_data, "Create Hematokrit", "http://localhost:3000/api/user/health_status/hematokrit", http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Unscoped().Delete(&hematokrit).Error; err != nil {
		healthstatusresponse.HematokritFailedResponse(c, strings.Title(err.Error()), hematokrit_data, "Create Hematokrit", "http://localhost:3000/api/user/health_status/hematokrit", http.StatusBadRequest)
	}

	hematokrit_data.Id = user.(uuid.UUID)
	hematokrit_data.Data = hematokrit.Data
	hematokrit_data.Notes = hematokrit.Notes
	hematokrit_data.Date = hematokrit.Date

	healthstatusresponse.HematokritSuccessResponse(c, "Success Delete Data", hematokrit_data, "http://localhost:3000/api/user/health_status/hematokrit", http.StatusOK)
}

func CreateTrombosit(c *gin.Context) {
	var body models.Trombosit
	user, _ := c.Get("user")

	var trombosit_data struct {
		Data  float32
		Notes string
		Date  string
	}

	if c.Bind(&body) != nil {
		healthstatusresponse.TrombositFailedResponse(c, "Failed to read body", trombosit_data, "Create Trombosit", "http://localhost:3000/api/user/health_status/trombosit", http.StatusBadRequest)
		return
	}

	trombosit_data.Data = body.Data
	trombosit_data.Notes = body.Notes
	trombosit_data.Date = body.Date

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make([]string, len(validationErrors))
		for i, fieldError := range validationErrors {
			errorMessages[i] = getValidationErrorTagMessage(fieldError)
		}

		healthstatusresponse.TrombositFailedResponse(c, strings.Join(errorMessages, ", "), trombosit_data, "Create Trombosit", "http://localhost:3000/api/user/health_status/trombosit", http.StatusBadRequest)
		return
	}

	newUUID := uuid.New()
	trombosit := models.Trombosit{
		ID:     newUUID,
		UserID: user.(uuid.UUID),
		Data:   body.Data,
		Notes:  body.Notes,
		Date:   body.Date,
	}

	if err := initializers.DB.Create(&trombosit).Error; err != nil {
		healthstatusresponse.TrombositFailedResponse(c, strings.Title(err.Error()), trombosit_data, "Create Trombosit", "http://localhost:3000/api/user/health_status/trombosit", http.StatusBadRequest)
	}

	healthstatusresponse.TrombositSuccessResponse(c, "Success Create Data", trombosit_data, "http://localhost:3000/api/user/health_status/trombosit", http.StatusCreated)
}

func GetTrombosit(c *gin.Context) {
	var trombosit []models.Trombosit
	user, _ := c.Get("user")

	var trombosit_data []struct {
		Id    uuid.UUID
		Data  float32
		Notes string
		Date  string
	}

	initializers.DB.Where("user_id = ?", user).Order("date asc").Find(&trombosit)

	if initializers.DB.Error != nil {
		healthstatusresponse.TrombositFailedResponse(c, "User Not Found", trombosit_data, "Create Trombosit", "http://localhost:3000/api/user/health_status/trombosit", http.StatusBadRequest)
		return
	}

	for _, item := range trombosit {
		trombosit_data = append(trombosit_data, struct {
			Id    uuid.UUID
			Data  float32
			Notes string
			Date  string
		}{
			Id:    item.ID,
			Data:  item.Data,
			Notes: item.Notes,
			Date:  item.Date,
		})
	}

	if trombosit_data == nil {
		healthstatusresponse.TrombositFailedResponse(c, "Data Not Found", trombosit_data, "Create Trombosit", "http://localhost:3000/api/user/health_status/trombosit", http.StatusNotFound)
		return

	}

	healthstatusresponse.TrombositSuccessResponse(c, "Success Get Data", trombosit_data, "http://localhost:3000/api/user/health_status/trombosit", http.StatusOK)
}

func UpdateTrombosit(c *gin.Context) {
	var body models.Trombosit
	trombosit_id := c.Param("trombosit_id")

	var trombosit_data struct {
		Id    uuid.UUID
		Data  float32
		Notes string
		Date  string
	}

	user, _ := c.Get("user")

	if c.Bind(&body) != nil {
		healthstatusresponse.TrombositFailedResponse(c, "Failed to read body", trombosit_data, "Create Trombosit", "http://localhost:3000/api/user/health_status/trombosit", http.StatusBadRequest)
		return
	}

	body.ID = uuid.MustParse(trombosit_id)
	body.UserID = user.(uuid.UUID)

	trombosit_data.Id = user.(uuid.UUID)
	trombosit_data.Data = body.Data
	trombosit_data.Notes = body.Notes
	trombosit_data.Date = body.Date

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make([]string, len(validationErrors))
		for i, fieldError := range validationErrors {
			errorMessages[i] = getValidationErrorTagMessage(fieldError)
		}

		healthstatusresponse.TrombositFailedResponse(c, strings.Join(errorMessages, ", "), trombosit_data, "Create Trombosit", "http://localhost:3000/api/user/health_status/trombosit", http.StatusBadRequest)
		return
	}

	var trombosit models.Trombosit
	if err := initializers.DB.First(&trombosit, "user_id = ? AND ID = ?", user, trombosit_id).Error; err != nil {
		healthstatusresponse.TrombositFailedResponse(c, strings.Title(err.Error()), trombosit_data, "Create Trombosit", "http://localhost:3000/api/user/health_status/trombosit", http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Save(&body).Error; err != nil {
		healthstatusresponse.TrombositFailedResponse(c, strings.Title(err.Error()), trombosit_data, "Create Trombosit", "http://localhost:3000/api/user/health_status/trombosit", http.StatusBadRequest)
		return
	}

	healthstatusresponse.TrombositSuccessResponse(c, "Success Update Data", trombosit_data, "http://localhost:3000/api/user/health_status/trombosit", http.StatusOK)
}

func DeleteTrombosit(c *gin.Context) {
	trombosit_id := c.Param("trombosit_id")
	user, _ := c.Get("user")

	var trombosit_data struct {
		Id    uuid.UUID
		Data  float32
		Notes string
		Date  string
	}

	var trombosit models.Trombosit
	if err := initializers.DB.First(&trombosit, "user_id = ? AND ID = ?", user, trombosit_id).Error; err != nil {
		healthstatusresponse.TrombositFailedResponse(c, strings.Title(err.Error()), trombosit_data, "Create Trombosit", "http://localhost:3000/api/user/health_status/trombosit", http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Unscoped().Delete(&trombosit).Error; err != nil {
		healthstatusresponse.TrombositFailedResponse(c, strings.Title(err.Error()), trombosit_data, "Create Trombosit", "http://localhost:3000/api/user/health_status/trombosit", http.StatusBadRequest)
		return
	}

	trombosit_data.Id = user.(uuid.UUID)
	trombosit_data.Data = trombosit.Data
	trombosit_data.Notes = trombosit.Notes
	trombosit_data.Date = trombosit.Date

	healthstatusresponse.TrombositSuccessResponse(c, "Success Delete Data", trombosit_data, "http://localhost:3000/api/user/health_status/trombosit", http.StatusOK)
}

func GetBcrAblPatient(c *gin.Context) {
	if !ParseWebToken(c) {
		return
	}

	var bcr_abl_data []models.BCR_ABL
	var response []models.HealthStatusData

	query := `
        SELECT ba.*
        FROM bcr_abls ba
        INNER JOIN (
            SELECT user_id, MAX(date) AS max_date
            FROM bcr_abls
            GROUP BY user_id
        ) subquery ON ba.user_id = subquery.user_id AND ba.date = subquery.max_date
        ORDER BY ba.date DESC
    `

	result := initializers.DB.Raw(query).Scan(&bcr_abl_data)
	if result.Error != nil {
		healthstatusresponse.HealthStatusWebsiteFailedResponse(c, "Failed to Get BCR ABL Data", "", http.StatusInternalServerError)
		return
	}

	for _, item := range bcr_abl_data {
		var user models.User
		initializers.DB.First(&user, "ID = ?", item.UserID)
		response = append(response, models.HealthStatusData{
			ID:          item.ID,
			UserID:      item.UserID,
			Name:        user.Name,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			Data:        item.Data,
			Notes:       item.Notes,
			Date:        item.Date,
		})
	}

	healthstatusresponse.HealthStatusWebsiteSuccessResponse(c, "Success to Get BCR ABL Data", response, http.StatusOK)
}

func GetLeukocytesPatient(c *gin.Context) {
	if !ParseWebToken(c) {
		return
	}

	var leukocytes_data []models.Leukocytes
	var response []models.HealthStatusDataGender

	query := `
        SELECT ba.*
        FROM leukocytes ba
        INNER JOIN (
            SELECT user_id, MAX(date) AS max_date
            FROM leukocytes
            GROUP BY user_id
        ) subquery ON ba.user_id = subquery.user_id AND ba.date = subquery.max_date
        ORDER BY ba.date DESC
    `

	result := initializers.DB.Raw(query).Scan(&leukocytes_data)
	if result.Error != nil {
		healthstatusresponse.HealthStatusWebsiteFailedResponse(c, "Failed to Get Leukocytes Data", "", http.StatusInternalServerError)
		return
	}

	for _, item := range leukocytes_data {
		var user models.User
		initializers.DB.First(&user, "ID = ?", item.UserID)
		response = append(response, models.HealthStatusDataGender{
			ID:          item.ID,
			UserID:      item.UserID,
			Name:        user.Name,
			Gender:      user.Gender,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			Data:        item.Data,
			Notes:       item.Notes,
			Date:        item.Date,
		})
	}

	healthstatusresponse.HealthStatusWebsiteSuccessResponse(c, "Success to Get Leukocytes Data", response, http.StatusOK)
}

func GetPotentialHydrogenPatient(c *gin.Context) {
	if !ParseWebToken(c) {
		return
	}

	var potential_hydrogen_data []models.PotentialHydrogen
	var response []models.HealthStatusData

	query := `
        SELECT ba.*
        FROM potential_hydrogens ba
        INNER JOIN (
            SELECT user_id, MAX(date) AS max_date
            FROM potential_hydrogens
            GROUP BY user_id
        ) subquery ON ba.user_id = subquery.user_id AND ba.date = subquery.max_date
        ORDER BY ba.date DESC
    `

	result := initializers.DB.Raw(query).Scan(&potential_hydrogen_data)
	if result.Error != nil {
		healthstatusresponse.HealthStatusWebsiteFailedResponse(c, "Failed to Get Potential Hydrogen Data", "", http.StatusInternalServerError)
		return
	}

	for _, item := range potential_hydrogen_data {
		var user models.User
		initializers.DB.First(&user, "ID = ?", item.UserID)
		response = append(response, models.HealthStatusData{
			ID:          item.ID,
			UserID:      item.UserID,
			Name:        user.Name,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			Data:        item.Data,
			Notes:       item.Notes,
			Date:        item.Date,
		})
	}

	healthstatusresponse.HealthStatusWebsiteSuccessResponse(c, "Success to Get Potential Hydrogen Data", response, http.StatusOK)
}

func GetHemoglobinPatient(c *gin.Context) {
	if !ParseWebToken(c) {
		return
	}

	var hemoglobin_data []models.Hemoglobin
	var response []models.HealthStatusDataGender

	query := `
        SELECT ba.*
        FROM hemoglobins ba
        INNER JOIN (
            SELECT user_id, MAX(date) AS max_date
            FROM hemoglobins
            GROUP BY user_id
        ) subquery ON ba.user_id = subquery.user_id AND ba.date = subquery.max_date
        ORDER BY ba.date DESC
    `

	result := initializers.DB.Raw(query).Scan(&hemoglobin_data)
	if result.Error != nil {
		healthstatusresponse.HealthStatusWebsiteFailedResponse(c, "Failed to Get Hemoglobin Data", "", http.StatusInternalServerError)
		return
	}

	for _, item := range hemoglobin_data {
		var user models.User
		initializers.DB.First(&user, "ID = ?", item.UserID)
		response = append(response, models.HealthStatusDataGender{
			ID:          item.ID,
			UserID:      item.UserID,
			Name:        user.Name,
			Gender:      user.Gender,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			Data:        item.Data,
			Notes:       item.Notes,
			Date:        item.Date,
		})
	}

	healthstatusresponse.HealthStatusWebsiteSuccessResponse(c, "Success to Get Hemoglobin Data", response, http.StatusOK)
}

func GetHeartRatePatient(c *gin.Context) {
	if !ParseWebToken(c) {
		return
	}

	var heart_rate_data []models.HeartRate
	var response []models.HealthStatusData

	query := `
        SELECT ba.*
        FROM heart_rates ba
        INNER JOIN (
            SELECT user_id, MAX(date) AS max_date
            FROM heart_rates
            GROUP BY user_id
        ) subquery ON ba.user_id = subquery.user_id AND ba.date = subquery.max_date
        ORDER BY ba.date DESC
    `

	result := initializers.DB.Raw(query).Scan(&heart_rate_data)
	if result.Error != nil {
		healthstatusresponse.HealthStatusWebsiteFailedResponse(c, "Failed to Get Heart Rate Data", "", http.StatusInternalServerError)
		return
	}

	for _, item := range heart_rate_data {
		var user models.User
		initializers.DB.First(&user, "ID = ?", item.UserID)
		response = append(response, models.HealthStatusData{
			ID:          item.ID,
			UserID:      item.UserID,
			Name:        user.Name,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			Data:        item.Data,
			Notes:       item.Notes,
			Date:        item.Date,
		})
	}

	healthstatusresponse.HealthStatusWebsiteSuccessResponse(c, "Success to Get Heart Rate Data", response, http.StatusOK)
}

func GetBloodPressurePatient(c *gin.Context) {
	if !ParseWebToken(c) {
		return
	}

	var blood_pressure_data []models.BloodPressure
	var response []models.HealthStatusDataBloodPressure

	query := `
        SELECT ba.*
        FROM blood_pressures ba
        INNER JOIN (
            SELECT user_id, MAX(date) AS max_date
            FROM blood_pressures
            GROUP BY user_id
        ) subquery ON ba.user_id = subquery.user_id AND ba.date = subquery.max_date
        ORDER BY ba.date DESC
    `

	result := initializers.DB.Raw(query).Scan(&blood_pressure_data)
	if result.Error != nil {
		healthstatusresponse.HealthStatusWebsiteFailedResponse(c, "Failed to Get Blood Pressure Data", "", http.StatusInternalServerError)
		return
	}

	for _, item := range blood_pressure_data {
		var user models.User
		initializers.DB.First(&user, "ID = ?", item.UserID)
		response = append(response, models.HealthStatusDataBloodPressure{
			ID:          item.ID,
			UserID:      item.UserID,
			Name:        user.Name,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			DataSys:     item.DataSys,
			DataDia:     item.DataDia,
			Notes:       item.Notes,
			Date:        item.Date,
		})
	}

	healthstatusresponse.HealthStatusWebsiteSuccessResponse(c, "Success to Get Blood Pressure Data", response, http.StatusOK)
}

func GetHematokritPatient(c *gin.Context) {
	if !ParseWebToken(c) {
		return
	}

	var hematokrit_data []models.Hematokrit
	var response []models.HealthStatusData

	query := `
        SELECT ba.*
        FROM hematokrits ba
        INNER JOIN (
            SELECT user_id, MAX(date) AS max_date
            FROM hematokrits
            GROUP BY user_id
        ) subquery ON ba.user_id = subquery.user_id AND ba.date = subquery.max_date
        ORDER BY ba.date DESC
    `

	result := initializers.DB.Raw(query).Scan(&hematokrit_data)
	if result.Error != nil {
		healthstatusresponse.HealthStatusWebsiteFailedResponse(c, "Failed to Get Hematokrit Data", "", http.StatusInternalServerError)
		return
	}

	for _, item := range hematokrit_data {
		var user models.User
		initializers.DB.First(&user, "ID = ?", item.UserID)
		response = append(response, models.HealthStatusData{
			ID:          item.ID,
			UserID:      item.UserID,
			Name:        user.Name,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			Data:        item.Data,
			Notes:       item.Notes,
			Date:        item.Date,
		})
	}

	healthstatusresponse.HealthStatusWebsiteSuccessResponse(c, "Success to Get Hematokrit Data", response, http.StatusOK)
}

func GetTrombositPatient(c *gin.Context) {
	if !ParseWebToken(c) {
		return
	}

	var trombosit_data []models.Trombosit
	var response []models.HealthStatusData

	query := `
        SELECT ba.*
        FROM trombosits ba
        INNER JOIN (
            SELECT user_id, MAX(date) AS max_date
            FROM trombosits
            GROUP BY user_id
        ) subquery ON ba.user_id = subquery.user_id AND ba.date = subquery.max_date
        ORDER BY ba.date DESC
    `

	result := initializers.DB.Raw(query).Scan(&trombosit_data)
	if result.Error != nil {
		healthstatusresponse.HealthStatusWebsiteFailedResponse(c, "Failed to Get Trombosit Data", "", http.StatusInternalServerError)
		return
	}

	for _, item := range trombosit_data {
		var user models.User
		initializers.DB.First(&user, "ID = ?", item.UserID)
		response = append(response, models.HealthStatusData{
			ID:          item.ID,
			UserID:      item.UserID,
			Name:        user.Name,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			Data:        item.Data,
			Notes:       item.Notes,
			Date:        item.Date,
		})
	}

	healthstatusresponse.HealthStatusWebsiteSuccessResponse(c, "Success to Get Trombosit Data", response, http.StatusOK)
}
