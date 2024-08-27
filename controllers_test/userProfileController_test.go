package controllers_test

import (
	"bytes"
	"elgeka-mobile/controllers"
	"elgeka-mobile/middleware"
	"elgeka-mobile/models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TestProfileData_Success(t *testing.T) {
	router := gin.Default()

	router.GET("/api/user/profile", middleware.RequireAuth, controllers.ProfileData)

	req, err := http.NewRequest("GET", "/api/user/profile", nil)

	if err != nil {
		t.Fatal(err)
	}

	req.AddCookie(CookieConfiguration())

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	type ExpectedResponse struct {
		Message string            `json:"Message"`
		Data    []models.UserData `json:"Data"`
		Link    []struct {
			Name string `json:"Name"`
			Link string `json:"Link"`
		} `json:"Link"`
	}

	var expectedBody ExpectedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		fmt.Println(rec.Body.String())
		t.Fatal(err)
	}

	if expectedBody.Message != "Success Get Profile" {
		t.Errorf("expected message %s but got %s", "Success Get Profile", expectedBody.Message)
	}
}

func TestProfileData_Failed(t *testing.T) {
	router := gin.Default()

	router.GET("/api/user/profile", middleware.RequireAuth, controllers.ProfileData)

	req, err := http.NewRequest("GET", "/api/user/profile", nil)

	if err != nil {
		t.Fatal(err)
	}

	expiresTime, _ := time.Parse(time.RFC1123, "Mon, 15 Apr 2023 17:00:20 GMT")
	req.AddCookie(
		&http.Cookie{
			Name:     "Authorization",
			Value:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ1.asJleHAiOjE3MTMyMDA0MjAsInN1YiI6IjE0YzlhNDQzLTAzZTUtNGJhNi05NjY0LTBmODIwYjE5ZDhhYiJ0.L9z84gPX0l_O3GeyRi0ZAhMGxoWzXVV7k9fXw6KpEo4",
			Path:     "/",
			HttpOnly: true,
			Expires:  expiresTime,
		},
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d but got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestEditProfile_Success(t *testing.T) {
	router := gin.Default()

	router.PUT("/api/user/profile/edit", middleware.RequireAuth, controllers.EditProfile)

	reqBody := models.User{
		Name:        "Ane",
		Province:    "Jawa Timur",
		District:    "Kota Surabaya",
		SubDistrict: "Buahmangga",
		Village:     "Margamati",
		Address:     "Ciuaua No.12",
		Gender:      "male",
		BirthDate:   "2023-12-02",
		BloodGroup:  "B",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("PUT", "/api/user/profile/edit", bytes.NewBuffer(reqJSON))
	req.AddCookie(CookieConfiguration())

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	type ExpectedResponse struct {
		Message string `json:"Message"`
		Data    []struct {
			ID string `json:"ID"`
		} `json:"Data"`
		Link []struct {
			Name string `json:"Name"`
			Link string `json:"Link"`
		} `json:"Link"`
	}

	var expectedBody ExpectedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success Update User" {
		t.Errorf("expected message body %s but got %s", "Success Update User", expectedBody.Message)
	}
}

func TestEditProfile_Failed(t *testing.T) {
	router := gin.Default()

	router.PUT("/api/user/profile/edit", middleware.RequireAuth, controllers.EditProfile)

	reqBody := models.User{
		Name:        "",
		Province:    "",
		District:    "",
		SubDistrict: "",
		Village:     "",
		Address:     "",
		Gender:      "",
		BirthDate:   "",
		BloodGroup:  "",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("PUT", "/api/user/profile/edit", bytes.NewBuffer(reqJSON))
	req.AddCookie(CookieConfiguration())

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d but got %d", http.StatusBadRequest, rec.Code)
	}

	type ExpectedResponse struct {
		ErrorMessage string `json:"ErrorMessage"`
		Data         []struct {
			ID string `json:"ID"`
		} `json:"Data"`
		Link []struct {
			Name string `json:"Name"`
			Link string `json:"Link"`
		} `json:"Link"`
	}

	var expectedBody ExpectedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Body Can't Null" {
		t.Errorf("expected message body %s but got %s", "Body Can't Null", expectedBody.ErrorMessage)
	}
}

func TestAddPersonalDoctor_Success(t *testing.T) {
	router := gin.Default()

	router.POST("/api/user/add/personal_doctor", middleware.RequireAuth, controllers.AddPersonalDoctor)

	reqBody := models.UserPersonalDoctor{
		DoctorID: uuid.MustParse("a8cedfc6-4e79-4e92-b871-bcd66c009800"),
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/api/user/add/personal_doctor", bytes.NewBuffer(reqJSON))
	req.AddCookie(CookieConfiguration())

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	type ExpectedResponse struct {
		Message string `json:"Message"`
		Data    []struct {
			ID string `json:"Personal Doctor ID"`
		} `json:"Data"`
		Link []struct {
			Name string `json:"Name"`
			Link string `json:"Link"`
		} `json:"Link"`
	}

	var expectedBody ExpectedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success Add Personal Doctor" {
		t.Errorf("expected message body %s but got %s", "Success Add Personal Doctor", expectedBody.Message)
	}
}

func TestAddPersonalDoctor_Failed(t *testing.T) {
	router := gin.Default()

	router.POST("/api/user/add/personal_doctor", middleware.RequireAuth, controllers.AddPersonalDoctor)

	reqBody := models.UserPersonalDoctor{
		DoctorID: uuid.MustParse("a8cedfc6-4e79-4e92-b871-bcd66c009801"),
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/api/user/add/personal_doctor", bytes.NewBuffer(reqJSON))
	req.AddCookie(CookieConfiguration())

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d but got %d", http.StatusBadRequest, rec.Code)
	}

	type ExpectedResponse struct {
		ErrorMessage string `json:"ErrorMessage"`
		Data         string `json:"Data"`
		Link         []struct {
			Name string `json:"Name"`
			Link string `json:"Link"`
		} `json:"Link"`
	}

	var expectedBody ExpectedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Doctor Not Found" {
		t.Errorf("expected message body %s but got %s", "Doctor Not Found", expectedBody.ErrorMessage)
	}
}

