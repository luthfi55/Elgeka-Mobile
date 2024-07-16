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

type successExpectedMedicineResponse struct {
	Message string `json:"Message"`
	Data    []struct {
		ID    string `json:"ID"`
		Name  string `json:"Name"`
		Stock int    `json:"Stock"`
	} `json:"Data"`
	Link []struct {
		Name string `json:"Name"`
		Link string `json:"Link"`
	} `json:"Link"`
}

type failedExpectedMedicineResponse struct {
	ErrorMessage string `json:"ErrorMessage"`
	Data         []struct {
		ID    string `json:"ID"`
		Name  string `json:"Name"`
		Stock int    `json:"Stock"`
	} `json:"Data"`
	Link []struct {
		Name string `json:"Name"`
		Link string `json:"Link"`
	} `json:"Link"`
}

type successExpectedMedicineScheduleResponse struct {
	Message string `json:"Message"`
	Data    []struct {
		ID     string `json:"ID"`
		Date   string `json:"Date"`
		Status bool   `json:"Status"`
	} `json:"Data"`
	Link []struct {
		Name string `json:"Name"`
		Link string `json:"Link"`
	} `json:"Link"`
}

type failedExpectedMedicineScheduleResponse struct {
	ErrorMessage string `json:"ErrorMessage"`
	Data         []struct {
		ID     string `json:"ID"`
		Date   string `json:"Date"`
		Status bool   `json:"Status"`
	} `json:"Data"`
	Link []struct {
		Name string `json:"Name"`
		Link string `json:"Link"`
	} `json:"Link"`
}

func TestAddMedicine_Success(t *testing.T) {
	router := gin.Default()

	router.POST("/api/user/medicine", middleware.RequireAuth, controllers.AddMedicine)

	reqBody := models.Medicine{
		Name:   "Generic Imatinib",
		Dosage: "200mg",
		Stock:  10,
	}
	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/api/user/medicine", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(CookieConfiguration())

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected status code %d but got %d", http.StatusCreated, rec.Code)
	}

	var expectedBody successExpectedMedicineResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success to Add Medicine Data" {
		t.Errorf("expected message %s but got %s", "Success to Add Medicine Data", expectedBody.Message)
	}
}

func TestAddMedicine_Failed(t *testing.T) {
	router := gin.Default()

	router.POST("/api/user/medicine", middleware.RequireAuth, controllers.AddMedicine)

	reqBody := models.Medicine{
		Name:   "",
		Dosage: "200mg",
		Stock:  10,
	}
	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/api/user/medicine", bytes.NewBuffer(reqJSON))
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

	var expectedBody failedExpectedMedicineResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Name or Stock Can't be Empty" {
		t.Errorf("expected message %s but got %s", "Name or Stock Can't be Empty", expectedBody.ErrorMessage)
	}
}

func TestListMedicine_Success(t *testing.T) {
	router := gin.Default()
	router.GET("/api/user/medicine", middleware.RequireAuth, controllers.ListMedicine)
	req, err := http.NewRequest("GET", "/api/user/medicine", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(CookieConfiguration())

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	var expectedBody successExpectedMedicineResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success to Get Medicine List" {
		t.Errorf("expected message %s but got %s", "Success to Get Medicine List", expectedBody.Message)
	}
}

