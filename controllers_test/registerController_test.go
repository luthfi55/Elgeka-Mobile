package controllers_test

import (
	"bytes"
	"elgeka-mobile/controllers"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

type successExpectedRegisterResponse struct {
	Message string `json:"Message"`
	Data    []struct {
		ID    string `json:"ID"`
		Email string `json:"Email"`
	} `json:"Data"`
	Link []struct {
		Name string `json:"Name"`
		Link string `json:"Link"`
	} `json:"Link"`
}

type failedExpectedRegisterResponse struct {
	ErrorMessage string `json:"ErrorMessage"`
	Data         []struct {
	} `json:"Data"`
	Link []struct {
		Name string `json:"Name"`
		Link string `json:"Link"`
	} `json:"Link"`
}

func TestUserRegister_Success(t *testing.T) {

	router := gin.Default()

	router.POST("/api/user/register", controllers.UserRegister)

	reqBody := []byte(`{
		"Name": "John Doe",
		"province": "Jawa Barat",
		"district": "Kota Bandung",
		"SubDistrict": "Buahbatu",
		"village": "Margasari",
		"address": "Ciwastra No.81",
		"Gender": "male",
		"BirthDate": "1990-01-01",
		"BloodGroup": "A",
		"diagnosisdate": "2023-12-02",
		"PhoneNumber": "6289533991178",
		"Email": "johnn12@example.com",
		"Password": "Password123*",
		"PasswordConfirmation": "Password123*"
	}`)

	req, err := http.NewRequest("POST", "/api/user/register", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected status code %d but got %d", http.StatusCreated, rec.Code)
	}

	var expectedBody successExpectedRegisterResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Register Success" {
		t.Errorf("expected message %s but got %s", "Register Success", expectedBody.Message)
	}
}

func TestUserRegister_Failed(t *testing.T) {

	router := gin.Default()

	router.POST("/api/user/register", controllers.UserRegister)

	reqBody := []byte(`{
		"Name": "John Doe",
		"province": "Jawa Barat",
		"district": "Kota Bandung",
		"SubDistrict": "Buahbatu",
		"village": "Margasari",
		"address": "Ciwastra No.81",
		"Gender": "male",
		"BirthDate": "1990-01-01",
		"BloodGroup": "A",
		"diagnosisdate": "2023-12-02",
		"PhoneNumber": "6289533991178",
		"Email": "johnn12@example.com",
		"Password": "Password123*",
		"PasswordConfirmation": "Password123*"
	}`)

	req, err := http.NewRequest("POST", "/api/user/register", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d but got %d", http.StatusBadRequest, rec.Code)
	}

	var expectedBody failedExpectedRegisterResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Email Already Use" {
		t.Errorf("expected ErrorMessagemessage %s but got %s", "Email Already Use", expectedBody.ErrorMessage)
	}
}

func TestDoctorRegister_Success(t *testing.T) {

	router := gin.Default()

	router.POST("/api/doctor/register", controllers.DoctorRegister)

	reqBody := []byte(`{
		"Name": "Dr. John Doe",
		"Address": "123 Main St",
		"Gender": "male",		
		"polyname": "kandungan",
    	"hospitalname": "mayapada",
		"PhoneNumber": "0895339988332",
		"Email": "drjonnn@example.com",
		"Password": "Password123*",
		"PasswordConfirmation": "Password123*"
	}`)

	req, err := http.NewRequest("POST", "/api/doctor/register", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected status code %d but got %d", http.StatusCreated, rec.Code)
	}

	var expectedBody successExpectedRegisterResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Register Success" {
		t.Errorf("expected message %s but got %s", "Register Success", expectedBody.Message)
	}
}

func TestDoctorRegister_Failed(t *testing.T) {

	router := gin.Default()

	router.POST("/api/doctor/register", controllers.DoctorRegister)

	reqBody := []byte(`{
		"Name": "Dr. John Doe",
		"Address": "123 Main St",
		"Gender": "male",		
		"polyname": "kandungan",
    	"hospitalname": "mayapada",
		"PhoneNumber": "1234567890",
		"Email": "drjohn@example.com",
		"Password": "Password123*",
		"PasswordConfirmation": "Password123*"
	}`)

	req, err := http.NewRequest("POST", "/api/doctor/register", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d but got %d", http.StatusBadRequest, rec.Code)
	}

	var expectedBody failedExpectedRegisterResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Email Already Use" {
		t.Errorf("expected ErrorMessagemessage %s but got %s", "Email Already Use", expectedBody.ErrorMessage)
	}
}
