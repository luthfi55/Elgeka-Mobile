package controllers

import (
	"elgeka-mobile/initializers"
	"elgeka-mobile/middleware"
	"elgeka-mobile/models"
	doctorresponse "elgeka-mobile/response/DoctorResponse"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func DoctorProfileController(r *gin.Engine) {
	r.GET("api/doctor/profile", middleware.RequireAuth, DoctorProfile)
	r.PUT("api/doctor/profile/edit", middleware.RequireAuth, EditDoctorProfile)

	r.GET("api/doctor/patient_request", middleware.RequireAuth, DoctorPatientRequest)
	r.PUT("api/doctor/patient_request/accept/:acceptance_id", middleware.RequireAuth, DoctorPatientAccept)
	r.PUT("api/doctor/patient_request/reject/:acceptance_id", middleware.RequireAuth, DoctorPatientReject)

	r.GET("api/doctor/patient/list", middleware.RequireAuth, DoctorPatientList)
	r.GET("api/doctor/patient/profile/:acceptance_id", middleware.RequireAuth, DoctorPatientProfile)
	r.GET("api/doctor/patient/health_status/:acceptance_id", middleware.RequireAuth, DoctorPatientHealthStatus)
	r.GET("api/doctor/patient/medicine/:acceptance_id", middleware.RequireAuth, DoctorPatientMedicine)

	r.GET("api/doctor/list/website", ListDoctorWebsite)
	r.GET("api/doctor/list_patient/website", ListPatientDoctorWebsite)
}

func DoctorCheck(c *gin.Context, doctor any) bool {
	var doctor_account models.Doctor
	if err := initializers.DB.First(&doctor_account, "id = ?", doctor).Error; err != nil {
		doctorresponse.CheckDoctorAccountFailedResponse(c)
		return false
	}
	return true
}

func DoctorProfile(c *gin.Context) {
	doctor, _ := c.Get("doctor")

	if !DoctorCheck(c, doctor) {
		return
	}

	var doctor_account models.Doctor
	if err := initializers.DB.First(&doctor_account, "id = ?", doctor).Error; err != nil {
		doctorresponse.CheckDoctorAccountFailedResponse(c)
	}

	var doctor_data models.DoctorProfile

	doctor_data.ID = doctor_account.ID
	doctor_data.Name = doctor_account.Name
	doctor_data.PhoneNumber = doctor_account.PhoneNumber
	doctor_data.Email = doctor_account.Email
	doctor_data.Gender = doctor_account.Gender
	doctor_data.HospitalName = doctor_account.HospitalName
	doctor_data.PolyName = doctor_account.PolyName

	doctorresponse.GetDoctorProfileSuccessResponse(c, "Success to Get Doctor Profile Data", doctor_data, http.StatusOK)
}

func EditDoctorProfile(c *gin.Context) {
	doctor, _ := c.Get("doctor")

	if !DoctorCheck(c, doctor) {
		return
	}

	var body models.Doctor

	if c.Bind(&body) != nil {
		doctorresponse.UpdateDoctorProfileFailedResponse(c, "Failed to read body", "", http.StatusBadRequest)
		return
	}

	var doctor_account models.Doctor
	if err := initializers.DB.First(&doctor_account, "id = ?", doctor).Error; err != nil {
		doctorresponse.CheckDoctorAccountFailedResponse(c)
	}

	if body.Name != "" {
		doctor_account.Name = body.Name
	}

	if body.PhoneNumber != "" {
		doctor_account.PhoneNumber = body.PhoneNumber
	}

	if body.Gender != "" {
		doctor_account.Gender = body.Gender
	}

	if body.PolyName != "" {
		doctor_account.PolyName = body.PolyName
	}

	if body.HospitalName != "" {
		doctor_account.HospitalName = body.HospitalName
	}

	if body.Name == "" && body.PhoneNumber == "" && body.Gender == "" && body.PolyName == "" && body.HospitalName == "" {
		doctorresponse.UpdateDoctorProfileFailedResponse(c, "Body Can't Null", "", http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Save(&doctor_account).Error; err != nil {
		doctorresponse.UpdateDoctorProfileFailedResponse(c, "Failed Update Doctor", "", http.StatusBadRequest)
		return
	}

	var doctor_data models.DoctorProfile

	doctor_data.ID = doctor_account.ID
	doctor_data.Name = doctor_account.Name
	doctor_data.PhoneNumber = doctor_account.PhoneNumber
	doctor_data.Email = doctor_account.Email
	doctor_data.Gender = doctor_account.Gender
	doctor_data.HospitalName = doctor_account.HospitalName
	doctor_data.PolyName = doctor_account.PolyName

	doctorresponse.UpdateDoctorProfileSuccessResponse(c, "Success to Update Doctor Profile Data", doctor_data, http.StatusOK)
}

func DoctorPatientRequest(c *gin.Context) {
	doctor, _ := c.Get("doctor")

	if !DoctorCheck(c, doctor) {
		return
	}

	var patient_request []models.UserPersonalDoctor
	if err := initializers.DB.Where("doctor_id = ? AND request = ?", doctor, "Pending").Order("created_at desc").Find(&patient_request).Error; err != nil {
		doctorresponse.ListAcceptancePatientFailedResponse(c, "Failed to Get List Acceptance Patient", "", http.StatusBadRequest)
		return
	}
	var patient []struct {
		AcceptanceID uuid.UUID
		PatientName  string
		PhoneNumber  string
	}
	for _, item := range patient_request {
		var user models.User
		initializers.DB.First(&user, "id = ?", item.UserID)

		patient = append(patient, struct {
			AcceptanceID uuid.UUID
			PatientName  string
			PhoneNumber  string
		}{
			AcceptanceID: item.ID,
			PatientName:  user.Name,
			PhoneNumber:  user.PhoneNumber,
		})
	}

	doctorresponse.ListAcceptancePatientSuccessResponse(c, "Success Get List Acceptance Patient", patient, http.StatusOK)
}

func DoctorPatientAccept(c *gin.Context) {
	doctor, _ := c.Get("doctor")
	acceptanceID := c.Param("acceptance_id")

	if !DoctorCheck(c, doctor) {
		return
	}

	var acceptance models.UserPersonalDoctor
	result := initializers.DB.First(&acceptance, "id = ? AND doctor_id = ? AND request = ?", acceptanceID, doctor, "Pending")
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			doctorresponse.DoctorPatientAcceptFailedResponse(c, "Acceptance Data Not Found", acceptanceID, "List Patient Acceptance", "http://localhost:3000/api/doctor/patient_request/accept/:acceptance_id", http.StatusNotFound)

			return
		} else {
			doctorresponse.DoctorPatientAcceptFailedResponse(c, "Database Error", acceptanceID, "List Patient Acceptance", "http://localhost:3000/api/doctor/patient_request/accept/:acceptance_id", http.StatusInternalServerError)
			return
		}
	}

	var personal_doctor_data models.UserPersonalDoctor
	if err := initializers.DB.First(&personal_doctor_data, "user_id = ? AND end_date = ? AND request = ?", acceptance.UserID, "", "Accepted").Error; err == nil {
		personal_doctor_data.EndDate = time.Now().Format("2006-01-02")
		if err := initializers.DB.Save(&personal_doctor_data).Error; err != nil {
			doctorresponse.DoctorPatientAcceptFailedResponse(c, "Failed Update Latest Personal Doctor End Date", acceptanceID, "List Patient Acceptance", "http://localhost:3000/api/doctor/patient_request/accept/:acceptance_id", http.StatusBadRequest)
			return
		}
	}

	currentTime := time.Now()
	startDate := currentTime.Format("2006-01-02")
	acceptance.Request = "Accepted"
	acceptance.StartDate = startDate

	if err := initializers.DB.Save(&acceptance).Error; err != nil {
		doctorresponse.DoctorPatientAcceptFailedResponse(c, "Failed to Accept Patient", acceptanceID, "List Patient Acceptance", "http://localhost:3000/api/doctor/patient_request/accept/:acceptance_id", http.StatusBadRequest)
		return
	}

	doctorresponse.DoctorPatientAcceptSuccessResponse(c, "Success to Accept Patient", acceptance, http.StatusOK)
}

func DoctorPatientReject(c *gin.Context) {
	doctor, _ := c.Get("doctor")
	acceptanceID := c.Param("acceptance_id")

	if !DoctorCheck(c, doctor) {
		return
	}

	var acceptance models.UserPersonalDoctor
	result := initializers.DB.First(&acceptance, "id = ? AND doctor_id = ? AND request = ?", acceptanceID, doctor, "Pending")
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			doctorresponse.DoctorPatientRejectFailedResponse(c, "Acceptance Data Not Found", acceptanceID, "List Patient Acceptance", "http://localhost:3000/api/doctor/patient_request/accept/:acceptance_id", http.StatusNotFound)
			return
		} else {
			doctorresponse.DoctorPatientRejectFailedResponse(c, "Database Error", acceptanceID, "List Patient Acceptance", "http://localhost:3000/api/doctor/patient_request/accept/:acceptance_id", http.StatusInternalServerError)
			return
		}
	}

	acceptance.Request = "Rejected"

	currentTime := time.Now()
	endDate := currentTime.Format("2006-01-02")
	acceptance.EndDate = endDate

	if err := initializers.DB.Save(&acceptance).Error; err != nil {
		doctorresponse.DoctorPatientRejectFailedResponse(c, "Failed to Reject Patient", acceptanceID, "List Patient Acceptance", "http://localhost:3000/api/doctor/patient_request/accept/:acceptance_id", http.StatusBadRequest)
		return
	}

	doctorresponse.DoctorPatientRejectSuccessResponse(c, "Success to Reject Patient", acceptance, http.StatusOK)
}