func TestGetListPersonalDoctor_Success(t *testing.T) {
	router := gin.Default()

	router.GET("/api/user/list/personal_doctor", middleware.RequireAuth, controllers.GetPersonalDoctor)

	req, err := http.NewRequest("GET", "/api/user/list/personal_doctor", nil)
	req.AddCookie(CookieConfiguration())

	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	type ExpectedResponse struct {
		Message string `json:"Message"`
		Data    []struct {
			DoctorName  string `json:"DoctorName"`
			PhoneNumber string `json:"PhoneNumber"`
			StartDate   string `json:"StartDate"`
			EndDate     string `json:"Email"`
		} `json:"Data"`
		Link []struct {
			Name string `json:"Name"`
			Link string `json:"Link"`
		} `json:"Link"`
	}

	var expectedBody ExpectedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success Get Personal Doctor" {
		t.Errorf("expected message body %s but got %s", "Success Get Personal Doctor", expectedBody.Message)
	}
}

func TestGetListPersonalDoctor_Failed(t *testing.T) {
	router := gin.Default()

	router.GET("/api/user/list/personal_doctor", middleware.RequireAuth, controllers.GetPersonalDoctor)

	req, err := http.NewRequest("GET", "/api/user/list/personal_doctor", nil)

	if err != nil {
		t.Fatal(err)
	}
	expiresTime, _ := time.Parse(time.RFC1123, "Mon, 15 Apr 2023 17:00:20 GMT")
	req.AddCookie(
		&http.Cookie{
			Name:     "Authorization",
			Value:    "ayJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ1.asJleHAiOjE3MTMyMDA0MjAsInN1YiI6IjE0YzlhNDQzLTAzZTUtNGJhNi05NjY0LTBmODIwYjE5ZDhhYiJ0.L9z84gPX0l_O3GeyRi0ZAhMGxoWzXVV7k9fXw6KpEo4",
			Path:     "/",
			HttpOnly: true,
			Expires:  expiresTime,
		},
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d but got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestGetListActivateDoctor_Success(t *testing.T) {
	router := gin.Default()

	router.GET("/api/user/list/activate_doctor", middleware.RequireAuth, controllers.ListActivateDoctor)

	req, err := http.NewRequest("GET", "/api/user/list/activate_doctor", nil)
	req.AddCookie(CookieConfiguration())

	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	type ExpectedResponse struct {
		Message string `json:"Message"`
		Data    []struct {
			DoctorName  string `json:"DoctorName"`
			PhoneNumber string `json:"PhoneNumber"`
			StartDate   string `json:"StartDate"`
			EndDate     string `json:"Email"`
		} `json:"Data"`
		Link []struct {
			Name string `json:"Name"`
			Link string `json:"Link"`
		} `json:"Link"`
	}

	var expectedBody ExpectedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success Get Active Doctor" {
		t.Errorf("expected message body %s but got %s", "Success Get Active Doctor", expectedBody.Message)
	}
}

func TestGetListActivateDoctor_Failed(t *testing.T) {
	router := gin.Default()

	router.GET("/api/user/list/activate_doctor", middleware.RequireAuth, controllers.ListActivateDoctor)

	req, err := http.NewRequest("GET", "/api/user/list/activate_doctor", nil)

	if err != nil {
		t.Fatal(err)
	}
	expiresTime, _ := time.Parse(time.RFC1123, "Mon, 15 Apr 2023 17:00:20 GMT")
	req.AddCookie(
		&http.Cookie{
			Name:     "Authorization",
			Value:    "ayJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ1.asJleHAiOjE3MTMyMDA0MjAsInN1YiI6IjE0YzlhNDQzLTAzZTUtNGJhNi05NjY0LTBmODIwYjE5ZDhhYiJ0.L9z84gPX0l_O3GeyRi0ZAhMGxoWzXVV7k9fXw6KpEo4",
			Path:     "/",
			HttpOnly: true,
			Expires:  expiresTime,
		},
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d but got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestListUserWebsite_Success(t *testing.T) {
	router := gin.Default()

	router.GET("/api/user/list/website", controllers.ListUserWebsite)

	req, err := http.NewRequest("GET", "/api/user/list/website", nil)

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.AddCookie(WebsiteBearierTokenConfiguration())
	bearerToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImlkIjoiNGUyZDY4NDgtM2FlMy00NjdjLTk5NzQtZTVkMTdhYWJhMGU3IiwidXNlcm5hbWUiOiJwZW5ndXJ1c2VsZ2VrYSIsImZ1bGxfbmFtZSI6IlBlbmd1cnVzIFV0YW1hIEVMR0VLQSBKQUJBUiIsImlzX2FjdGl2ZSI6dHJ1ZSwic3VwZXJVc2VyIjp0cnVlfSwiaWF0IjoxNzE4MDkwOTE5LCJleHAiOjE3MTgxMTI1MTl9.uZKwT15qe6Q6xytlbLtFQNysSoKd4VR0MI5diYqqge8"
	req.Header.Set("Authorization", "Bearer "+bearerToken)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	var expectedBody ExpectedSuccessResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success to Get Patient List" {
		t.Errorf("expected message body %s but got %s", "Success to Get Patient List", expectedBody.Message)
	}
}

func TestListUserWebsite_Failed(t *testing.T) {
	router := gin.Default()

	router.GET("/api/user/list/website", controllers.ListUserWebsite)

	req, err := http.NewRequest("GET", "/api/user/list/website", nil)

	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d but got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestEditUserPassword_Success(t *testing.T) {
	router := gin.Default()

	router.PUT("/api/user/profile/password/edit", middleware.RequireAuth, controllers.EditUserPassword)

	var body struct {
		OldPassword          string
		Password             string
		PasswordConfirmation string
	}

	body = struct {
		OldPassword          string
		Password             string
		PasswordConfirmation string
	}{
		OldPassword:          "Arinnnn1*",
		Password:             "Arinnnn2*",
		PasswordConfirmation: "Arinnnn2*",
	}

	reqJSON, _ := json.Marshal(body)

	req, err := http.NewRequest("PUT", "/api/user/profile/password/edit", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.AddCookie(CookieConfiguration())

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	var expectedBody ExpectedSuccessResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success Update User Password" {
		t.Errorf("expected message body %s but got %s", "Success Update User Password", expectedBody.Message)
	}
}

func TestEditUserPassword_Failed(t *testing.T) {
	router := gin.Default()

	router.PUT("/api/user/profile/password/edit", middleware.RequireAuth, controllers.EditUserPassword)

	var body struct {
		OldPassword          string
		Password             string
		PasswordConfirmation string
	}

	body = struct {
		OldPassword          string
		Password             string
		PasswordConfirmation string
	}{
		OldPassword:          "Arinnnn1*",
		Password:             "Arinnnn2*",
		PasswordConfirmation: "Arinnnn2*",
	}

	reqJSON, _ := json.Marshal(body)

	req, err := http.NewRequest("PUT", "/api/user/profile/password/edit", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.AddCookie(CookieConfiguration())

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d but got %d", http.StatusBadRequest, rec.Code)
	}

	var expectedBody ExpectedFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Wrong Old Password" {
		t.Errorf("expected message body %s but got %s", "Wrong Old Password", expectedBody.ErrorMessage)
	}
}

func TestEditUserInformation_Success(t *testing.T) {
	router := gin.Default()

	router.PUT("/api/user/profile/information/edit", middleware.RequireAuth, controllers.EditUserInformation)

	reqBody := models.UserInformation{
		PcrLevel:          "<10%",
		TherapyActive:     false,
		TreatmentFree:     true,
		TreatmentFreeDate: "2024-04-12",
		MonitoringPlace:   "Komunitas Rumah Sakit",
		PcrFrequent:       "Sekali dalam 4 bulan",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("PUT", "/api/user/profile/information/edit", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.AddCookie(CookieConfiguration())

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	var expectedBody ExpectedSuccessResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success Update User Information" {
		t.Errorf("expected message body %s but got %s", "Success Update User Information", expectedBody.Message)
	}
}

func TestEditUserInformation_Failed(t *testing.T) {
	router := gin.Default()

	router.PUT("/api/user/profile/information/edit", middleware.RequireAuth, controllers.EditUserInformation)

	reqBody := models.UserInformation{
		PcrLevel:          "",
		TherapyActive:     false,
		TreatmentFree:     true,
		TreatmentFreeDate: "2024-04-12",
		MonitoringPlace:   "Komunitas Rumah Sakit",
		PcrFrequent:       "Sekali dalam 4 bulan",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("PUT", "/api/user/profile/information/edit", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.AddCookie(CookieConfiguration())

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d but got %d", http.StatusBadRequest, rec.Code)
	}

	var expectedBody ExpectedFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Field 'PcrLevel' is required." {
		t.Errorf("expected message body %s but got %s", "Field 'PcrLevel' is required.", expectedBody.ErrorMessage)
	}
}

func TestGetDoctorData_Success(t *testing.T) {
	router := gin.Default()

	router.GET("/api/user/get/personal_doctor/:doctor_id", middleware.RequireAuth, controllers.GetDoctorData)

	doctor_id := "5fae0b5d-01d0-48fb-9e23-ca1e9aca2f20"

	req, err := http.NewRequest("GET", "/api/user/get/personal_doctor/"+doctor_id, nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.AddCookie(CookieConfiguration())

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	var expectedBody ExpectedSuccessResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success Get Doctor Data" {
		t.Errorf("expected message body %s but got %s", "Success Get Doctor Data", expectedBody.Message)
	}
}

func TestGetDoctorData_Failed(t *testing.T) {
	router := gin.Default()

	router.GET("/api/user/get/personal_doctor/:doctor_id", middleware.RequireAuth, controllers.GetDoctorData)

	doctor_id := "5fae0b3e-01d0-48fz-9e23-cv1e9aca2f20"

	req, err := http.NewRequest("GET", "/api/user/get/personal_doctor/"+doctor_id, nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.AddCookie(CookieConfiguration())

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d but got %d", http.StatusBadRequest, rec.Code)
	}

	var expectedBody ExpectedFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Doctor Not Found" {
		t.Errorf("expected message body %s but got %s", "Doctor Not Found", expectedBody.ErrorMessage)
	}
}
