package controllers

import (
	"elgeka-mobile/initializers"
	"elgeka-mobile/middleware"
	"elgeka-mobile/models"
	userresponse "elgeka-mobile/response/UserResponse"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UserProfileController(r *gin.Engine) {
	r.GET("api/user/profile", middleware.RequireAuth, ProfileData)
	r.PUT("api/user/profile/edit", middleware.RequireAuth, EditProfile)
	r.POST("api/user/add/personal_doctor", middleware.RequireAuth, AddPersonalDoctor)
	r.GET("api/user/list/personal_doctor", middleware.RequireAuth, GetPersonalDoctor)
	r.GET("api/user/list/activate_doctor", middleware.RequireAuth, ListActivateDoctor)
}

func ProfileData(c *gin.Context) {
	user, _ := c.Get("user")

	var user_data models.User
	if err := initializers.DB.First(&user_data, "id = ?", user).Error; err != nil {
		userresponse.GetProfileFailedResponse(c, "Failed To Find User", "", "Get Profile", "http://localhost:3000/api/user/profile", http.StatusBadRequest)
		return
	}

	profile_data := models.UserData{
		ID:          user_data.ID,
		Name:        user_data.Name,
		Email:       user_data.Email,
		Address:     user_data.Address,
		Gender:      user_data.Gender,
		BirthDate:   user_data.BirthDate,
		BloodGroup:  user_data.BloodGroup,
		PhoneNumber: user_data.PhoneNumber,
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

	if body.Gender != "" {
		if body.Gender == "male" || body.Gender == "female" {
			user_data.Gender = body.Gender
		}
	}

	if body.BirthDate != "" {
		user_data.BirthDate = body.BirthDate
	}

	if body.BloodGroup != "" {
		if body.BloodGroup == "A" || body.BloodGroup == "B" || body.BloodGroup == "AB" || body.BloodGroup == "O" {
			user_data.BloodGroup = body.BloodGroup
		}
	}

	if body.Name == "" && body.Address == "" && body.Gender == "" && body.BirthDate == "" && body.BloodGroup == "" {
		userresponse.UpdateUserProfileFailedResponse(c, "Body Can't Null", body, "Edit Profile", "http://localhost:3000/api/user/profile/edit", http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Save(&user_data).Error; err != nil {
		userresponse.UpdateUserProfileFailedResponse(c, "Failed Update User", body, "Edit Profile", "http://localhost:3000/api/user/profile/edit", http.StatusBadRequest)
		return
	}

	userresponse.UpdateUserProfileSuccessResponse(c, "Success Update User", user_data.ID, "http://localhost:3000/api/user/profile/edit", http.StatusOK)

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

	personal_doctor_data := models.UserPersonalDoctor{}
	if err := initializers.DB.First(&personal_doctor_data, "user_id = ? AND end_date = ?", user, "").Error; err == nil {
		personal_doctor_data.EndDate = time.Now().Format("2006-01-02")
		if err := initializers.DB.Save(&personal_doctor_data).Error; err != nil {
			userresponse.AddPersonalDoctorFailedResponse(c, "Failed To Update Latest Personal Doctor End Date", body.DoctorID, "Add Personal Doctor", "http://localhost:3000/api/user/add/doctor", http.StatusBadRequest)
			return
		}
	}

	newUUID := uuid.New()
	currentTime := time.Now()
	startDate := currentTime.Format("2006-01-02")

	personal_doctor := models.UserPersonalDoctor{ID: newUUID, UserID: user.(uuid.UUID), DoctorID: doctorUUID, StartDate: startDate}

	if err := initializers.DB.Create(&personal_doctor).Error; err != nil {
		userresponse.AddPersonalDoctorFailedResponse(c, "Failed To Add Personal Doctor", body.DoctorID, "Add Personal Doctor", "http://localhost:3000/api/user/add/doctor", http.StatusBadRequest)
		return
	}

	userresponse.AddPersonalDoctorSuccessResponse(c, "Success Add Personal Doctor", body.DoctorID, "http://localhost:3000/api/user/add/doctor", http.StatusOK)

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
		DoctorName string
		StartDate  string
		EndDate    string
	}

	var personal_doctor_data []models.UserPersonalDoctor
	if err := initializers.DB.Where("user_id = ?", user).Order("start_date asc").Find(&personal_doctor_data).Error; err != nil {
		userresponse.GetPersonalDoctorFailedResponse(c, "Failed To Get Personal Doctor", "", "Get Personal Doctor", "http://localhost:3000/api/user/list/personal_doctor", http.StatusBadRequest)
		return
	}
	for _, item := range personal_doctor_data {
		var doctor models.Doctor
		initializers.DB.First(&doctor, "id = ?", item.DoctorID)

		personal_doctor = append(personal_doctor, struct {
			DoctorName string
			StartDate  string
			EndDate    string
		}{
			DoctorName: doctor.Name,
			StartDate:  item.StartDate,
			EndDate:    item.EndDate,
		})
	}

	userresponse.GetPersonalDoctorSuccessResponse(c, "Success Get Personal Doctor", personal_doctor, "http://localhost:3000/api/user/list/personal_doctor", http.StatusOK)
}