func TestListMedicine_Failed(t *testing.T) {
	router := gin.Default()
	router.GET("/api/user/medicine", middleware.RequireAuth, controllers.ListMedicine)
	req, err := http.NewRequest("GET", "/api/user/medicine", nil)
	if err != nil {
		t.Fatal(err)
	}
	expiresTime, _ := time.Parse(time.RFC1123, "Tue, 13 Aug 2024 04:56:45 GMT")
	req.AddCookie(
		&http.Cookie{
			Name:     "Authorization",
			Value:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ8.eyJleHAiOjE3MTMyMDA0MjAsInN1YiI6IjE0YzlhNDQzLTAzZTUtNGJhNi05NjY0LTBmODIwYjE5ZDhhYiJ9.L9z84gPX0l_O3GeyRi0ZAhMGxoWzXVV7k9fXw6KpEo4",
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

func TestUpdateMedicine_Success(t *testing.T) {
	router := gin.Default()
	medicine_id := "9f3332d6-a09d-4da8-a414-33f4a264cda2"

	router.PUT("api/user/medicine/:medicine_id", middleware.RequireAuth, controllers.UpdateMedicine)

	reqBody := models.Medicine{
		Name:   "Generic Imatinib",
		Dosage: "300mg",
		Stock:  100,
	}
	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("PUT", "/api/user/medicine/"+medicine_id, bytes.NewBuffer(reqJSON))
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

	var expectedBody successExpectedMedicineResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success to Update Medicine Data" {
		t.Errorf("expected message body %s but got %s", "Success to Update Medicine Data", expectedBody.Message)
	}
}

func TestUpdateMedicine_Failed(t *testing.T) {
	router := gin.Default()
	medicine_id := "d325b785-a18e-4e0f-a264-20135c2a1b63"

	router.PUT("api/user/medicine/:medicine_id", middleware.RequireAuth, controllers.UpdateMedicine)

	reqBody := models.Medicine{
		Name:   "",
		Dosage: "100mg",
		Stock:  100,
	}
	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("PUT", "/api/user/medicine/"+medicine_id, bytes.NewBuffer(reqJSON))
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

	var expectedBody failedExpectedMedicineResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Name Can't be Empty" {
		t.Errorf("expected message body %s but got %s", "Name Can't be Empty", expectedBody.ErrorMessage)
	}
}

func TestDeleteMedicine_Success(t *testing.T) {
	router := gin.Default()
	medicine_id := "a35ac9a3-a3f2-4e29-a696-7a87ad419faf"

	router.DELETE("api/user/medicine/:medicine_id", middleware.RequireAuth, controllers.DeleteMedicine)

	req, err := http.NewRequest("DELETE", "/api/user/medicine/"+medicine_id, nil)
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

	var expectedBody successExpectedMedicineResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success to Delete Medicine Data" {
		t.Errorf("expected message body %s but got %s", "Success to Delete Medicine Data", expectedBody.Message)
	}
}

func TestDeleteMedicine_Failed(t *testing.T) {
	router := gin.Default()
	medicine_id := "d325b785-a18e-4e0f-a264-20135c2a1b63"

	router.DELETE("api/user/medicine/:medicine_id", middleware.RequireAuth, controllers.DeleteMedicine)

	req, err := http.NewRequest("DELETE", "/api/user/medicine/"+medicine_id, nil)
	req.AddCookie(CookieConfiguration())

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status code %d but got %d", http.StatusNotFound, rec.Code)
	}

	var expectedBody failedExpectedMedicineResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Medicine Not Found" {
		t.Errorf("expected message body %s but got %s", "Medicine Not Found", expectedBody.ErrorMessage)
	}
}

func TestAddMedicineSchedule_Success(t *testing.T) {
	router := gin.Default()

	medicine_id := "1bb31639-b6c4-4f29-9b96-f773d4c99340"

	router.POST("/api/user/medicine/schedule/:medicine_id", middleware.RequireAuth, controllers.AddMedicineSchedule)

	reqBody := models.MedicineSchedule{
		Date: "2024-07-25",
	}
	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/api/user/medicine/schedule/"+medicine_id, bytes.NewBuffer(reqJSON))
	req.AddCookie(CookieConfiguration())

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected status code %d but got %d", http.StatusCreated, rec.Code)
	}

	var expectedBody successExpectedMedicineScheduleResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success to Add Medicine Schedule" {
		t.Errorf("expected message %s but got %s", "Success to Add Medicine Schedule", expectedBody.Message)
	}
}

