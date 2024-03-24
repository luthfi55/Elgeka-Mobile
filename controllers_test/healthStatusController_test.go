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

type ExpectedHSSuccessResponse struct {
	Message string `json:"Message"`
}

type ExpectedHSFailedResponse struct {
	ErrorMessage string `json:"ErrorMessage"`
}

//BCR ABL

func TestCreateBcrAbl_Success(t *testing.T) {

	router := gin.Default()

	router.POST("/api/user/health_status/bcr_abl", middleware.RequireAuth, controllers.CreateBcrAbl)

	reqBody := models.BCR_ABL{
		Data:  3.2,
		Notes: "tes",
		Date:  "2024-02-23",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/api/user/health_status/bcr_abl", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.AddCookie(CookieConfiguration())

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected status code %d but got %d", http.StatusCreated, rec.Code)
	}

	var expectedBody ExpectedHSSuccessResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success Create Data" {
		t.Errorf("expected message body %s but got %s", "Success Create Data", expectedBody.Message)
	}
}

func TestCreateBcrAbl_Failed(t *testing.T) {

	router := gin.Default()

	router.POST("/api/user/health_status/bcr_abl", middleware.RequireAuth, controllers.CreateBcrAbl)

	reqBody := models.BCR_ABL{
		Data:  3.2,
		Notes: "",
		Date:  "2024-02-23",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/api/user/health_status/bcr_abl", bytes.NewBuffer(reqJSON))
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

	var expectedBody ExpectedHSFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Notes Cant Be Empty" {
		t.Errorf("expected message body %s but got %s", "Success Create Data", expectedBody.ErrorMessage)
	}
}

func TestGetBcrAbl_Success(t *testing.T) {
	router := gin.Default()

	router.GET("/api/user/health_status/bcr_abl", middleware.RequireAuth, controllers.GetBcrAbl)

	req, err := http.NewRequest("GET", "/api/user/health_status/bcr_abl", nil)
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

	var expectedBody ExpectedHSSuccessResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success Get Data" {
		t.Errorf("expected message body %s but got %s", "Success Get Data", expectedBody.Message)
	}
}

func TestGetBcrAbl_Failed(t *testing.T) {
	router := gin.Default()

	router.GET("/api/user/health_status/bcr_abl", middleware.RequireAuth, controllers.GetBcrAbl)

	req, err := http.NewRequest("GET", "/api/user/health_status/bcr_abl", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	expiresTime, _ := time.Parse(time.RFC1123, "Mon, 15 Apr 2024 17:00:20 GMT")
	req.AddCookie(&http.Cookie{
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

func TestUpdateBcrAbl_Success(t *testing.T) {
	router := gin.Default()

	router.PUT("/api/user/health_status/bcr_abl/:bcr_abl_id", middleware.RequireAuth, controllers.UpdateBcrAbl)

	bcr_abl_id := "e3426e5e-d0cd-4078-ab1b-be1cfc7b2768"

	reqBody := models.BCR_ABL{
		Data:  1.2,
		Notes: "Halo",
		Date:  "2024-02-22",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("PUT", "/api/user/health_status/bcr_abl/"+bcr_abl_id, bytes.NewBuffer(reqJSON))
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

	var expectedBody ExpectedHSSuccessResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success Update Data" {
		t.Errorf("expected message body %s but got %s", "Success Update Data", expectedBody.Message)
	}
}

func TestUpdateBcrAbl_Failed(t *testing.T) {
	router := gin.Default()

	router.PUT("/api/user/health_status/bcr_abl/:bcr_abl_id", middleware.RequireAuth, controllers.UpdateBcrAbl)

	bcr_abl_id := "e3426e5e-d0cd-4078-ab1b-be1cfc7b2768"

	reqBody := models.BCR_ABL{
		Data:  1.2,
		Notes: "",
		Date:  "2024-02-22",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("PUT", "/api/user/health_status/bcr_abl/"+bcr_abl_id, bytes.NewBuffer(reqJSON))
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

	var expectedBody ExpectedHSFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Notes Cant Be Empty" {
		t.Errorf("expected message body %s but got %s", "Notes Cant Be Empty", expectedBody.ErrorMessage)
	}
}

func TestDeleteBcrAbl_Success(t *testing.T) {

	router := gin.Default()

	router.DELETE("/api/user/health_status/bcr_abl/:bcr_abl_id", middleware.RequireAuth, controllers.DeleteBcrAbl)

	bcr_abl_id := "92e66f76-332e-45a8-a9d3-a919f01abdd3"

	req, err := http.NewRequest("DELETE", "/api/user/health_status/bcr_abl/"+bcr_abl_id, nil)

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

	var expectedBody ExpectedHSSuccessResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)

	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success Delete Data" {
		t.Errorf("expected message body %s but got %s", "Success Delete Data", expectedBody.Message)
	}
}

func TestDeleteBcrAbl_Failed(t *testing.T) {

	router := gin.Default()

	router.DELETE("/api/user/health_status/bcr_abl/:bcr_abl_id", middleware.RequireAuth, controllers.DeleteBcrAbl)

	bcr_abl_id := "92e66f76-332e-45a8-a9d3-a919f01abde6"

	req, err := http.NewRequest("DELETE", "/api/user/health_status/bcr_abl/"+bcr_abl_id, nil)

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

	var expectedBody ExpectedHSFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)

	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Record Not Found" {
		t.Errorf("expected message body %s but got %s", "Record Not Found", expectedBody.ErrorMessage)
	}
}