func DoctorPatientList(c *gin.Context) {
	doctor, _ := c.Get("doctor")

	if !DoctorCheck(c, doctor) {
		return
	}

	var patient_request []models.UserPersonalDoctor
	if err := initializers.DB.Where("doctor_id = ? AND request = ? AND end_date = ?", doctor, "Accepted", "").Order("created_at desc").Find(&patient_request).Error; err != nil {
		doctorresponse.ListDoctorPatientFailedResponse(c, "Failed to Get List Patient", "", http.StatusBadRequest)
		return
	}
	var patient []struct {
		AcceptanceID uuid.UUID
		UserID       uuid.UUID
		PatientName  string
		PhoneNumber  string
	}
	for _, item := range patient_request {
		var user models.User
		initializers.DB.First(&user, "id = ?", item.UserID)

		patient = append(patient, struct {
			AcceptanceID uuid.UUID
			UserID       uuid.UUID
			PatientName  string
			PhoneNumber  string
		}{
			AcceptanceID: item.ID,
			UserID:       user.ID,
			PatientName:  user.Name,
			PhoneNumber:  user.PhoneNumber,
		})
	}
	doctorresponse.ListDoctorPatientSuccessResponse(c, "Success to Get List Patient", patient, http.StatusOK)
}

func DoctorPatientProfile(c *gin.Context) {
	doctor, _ := c.Get("doctor")
	acceptanceID := c.Param("acceptance_id")

	if !DoctorCheck(c, doctor) {
		return
	}

	var patient_profile models.UserPersonalDoctor
	if err := initializers.DB.First(&patient_profile, "id = ? AND doctor_id = ? AND request = ? AND end_date = ?", acceptanceID, doctor, "Accepted", "").Error; err != nil {
		doctorresponse.DoctorPatientProfileFailedResponse(c, "Patient Profile Not Found", "", http.StatusNotFound)
		return
	}

	var patient_data models.User
	if err := initializers.DB.First(&patient_data, "id = ?", patient_profile.UserID).Error; err != nil {
		doctorresponse.DoctorPatientProfileFailedResponse(c, "Failed to Get Patient Data", "", http.StatusNotFound)
		return
	}

	var patient_information models.UserInformation
	if err := initializers.DB.First(&patient_information, "user_id = ?", patient_profile.UserID).Error; err != nil {
		doctorresponse.DoctorPatientProfileFailedResponse(c, "Failed to Get Patient Information Data", "", http.StatusNotFound)
		return
	}

	data := models.UserInformationData{
		ID:                patient_data.ID,
		Name:              patient_data.Name,
		Email:             patient_data.Email,
		Address:           patient_data.Address,
		Province:          patient_data.Province,
		District:          patient_data.District,
		SubDistrict:       patient_data.SubDistrict,
		Village:           patient_data.Village,
		Gender:            patient_data.Gender,
		BirthDate:         patient_data.BirthDate,
		BloodGroup:        patient_data.BloodGroup,
		PhoneNumber:       patient_data.PhoneNumber,
		PcrLevel:          patient_information.PcrLevel,
		TherapyActive:     patient_information.TherapyActive,
		TreatmentFree:     patient_information.TreatmentFree,
		TreatmentFreeDate: patient_information.TreatmentFreeDate,
		MonitoringPlace:   patient_information.MonitoringPlace,
		PcrFrequent:       patient_information.PcrFrequent,
	}

	doctorresponse.DoctorPatientProfileSuccessResponse(c, "Success to Get Patient Data", data, http.StatusOK)
}

