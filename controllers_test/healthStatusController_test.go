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

type ExpectedSuccessResponse struct {
	Message string `json:"Message"`
}

type ExpectedFailedResponse struct {
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

	var expectedBody ExpectedSuccessResponse
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

	var expectedBody ExpectedFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Field 'Notes' is required." {
		t.Errorf("expected message body %s but got %s", "Field 'Notes' is required.", expectedBody.ErrorMessage)
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

	var expectedBody ExpectedSuccessResponse
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

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d but got %d", http.StatusBadRequest, rec.Code)
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

	var expectedBody ExpectedSuccessResponse
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

	var expectedBody ExpectedFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Field 'Notes' is required." {
		t.Errorf("expected message body %s but got %s", "Field 'Notes' is required.", expectedBody.ErrorMessage)
	}
}

func TestDeleteBcrAbl_Success(t *testing.T) {

	router := gin.Default()

	router.DELETE("/api/user/health_status/bcr_abl/:bcr_abl_id", middleware.RequireAuth, controllers.DeleteBcrAbl)

	bcr_abl_id := "4bab13c5-5c5f-4d3c-b274-edf19a19310d"

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

	var expectedBody ExpectedSuccessResponse
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

	var expectedBody ExpectedFailedResponse
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

	var expectedBody ExpectedSuccessResponse
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

	var expectedBody ExpectedFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Field 'Notes' is required." {
		t.Errorf("expected message body %s but got %s", "Field 'Notes' is required.", expectedBody.ErrorMessage)
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

	var expectedBody ExpectedSuccessResponse
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

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d but got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestUpdateLeukocytes_Success(t *testing.T) {
	router := gin.Default()

	router.PUT("/api/user/health_status/leukocytes/:leukocytes_id", middleware.RequireAuth, controllers.UpdateLeukocytes)

	leukocytes_id := "968af2b8-f428-486b-a065-157877eee8b5"

	reqBody := models.Leukocytes{
		Data:  20.1,
		Notes: "Leukocytes 3",
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

	var expectedBody ExpectedSuccessResponse
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

	leukocytes_id := "968af2b8-f428-486b-a065-157877eee8b5"

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

	var expectedBody ExpectedFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Field 'Notes' is required." {
		t.Errorf("expected message body %s but got %s", "Field 'Notes' is required.", expectedBody.ErrorMessage)
	}
}

func TestDeleteLeukocytes_Success(t *testing.T) {

	router := gin.Default()

	router.DELETE("/api/user/health_status/leukocytes/:leukocytes_id", middleware.RequireAuth, controllers.DeleteLeukocytes)

	leukocytes_id := "d57e5814-f9b9-4f5b-9992-fd10b57c6e64"

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

	var expectedBody ExpectedSuccessResponse
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

	var expectedBody ExpectedFailedResponse
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

	router.POST("/api/user/health_status/potential_hydrogen", middleware.RequireAuth, controllers.CreatePotentialHydrogen)

	reqBody := models.PotentialHydrogen{
		Data:  5.2,
		Notes: "PH+",
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

	var expectedBody ExpectedSuccessResponse
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

	var expectedBody ExpectedFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Field 'Notes' is required." {
		t.Errorf("expected message body %s but got %s", "Field 'Notes' is required.", expectedBody.ErrorMessage)
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

	var expectedBody ExpectedSuccessResponse
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

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d but got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestUpdatePotentialHydrogen_Success(t *testing.T) {
	router := gin.Default()

	router.PUT("/api/user/health_status/potential_hydrogen/:potential_hydrogen_id", middleware.RequireAuth, controllers.UpdatePotentialHydrogen)

	potential_hydrogen_id := "4d7d43b4-67d6-499f-9f90-976012e33f0d"

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

	var expectedBody ExpectedSuccessResponse
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

	potential_hydrogen_id := "4d7d43b4-67d6-499f-9f90-976012e33f0d"

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

	var expectedBody ExpectedFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Field 'Notes' is required." {
		t.Errorf("expected message body %s but got %s", "Field 'Notes' is required.", expectedBody.ErrorMessage)
	}
}

func TestDeletePotentialHydrogen_Success(t *testing.T) {

	router := gin.Default()

	router.DELETE("/api/user/health_status/potential_hydrogen/:potential_hydrogen_id", middleware.RequireAuth, controllers.DeletePotentialHydrogen)

	potential_hydrogen_id := "4d7d43b4-67d6-499f-9f90-976012e33f0d"

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

	var expectedBody ExpectedSuccessResponse
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

	var expectedBody ExpectedFailedResponse
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

	var expectedBody ExpectedSuccessResponse
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

	var expectedBody ExpectedFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Field 'Notes' is required." {
		t.Errorf("expected message body %s but got %s", "Field 'Notes' is required.", expectedBody.ErrorMessage)
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

	var expectedBody ExpectedSuccessResponse
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

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d but got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestUpdateHemoglobin_Success(t *testing.T) {
	router := gin.Default()

	router.PUT("/api/user/health_status/hemoglobin/:hemoglobin_id", middleware.RequireAuth, controllers.UpdateHemoglobin)

	hemoglobin_id := "05dd5b96-46e2-4da5-9996-9fe507a18bfa"

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

	var expectedBody ExpectedSuccessResponse
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

	hemoglobin_id := "05dd5b96-46e2-4da5-9996-9fe507a18bfa"

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

	var expectedBody ExpectedFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Field 'Notes' is required." {
		t.Errorf("expected message body %s but got %s", "Field 'Notes' is required.", expectedBody.ErrorMessage)
	}
}

func TestDeleteHemoglobin_Success(t *testing.T) {

	router := gin.Default()

	router.DELETE("/api/user/health_status/hemoglobin/:hemoglobin_id", middleware.RequireAuth, controllers.DeleteHemoglobin)

	hemoglobin_id := "05dd5b96-46e2-4da5-9996-9fe507a18bfa"

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

	var expectedBody ExpectedSuccessResponse
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

	var expectedBody ExpectedFailedResponse
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

	var expectedBody ExpectedSuccessResponse
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

	var expectedBody ExpectedFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Field 'Notes' is required." {
		t.Errorf("expected message body %s but got %s", "Field 'Notes' is required.", expectedBody.ErrorMessage)
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

	var expectedBody ExpectedSuccessResponse
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

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d but got %d", http.StatusBadRequest, rec.Code)
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

	var expectedBody ExpectedSuccessResponse
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

	var expectedBody ExpectedFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Field 'Notes' is required." {
		t.Errorf("expected message body %s but got %s", "Field 'Notes' is required.", expectedBody.ErrorMessage)
	}
}

func TestDeleteHeartRate_Success(t *testing.T) {

	router := gin.Default()

	router.DELETE("/api/user/health_status/heart_rate/:heart_rate_id", middleware.RequireAuth, controllers.DeleteHeartRate)

	heart_rate_id := "aa85f80b-59c0-4323-918e-ce6618068555"

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

	var expectedBody ExpectedSuccessResponse
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

	var expectedBody ExpectedFailedResponse
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

	var expectedBody ExpectedSuccessResponse
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

	var expectedBody ExpectedFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Field 'Notes' is required." {
		t.Errorf("expected message body %s but got %s", "Field 'Notes' is required.", expectedBody.ErrorMessage)
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

	var expectedBody ExpectedSuccessResponse
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

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d but got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestUpdateBloodPressure_Success(t *testing.T) {
	router := gin.Default()

	router.PUT("/api/user/health_status/blood_pressure/:blood_pressure_id", middleware.RequireAuth, controllers.UpdateBloodPressure)

	blood_pressure_id := "aa0e2319-b7a3-4e45-bbd5-ed6aa322c60d"

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

	var expectedBody ExpectedSuccessResponse
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

	blood_pressure_id := "aa0e2319-b7a3-4e45-bbd5-ed6aa322c60d"

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

	var expectedBody ExpectedFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Field 'Notes' is required." {
		t.Errorf("expected message body %s but got %s", "Field 'Notes' is required.", expectedBody.ErrorMessage)
	}
}

func TestDeleteBloodPressure_Success(t *testing.T) {

	router := gin.Default()

	router.DELETE("/api/user/health_status/blood_pressure/:blood_pressure_id", middleware.RequireAuth, controllers.DeleteBloodPressure)

	blood_pressure_id := "aa0e2319-b7a3-4e45-bbd5-ed6aa322c60d"

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

	var expectedBody ExpectedSuccessResponse
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

	var expectedBody ExpectedFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)

	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Record Not Found" {
		t.Errorf("expected message body %s but got %s", "Record Not Found", expectedBody.ErrorMessage)
	}
}