// Leukocytes

func TestCreateLeukocytes_Success(t *testing.T) {

	router := gin.Default()

	router.POST("/api/user/health_status/leukocytes", middleware.RequireAuth, controllers.CreateLeukocytes)

	reqBody := models.Leukocytes{
		Data:  7.2,
		Notes: "tes",
		Date:  "2024-02-23",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/api/user/health_status/leukocytes", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.AddCookie(CookieConfiguration())

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected status code %d but got %d", http.StatusCreated, rec.Code)
	}

	var expectedBody ExpectedHSSuccessResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success Create Data" {
		t.Errorf("expected message body %s but got %s", "Success Create Data", expectedBody.Message)
	}
}

func TestCreateLeukocytes_Failed(t *testing.T) {

	router := gin.Default()

	router.POST("/api/user/health_status/leukocytes", middleware.RequireAuth, controllers.CreateLeukocytes)

	reqBody := models.Leukocytes{
		Data:  3.2,
		Notes: "",
		Date:  "2024-02-23",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/api/user/health_status/leukocytes", bytes.NewBuffer(reqJSON))
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

	var expectedBody ExpectedHSFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Notes Cant Be Empty" {
		t.Errorf("expected message body %s but got %s", "Success Create Data", expectedBody.ErrorMessage)
	}
}

func TestGetLeukocytes_Success(t *testing.T) {
	router := gin.Default()

	router.GET("/api/user/health_status/leukocytes", middleware.RequireAuth, controllers.GetLeukocytes)

	req, err := http.NewRequest("GET", "/api/user/health_status/leukocytes", nil)
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

	var expectedBody ExpectedHSSuccessResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success Get Data" {
		t.Errorf("expected message body %s but got %s", "Success Get Data", expectedBody.Message)
	}
}

