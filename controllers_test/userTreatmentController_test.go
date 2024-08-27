package controllers_test

import (
	"bytes"
	"elgeka-mobile/controllers"
	"elgeka-mobile/middleware"
	"elgeka-mobile/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestGetUserTreatment_Success(t *testing.T) {
	router := gin.Default()

	router.GET("/api/user/profile/treatment", middleware.RequireAuth, controllers.GetTreatmentData)

	req, err := http.NewRequest("GET", "/api/user/profile/treatment", nil)
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

	if expectedBody.Message != "Success to Get Treatment Data" {
		t.Errorf("expected message body %s but got %s", "Success to Get Treatment Data", expectedBody.Message)
	}
}

func TestGetUserTreatment_Failed(t *testing.T) {
	router := gin.Default()

	router.GET("/api/user/profile/treatment", middleware.RequireAuth, controllers.GetTreatmentData)

	req, err := http.NewRequest("GET", "/api/user/profile/treatment", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	expiresTime, _ := time.Parse(time.RFC1123, "Mon, 15 Apr 2024 17:00:20 GMT")
	req.AddCookie(&http.Cookie{
		Name:     "Authorization",
		Value:    "eyJhbGciOiKIUDI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTcwNTgxNjQsInN1YiI6IjVmYWUwYjVkLTAxZDAtNDhmYi05ZTIzLWNhMWU5YWNhMmYyMCJ9.j4kcnBLIyOPnXl5Ok1gZKhqsAysXs2MTsEmzN23sVi0",
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

func TestUpdateUserTreatment_Success(t *testing.T) {
	router := gin.Default()
	treatment_id := "551924fa-1f2b-4ffc-87d5-8f8e7dce431c"

	router.PUT("/api/user/profile/treatment/edit/:treatment_id", middleware.RequireAuth, controllers.EditTreatmentData)

	reqBody := models.UserTreatment{
		FirstTreatment:  "Imatinib (Glivec)",
		SecondTreatment: "Bosutinib (Bosulif)",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("PUT", "/api/user/profile/treatment/edit/"+treatment_id, bytes.NewBuffer(reqJSON))
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

	if expectedBody.Message != "Success To Update Treatment Data" {
		t.Errorf("expected message body %s but got %s", "Success To Update Treatment Data", expectedBody.Message)
	}
}

func TestUpdateUserTreatment_Failed(t *testing.T) {
	router := gin.Default()
	treatment_id := "551924fa-1f2b-4ffc-87d5-8f8e7dce431c"

	router.PUT("/api/user/profile/treatment/edit/:treatment_id", middleware.RequireAuth, controllers.EditTreatmentData)

	reqBody := models.UserTreatment{
		FirstTreatment:  "Imatinib (Glivec)",
		SecondTreatment: "Makan",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("PUT", "/api/user/profile/treatment/edit/"+treatment_id, bytes.NewBuffer(reqJSON))
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

	if expectedBody.ErrorMessage != "Unknown 'Makan' as Treatment Data" {
		t.Errorf("expected message body %s but got %s", "Unknown 'Makan' as Treatment Data", expectedBody.ErrorMessage)
	}
}
