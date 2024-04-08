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
		Name:  "Sanmol",
		Stock: 10,
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
		Name:  "",
		Stock: 10,
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

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	if rec.Code != http.StatusUnauthorized {
		t.Errorf("expected status code %d but got %d", http.StatusUnauthorized, rec.Code)
	}
}

func TestUpdateMedicine_Success(t *testing.T) {
	router := gin.Default()
	medicine_id := "d325b785-a18e-4e0f-a264-20135c2a1b63"

	router.PUT("api/user/medicine/:medicine_id", middleware.RequireAuth, controllers.UpdateMedicine)

	reqBody := models.Medicine{
		Name:  "Paracetamol",
		Stock: 100,
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
		Name:  "",
		Stock: 100,
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

	medicine_id := "a35ac9a3-a3f2-4e29-a696-7a87ad419faf"

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

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("expected status code %d but got %d", http.StatusUnauthorized, rec.Code)
	}

}