func TestGetLeukocytes_Failed(t *testing.T) {
	router := gin.Default()

	router.GET("/api/user/health_status/leukocytes", middleware.RequireAuth, controllers.GetLeukocytes)

	req, err := http.NewRequest("GET", "/api/user/health_status/leukocytes", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	expiresTime, _ := time.Parse(time.RFC1123, "Mon, 15 Apr 2024 17:00:20 GMT")
	req.AddCookie(&http.Cookie{
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

func TestUpdateLeukocytes_Success(t *testing.T) {
	router := gin.Default()

	router.PUT("/api/user/health_status/leukocytes/:leukocytes_id", middleware.RequireAuth, controllers.UpdateLeukocytes)

	leukocytes_id := "8ca4dd68-a02e-424d-858d-540b3b808573"

	reqBody := models.Leukocytes{
		Data:  1.2,
		Notes: "Halo",
		Date:  "2024-02-22",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("PUT", "/api/user/health_status/leukocytes/"+leukocytes_id, bytes.NewBuffer(reqJSON))
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

	var expectedBody ExpectedHSSuccessResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success Update Data" {
		t.Errorf("expected message body %s but got %s", "Success Update Data", expectedBody.Message)
	}
}

func TestUpdateLeukocytes_Failed(t *testing.T) {
	router := gin.Default()

	router.PUT("/api/user/health_status/leukocytes/:leukocytes_id", middleware.RequireAuth, controllers.UpdateLeukocytes)

	leukocytes_id := "be622800-dbe1-42ec-a53e-c0179175aa7c"

	reqBody := models.Leukocytes{
		Data:  1.2,
		Notes: "",
		Date:  "2024-02-22",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("PUT", "/api/user/health_status/leukocytes/"+leukocytes_id, bytes.NewBuffer(reqJSON))
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

	var expectedBody ExpectedHSFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Notes Cant Be Empty" {
		t.Errorf("expected message body %s but got %s", "Notes Cant Be Empty", expectedBody.ErrorMessage)
	}
}

func TestDeleteLeukocytes_Success(t *testing.T) {

	router := gin.Default()

	router.DELETE("/api/user/health_status/leukocytes/:leukocytes_id", middleware.RequireAuth, controllers.DeleteLeukocytes)

	leukocytes_id := "be622800-dbe1-42ec-a53e-c0179175aa7c"

	req, err := http.NewRequest("DELETE", "/api/user/health_status/leukocytes/"+leukocytes_id, nil)

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

	var expectedBody ExpectedHSSuccessResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)

	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success Delete Data" {
		t.Errorf("expected message body %s but got %s", "Success Delete Data", expectedBody.Message)
	}
}

func TestDeleteLeukocytes_Failed(t *testing.T) {

	router := gin.Default()

	router.DELETE("/api/user/health_status/leukocytes/:leukocytes_id", middleware.RequireAuth, controllers.DeleteLeukocytes)

	leukocytes_id := "92e66f76-332e-45a8-a9d3-a919f01abde6"

	req, err := http.NewRequest("DELETE", "/api/user/health_status/leukocytes/"+leukocytes_id, nil)

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

	var expectedBody ExpectedHSFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)

	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Record Not Found" {
		t.Errorf("expected message body %s but got %s", "Record Not Found", expectedBody.ErrorMessage)
	}
}

// Potential Hydrogen

func TestCreatePotentialHydrogen_Success(t *testing.T) {

	router := gin.Default()

	router.POST("/api/user/health_status/blood_pressure", middleware.RequireAuth, controllers.CreatePotentialHydrogen)

	reqBody := models.PotentialHydrogen{
		Data:  7.2,
		Notes: "tes",
		Date:  "2024-02-23",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/api/user/health_status/potential_hydrogen", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.AddCookie(CookieConfiguration())

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected status code %d but got %d", http.StatusCreated, rec.Code)
	}

	var expectedBody ExpectedHSSuccessResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success Create Data" {
		t.Errorf("expected message body %s but got %s", "Success Create Data", expectedBody.Message)
	}
}

func TestCreatePotentialHydrogen_Failed(t *testing.T) {

	router := gin.Default()

	router.POST("/api/user/health_status/potential_hydrogen", middleware.RequireAuth, controllers.CreatePotentialHydrogen)

	reqBody := models.PotentialHydrogen{
		Data:  3.2,
		Notes: "",
		Date:  "2024-02-23",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/api/user/health_status/potential_hydrogen", bytes.NewBuffer(reqJSON))
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

	var expectedBody ExpectedHSFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Notes Cant Be Empty" {
		t.Errorf("expected message body %s but got %s", "Success Create Data", expectedBody.ErrorMessage)
	}
}