// Hematokrit

func TestCreateHematokrit_Success(t *testing.T) {

	router := gin.Default()

	router.POST("/api/user/health_status/hematokrit", middleware.RequireAuth, controllers.CreateHematokrit)

	reqBody := models.Hematokrit{
		Data:  7.2,
		Notes: "tes",
		Date:  "2024-02-23",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/api/user/health_status/hematokrit", bytes.NewBuffer(reqJSON))
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

	var expectedBody ExpectedSuccessResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success Create Data" {
		t.Errorf("expected message body %s but got %s", "Success Create Data", expectedBody.Message)
	}
}

func TestCreateHematokrit_Failed(t *testing.T) {

	router := gin.Default()

	router.POST("/api/user/health_status/hematokrit", middleware.RequireAuth, controllers.CreateHematokrit)

	reqBody := models.Hematokrit{
		Data:  7.2,
		Notes: "",
		Date:  "2024-02-23",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/api/user/health_status/hematokrit", bytes.NewBuffer(reqJSON))
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

	if expectedBody.ErrorMessage != "Field 'Notes' is required." {
		t.Errorf("expected message body %s but got %s", "Field 'Notes' is required.", expectedBody.ErrorMessage)
	}
}

func TestGetHematokrit_Success(t *testing.T) {
	router := gin.Default()

	router.GET("/api/user/health_status/hematokrit", middleware.RequireAuth, controllers.GetHematokrit)

	req, err := http.NewRequest("GET", "/api/user/health_status/hematokrit", nil)
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

	if expectedBody.Message != "Success Get Data" {
		t.Errorf("expected message body %s but got %s", "Success Get Data", expectedBody.Message)
	}
}

func TestGetHematokrit_Failed(t *testing.T) {
	router := gin.Default()

	router.GET("/api/user/health_status/hematokrit", middleware.RequireAuth, controllers.GetHematokrit)

	req, err := http.NewRequest("GET", "/api/user/health_status/hematokrit", nil)
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

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d but got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestUpdateHematokrit_Success(t *testing.T) {
	router := gin.Default()

	router.PUT("/api/user/health_status/hematokrit/:hematokrit_id", middleware.RequireAuth, controllers.UpdateHematokrit)

	hematokrit_id := "05dd5b96-46e2-4da5-9996-9fe507a18bfa"

	reqBody := models.Hematokrit{
		Data:  1.2,
		Notes: "Halo",
		Date:  "2024-02-22",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("PUT", "/api/user/health_status/hematokrit/"+hematokrit_id, bytes.NewBuffer(reqJSON))
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

	if expectedBody.Message != "Success Update Data" {
		t.Errorf("expected message body %s but got %s", "Success Update Data", expectedBody.Message)
	}
}

func TestUpdateHematokrit_Failed(t *testing.T) {
	router := gin.Default()

	router.PUT("/api/user/health_status/hematokrit/:hematokrit_id", middleware.RequireAuth, controllers.UpdateHematokrit)

	hematokrit_id := "05dd5b96-46e2-4da5-9996-9fe507a18bfa"

	reqBody := models.Hematokrit{
		Data:  7.2,
		Notes: "",
		Date:  "2024-02-22",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("PUT", "/api/user/health_status/hematokrit/"+hematokrit_id, bytes.NewBuffer(reqJSON))
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

	if expectedBody.ErrorMessage != "Field 'Notes' is required." {
		t.Errorf("expected message body %s but got %s", "Field 'Notes' is required.", expectedBody.ErrorMessage)
	}
}

func TestDeleteHematokrit_Success(t *testing.T) {

	router := gin.Default()

	router.DELETE("/api/user/health_status/hematokrit/:hematokrit_id", middleware.RequireAuth, controllers.DeleteHematokrit)

	hematokrit_id := "05dd5b96-46e2-4da5-9996-9fe507a18bfa"

	req, err := http.NewRequest("DELETE", "/api/user/health_status/hematokrit/"+hematokrit_id, nil)

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

	if expectedBody.Message != "Success Delete Data" {
		t.Errorf("expected message body %s but got %s", "Success Delete Data", expectedBody.Message)
	}
}

func TestDeleteHematokrit_Failed(t *testing.T) {

	router := gin.Default()

	router.DELETE("/api/user/health_status/hematokrit/:hematokrit_id", middleware.RequireAuth, controllers.DeleteHematokrit)

	hematokrit_id := "92e66f76-332e-45a8-a9d3-a919f01abde6"

	req, err := http.NewRequest("DELETE", "/api/user/health_status/hematokrit/"+hematokrit_id, nil)

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

	if expectedBody.ErrorMessage != "Record Not Found" {
		t.Errorf("expected message body %s but got %s", "Record Not Found", expectedBody.ErrorMessage)
	}
}

// Trombosit

func TestCreateTrombosit_Success(t *testing.T) {

	router := gin.Default()

	router.POST("/api/user/health_status/trombosit", middleware.RequireAuth, controllers.CreateTrombosit)

	reqBody := models.Trombosit{
		Data:  7.2,
		Notes: "tes",
		Date:  "2024-02-23",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/api/user/health_status/trombosit", bytes.NewBuffer(reqJSON))
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

	var expectedBody ExpectedSuccessResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success Create Data" {
		t.Errorf("expected message body %s but got %s", "Success Create Data", expectedBody.Message)
	}
}

func TestCreateTrombosit_Failed(t *testing.T) {

	router := gin.Default()

	router.POST("/api/user/health_status/trombosit", middleware.RequireAuth, controllers.CreateTrombosit)

	reqBody := models.Trombosit{
		Data:  7.2,
		Notes: "",
		Date:  "2024-02-23",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/api/user/health_status/trombosit", bytes.NewBuffer(reqJSON))
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

	if expectedBody.ErrorMessage != "Field 'Notes' is required." {
		t.Errorf("expected message body %s but got %s", "Field 'Notes' is required.", expectedBody.ErrorMessage)
	}
}

func TestGetTrombosit_Success(t *testing.T) {
	router := gin.Default()

	router.GET("/api/user/health_status/trombosit", middleware.RequireAuth, controllers.GetTrombosit)

	req, err := http.NewRequest("GET", "/api/user/health_status/trombosit", nil)
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

	if expectedBody.Message != "Success Get Data" {
		t.Errorf("expected message body %s but got %s", "Success Get Data", expectedBody.Message)
	}
}

func TestGetTrombosit_Failed(t *testing.T) {
	router := gin.Default()

	router.GET("/api/user/health_status/trombosit", middleware.RequireAuth, controllers.GetTrombosit)

	req, err := http.NewRequest("GET", "/api/user/health_status/trombosit", nil)
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

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d but got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestUpdateTrombosit_Success(t *testing.T) {
	router := gin.Default()

	router.PUT("/api/user/health_status/trombosit/:trombosit_id", middleware.RequireAuth, controllers.UpdateTrombosit)

	trombosit_id := "077a0e36-9ec7-4050-ae40-0c0af7245e7e"

	reqBody := models.Trombosit{
		Data:  1.2,
		Notes: "Halo",
		Date:  "2024-02-22",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("PUT", "/api/user/health_status/trombosit/"+trombosit_id, bytes.NewBuffer(reqJSON))
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

	if expectedBody.Message != "Success Update Data" {
		t.Errorf("expected message body %s but got %s", "Success Update Data", expectedBody.Message)
	}
}

func TestUpdateTrombosit_Failed(t *testing.T) {
	router := gin.Default()

	router.PUT("/api/user/health_status/trombosit/:trombosit_id", middleware.RequireAuth, controllers.UpdateTrombosit)

	trombosit_id := "077a0e36-9ec7-4050-ae40-0c0af7245e7e"

	reqBody := models.Trombosit{
		Data:  7.2,
		Notes: "",
		Date:  "2024-02-22",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("PUT", "/api/user/health_status/trombosit/"+trombosit_id, bytes.NewBuffer(reqJSON))
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

	if expectedBody.ErrorMessage != "Field 'Notes' is required." {
		t.Errorf("expected message body %s but got %s", "Field 'Notes' is required.", expectedBody.ErrorMessage)
	}
}

func TestDeleteTrombosit_Success(t *testing.T) {

	router := gin.Default()

	router.DELETE("/api/user/health_status/trombosit/:trombosit_id", middleware.RequireAuth, controllers.DeleteTrombosit)

	trombosit_id := "077a0e36-9ec7-4050-ae40-0c0af7245e7e"

	req, err := http.NewRequest("DELETE", "/api/user/health_status/trombosit/"+trombosit_id, nil)

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

	if expectedBody.Message != "Success Delete Data" {
		t.Errorf("expected message body %s but got %s", "Success Delete Data", expectedBody.Message)
	}
}

func TestDeleteTrombosit_Failed(t *testing.T) {

	router := gin.Default()

	router.DELETE("/api/user/health_status/trombosit/:trombosit_id", middleware.RequireAuth, controllers.DeleteTrombosit)

	trombosit_id := "92e66f76-332e-45a8-a9d3-a919f01abde6"

	req, err := http.NewRequest("DELETE", "/api/user/health_status/trombosit/"+trombosit_id, nil)

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

	if expectedBody.ErrorMessage != "Record Not Found" {
		t.Errorf("expected message body %s but got %s", "Record Not Found", expectedBody.ErrorMessage)
	}
}
