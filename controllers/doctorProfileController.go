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
	r.GET("api/doctor/patient_request", middleware.RequireAuth, DoctorPatientRequest)
	r.PUT("api/doctor/patient_request/accept/:acceptance_id", middleware.RequireAuth, DoctorPatientAccept)
	r.PUT("api/doctor/patient_request/reject/:acceptance_id", middleware.RequireAuth, DoctorPatientReject)

	r.GET("api/doctor/patient/list", middleware.RequireAuth, DoctorPatientList)
	r.GET("api/doctor/patient/profile/:acceptance_id", middleware.RequireAuth, DoctorPatientProfile)
}

func DoctorCheck(c *gin.Context, doctor any) bool {
	var doctor_account models.Doctor
	if err := initializers.DB.First(&doctor_account, "id = ?", doctor).Error; err != nil {
		doctorresponse.CheckDoctorAccountFailedResponse(c)
		return false
	}
	return true
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

	data := models.UserData{
		ID:          patient_data.ID,
		Name:        patient_data.Name,
		Email:       patient_data.Email,
		Address:     patient_data.Address,
		Province:    patient_data.Province,
		District:    patient_data.District,
		SubDistrict: patient_data.SubDistrict,
		Village:     patient_data.Village,
		Gender:      patient_data.Gender,
		BirthDate:   patient_data.BirthDate,
		BloodGroup:  patient_data.BloodGroup,
		PhoneNumber: patient_data.PhoneNumber,
	}

	doctorresponse.DoctorPatientProfileSuccessResponse(c, "Success to Get Patient Data", data, http.StatusOK)
}