func TestGetPotentialHydrogen_Success(t *testing.T) {
	router := gin.Default()

	router.GET("/api/user/health_status/potential_hydrogen", middleware.RequireAuth, controllers.GetPotentialHydrogen)

	req, err := http.NewRequest("GET", "/api/user/health_status/potential_hydrogen", nil)
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

	var expectedBody ExpectedHSSuccessResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success Get Data" {
		t.Errorf("expected message body %s but got %s", "Success Get Data", expectedBody.Message)
	}
}

func TestGetPotentialHydrogen_Failed(t *testing.T) {
	router := gin.Default()

	router.GET("/api/user/health_status/potential_hydrogen", middleware.RequireAuth, controllers.GetPotentialHydrogen)

	req, err := http.NewRequest("GET", "/api/user/health_status/potential_hydrogen", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	expiresTime, _ := time.Parse(time.RFC1123, "Mon, 15 Apr 2024 17:00:20 GMT")
	req.AddCookie(&http.Cookie{
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

func TestUpdatePotentialHydrogen_Success(t *testing.T) {
	router := gin.Default()

	router.PUT("/api/user/health_status/potential_hydrogen/:potential_hydrogen_id", middleware.RequireAuth, controllers.UpdatePotentialHydrogen)

	potential_hydrogen_id := "3fbdfc21-4548-4c33-8f9e-80736e6d18f1"

	reqBody := models.PotentialHydrogen{
		Data:  1.2,
		Notes: "Halo",
		Date:  "2024-02-22",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("PUT", "/api/user/health_status/potential_hydrogen/"+potential_hydrogen_id, bytes.NewBuffer(reqJSON))
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

	var expectedBody ExpectedHSSuccessResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success Update Data" {
		t.Errorf("expected message body %s but got %s", "Success Update Data", expectedBody.Message)
	}
}

func TestUpdatePotentialHydrogen_Failed(t *testing.T) {
	router := gin.Default()

	router.PUT("/api/user/health_status/potential_hydrogen/:potential_hydrogen_id", middleware.RequireAuth, controllers.UpdatePotentialHydrogen)

	potential_hydrogen_id := "be622800-dbe1-42ec-a53e-c0179175aa7c"

	reqBody := models.PotentialHydrogen{
		Data:  1.2,
		Notes: "",
		Date:  "2024-02-22",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("PUT", "/api/user/health_status/potential_hydrogen/"+potential_hydrogen_id, bytes.NewBuffer(reqJSON))
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

	var expectedBody ExpectedHSFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Notes Cant Be Empty" {
		t.Errorf("expected message body %s but got %s", "Notes Cant Be Empty", expectedBody.ErrorMessage)
	}
}

func TestDeletePotentialHydrogen_Success(t *testing.T) {

	router := gin.Default()

	router.DELETE("/api/user/health_status/potential_hydrogen/:potential_hydrogen_id", middleware.RequireAuth, controllers.DeletePotentialHydrogen)

	potential_hydrogen_id := "90469cd0-c309-4998-b25e-5b94c259e7d1"

	req, err := http.NewRequest("DELETE", "/api/user/health_status/potential_hydrogen/"+potential_hydrogen_id, nil)

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

	var expectedBody ExpectedHSSuccessResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)

	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success Delete Data" {
		t.Errorf("expected message body %s but got %s", "Success Delete Data", expectedBody.Message)
	}
}

func TestDeletePotentialHydrogen_Failed(t *testing.T) {

	router := gin.Default()

	router.DELETE("/api/user/health_status/potential_hydrogen/:potential_hydrogen_id", middleware.RequireAuth, controllers.DeletePotentialHydrogen)

	potential_hydrogen_id := "92e66f76-332e-45a8-a9d3-a919f01abde6"

	req, err := http.NewRequest("DELETE", "/api/user/health_status/potential_hydrogen/"+potential_hydrogen_id, nil)

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

	var expectedBody ExpectedHSFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)

	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Record Not Found" {
		t.Errorf("expected message body %s but got %s", "Record Not Found", expectedBody.ErrorMessage)
	}
}

