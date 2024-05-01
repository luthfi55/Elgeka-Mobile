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