func TestAddMedicineSchedule_Failed(t *testing.T) {
	router := gin.Default()

	medicine_id := "a35ac9a3-a3f2-4e29-a696-7a87ad419faz"

	router.POST("/api/user/medicine/schedule/:medicine_id", middleware.RequireAuth, controllers.AddMedicineSchedule)

	reqBody := models.MedicineSchedule{
		Date: "2002",
	}
	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/api/user/medicine/schedule/"+medicine_id, bytes.NewBuffer(reqJSON))
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

	var expectedBody failedExpectedMedicineScheduleResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Invalid Medicine ID" {
		t.Errorf("expected message %s but got %s", "Success to Add Medicine Schedule", expectedBody.ErrorMessage)
	}
}

func TestListMedicineSchedule_Success(t *testing.T) {
	router := gin.Default()

	router.GET("/api/user/medicine/schedule", middleware.RequireAuth, controllers.ListMedicineSchedule)

	req, err := http.NewRequest("GET", "/api/user/medicine/schedule", nil)
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

	var expectedBody successExpectedMedicineScheduleResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success to Get Medicine Schedule List" {
		t.Errorf("expected message %s but got %s", "Success to Get Medicine Schedule List", expectedBody.Message)
	}
}

func TestListMedicineSchedule_Failed(t *testing.T) {
	router := gin.Default()

	router.GET("/api/user/medicine/schedule", middleware.RequireAuth, controllers.ListMedicineSchedule)

	req, err := http.NewRequest("GET", "/api/user/medicine/schedule", nil)
	expiresTime, _ := time.Parse(time.RFC1123, "Mon, 15 Apr 2024 17:00:20 GMT")
	req.AddCookie(
		&http.Cookie{
			Name:     "Authorization",
			Value:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ8.eyJleHAiOjE3MTMyMDA0MjAsInN1YiI6IjE0YzlhNDQzLTAzZTUtNGJhNi05NjY0LTBmODIwYjE5ZDhhYiJ9.L9z84gPX0l_O3GeyRi0ZAhMGxoWzXVV7k9fXw6KpEo4",
			Path:     "/",
			HttpOnly: true,
			Expires:  expiresTime,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d but got %d", http.StatusBadRequest, rec.Code)
	}

}

func TestUpdateMedicineSchedule_Success(t *testing.T) {
	router := gin.Default()

	router.PUT("/api/user/medicine/schedule/:schedule_id", middleware.RequireAuth, controllers.UpdateMedicineSchedule)

	schedule_id := "7f902be6-3590-474f-9d66-777af8a973f9"

	req, err := http.NewRequest("PUT", "/api/user/medicine/schedule/"+schedule_id, bytes.NewBuffer(nil))
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

	var expectedBody successExpectedMedicineScheduleResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success to Update Medicine Schedule" {
		t.Errorf("expected message %s but got %s", "Success to Update Medicine Schedule", expectedBody.Message)
	}
}

func TestUpdateMedicineSchedule_Failed(t *testing.T) {
	router := gin.Default()

	router.PUT("/api/user/medicine/schedule/:schedule_id", middleware.RequireAuth, controllers.UpdateMedicineSchedule)

	schedule_id := "df636bd0-cb0e-4964-9b0a-285fcfd36d21"

	req, err := http.NewRequest("PUT", "/api/user/medicine/schedule/"+schedule_id, bytes.NewBuffer(nil))
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

	var expectedBody failedExpectedMedicineScheduleResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Failed to Find Medicine Schedule" {
		t.Errorf("expected message %s but got %s", "Failed to Find Medicine Schedule", expectedBody.ErrorMessage)
	}
}

func TestDeleteMedicineSchedule_Success(t *testing.T) {
	router := gin.Default()

	router.DELETE("/api/user/medicine/schedule/:schedule_id", middleware.RequireAuth, controllers.DeleteMedicineSchedule)

	schedule_id := "25006641-4f28-4248-8b58-8de54666e210"

	req, err := http.NewRequest("DELETE", "/api/user/medicine/schedule/"+schedule_id, bytes.NewBuffer(nil))
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

	var expectedBody successExpectedMedicineScheduleResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success to Delete Medicine Schedule" {
		t.Errorf("expected message %s but got %s", "Success to Delete Medicine Schedule", expectedBody.Message)
	}
}