// Hemoglobin

func TestCreateHemoglobin_Success(t *testing.T) {

	router := gin.Default()

	router.POST("/api/user/health_status/hemoglobin", middleware.RequireAuth, controllers.CreateHemoglobin)

	reqBody := models.Hemoglobin{
		Data:  7.2,
		Notes: "tes",
		Date:  "2024-02-23",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/api/user/health_status/hemoglobin", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.AddCookie(CookieConfiguration())

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected status code %d but got %d", http.StatusCreated, rec.Code)
	}

	var expectedBody ExpectedHSSuccessResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success Create Data" {
		t.Errorf("expected message body %s but got %s", "Success Create Data", expectedBody.Message)
	}
}

func TestCreateHemoglobin_Failed(t *testing.T) {

	router := gin.Default()

	router.POST("/api/user/health_status/hemoglobin", middleware.RequireAuth, controllers.CreateHemoglobin)

	reqBody := models.Hemoglobin{
		Data:  3.2,
		Notes: "",
		Date:  "2024-02-23",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/api/user/health_status/hemoglobin", bytes.NewBuffer(reqJSON))
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

	var expectedBody ExpectedHSFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Notes Cant Be Empty" {
		t.Errorf("expected message body %s but got %s", "Success Create Data", expectedBody.ErrorMessage)
	}
}

func TestGetHemoglobin_Success(t *testing.T) {
	router := gin.Default()

	router.GET("/api/user/health_status/hemoglobin", middleware.RequireAuth, controllers.GetHemoglobin)

	req, err := http.NewRequest("GET", "/api/user/health_status/hemoglobin", nil)
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

	var expectedBody ExpectedHSSuccessResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success Get Data" {
		t.Errorf("expected message body %s but got %s", "Success Get Data", expectedBody.Message)
	}
}

func TestGetHemoglobin_Failed(t *testing.T) {
	router := gin.Default()

	router.GET("/api/user/health_status/hemoglobin", middleware.RequireAuth, controllers.GetHemoglobin)

	req, err := http.NewRequest("GET", "/api/user/health_status/hemoglobin", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	expiresTime, _ := time.Parse(time.RFC1123, "Mon, 15 Apr 2024 17:00:20 GMT")
	req.AddCookie(&http.Cookie{
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

func TestUpdateHemoglobin_Success(t *testing.T) {
	router := gin.Default()

	router.PUT("/api/user/health_status/hemoglobin/:hemoglobin_id", middleware.RequireAuth, controllers.UpdateHemoglobin)

	hemoglobin_id := "a93592d0-a02c-4be8-8bb7-da1a92f0113e"

	reqBody := models.Hemoglobin{
		Data:  1.2,
		Notes: "Halo",
		Date:  "2024-02-22",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("PUT", "/api/user/health_status/hemoglobin/"+hemoglobin_id, bytes.NewBuffer(reqJSON))
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

	var expectedBody ExpectedHSSuccessResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success Update Data" {
		t.Errorf("expected message body %s but got %s", "Success Update Data", expectedBody.Message)
	}
}

func TestUpdateHemoglobin_Failed(t *testing.T) {
	router := gin.Default()

	router.PUT("/api/user/health_status/hemoglobin/:hemoglobin_id", middleware.RequireAuth, controllers.UpdateHemoglobin)

	hemoglobin_id := "a93592d0-a02c-4be8-8bb7-da1a92f0113e"

	reqBody := models.Hemoglobin{
		Data:  1.2,
		Notes: "",
		Date:  "2024-02-22",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("PUT", "/api/user/health_status/hemoglobin/"+hemoglobin_id, bytes.NewBuffer(reqJSON))
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

	var expectedBody ExpectedHSFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Notes Cant Be Empty" {
		t.Errorf("expected message body %s but got %s", "Notes Cant Be Empty", expectedBody.ErrorMessage)
	}
}

func TestDeleteHemoglobin_Success(t *testing.T) {

	router := gin.Default()

	router.DELETE("/api/user/health_status/hemoglobin/:hemoglobin_id", middleware.RequireAuth, controllers.DeleteHemoglobin)

	hemoglobin_id := "77082cda-5f2a-4ae1-ae56-586dc1d24a08"

	req, err := http.NewRequest("DELETE", "/api/user/health_status/hemoglobin/"+hemoglobin_id, nil)

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

	var expectedBody ExpectedHSSuccessResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)

	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success Delete Data" {
		t.Errorf("expected message body %s but got %s", "Success Delete Data", expectedBody.Message)
	}
}

func TestDeleteHemoglobin_Failed(t *testing.T) {

	router := gin.Default()

	router.DELETE("/api/user/health_status/hemoglobin/:hemoglobin_id", middleware.RequireAuth, controllers.DeleteHemoglobin)

	hemoglobin_id := "92e66f76-332e-45a8-a9d3-a919f01abde6"

	req, err := http.NewRequest("DELETE", "/api/user/health_status/hemoglobin/"+hemoglobin_id, nil)

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

	var expectedBody ExpectedHSFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)

	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Record Not Found" {
		t.Errorf("expected message body %s but got %s", "Record Not Found", expectedBody.ErrorMessage)
	}
}

