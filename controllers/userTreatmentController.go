package controllers

import (
	"elgeka-mobile/initializers"
	"elgeka-mobile/middleware"
	"elgeka-mobile/models"
	userresponse "elgeka-mobile/response/UserResponse"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UserTreatmentController(r *gin.Engine) {
	r.GET("api/user/profile/treatment", middleware.RequireAuth, GetTreatmentData)
	r.PUT("api/user/profile/treatment/edit/:treatment_id", middleware.RequireAuth, EditTreatmentData)
}

func GetTreatmentData(c *gin.Context) {
	user, _ := c.Get("user")

	var user_data models.User
	if err := initializers.DB.First(&user_data, "id = ?", user).Error; err != nil {
		userresponse.GetTreatmentDataFailedResponse(c, "Failed To Find User", "", http.StatusBadRequest)
		return
	}

	var treatment models.UserTreatment
	if err := initializers.DB.First(&treatment, "user_id = ?", user).Error; err != nil {
		newUUID := uuid.New()
		treatment.ID = newUUID
		treatment.UserID = user_data.ID
		if err := initializers.DB.Create(&treatment).Error; err != nil {
			userresponse.GetTreatmentDataFailedResponse(c, "Failed To Create Treatment Data", "", http.StatusInternalServerError)
			return
		}
	}

	var Data struct {
		ID              uuid.UUID
		FirstTreatment  string
		SecondTreatment string
	}

	Data.ID = treatment.ID
	Data.FirstTreatment = treatment.FirstTreatment
	Data.SecondTreatment = treatment.SecondTreatment

	userresponse.GetTreatmentDataSuccessResponse(c, "Success to Get Treatment Data", Data, http.StatusOK)

}

func EditTreatmentData(c *gin.Context) {
	user, _ := c.Get("user")
	treatmentID := c.Param("treatment_id")

	var body models.UserTreatment

	if c.Bind(&body) != nil {
		userresponse.GetTreatmentDataFailedResponse(c, "Failed To Read Body", "", http.StatusBadRequest)
		return
	}

	if body.FirstTreatment == body.SecondTreatment {
		userresponse.GetTreatmentDataFailedResponse(c, "First Treatment And Second Treatment Can't Be The Same", "", http.StatusBadRequest)
		return
	}

	if body.FirstTreatment == "" {
		body.FirstTreatment = ""
	} else if body.FirstTreatment != "Imatinib (Glivec)" && body.FirstTreatment != "Generic Imatinib" && body.FirstTreatment != "Nilotinib (Tasigna)" && body.FirstTreatment != "Generic Nilotinib" && body.FirstTreatment != "Dasatinib (Sprycel)" && body.FirstTreatment != "Generic Dasatinib" && body.FirstTreatment != "Bosutinib (Bosulif)" && body.FirstTreatment != "Ponatinib (Iclusig)" && body.FirstTreatment != "Radotinib (Supect)" && body.FirstTreatment != "Hydroxyurea" && body.FirstTreatment != "Interferon alfa" && body.FirstTreatment != "Interferon beta" && body.FirstTreatment != "Bone marrow transplantation" {
		userresponse.GetTreatmentDataFailedResponse(c, "Unknown '"+body.FirstTreatment+"' as Treatment Data", "", http.StatusBadRequest)
		return
	}

	if body.SecondTreatment == "" {
		body.SecondTreatment = ""
	} else if body.SecondTreatment != "Imatinib (Glivec)" && body.SecondTreatment != "Generic Imatinib" && body.SecondTreatment != "Nilotinib (Tasigna)" && body.SecondTreatment != "Generic Nilotinib" && body.SecondTreatment != "Dasatinib (Sprycel)" && body.SecondTreatment != "Generic Dasatinib" && body.SecondTreatment != "Bosutinib (Bosulif)" && body.SecondTreatment != "Ponatinib (Iclusig)" && body.SecondTreatment != "Radotinib (Supect)" && body.SecondTreatment != "Hydroxyurea" && body.SecondTreatment != "Interferon alfa" && body.SecondTreatment != "Interferon beta" && body.SecondTreatment != "Bone marrow transplantation" {
		userresponse.GetTreatmentDataFailedResponse(c, "Unknown '"+body.SecondTreatment+"' as Treatment Data", "", http.StatusBadRequest)
		return
	}

	var user_data models.User
	if err := initializers.DB.First(&user_data, "id = ?", user).Error; err != nil {
		userresponse.GetTreatmentDataFailedResponse(c, "Failed To Find User", "", http.StatusBadRequest)
		return
	}

	var treatment models.UserTreatment
	if err := initializers.DB.First(&treatment, "id = ?", treatmentID).Error; err != nil {
		userresponse.GetTreatmentDataFailedResponse(c, "Failed To Find Treatment Data", "", http.StatusBadRequest)
		return
	}

	treatment.FirstTreatment = body.FirstTreatment
	treatment.SecondTreatment = body.SecondTreatment

	if body.FirstTreatment == "" && body.SecondTreatment == "" {
		userresponse.GetTreatmentDataFailedResponse(c, "Body Can't Null", "", http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Save(&treatment).Error; err != nil {
		userresponse.GetTreatmentDataFailedResponse(c, "Failed To Update Treatment Data", "", http.StatusInternalServerError)
		return
	}

	var Data struct {
		ID              uuid.UUID
		FirstTreatment  string
		SecondTreatment string
	}

	Data.ID = treatment.ID
	Data.FirstTreatment = treatment.FirstTreatment
	Data.SecondTreatment = treatment.SecondTreatment

	userresponse.GetTreatmentDataSuccessResponse(c, "Success To Update Treatment Data", Data, http.StatusOK)
}
