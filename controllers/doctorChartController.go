package controllers

import (
	"elgeka-mobile/initializers"
	"elgeka-mobile/middleware"
	"elgeka-mobile/models"
	doctorresponse "elgeka-mobile/response/DoctorResponse"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func DoctorChartController(r *gin.Engine) {
	r.GET("api/doctor/patient/data/gender", middleware.RequireAuth, DataByGender)
	r.GET("api/doctor/patient/data/age", middleware.RequireAuth, DataByAge)
	r.GET("api/doctor/patient/data/diagnosis_date", middleware.RequireAuth, DataByDiagnosisDate)
	r.GET("api/doctor/patient/data/treatment", middleware.RequireAuth, DataByTreatment)

	r.GET("api/doctor/patient/data/symptom/:type/:user_id", middleware.RequireAuth, GetSymptomUserData)
}

func DataByGender(c *gin.Context) {
	doctor, _ := c.Get("doctor")

	if !DoctorCheck(c, doctor) {
		return
	}

	var male_user []models.User
	result := initializers.DB.Where("gender = ?", "male").Find(&male_user)
	if result.Error != nil {
		doctorresponse.GenderChartFailedResponse(c, "Database Error", "", http.StatusInternalServerError)
		return
	}
	numMaleUsers := len(male_user)

	var female_user []models.User
	result = initializers.DB.Where("gender = ?", "female").Find(&female_user)
	if result.Error != nil {
		doctorresponse.GenderChartFailedResponse(c, "Database Error", "", http.StatusInternalServerError)
		return
	}
	numFemaleUsers := len(female_user)

	var Data struct {
		Male          int
		Female        int
		TotalPatient  int
		MalePercent   float32
		FemalePercent float32
	}

	Data.Male = numMaleUsers
	Data.Female = numFemaleUsers
	Data.TotalPatient = numMaleUsers + numFemaleUsers
	Data.MalePercent = float32(numMaleUsers) / float32(Data.TotalPatient) * 100
	Data.FemalePercent = float32(numFemaleUsers) / float32(Data.TotalPatient) * 100

	doctorresponse.GenderChartSuccessResponse(c, "Succes to Get Patient Data by Gender", Data, http.StatusOK)
}

func DataByAge(c *gin.Context) {
	doctor, _ := c.Get("doctor")

	if !DoctorCheck(c, doctor) {
		return
	}

	var users []models.User
	result := initializers.DB.Find(&users)
	if result.Error != nil {
		doctorresponse.AgeChartFailedResponse(c, "Database Error", "", http.StatusInternalServerError)
		return
	}

	var undereighteen int
	var underthirtynine int
	var underfiftynine int
	var underseventyfive int
	var overseventyfive int
	for _, user := range users {
		birthdate, err := time.Parse("2006-01-02", user.BirthDate)
		if err != nil {
			continue
		}
		age := calculateAge(birthdate, time.Now())
		if age <= 18 {
			undereighteen++
		} else if age >= 19 && age <= 39 {
			underthirtynine++
		} else if age >= 40 && age <= 59 {
			underfiftynine++
		} else if age >= 60 && age <= 75 {
			underseventyfive++
		} else if age > 75 {
			overseventyfive++
		}
	}

	var Data struct {
		UnderEighTeen           int
		UnderThirtyNine         int
		UnderFiftyNine          int
		UnderSeventyFive        int
		OverSeventyFive         int
		TotalPatient            int
		UnderEighTeenPercent    float32
		UnderThirtyNinePercent  float32
		UnderFiftyNinePercent   float32
		UnderSeventyFivePercent float32
		OverSeventyFivePercent  float32
	}

	Data.UnderEighTeen = undereighteen
	Data.UnderThirtyNine = underthirtynine
	Data.UnderFiftyNine = underfiftynine
	Data.UnderSeventyFive = underseventyfive
	Data.OverSeventyFive = overseventyfive
	Data.TotalPatient = undereighteen + underthirtynine + underfiftynine + underseventyfive + overseventyfive
	Data.UnderEighTeenPercent = float32(undereighteen) / float32(Data.TotalPatient) * 100
	Data.UnderThirtyNinePercent = float32(underthirtynine) / float32(Data.TotalPatient) * 100
	Data.UnderFiftyNinePercent = float32(underfiftynine) / float32(Data.TotalPatient) * 100
	Data.UnderSeventyFivePercent = float32(underseventyfive) / float32(Data.TotalPatient) * 100
	Data.OverSeventyFivePercent = float32(overseventyfive) / float32(Data.TotalPatient) * 100

	doctorresponse.AgeChartSuccessResponse(c, "Succes to Get Patient Data by Age", Data, http.StatusOK)
}

func calculateAge(birthdate, currentDate time.Time) int {
	years := currentDate.Year() - birthdate.Year()

	if currentDate.YearDay() < birthdate.YearDay() {
		years--
	}

	return years
}

func DataByDiagnosisDate(c *gin.Context) {
	doctor, _ := c.Get("doctor")

	if !DoctorCheck(c, doctor) {
		return
	}

	var users []models.User
	result := initializers.DB.Find(&users)
	if result.Error != nil {
		doctorresponse.DiagnosisChartFailedResponse(c, "Database Error", "", http.StatusInternalServerError)
		return
	}

	var underOneYear int
	var underThreeYear int
	var underTenYear int
	var overTenYear int
	for _, user := range users {
		diagnosisDate, err := time.Parse("2006-01-02", user.DiagnosisDate)
		if err != nil {
			continue
		}
		timeDiagnosis := calculateAge(diagnosisDate, time.Now())
		if timeDiagnosis < 1 {
			underOneYear++
		} else if timeDiagnosis >= 1 && timeDiagnosis < 3 {
			underThreeYear++
		} else if timeDiagnosis >= 3 && timeDiagnosis < 10 {
			underTenYear++
		} else if timeDiagnosis >= 10 {
			overTenYear++
		}
	}

	var Data struct {
		UnderOneYear          int
		UnderThreeYear        int
		UnderTenYear          int
		OverTenYear           int
		TotalPatient          int
		UnderOneYearPercent   float32
		UnderThreeYearPercent float32
		UnderTenYearPercent   float32
		OverTenYearPercent    float32
	}

	Data.UnderOneYear = underOneYear
	Data.UnderThreeYear = underThreeYear
	Data.UnderTenYear = underTenYear
	Data.OverTenYear = overTenYear
	Data.TotalPatient = underOneYear + underThreeYear + underTenYear + overTenYear
	Data.UnderOneYearPercent = float32(underOneYear) / float32(Data.TotalPatient) * 100
	Data.UnderThreeYearPercent = float32(underThreeYear) / float32(Data.TotalPatient) * 100
	Data.UnderTenYearPercent = float32(underTenYear) / float32(Data.TotalPatient) * 100
	Data.OverTenYearPercent = float32(overTenYear) / float32(Data.TotalPatient) * 100

	doctorresponse.DiagnosisChartSuccessResponse(c, "Succes to Get Patient Data by Diagnonsis Date", Data, http.StatusOK)

}

func DataByTreatment(c *gin.Context) {
	doctor, _ := c.Get("doctor")
	if !DoctorCheck(c, doctor) {
		return
	}

	var treatment []models.UserTreatment
	initializers.DB.Find(&treatment)
	totalPatient := len(treatment)

	treatments := []string{
		"Imatinib (Glivec)", "Generic Imatinib", "Nilotinib (Tasigna)", "Generic Nilotinib",
		"Dasatinib (Sprycel)", "Generic Dasatinib", "Bosutinib (Bosulif)", "Ponatinib (Iclusig)",
		"Radotinib (Supect)", "Hydroxyurea", "Interferon alfa", "Interferon beta",
		"Bone marrow transplantation", "Radotinib", "Olveramtinib",
	}

	data := make(map[string]int)
	for _, t := range treatments {
		var userTreatments []models.UserTreatment
		result := initializers.DB.Where("first_treatment = ? OR second_treatment = ?", t, t).Find(&userTreatments)
		if result.Error != nil {
			doctorresponse.TreatmentChartFailedResponse(c, "Database Error", "", http.StatusInternalServerError)
			return
		}
		data[t] = len(userTreatments)
	}

	response := make(map[string]interface{})
	for t, count := range data {
		response[t] = count
		response[t+" Percent"] = float32(count) / float32(totalPatient) / 2 * 100
	}
	response["TotalPatient"] = totalPatient

	doctorresponse.TreatmentChartSuccessResponse(c, "Success to Get Patient Data by Treatment", response, http.StatusOK)
}

func GetSymptomUserData(c *gin.Context) {
	doctor, _ := c.Get("doctor")
	userID := c.Param("user_id")
	typeSymptom := c.Param("type")
	if !DoctorCheck(c, doctor) {
		return
	}

	var user models.User
	result := initializers.DB.First(&user, "ID = ?", userID)

	if result.Error != nil {
		doctorresponse.SymptomAnswerDoctorFailedResponse(c, "Failed to Get Symptom User Data", "", http.StatusInternalServerError)
	}

	var symptom []models.SymptomAnswer
	var response []models.SymptomAnswerData

	if typeSymptom == "Oral" {
		result := initializers.DB.Where("user_id = ? AND type = ?", userID, "Oral").Find(&symptom)
		if result.Error != nil {
			doctorresponse.SymptomAnswerDoctorFailedResponse(c, "Failed to Get Data", "", http.StatusInternalServerError)
		}

		for _, symptom_data := range symptom {
			response = append(response, models.SymptomAnswerData{
				ID:         symptom_data.ID,
				Type:       symptom_data.Type,
				Date:       symptom_data.Date,
				WordAnswer: symptom_data.WordAnswer,
			})
		}

		doctorresponse.SymptomAnswerDoctorSuccessResponse(c, "Success to Get Oral Symptom User Data", response, http.StatusOK)
	} else if typeSymptom == "Digestive" {
		result := initializers.DB.Where("user_id = ? AND type = ?", userID, "Digestive").Find(&symptom)
		if result.Error != nil {
			doctorresponse.SymptomAnswerDoctorFailedResponse(c, "Failed to Get Data", "", http.StatusInternalServerError)
		}

		for _, symptom_data := range symptom {
			response = append(response, models.SymptomAnswerData{
				ID:         symptom_data.ID,
				Type:       symptom_data.Type,
				Date:       symptom_data.Date,
				WordAnswer: symptom_data.WordAnswer,
			})
		}

		doctorresponse.SymptomAnswerDoctorSuccessResponse(c, "Success to Get Digestive Symptom User Data", response, http.StatusOK)
	} else if typeSymptom == "Respiratory" {
		result := initializers.DB.Where("user_id = ? AND type = ?", userID, "Respiratory").Find(&symptom)
		if result.Error != nil {
			doctorresponse.SymptomAnswerDoctorFailedResponse(c, "Failed to Get Data", "", http.StatusInternalServerError)
		}
		for _, symptom_data := range symptom {
			response = append(response, models.SymptomAnswerData{
				ID:         symptom_data.ID,
				Type:       symptom_data.Type,
				Date:       symptom_data.Date,
				WordAnswer: symptom_data.WordAnswer,
			})
		}
		doctorresponse.SymptomAnswerDoctorSuccessResponse(c, "Success to Get Respiratory Symptom User Data", response, http.StatusOK)
	} else if typeSymptom == "Skin" {
		result := initializers.DB.Where("user_id = ? AND type = ?", userID, "Skin").Find(&symptom)
		if result.Error != nil {
			doctorresponse.SymptomAnswerDoctorFailedResponse(c, "Failed to Get Data", "", http.StatusInternalServerError)
		}
		for _, symptom_data := range symptom {
			response = append(response, models.SymptomAnswerData{
				ID:         symptom_data.ID,
				Type:       symptom_data.Type,
				Date:       symptom_data.Date,
				WordAnswer: symptom_data.WordAnswer,
			})
		}
		doctorresponse.SymptomAnswerDoctorSuccessResponse(c, "Success to Get Skin Symptom User Data", response, http.StatusOK)
	} else if typeSymptom == "Hair" {
		result := initializers.DB.Where("user_id = ? AND type = ?", userID, "Hair").Find(&symptom)
		if result.Error != nil {
			doctorresponse.SymptomAnswerDoctorFailedResponse(c, "Failed to Get Data", "", http.StatusInternalServerError)
		}
		for _, symptom_data := range symptom {
			response = append(response, models.SymptomAnswerData{
				ID:         symptom_data.ID,
				Type:       symptom_data.Type,
				Date:       symptom_data.Date,
				WordAnswer: symptom_data.WordAnswer,
			})
		}
		doctorresponse.SymptomAnswerDoctorSuccessResponse(c, "Success to Get Hair Symptom User Data", response, http.StatusOK)
	} else if typeSymptom == "Nails" {
		result := initializers.DB.Where("user_id = ? AND type = ?", userID, "Nails").Find(&symptom)
		if result.Error != nil {
			doctorresponse.SymptomAnswerDoctorFailedResponse(c, "Failed to Get Data", "", http.StatusInternalServerError)
		}
		for _, symptom_data := range symptom {
			response = append(response, models.SymptomAnswerData{
				ID:         symptom_data.ID,
				Type:       symptom_data.Type,
				Date:       symptom_data.Date,
				WordAnswer: symptom_data.WordAnswer,
			})
		}
		doctorresponse.SymptomAnswerDoctorSuccessResponse(c, "Success to Get Nails Symptom User Data", response, http.StatusOK)
	} else if typeSymptom == "Swelling" {
		result := initializers.DB.Where("user_id = ? AND type = ?", userID, "Swelling").Find(&symptom)
		if result.Error != nil {
			doctorresponse.SymptomAnswerDoctorFailedResponse(c, "Failed to Get Data", "", http.StatusInternalServerError)
		}
		for _, symptom_data := range symptom {
			response = append(response, models.SymptomAnswerData{
				ID:         symptom_data.ID,
				Type:       symptom_data.Type,
				Date:       symptom_data.Date,
				WordAnswer: symptom_data.WordAnswer,
			})
		}
		doctorresponse.SymptomAnswerDoctorSuccessResponse(c, "Success to Get Swelling Symptom User Data", response, http.StatusOK)
	} else if typeSymptom == "Senses" {
		result := initializers.DB.Where("user_id = ? AND type = ?", userID, "Senses").Find(&symptom)
		if result.Error != nil {
			doctorresponse.SymptomAnswerDoctorFailedResponse(c, "Failed to Get Data", "", http.StatusInternalServerError)
		}
		for _, symptom_data := range symptom {
			response = append(response, models.SymptomAnswerData{
				ID:         symptom_data.ID,
				Type:       symptom_data.Type,
				Date:       symptom_data.Date,
				WordAnswer: symptom_data.WordAnswer,
			})
		}
		doctorresponse.SymptomAnswerDoctorSuccessResponse(c, "Success to Get Senses Symptom User Data", response, http.StatusOK)
	} else if typeSymptom == "Moods" {
		result := initializers.DB.Where("user_id = ? AND type = ?", userID, "Moods").Find(&symptom)
		if result.Error != nil {
			doctorresponse.SymptomAnswerDoctorFailedResponse(c, "Failed to Get Data", "", http.StatusInternalServerError)
		}
		for _, symptom_data := range symptom {
			response = append(response, models.SymptomAnswerData{
				ID:         symptom_data.ID,
				Type:       symptom_data.Type,
				Date:       symptom_data.Date,
				WordAnswer: symptom_data.WordAnswer,
			})
		}
		doctorresponse.SymptomAnswerDoctorSuccessResponse(c, "Success to Get Moods Symptom User Data", response, http.StatusOK)
	} else if typeSymptom == "Pain" {
		result := initializers.DB.Where("user_id = ? AND type = ?", userID, "Pain").Find(&symptom)
		if result.Error != nil {
			doctorresponse.SymptomAnswerDoctorFailedResponse(c, "Failed to Get Data", "", http.StatusInternalServerError)
		}
		for _, symptom_data := range symptom {
			response = append(response, models.SymptomAnswerData{
				ID:         symptom_data.ID,
				Type:       symptom_data.Type,
				Date:       symptom_data.Date,
				WordAnswer: symptom_data.WordAnswer,
			})
		}
		doctorresponse.SymptomAnswerDoctorSuccessResponse(c, "Success to Get Pain Symptom User Data", response, http.StatusOK)
	} else if typeSymptom == "Cognitive" {
		result := initializers.DB.Where("user_id = ? AND type = ?", userID, "Cognitive").Find(&symptom)
		if result.Error != nil {
			doctorresponse.SymptomAnswerDoctorFailedResponse(c, "Failed to Get Data", "", http.StatusInternalServerError)
		}
		for _, symptom_data := range symptom {
			response = append(response, models.SymptomAnswerData{
				ID:         symptom_data.ID,
				Type:       symptom_data.Type,
				Date:       symptom_data.Date,
				WordAnswer: symptom_data.WordAnswer,
			})
		}
		doctorresponse.SymptomAnswerDoctorSuccessResponse(c, "Success to Get Cognitive Symptom User Data", response, http.StatusOK)
	} else if typeSymptom == "Urinary" {
		result := initializers.DB.Where("user_id = ? AND type = ?", userID, "Urinary").Find(&symptom)
		if result.Error != nil {
			doctorresponse.SymptomAnswerDoctorFailedResponse(c, "Failed to Get Data", "", http.StatusInternalServerError)
		}
		for _, symptom_data := range symptom {
			response = append(response, models.SymptomAnswerData{
				ID:         symptom_data.ID,
				Type:       symptom_data.Type,
				Date:       symptom_data.Date,
				WordAnswer: symptom_data.WordAnswer,
			})
		}
		doctorresponse.SymptomAnswerDoctorSuccessResponse(c, "Success to Get Urinary Symptom User Data", response, http.StatusOK)
	} else if typeSymptom == "Genitals" {
		result := initializers.DB.Where("user_id = ? AND type = ?", userID, "Genitals").Find(&symptom)
		if result.Error != nil {
			doctorresponse.SymptomAnswerDoctorFailedResponse(c, "Failed to Get Data", "", http.StatusInternalServerError)
		}
		for _, symptom_data := range symptom {
			response = append(response, models.SymptomAnswerData{
				ID:         symptom_data.ID,
				Type:       symptom_data.Type,
				Date:       symptom_data.Date,
				WordAnswer: symptom_data.WordAnswer,
			})
		}
		doctorresponse.SymptomAnswerDoctorSuccessResponse(c, "Success to Get Genitals Symptom User Data", response, http.StatusOK)
	} else if typeSymptom == "Reproductive" {
		result := initializers.DB.Where("user_id = ? AND type = ?", userID, "Reproductive").Find(&symptom)
		if result.Error != nil {
			doctorresponse.SymptomAnswerDoctorFailedResponse(c, "Failed to Get Data", "", http.StatusInternalServerError)
		}
		for _, symptom_data := range symptom {
			response = append(response, models.SymptomAnswerData{
				ID:         symptom_data.ID,
				Type:       symptom_data.Type,
				Date:       symptom_data.Date,
				WordAnswer: symptom_data.WordAnswer,
			})
		}
		doctorresponse.SymptomAnswerDoctorSuccessResponse(c, "Success to Get Reproductive Symptom User Data", response, http.StatusOK)
	} else {
		doctorresponse.SymptomAnswerDoctorFailedResponse(c, "Symptom Type Not Found", "", http.StatusBadRequest)
	}
}