// Heart Rate

func TestCreateHeartRate_Success(t *testing.T) {

	router := gin.Default()

	router.POST("/api/user/health_status/heart_rate", middleware.RequireAuth, controllers.CreateHeartRate)

	reqBody := models.HeartRate{
		Data:  7.2,
		Notes: "tes",
		Date:  "2024-02-23",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/api/user/health_status/heart_rate", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.AddCookie(CookieConfiguration())

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected status code %d but got %d", http.StatusCreated, rec.Code)
	}

	var expectedBody ExpectedHSSuccessResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success Create Data" {
		t.Errorf("expected message body %s but got %s", "Success Create Data", expectedBody.Message)
	}
}

func TestCreateHeartRate_Failed(t *testing.T) {

	router := gin.Default()

	router.POST("/api/user/health_status/heart_rate", middleware.RequireAuth, controllers.CreateHeartRate)

	reqBody := models.HeartRate{
		Data:  3.2,
		Notes: "",
		Date:  "2024-02-23",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/api/user/health_status/heart_rate", bytes.NewBuffer(reqJSON))
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

	var expectedBody ExpectedHSFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Notes Cant Be Empty" {
		t.Errorf("expected message body %s but got %s", "Success Create Data", expectedBody.ErrorMessage)
	}
}

func TestGetHeartRate_Success(t *testing.T) {
	router := gin.Default()

	router.GET("/api/user/health_status/heart_rate", middleware.RequireAuth, controllers.GetHeartRate)

	req, err := http.NewRequest("GET", "/api/user/health_status/heart_rate", nil)
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

	var expectedBody ExpectedHSSuccessResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success Get Data" {
		t.Errorf("expected message body %s but got %s", "Success Get Data", expectedBody.Message)
	}
}

