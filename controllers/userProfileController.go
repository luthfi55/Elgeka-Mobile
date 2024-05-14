package controllers

import (
	"elgeka-mobile/initializers"
	"elgeka-mobile/middleware"
	"elgeka-mobile/models"
	userresponse "elgeka-mobile/response/UserResponse"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func UserProfileController(r *gin.Engine) {
	r.GET("api/user/profile", middleware.RequireAuth, ProfileData)
	r.PUT("api/user/profile/edit", middleware.RequireAuth, EditProfile)
	r.PUT("api/user/profile/information/edit", middleware.RequireAuth, EditUserInformation)
	r.POST("api/user/add/personal_doctor", middleware.RequireAuth, AddPersonalDoctor)
	r.GET("api/user/list/personal_doctor", middleware.RequireAuth, GetPersonalDoctor)
	r.GET("api/user/list/activate_doctor", middleware.RequireAuth, ListActivateDoctor)
	r.GET("api/user/list/website", ListUserWebsite)
}

func ProfileData(c *gin.Context) {
	user, _ := c.Get("user")

	var user_data models.User
	if err := initializers.DB.First(&user_data, "id = ?", user).Error; err != nil {
		userresponse.GetProfileFailedResponse(c, "Failed To Find User", "", "Get Profile", "http://localhost:3000/api/user/profile", http.StatusBadRequest)
		return
	}

	var patient_information models.UserInformation
	if err := initializers.DB.First(&patient_information, "user_id = ?", user).Error; err != nil {
		userresponse.GetProfileFailedResponse(c, "Failed To Find User Information", "", "Get Profile", "http://localhost:3000/api/user/profile", http.StatusBadRequest)
		return
	}

	profile_data := models.UserInformationData{
		ID:                user_data.ID,
		Name:              user_data.Name,
		Email:             user_data.Email,
		Address:           user_data.Address,
		Province:          user_data.Province,
		District:          user_data.District,
		SubDistrict:       user_data.SubDistrict,
		Village:           user_data.Village,
		Gender:            user_data.Gender,
		BirthDate:         user_data.BirthDate,
		BloodGroup:        user_data.BloodGroup,
		DiagnosisDate:     user_data.DiagnosisDate,
		PhoneNumber:       user_data.PhoneNumber,
		PcrLevel:          patient_information.PcrLevel,
		TherapyActive:     patient_information.TherapyActive,
		TreatmentFree:     patient_information.TreatmentFree,
		TreatmentFreeDate: patient_information.TreatmentFreeDate,
		MonitoringPlace:   patient_information.MonitoringPlace,
		PcrFrequent:       patient_information.PcrFrequent,
	}

	userresponse.GetProfileSuccessResponse(c, "Success Get Profile", profile_data, "http://localhost:3000/api/user/profile", http.StatusOK)
}

func EditProfile(c *gin.Context) {
	user, _ := c.Get("user")

	var body models.User

	if c.Bind(&body) != nil {
		userresponse.UpdateUserProfileFailedResponse(c, "Failed to read body", body, "Edit Profile", "http://localhost:3000/api/user/profile/edit", http.StatusBadRequest)
		return
	}

	var user_data models.User
	if err := initializers.DB.First(&user_data, "id = ?", user).Error; err != nil {
		userresponse.UpdateUserProfileFailedResponse(c, "Failed To Find User", body, "Edit Profile", "http://localhost:3000/api/user/profile/edit", http.StatusBadRequest)
		return
	}

	if body.Name != "" {
		user_data.Name = body.Name
	}

	if body.Address != "" {
		user_data.Address = body.Address
	}

	if body.Province != "" {
		user_data.Province = body.Province
	}

	if body.District != "" {
		user_data.District = body.District
	}

	if body.SubDistrict != "" {
		user_data.SubDistrict = body.SubDistrict
	}

	if body.Village != "" {
		user_data.Village = body.Village
	}

	if body.Gender != "" {
		if body.Gender == "male" || body.Gender == "female" {
			user_data.Gender = body.Gender
		}
	}

	if body.BirthDate != "" {
		user_data.BirthDate = body.BirthDate
	}

	if body.DiagnosisDate != "" {
		user_data.DiagnosisDate = body.DiagnosisDate
	}

	if body.BloodGroup != "" {
		if body.BloodGroup == "A" || body.BloodGroup == "B" || body.BloodGroup == "AB" || body.BloodGroup == "O" {
			user_data.BloodGroup = body.BloodGroup
		}
	}

	if body.Name == "" && body.Address == "" && body.Province == "" && body.District == "" && body.SubDistrict == "" && body.Village == "" && body.Gender == "" && body.BirthDate == "" && body.BloodGroup == "" && body.DiagnosisDate == "" {
		userresponse.UpdateUserProfileFailedResponse(c, "Body Can't Null", body, "Edit Profile", "http://localhost:3000/api/user/profile/edit", http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Save(&user_data).Error; err != nil {
		userresponse.UpdateUserProfileFailedResponse(c, "Failed Update User", body, "Edit Profile", "http://localhost:3000/api/user/profile/edit", http.StatusBadRequest)
		return
	}

	userresponse.UpdateUserProfileSuccessResponse(c, "Success Update User", user_data.ID, "http://localhost:3000/api/user/profile/edit", http.StatusOK)

}

func EditUserInformation(c *gin.Context) {
	user, _ := c.Get("user")

	var body models.UserInformation

	if c.Bind(&body) != nil {
		userresponse.UpdateUserInformationProfileFailedResponse(c, "Failed to read body", body, "Edit Profile", "http://localhost:3000/api/user/profile/edit", http.StatusBadRequest)
		return
	}

	var user_data models.UserInformation
	if err := initializers.DB.First(&user_data, "user_id = ?", user).Error; err != nil {
		userresponse.UpdateUserInformationProfileFailedResponse(c, "Failed To Find User Information", body, "Edit Profile", "http://localhost:3000/api/user/profile/edit", http.StatusBadRequest)
		return
	}

	if body.PcrLevel != "" {
		user_data.PcrLevel = body.PcrLevel
	}

	user_data.TherapyActive = body.TherapyActive
	user_data.TreatmentFree = body.TreatmentFree

	if body.TreatmentFreeDate != "" {
		user_data.TreatmentFreeDate = body.TreatmentFreeDate
	}

	if body.MonitoringPlace != "" {
		user_data.MonitoringPlace = body.MonitoringPlace
	}

	if body.PcrFrequent != "" {
		user_data.PcrFrequent = body.PcrFrequent
	}

	if body.PcrLevel == "" && body.TreatmentFreeDate == "" && body.MonitoringPlace == "" && body.PcrFrequent == "" {
		userresponse.UpdateUserInformationProfileFailedResponse(c, "Body Can't Null", body, "Edit Profile", "http://localhost:3000/api/user/profile/edit", http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Save(&user_data).Error; err != nil {
		userresponse.UpdateUserInformationProfileFailedResponse(c, "Failed Update User", body, "Edit Profile", "http://localhost:3000/api/user/profile/edit", http.StatusBadRequest)
		return
	}

	userresponse.UpdateUserProfileSuccessResponse(c, "Success Update User Information", user_data.ID, "http://localhost:3000/api/user/profile/edit", http.StatusOK)

}

func EditUserPassword(c *gin.Context) {
	user, _ := c.Get("user")

	var body models.UserInformation

	if c.Bind(&body) != nil {
		userresponse.UpdateUserInformationProfileFailedResponse(c, "Failed to read body", body, "Edit Profile", "http://localhost:3000/api/user/profile/edit", http.StatusBadRequest)
		return
	}

	var user_data models.UserInformation
	if err := initializers.DB.First(&user_data, "user_id = ?", user).Error; err != nil {
		userresponse.UpdateUserInformationProfileFailedResponse(c, "Failed To Find User Information", body, "Edit Profile", "http://localhost:3000/api/user/profile/edit", http.StatusBadRequest)
		return
	}

	if body.PcrLevel != "" {
		user_data.PcrLevel = body.PcrLevel
	}

	user_data.TherapyActive = body.TherapyActive
	user_data.TreatmentFree = body.TreatmentFree

	if body.TreatmentFreeDate != "" {
		user_data.TreatmentFreeDate = body.TreatmentFreeDate
	}

	if body.MonitoringPlace != "" {
		user_data.MonitoringPlace = body.MonitoringPlace
	}

	if body.PcrFrequent != "" {
		user_data.PcrFrequent = body.PcrFrequent
	}

	if body.PcrLevel == "" && body.TreatmentFreeDate == "" && body.MonitoringPlace == "" && body.PcrFrequent == "" {
		userresponse.UpdateUserInformationProfileFailedResponse(c, "Body Can't Null", body, "Edit Profile", "http://localhost:3000/api/user/profile/edit", http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Save(&user_data).Error; err != nil {
		userresponse.UpdateUserInformationProfileFailedResponse(c, "Failed Update User", body, "Edit Profile", "http://localhost:3000/api/user/profile/edit", http.StatusBadRequest)
		return
	}

	userresponse.UpdateUserProfileSuccessResponse(c, "Success Update User Information", user_data.ID, "http://localhost:3000/api/user/profile/edit", http.StatusOK)

}

func AddPersonalDoctor(c *gin.Context) {
	user, _ := c.Get("user")

	var body struct {
		DoctorID string `json:"DoctorID"`
	}

	if c.Bind(&body) != nil {
		userresponse.AddPersonalDoctorFailedResponse(c, "Failed to read body", body.DoctorID, "Add Personal Doctor", "http://localhost:3000/api/user/add/doctor", http.StatusBadRequest)
		return
	}

	var user_data models.User

	if err := initializers.DB.First(&user_data, "id = ?", user).Error; err != nil {
		userresponse.AddPersonalDoctorFailedResponse(c, "User Not Found", body.DoctorID, "Add Personal Doctor", "http://localhost:3000/api/user/add/doctor", http.StatusBadRequest)
		return
	}

	var doctor_data models.Doctor

	if err := initializers.DB.First(&doctor_data, "id = ?", body.DoctorID).Error; err != nil {
		userresponse.AddPersonalDoctorFailedResponse(c, "Doctor Not Found", body.DoctorID, "Add Personal Doctor", "http://localhost:3000/api/user/add/doctor", http.StatusBadRequest)
		return
	}

	doctorUUID, err := uuid.Parse(body.DoctorID)
	if err != nil {
		userresponse.AddPersonalDoctorFailedResponse(c, "Invalid Doctor ID", body.DoctorID, "Add Personal Doctor", "http://localhost:3000/api/user/add/doctor", http.StatusBadRequest)
		return
	}

	var patient_request []models.UserPersonalDoctor
	if err := initializers.DB.Where("user_id = ? AND request = ?", user, "Pending").First(&patient_request).Error; err != nil {
		//Pending Not Found
		if errors.Is(err, gorm.ErrRecordNotFound) {

			newUUID := uuid.New()
			// currentTime := time.Now()
			// startDate := currentTime.Format("2006-01-02")

			personal_doctor := models.UserPersonalDoctor{ID: newUUID, UserID: user.(uuid.UUID), DoctorID: doctorUUID, Request: "Pending"}

			if err := initializers.DB.Create(&personal_doctor).Error; err != nil {
				userresponse.AddPersonalDoctorFailedResponse(c, "Failed To Add Personal Doctor", body.DoctorID, "Add Personal Doctor", "http://localhost:3000/api/user/add/doctor", http.StatusBadRequest)
				return
			}

			userresponse.AddPersonalDoctorSuccessResponse(c, "Success Add Personal Doctor", body.DoctorID, "http://localhost:3000/api/user/add/doctor", http.StatusOK)
			return
		}
		// Another Error
	}
	//Pending found
	userresponse.AddPersonalDoctorFailedResponse(c, "Can't Add the Doctor. Please Wait Until The Doctor Accepts the Request First.", body.DoctorID, "Add Personal Doctor", "http://localhost:3000/api/user/add/doctor", http.StatusBadRequest)
}

func ListActivateDoctor(c *gin.Context) {
	var activate_doctor []struct {
		DoctorID   uuid.UUID
		DoctorName string
	}

	var activate_doctor_data []models.Doctor
	if err := initializers.DB.Where("is_active = ? ", true).Order("name asc").Find(&activate_doctor_data).Error; err != nil {
		userresponse.GetPersonalDoctorFailedResponse(c, "Failed To Get Active Doctor", "", "Get Activate Doctor", "http://localhost:3000/api/user/list/activate_doctor", http.StatusBadRequest)
		return
	}
	for _, item := range activate_doctor_data {
		var doctor models.Doctor
		initializers.DB.First(&doctor, "id = ?", item.ID)

		activate_doctor = append(activate_doctor, struct {
			DoctorID   uuid.UUID
			DoctorName string
		}{
			DoctorID:   doctor.ID,
			DoctorName: doctor.Name,
		})
	}

	userresponse.GetPersonalDoctorSuccessResponse(c, "Success Get Active Doctor", activate_doctor, "http://localhost:3000/api/user/list/activate_doctor", http.StatusOK)
}

func GetPersonalDoctor(c *gin.Context) {
	user, _ := c.Get("user")

	var personal_doctor []struct {
		DoctorName   string
		PhoneNumber  string
		StartDate    string
		EndDate      string
		DoctorStatus string
	}

	var personal_doctor_data []models.UserPersonalDoctor
	if err := initializers.DB.Where("user_id = ? AND (request = ? OR request = ?)", user, "Accepted", "Pending").Order("created_at desc").Find(&personal_doctor_data).Error; err != nil {
		userresponse.GetPersonalDoctorFailedResponse(c, "Failed To Get Personal Doctor", "", "Get Personal Doctor", "http://localhost:3000/api/user/list/personal_doctor", http.StatusBadRequest)
		return
	}
	for _, item := range personal_doctor_data {
		var doctor models.Doctor
		initializers.DB.First(&doctor, "id = ?", item.DoctorID)
		doctor_status := "Before"
		if item.StartDate == "" && item.EndDate == "" {
			doctor_status = "Waiting"
		} else if item.EndDate == "" {
			doctor_status = "Now"
		}
		personal_doctor = append(personal_doctor, struct {
			DoctorName   string
			PhoneNumber  string
			StartDate    string
			EndDate      string
			DoctorStatus string
		}{
			DoctorName:   doctor.Name,
			PhoneNumber:  doctor.PhoneNumber,
			StartDate:    item.StartDate,
			EndDate:      item.EndDate,
			DoctorStatus: doctor_status,
		})
	}

	userresponse.GetPersonalDoctorSuccessResponse(c, "Success Get Personal Doctor", personal_doctor, "http://localhost:3000/api/user/list/personal_doctor", http.StatusOK)
}

func ListUserWebsite(c *gin.Context) {
	if !ParseWebToken(c) {
		return
	}

	var user []models.User
	result := initializers.DB.Where("is_active = ?", true).Find(&user)
	if result.Error != nil {
		activationLink := "http://localhost:3000/api/user/list/website"
		userresponse.ListUserWebsiteFailedResponse(c, "Failed to Get Patient List", "", "Patient List", activationLink, http.StatusInternalServerError)
		return
	}

	var user_data []models.UserDataAge

	for _, item := range user {
		birthdate, err := time.Parse("2006-01-02", item.BirthDate)
		if err != nil {
			continue
		}
		age := time.Now().Year() - birthdate.Year()
		user_data = append(user_data, models.UserDataAge{
			ID:          item.ID,
			Name:        item.Name,
			Email:       item.Email,
			Gender:      item.Gender,
			BirthDate:   item.BirthDate,
			Age:         age,
			BloodGroup:  item.BloodGroup,
			PhoneNumber: item.PhoneNumber,
			Province:    item.Province,
			District:    item.District,
			SubDistrict: item.SubDistrict,
			Village:     item.Village,
			Address:     item.Address,
		})
	}

	userresponse.ListUserWebsiteSuccessResponse(c, "Success to Get Patient List", user_data, http.StatusOK)
}