func TestDeleteMedicineSchedule_Failed(t *testing.T) {
	router := gin.Default()

	router.DELETE("/api/user/medicine/schedule/:schedule_id", middleware.RequireAuth, controllers.DeleteMedicineSchedule)

	schedule_id := "df636bd0-cb0e-4964-9b0a-285fcfd36d23"

	req, err := http.NewRequest("DELETE", "/api/user/medicine/schedule/"+schedule_id, bytes.NewBuffer(nil))
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

	var expectedBody failedExpectedMedicineScheduleResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Failed to Find Medicine Schedule" {
		t.Errorf("expected message %s but got %s", "Failed to Find Medicine Schedule", expectedBody.ErrorMessage)
	}
}

func TestListMedicineWebsite_Success(t *testing.T) {
	router := gin.Default()

	router.GET("/api/user/medicine/list/website", controllers.ListMedicineWebsite)

	req, err := http.NewRequest("GET", "/api/user/medicine/list/website", nil)

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.AddCookie(WebsiteBearierTokenConfiguration())
	bearerToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImlkIjoiNGUyZDY4NDgtM2FlMy00NjdjLTk5NzQtZTVkMTdhYWJhMGU3IiwidXNlcm5hbWUiOiJwZW5ndXJ1c2VsZ2VrYSIsImZ1bGxfbmFtZSI6IlBlbmd1cnVzIFV0YW1hIEVMR0VLQSBKYXdhIEJhcmF0IiwiaXNfYWN0aXZlIjp0cnVlLCJzdXBlclVzZXIiOnRydWV9LCJpYXQiOjE3MjA5MzQ3NTYsImV4cCI6MTcyMDk1NjM1Nn0.LAOceBARcHrhL3JXtts4WF7TOX7uGJBnYOh_t8GQAiM"
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

	if expectedBody.Message != "Success to Get Medicine List Website" {
		t.Errorf("expected message body %s but got %s", "Success to Get Medicine List Website", expectedBody.Message)
	}
}

func TestListMedicineWebsite_Failed(t *testing.T) {
	router := gin.Default()

	router.GET("/api/user/medicine/list/website", controllers.ListMedicineWebsite)

	req, err := http.NewRequest("GET", "/api/user/medicine/list/website", nil)

	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d but got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestListMedicinePatientWebsite_Success(t *testing.T) {
	router := gin.Default()

	router.GET("/api/user/medicine/list_patient/website", controllers.ListPatientMedicineWebsite)

	req, err := http.NewRequest("GET", "/api/user/medicine/list_patient/website", nil)

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.AddCookie(WebsiteBearierTokenConfiguration())
	bearerToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImlkIjoiNGUyZDY4NDgtM2FlMy00NjdjLTk5NzQtZTVkMTdhYWJhMGU3IiwidXNlcm5hbWUiOiJwZW5ndXJ1c2VsZ2VrYSIsImZ1bGxfbmFtZSI6IlBlbmd1cnVzIFV0YW1hIEVMR0VLQSBKYXdhIEJhcmF0IiwiaXNfYWN0aXZlIjp0cnVlLCJzdXBlclVzZXIiOnRydWV9LCJpYXQiOjE3MjA5MzQ3NTYsImV4cCI6MTcyMDk1NjM1Nn0.LAOceBARcHrhL3JXtts4WF7TOX7uGJBnYOh_t8GQAiM"
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

	if expectedBody.Message != "Success to Get Patient Medicine List Website" {
		t.Errorf("expected message body %s but got %s", "Success to Get Patient Medicine List Website", expectedBody.Message)
	}
}

func TestListMedicinePatientWebsite_Failed(t *testing.T) {
	router := gin.Default()

	router.GET("/api/user/medicine/list_patient/website", controllers.ListPatientMedicineWebsite)

	req, err := http.NewRequest("GET", "/api/user/medicine/list_patient/website", nil)

	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d but got %d", http.StatusBadRequest, rec.Code)
	}
}