func TestGetHeartRate_Failed(t *testing.T) {
	router := gin.Default()

	router.GET("/api/user/health_status/heart_rate", middleware.RequireAuth, controllers.GetHeartRate)

	req, err := http.NewRequest("GET", "/api/user/health_status/heart_rate", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	expiresTime, _ := time.Parse(time.RFC1123, "Mon, 15 Apr 2024 17:00:20 GMT")
	req.AddCookie(&http.Cookie{
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

func TestUpdateHeartRate_Success(t *testing.T) {
	router := gin.Default()

	router.PUT("/api/user/health_status/heart_rate/:heart_rate_id", middleware.RequireAuth, controllers.UpdateHeartRate)

	heart_rate_id := "aa85f80b-59c0-4323-918e-ce6618068555"

	reqBody := models.HeartRate{
		Data:  1.2,
		Notes: "Halo",
		Date:  "2024-02-22",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("PUT", "/api/user/health_status/heart_rate/"+heart_rate_id, bytes.NewBuffer(reqJSON))
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

	var expectedBody ExpectedHSSuccessResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success Update Data" {
		t.Errorf("expected message body %s but got %s", "Success Update Data", expectedBody.Message)
	}
}

func TestUpdateHeartRate_Failed(t *testing.T) {
	router := gin.Default()

	router.PUT("/api/user/health_status/heart_rate/:heart_rate_id", middleware.RequireAuth, controllers.UpdateHeartRate)

	heart_rate_id := "aa85f80b-59c0-4323-918e-ce6618068555"

	reqBody := models.HeartRate{
		Data:  1.2,
		Notes: "",
		Date:  "2024-02-22",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("PUT", "/api/user/health_status/heart_rate/"+heart_rate_id, bytes.NewBuffer(reqJSON))
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

	var expectedBody ExpectedHSFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Notes Cant Be Empty" {
		t.Errorf("expected message body %s but got %s", "Notes Cant Be Empty", expectedBody.ErrorMessage)
	}
}

func TestDeleteHeartRate_Success(t *testing.T) {

	router := gin.Default()

	router.DELETE("/api/user/health_status/heart_rate/:heart_rate_id", middleware.RequireAuth, controllers.DeleteHeartRate)

	heart_rate_id := "8a42b0b7-940f-44c6-b727-71f7a5e1354f"

	req, err := http.NewRequest("DELETE", "/api/user/health_status/heart_rate/"+heart_rate_id, nil)

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

	var expectedBody ExpectedHSSuccessResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)

	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success Delete Data" {
		t.Errorf("expected message body %s but got %s", "Success Delete Data", expectedBody.Message)
	}
}

func TestDeleteHeartRate_Failed(t *testing.T) {

	router := gin.Default()

	router.DELETE("/api/user/health_status/heart_rate/:heart_rate_id", middleware.RequireAuth, controllers.DeleteHeartRate)

	heart_rate_id := "92e66f76-332e-45a8-a9d3-a919f01abde6"

	req, err := http.NewRequest("DELETE", "/api/user/health_status/heart_rate/"+heart_rate_id, nil)

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

	var expectedBody ExpectedHSFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)

	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Record Not Found" {
		t.Errorf("expected message body %s but got %s", "Record Not Found", expectedBody.ErrorMessage)
	}
}

// Blood Pressure

func TestCreateBloodPressure_Success(t *testing.T) {

	router := gin.Default()

	router.POST("/api/user/health_status/blood_pressure", middleware.RequireAuth, controllers.CreateBloodPressure)

	reqBody := models.BloodPressure{
		DataSys: 7.2,
		DataDia: 6.3,
		Notes:   "tes",
		Date:    "2024-02-23",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/api/user/health_status/blood_pressure", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.AddCookie(CookieConfiguration())

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected status code %d but got %d", http.StatusCreated, rec.Code)
	}

	var expectedBody ExpectedHSSuccessResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success Create Data" {
		t.Errorf("expected message body %s but got %s", "Success Create Data", expectedBody.Message)
	}
}

func TestCreateBloodPressure_Failed(t *testing.T) {

	router := gin.Default()

	router.POST("/api/user/health_status/blood_pressure", middleware.RequireAuth, controllers.CreateBloodPressure)

	reqBody := models.BloodPressure{
		DataSys: 7.2,
		DataDia: 6.3,
		Notes:   "",
		Date:    "2024-02-23",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/api/user/health_status/blood_pressure", bytes.NewBuffer(reqJSON))
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

	var expectedBody ExpectedHSFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Notes Cant Be Empty" {
		t.Errorf("expected message body %s but got %s", "Success Create Data", expectedBody.ErrorMessage)
	}
}