func DoctorPatientHealthStatus(c *gin.Context) {
	doctor, _ := c.Get("doctor")
	acceptanceID := c.Param("acceptance_id")

	if !DoctorCheck(c, doctor) {
		return
	}

	var patient_profile models.UserPersonalDoctor
	if err := initializers.DB.First(&patient_profile, "id = ? AND doctor_id = ? AND request = ? AND end_date = ?", acceptanceID, doctor, "Accepted", "").Error; err != nil {
		doctorresponse.DoctorPatientHealthStatusFailedResponse(c, "Patient Profile Not Found", "", http.StatusNotFound)
		return
	}

	var patient_data models.User
	if err := initializers.DB.First(&patient_data, "id = ?", patient_profile.UserID).Error; err != nil {
		doctorresponse.DoctorPatientHealthStatusFailedResponse(c, "Failed to Get Patient Data", "", http.StatusNotFound)
		return
	}

	var bcr_abl []models.BCR_ABL
	if err := initializers.DB.Where("user_id = ?", patient_profile.UserID).Order("date asc").Find(&bcr_abl).Error; err != nil {
		doctorresponse.DoctorPatientHealthStatusFailedResponse(c, "Failed to Get BCR ABL Data", "", http.StatusNotFound)
		return
	}

	var bcr_abl_data []struct {
		Id    uuid.UUID
		Data  float32
		Notes string
		Date  string
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

	var leukocytes []models.Leukocytes
	if err := initializers.DB.Where("user_id = ?", patient_profile.UserID).Order("date asc").Find(&leukocytes).Error; err != nil {
		doctorresponse.DoctorPatientHealthStatusFailedResponse(c, "Failed to Get Leukocytes Data", "", http.StatusNotFound)
		return
	}

	var leukocytes_data []struct {
		Id    uuid.UUID
		Data  float32
		Notes string
		Date  string
	}

	for _, item := range leukocytes {
		leukocytes_data = append(leukocytes_data, struct {
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

	var potential_hydrogen []models.PotentialHydrogen
	if err := initializers.DB.Where("user_id = ?", patient_profile.UserID).Order("date asc").Find(&potential_hydrogen).Error; err != nil {
		doctorresponse.DoctorPatientHealthStatusFailedResponse(c, "Failed to Get Potential Hydrogen Data", "", http.StatusNotFound)
		return
	}

	var potential_hydrogen_data []struct {
		Id    uuid.UUID
		Data  float32
		Notes string
		Date  string
	}

	for _, item := range potential_hydrogen {
		potential_hydrogen_data = append(potential_hydrogen_data, struct {
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

	var hemoglobin []models.Hemoglobin
	if err := initializers.DB.Where("user_id = ?", patient_profile.UserID).Order("date asc").Find(&hemoglobin).Error; err != nil {
		doctorresponse.DoctorPatientHealthStatusFailedResponse(c, "Failed to Get Hemoglobin Data", "", http.StatusNotFound)
		return
	}

	var hemoglobin_data []struct {
		Id    uuid.UUID
		Data  float32
		Notes string
		Date  string
	}

	for _, item := range hemoglobin {
		hemoglobin_data = append(hemoglobin_data, struct {
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

	var blood_pressure []models.BloodPressure
	if err := initializers.DB.Where("user_id = ?", patient_profile.UserID).Order("date asc").Find(&blood_pressure).Error; err != nil {
		doctorresponse.DoctorPatientHealthStatusFailedResponse(c, "Failed to Get Blood Pressure Data", "", http.StatusNotFound)
		return
	}

	var blood_pressure_data []struct {
		Id      uuid.UUID
		DataSys float32
		DataDia float32
		Notes   string
		Date    string
	}

	for _, item := range blood_pressure {
		blood_pressure_data = append(blood_pressure_data, struct {
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

	var heart_rate []models.HeartRate
	if err := initializers.DB.Where("user_id = ?", patient_profile.UserID).Order("date asc").Find(&heart_rate).Error; err != nil {
		doctorresponse.DoctorPatientHealthStatusFailedResponse(c, "Failed to Get Heart Rate Data", "", http.StatusNotFound)
		return
	}

	var heart_rate_data []struct {
		Id    uuid.UUID
		Data  float32
		Notes string
		Date  string
	}

	for _, item := range heart_rate {
		heart_rate_data = append(heart_rate_data, struct {
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

	var health_status []struct {
		Name              string
		BCR_ABL           interface{}
		Leukocytes        interface{}
		PotentialHydrogen interface{}
		Hemoglobin        interface{}
		BloodPressure     interface{}
		HeartRate         interface{}
	}

	health_status = append(health_status, struct {
		Name              string
		BCR_ABL           interface{}
		Leukocytes        interface{}
		PotentialHydrogen interface{}
		Hemoglobin        interface{}
		BloodPressure     interface{}
		HeartRate         interface{}
	}{
		Name:              patient_data.Name,
		BCR_ABL:           bcr_abl_data,
		Leukocytes:        leukocytes_data,
		PotentialHydrogen: potential_hydrogen_data,
		Hemoglobin:        hemoglobin_data,
		BloodPressure:     blood_pressure_data,
		HeartRate:         heart_rate_data,
	})

	doctorresponse.DoctorPatientHealthStatusSuccessResponse(c, "Success to Get Health Status Data", health_status, http.StatusOK)
}

func ListDoctorWebsite(c *gin.Context) {
	if !ParseWebToken(c) {
		return
	}

	var doctor []models.Doctor

	result := initializers.DB.Where("is_active = ?", true).Find(&doctor)
	if result.Error != nil {
		doctorresponse.ListDoctorWebsiteFailedResponse(c, "Failed to Get Doctor List", "", http.StatusInternalServerError)
		return
	}

	var doctor_data []models.DoctorProfile

	for _, item := range doctor {
		doctor_data = append(doctor_data, models.DoctorProfile{
			ID:           item.ID,
			Name:         item.Name,
			PhoneNumber:  item.PhoneNumber,
			Email:        item.Email,
			Gender:       item.Gender,
			PolyName:     item.PolyName,
			HospitalName: item.HospitalName,
		})
	}

	doctorresponse.ListDoctorWebsiteSuccessResponse(c, "Success to Get Doctor List", doctor_data, http.StatusOK)
}

func ListPatientDoctorWebsite(c *gin.Context) {
	if !ParseWebToken(c) {
		return
	}

	var doctor []models.Doctor

	result := initializers.DB.Where("is_active = ?", true).Find(&doctor)
	if result.Error != nil {
		doctorresponse.ListPatientDoctorWebsiteFailedResponse(c, "Failed to Get Patient Doctor List", "", http.StatusInternalServerError)
		return
	}

	var doctor_data []models.DoctorPatientData

	for _, item := range doctor {
		var patient []models.UserPersonalDoctor
		var patient_list []models.UserPersonalDoctorData

		initializers.DB.Where("doctor_id = ? AND request = ? AND end_date = ?", item.ID, "Accepted", "").Find(&patient)
		for _, second_item := range patient {
			var patient_profile models.User
			initializers.DB.First(&patient_profile, "id = ?", second_item.UserID)
			patient_list = append(patient_list, models.UserPersonalDoctorData{
				UserID:        patient_profile.ID,
				Name:          patient_profile.Name,
				Email:         patient_profile.Email,
				PhoneNumber:   patient_profile.PhoneNumber,
				Gender:        patient_profile.Gender,
				BloodGroup:    patient_profile.BloodGroup,
				DiagnosisDate: patient_profile.DiagnosisDate,
			})
		}

		doctor_data = append(doctor_data, models.DoctorPatientData{
			ID:          item.ID,
			DoctorName:  item.Name,
			PatientData: patient_list,
		})

	}

	doctorresponse.ListPatientDoctorWebsiteSuccessResponse(c, "Success to Get Patient Doctor List", doctor_data, http.StatusOK)
}

func DoctorPatientMedicine(c *gin.Context) {
	doctor, _ := c.Get("doctor")
	acceptanceID := c.Param("acceptance_id")

	if !DoctorCheck(c, doctor) {
		return
	}

	var patient_profile models.UserPersonalDoctor
	if err := initializers.DB.First(&patient_profile, "id = ? AND doctor_id = ? AND request = ? AND end_date = ?", acceptanceID, doctor, "Accepted", "").Error; err != nil {
		doctorresponse.DoctorPatientHealthStatusFailedResponse(c, "Patient Profile Not Found", "", http.StatusNotFound)
		return
	}

	var patient_data models.User
	if err := initializers.DB.First(&patient_data, "id = ?", patient_profile.UserID).Error; err != nil {
		doctorresponse.DoctorPatientHealthStatusFailedResponse(c, "Failed to Get Patient Data", "", http.StatusNotFound)
		return
	}

	var medicine_data []models.Medicine
	if err := initializers.DB.Where("user_id = ?", patient_data.ID).Find(&medicine_data).Error; err != nil {
		doctorresponse.DoctorPatientHealthStatusFailedResponse(c, "Failed to Get Medicine Data", "", http.StatusNotFound)
		return
	}

	var medicine models.MedicineDataResponse

	var medicine_list []models.MedicineData

	for _, item := range medicine_data {
		medicine_list = append(medicine_list, models.MedicineData{
			ID:     item.ID,
			Name:   item.Name,
			Dosage: item.Dosage,
			Stock:  item.Stock,
		})
	}

	medicine.UserID = patient_data.ID
	medicine.Name = patient_data.Name
	medicine.Medicine = medicine_list

	doctorresponse.DoctorPatientHealthStatusSuccessResponse(c, "Success to Get Medicine Data", medicine, http.StatusOK)
}