func TestGetBloodPressure_Success(t *testing.T) {
	router := gin.Default()

	router.GET("/api/user/health_status/blood_pressure", middleware.RequireAuth, controllers.GetBloodPressure)

	req, err := http.NewRequest("GET", "/api/user/health_status/blood_pressure", nil)
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

	var expectedBody ExpectedHSSuccessResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success Get Data" {
		t.Errorf("expected message body %s but got %s", "Success Get Data", expectedBody.Message)
	}
}

func TestGetBloodPressure_Failed(t *testing.T) {
	router := gin.Default()

	router.GET("/api/user/health_status/blood_pressure", middleware.RequireAuth, controllers.GetBloodPressure)

	req, err := http.NewRequest("GET", "/api/user/health_status/blood_pressure", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	expiresTime, _ := time.Parse(time.RFC1123, "Mon, 15 Apr 2024 17:00:20 GMT")
	req.AddCookie(&http.Cookie{
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

func TestUpdateBloodPressure_Success(t *testing.T) {
	router := gin.Default()

	router.PUT("/api/user/health_status/blood_pressure/:blood_pressure_id", middleware.RequireAuth, controllers.UpdateBloodPressure)

	blood_pressure_id := "46142432-36bc-4fd2-8f56-aa5d9804bfdd"

	reqBody := models.BloodPressure{
		DataSys: 1.2,
		DataDia: 3.3,
		Notes:   "Halo",
		Date:    "2024-02-22",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("PUT", "/api/user/health_status/blood_pressure/"+blood_pressure_id, bytes.NewBuffer(reqJSON))
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

	var expectedBody ExpectedHSSuccessResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success Update Data" {
		t.Errorf("expected message body %s but got %s", "Success Update Data", expectedBody.Message)
	}
}

func TestUpdateBloodPressure_Failed(t *testing.T) {
	router := gin.Default()

	router.PUT("/api/user/health_status/blood_pressure/:blood_pressure_id", middleware.RequireAuth, controllers.UpdateBloodPressure)

	blood_pressure_id := "be622800-dbe1-42ec-a53e-c0179175aa7c"

	reqBody := models.BloodPressure{
		DataSys: 7.2,
		DataDia: 6.3,
		Notes:   "",
		Date:    "2024-02-22",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("PUT", "/api/user/health_status/blood_pressure/"+blood_pressure_id, bytes.NewBuffer(reqJSON))
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

	var expectedBody ExpectedHSFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Notes Cant Be Empty" {
		t.Errorf("expected message body %s but got %s", "Notes Cant Be Empty", expectedBody.ErrorMessage)
	}
}

func TestDeleteBloodPressure_Success(t *testing.T) {

	router := gin.Default()

	router.DELETE("/api/user/health_status/blood_pressure/:blood_pressure_id", middleware.RequireAuth, controllers.DeleteBloodPressure)

	blood_pressure_id := "34004c60-3a2e-41a6-825a-55c788f9365e"

	req, err := http.NewRequest("DELETE", "/api/user/health_status/blood_pressure/"+blood_pressure_id, nil)

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

	var expectedBody ExpectedHSSuccessResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)

	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success Delete Data" {
		t.Errorf("expected message body %s but got %s", "Success Delete Data", expectedBody.Message)
	}
}

func TestDeleteBloodPressure_Failed(t *testing.T) {

	router := gin.Default()

	router.DELETE("/api/user/health_status/blood_pressure/:blood_pressure_id", middleware.RequireAuth, controllers.DeleteBloodPressure)

	blood_pressure_id := "92e66f76-332e-45a8-a9d3-a919f01abde6"

	req, err := http.NewRequest("DELETE", "/api/user/health_status/blood_pressure/"+blood_pressure_id, nil)

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

	var expectedBody ExpectedHSFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)

	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Record Not Found" {
		t.Errorf("expected message body %s but got %s", "Record Not Found", expectedBody.ErrorMessage)
	}
}
