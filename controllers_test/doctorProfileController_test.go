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

func TestDoctorProfile_Success(t *testing.T) {
	router := gin.Default()

	router.GET("/api/doctor/profile", middleware.RequireAuth, controllers.DoctorProfile)

	req, err := http.NewRequest("GET", "/api/doctor/profile", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.AddCookie(DoctorCookieConfiguration())

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

	if expectedBody.Message != "Success to Get Doctor Profile Data" {
		t.Errorf("expected message body %s but got %s", "Success to Get Doctor Profile Data", expectedBody.Message)
	}
}

func TestDoctorProfile_Failed(t *testing.T) {
	router := gin.Default()

	router.GET("/api/doctor/profile", middleware.RequireAuth, controllers.DoctorProfile)

	req, err := http.NewRequest("GET", "/api/doctor/profile", nil)
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

func TestEditDoctorProfile_Success(t *testing.T) {
	router := gin.Default()

	router.PUT("/api/doctor/profile/edit", middleware.RequireAuth, controllers.EditDoctorProfile)

	reqBody := models.Doctor{
		Name:         "Ruli",
		PhoneNumber:  "6281392443131",
		Gender:       "male",
		Specialist:   "SP1",
		HospitalName: "RSUD Al Ihsan",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("PUT", "/api/doctor/profile/edit", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.AddCookie(DoctorCookieConfiguration())

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

	if expectedBody.Message != "Success to Update Doctor Profile Data" {
		t.Errorf("expected message body %s but got %s", "Success to Update Doctor Profile Data", expectedBody.Message)
	}
}

func TestEditDoctorProfile_Failed(t *testing.T) {
	router := gin.Default()

	router.PUT("/api/doctor/profile/edit", middleware.RequireAuth, controllers.EditDoctorProfile)

	req, err := http.NewRequest("PUT", "/api/doctor/profile/edit", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.AddCookie(DoctorCookieConfiguration())

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

	if expectedBody.ErrorMessage != "Failed to read body" {
		t.Errorf("expected message body %s but got %s", "Failed to read body", expectedBody.ErrorMessage)
	}
}

func TestEditDoctorPassword_Success(t *testing.T) {
	router := gin.Default()

	router.PUT("/api/doctor/profile/password/edit", middleware.RequireAuth, controllers.EditDoctorPassword)

	reqBody := []byte(`{
		"OldPassword":          "L12345+678",
		"Password":             "L12345.678",
		"PasswordConfirmation": "L12345.678"
	}`)

	req, err := http.NewRequest("PUT", "/api/doctor/profile/password/edit", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.AddCookie(DoctorCookieConfiguration())

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

	if expectedBody.Message != "Success Update Doctor Password" {
		t.Errorf("expected message body %s but got %s", "Success Update Doctor Password", expectedBody.Message)
	}
}

func TestEditDoctorPassword_Failed(t *testing.T) {
	router := gin.Default()

	router.PUT("/api/doctor/profile/password/edit", middleware.RequireAuth, controllers.EditDoctorPassword)

	req, err := http.NewRequest("PUT", "/api/doctor/profile/password/edit", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.AddCookie(DoctorCookieConfiguration())

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

	if expectedBody.ErrorMessage != "Failed to read body" {
		t.Errorf("expected message body %s but got %s", "Failed to read body", expectedBody.ErrorMessage)
	}
}
func TestDoctorPatientRequest_Success(t *testing.T) {
	router := gin.Default()

	router.GET("/api/doctor/patient_request", middleware.RequireAuth, controllers.DoctorPatientRequest)

	req, err := http.NewRequest("GET", "/api/doctor/patient_request", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.AddCookie(DoctorCookieConfiguration())

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

	if expectedBody.Message != "Success Get List Acceptance Patient" {
		t.Errorf("expected message body %s but got %s", "Success Get List Acceptance Patient", expectedBody.Message)
	}
}

func TestDoctorPatientRequest_Failed(t *testing.T) {
	router := gin.Default()

	router.GET("/api/doctor/patient_request", middleware.RequireAuth, controllers.DoctorPatientRequest)

	req, err := http.NewRequest("GET", "/api/doctor/patient_request", nil)
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

func TestDoctorPatientAccept_Success(t *testing.T) {
	router := gin.Default()
	acceptance_id := "8fa51ee8-29e3-4579-a6d9-a2f00e2dde6b"

	router.PUT("/api/doctor/patient_request/accept/:acceptance_id", middleware.RequireAuth, controllers.DoctorPatientAccept)

	req, err := http.NewRequest("PUT", "/api/doctor/patient_request/accept/"+acceptance_id, nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.AddCookie(DoctorCookieConfiguration())

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

	if expectedBody.Message != "Success to Accept Patient" {
		t.Errorf("expected message body %s but got %s", "Success to Accept Patient", expectedBody.Message)
	}
}

func TestDoctorPatientAccept_Failed(t *testing.T) {
	router := gin.Default()
	acceptance_id := "c6fa5126-ba8a-4a12-b406-b80c06486767"

	router.PUT("/api/doctor/patient_request/accept/:acceptance_id", middleware.RequireAuth, controllers.DoctorPatientAccept)

	req, err := http.NewRequest("PUT", "/api/doctor/patient_request/accept/"+acceptance_id, nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.AddCookie(DoctorCookieConfiguration())

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status code %d but got %d", http.StatusNotFound, rec.Code)
	}

	var expectedBody ExpectedFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Acceptance Data Not Found" {
		t.Errorf("expected message body %s but got %s", "Acceptance Data Not Found", expectedBody.ErrorMessage)
	}
}

func TestDoctorPatientReject_Success(t *testing.T) {
	router := gin.Default()
	acceptance_id := "8fa51ee8-29e3-4579-a6d9-a2f00e2dde6b"

	router.PUT("/api/doctor/patient_request/reject/:acceptance_id", middleware.RequireAuth, controllers.DoctorPatientReject)

	req, err := http.NewRequest("PUT", "/api/doctor/patient_request/reject/"+acceptance_id, nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.AddCookie(DoctorCookieConfiguration())

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

	if expectedBody.Message != "Success to Reject Patient" {
		t.Errorf("expected message body %s but got %s", "Success to Reject Patient", expectedBody.Message)
	}
}

func TestDoctorPatientReject_Failed(t *testing.T) {
	router := gin.Default()
	acceptance_id := "c6fa5126-ba8a-4a12-b406-b80c06486767"

	router.PUT("/api/doctor/patient_request/reject/:acceptance_id", middleware.RequireAuth, controllers.DoctorPatientReject)

	req, err := http.NewRequest("PUT", "/api/doctor/patient_request/reject/"+acceptance_id, nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.AddCookie(DoctorCookieConfiguration())

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status code %d but got %d", http.StatusNotFound, rec.Code)
	}

	var expectedBody ExpectedFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Acceptance Data Not Found" {
		t.Errorf("expected message body %s but got %s", "Acceptance Data Not Found", expectedBody.ErrorMessage)
	}
}

func TestDoctorPatientListSuccess(t *testing.T) {
	router := gin.Default()

	router.GET("/api/doctor/patient/list", middleware.RequireAuth, controllers.DoctorPatientList)

	req, err := http.NewRequest("GET", "/api/doctor/patient/list", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.AddCookie(DoctorCookieConfiguration())

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

	if expectedBody.Message != "Success to Get List Patient" {
		t.Errorf("expected message body %s but got %s", "Success to Get List Patient", expectedBody.Message)
	}
}

func TestDoctorPatientList_Failed(t *testing.T) {
	router := gin.Default()

	router.GET("/api/doctor/patient/list", middleware.RequireAuth, controllers.DoctorPatientList)

	req, err := http.NewRequest("GET", "/api/doctor/patient/list", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	expiresTime, _ := time.Parse(time.RFC1123, "Mon, 15 Apr 2024 17:00:20 GMT")
	req.AddCookie(&http.Cookie{
		Name:     "Authorization",
		Value:    "ekJhbGciOiKIUZI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTcwNTgxNjQsInN1YiI6IjVmYWUwYjVkLTAxZDAtNDhmYi05ZTIzLWNhMWU5YWNhMmYyMCJ9.j4kcnBLIyOPnXl5Ok1gZKhqsAysXs2MTsEmzN23sVi0",
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

func TestDoctorPatientProfile_Success(t *testing.T) {
	router := gin.Default()
	acceptance_id := "c6fa5121-ba8a-4ae4-b406-b80c06486767"

	router.GET("/api/doctor/patient/profile/:acceptance_id", middleware.RequireAuth, controllers.DoctorPatientProfile)

	req, err := http.NewRequest("GET", "/api/doctor/patient/profile/"+acceptance_id, nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.AddCookie(DoctorCookieConfiguration())

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

	if expectedBody.Message != "Success to Get Patient Data" {
		t.Errorf("expected message body %s but got %s", "Success to Get Patient Data", expectedBody.Message)
	}
}

func TestDoctorPatientProfile_Failed(t *testing.T) {
	router := gin.Default()
	acceptance_id := "3c291684-c68b-4848-a3d6-a7p4c246d685"

	router.GET("/api/doctor/patient/profile/:acceptance_id", middleware.RequireAuth, controllers.DoctorPatientProfile)

	req, err := http.NewRequest("GET", "/api/doctor/patient/profile/"+acceptance_id, nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.AddCookie(DoctorCookieConfiguration())

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status code %d but got %d", http.StatusNotFound, rec.Code)
	}

	var expectedBody ExpectedFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Patient Profile Not Found" {
		t.Errorf("expected message body %s but got %s", "Patient Profile Not Found", expectedBody.ErrorMessage)
	}
}

func TestDoctorPatientHealthStatus_Success(t *testing.T) {
	router := gin.Default()
	acceptance_id := "c6fa5121-ba8a-4ae4-b406-b80c06486767"

	router.GET("/api/doctor/patient/health_status/:acceptance_id", middleware.RequireAuth, controllers.DoctorPatientHealthStatus)

	req, err := http.NewRequest("GET", "/api/doctor/patient/health_status/"+acceptance_id, nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.AddCookie(DoctorCookieConfiguration())

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

	if expectedBody.Message != "Success to Get Health Status Data" {
		t.Errorf("expected message body %s but got %s", "Success to Get Health Status Data", expectedBody.Message)
	}
}

func TestDoctorPatientHealthStatus_Failed(t *testing.T) {
	router := gin.Default()
	acceptance_id := "3c297684-c68b-4828-a3d6-a734c246d681"

	router.GET("/api/doctor/patient/health_status/:acceptance_id", middleware.RequireAuth, controllers.DoctorPatientHealthStatus)

	req, err := http.NewRequest("GET", "/api/doctor/patient/health_status/"+acceptance_id, nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.AddCookie(DoctorCookieConfiguration())

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status code %d but got %d", http.StatusNotFound, rec.Code)
	}

	var expectedBody ExpectedFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Patient Profile Not Found" {
		t.Errorf("expected message body %s but got %s", "Patient Profile Not Found", expectedBody.ErrorMessage)
	}
}

func TestDoctorPatientMedicine_Success(t *testing.T) {
	router := gin.Default()
	acceptance_id := "c6fa5121-ba8a-4ae4-b406-b80c06486767"

	router.GET("/api/doctor/patient/medicine/:acceptance_id", middleware.RequireAuth, controllers.DoctorPatientMedicine)

	req, err := http.NewRequest("GET", "/api/doctor/patient/medicine/"+acceptance_id, nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.AddCookie(DoctorCookieConfiguration())

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

	if expectedBody.Message != "Success to Get Medicine Data" {
		t.Errorf("expected message body %s but got %s", "Success to Get Medicine Data", expectedBody.Message)
	}
}

func TestDoctorPatientMedicine_Failed(t *testing.T) {
	router := gin.Default()
	acceptance_id := "3c297684-c68b-4828-a3d6-a734c246d681"

	router.GET("/api/doctor/patient/medicine/:acceptance_id", middleware.RequireAuth, controllers.DoctorPatientMedicine)

	req, err := http.NewRequest("GET", "/api/doctor/patient/medicine/"+acceptance_id, nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.AddCookie(DoctorCookieConfiguration())

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status code %d but got %d", http.StatusNotFound, rec.Code)
	}

	var expectedBody ExpectedFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Patient Profile Not Found" {
		t.Errorf("expected message body %s but got %s", "Patient Profile Not Found", expectedBody.ErrorMessage)
	}
}

func TestDoctorPatientSymptom_Success(t *testing.T) {
	router := gin.Default()
	type_symptom := "Oral"
	user_id := "14c9a443-03e5-4ba6-9664-0f820b19d8ab"

	router.GET("/api/doctor/patient/data/symptom/:type/:user_id", middleware.RequireAuth, controllers.GetSymptomUserData)

	req, err := http.NewRequest("GET", "/api/doctor/patient/data/symptom/"+type_symptom+"/"+user_id, nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.AddCookie(DoctorCookieConfiguration())

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

	if expectedBody.Message != "Success to Get Oral Symptom User Data" {
		t.Errorf("expected message body %s but got %s", "Success to Get Oral Symptom User Data", expectedBody.Message)
	}
}

func TestDoctorPatientSymptom_Failed(t *testing.T) {
	router := gin.Default()
	type_symptom := "Oral"
	user_id := "14c9a443-03e5-4ba6-9664-0f820b19d8zb"

	router.GET("/api/doctor/patient/data/symptom/:type/:user_id", middleware.RequireAuth, controllers.GetSymptomUserData)

	req, err := http.NewRequest("GET", "/api/doctor/patient/data/symptom/"+type_symptom+"/"+user_id, nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.AddCookie(DoctorCookieConfiguration())

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("expected status code %d but got %d", http.StatusInternalServerError, rec.Code)
	}

	var expectedBody ExpectedFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Failed to Get Symptom User Data" {
		t.Errorf("expected message body %s but got %s", "Failed to Get Symptom User Data", expectedBody.ErrorMessage)
	}
}

func TestListDoctorWebsite_Success(t *testing.T) {
	router := gin.Default()

	router.GET("/api/doctor/list/website", controllers.ListDoctorWebsite)

	req, err := http.NewRequest("GET", "/api/doctor/list/website", nil)

	if err != nil {
		t.Fatal(err)
	}

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

	if expectedBody.Message != "Success to Get Doctor List" {
		t.Errorf("expected message body %s but got %s", "Success to Get Doctor List", expectedBody.Message)
	}
}

func TestListDoctorWebsite_Failed(t *testing.T) {
	router := gin.Default()

	router.GET("/api/doctor/list/website", middleware.RequireAuth, controllers.ListDoctorWebsite)

	req, err := http.NewRequest("GET", "/api/doctor/list/website", nil)

	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d but got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestListPatientDoctorWebsite_Success(t *testing.T) {
	router := gin.Default()

	router.GET("/api/doctor/list_patient/website", controllers.ListPatientDoctorWebsite)

	req, err := http.NewRequest("GET", "/api/doctor/list_patient/website", nil)

	if err != nil {
		t.Fatal(err)
	}

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

	if expectedBody.Message != "Success to Get Patient Doctor List" {
		t.Errorf("expected message body %s but got %s", "Success to Get Patient Doctor List", expectedBody.Message)
	}
}

func TestListPatientDoctorWebsite_Failed(t *testing.T) {
	router := gin.Default()

	router.GET("/api/doctor/list_patient/website", middleware.RequireAuth, controllers.ListPatientDoctorWebsite)

	req, err := http.NewRequest("GET", "/api/doctor/list_patient/website", nil)

	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d but got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestDeleteDoctorAccountWebsite_Success(t *testing.T) {
	router := gin.Default()

	router.POST("/api/doctor/delete/account/website/:doctor_id", controllers.DeactivateDoctorAccountWebsite)

	doctor_id := "e7d8712b-a456-4014-ad14-2b93c7bc1a6f"
	req, err := http.NewRequest("POST", "/api/doctor/delete/account/website/"+doctor_id, nil)

	if err != nil {
		t.Fatal(err)
	}

	req.AddCookie(WebsiteBearierTokenConfiguration())

	bearerToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImlkIjoiNGUyZDY4NDgtM2FlMy00NjdjLTk5NzQtZTVkMTdhYWJhMGU3IiwidXNlcm5hbWUiOiJwZW5ndXJ1c2VsZ2VrYSIsImZ1bGxfbmFtZSI6IlBlbmd1cnVzIFV0YW1hIEVMR0VLQSBKYXdhIEJhcmF0IiwiaXNfYWN0aXZlIjp0cnVlLCJzdXBlclVzZXIiOnRydWV9LCJpYXQiOjE3MjA5NDM5NzgsImV4cCI6MTcyMDk2NTU3OH0.F3eZSmWHPw6Udm_g3rXwJEYrBm_hvCWvHxJCSzyvVQI"
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

	if expectedBody.Message != "Success to Deactivate Doctor Account" {
		t.Errorf("expected message body %s but got %s", "Success to Deactivate Doctor Account", expectedBody.Message)
	}
}

func TestDeleteDoctorAccountWebsite_Failed(t *testing.T) {
	router := gin.Default()

	router.POST("/api/doctor/delete/account/website/:doctor_id", controllers.DeactivateDoctorAccountWebsite)

	doctor_id := "b79a2300-f9bd-4c57-bed5-f18c5fd47fzl"
	req, err := http.NewRequest("POST", "/api/doctor/delete/account/website/"+doctor_id, nil)

	if err != nil {
		t.Fatal(err)
	}

	req.AddCookie(WebsiteBearierTokenConfiguration())

	bearerToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImlkIjoiNGUyZDY4NDgtM2FlMy00NjdjLTk5NzQtZTVkMTdhYWJhMGU3IiwidXNlcm5hbWUiOiJwZW5ndXJ1c2VsZ2VrYSIsImZ1bGxfbmFtZSI6IlBlbmd1cnVzIFV0YW1hIEVMR0VLQSBKYXdhIEJhcmF0IiwiaXNfYWN0aXZlIjp0cnVlLCJzdXBlclVzZXIiOnRydWV9LCJpYXQiOjE3MjA5NDM5NzgsImV4cCI6MTcyMDk2NTU3OH0.F3eZSmWHPw6Udm_g3rXwJEYrBm_hvCWvHxJCSzyvVQI"
	req.Header.Set("Authorization", "Bearer "+bearerToken)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status code %d but got %d", http.StatusNotFound, rec.Code)
	}

	var expectedBody ExpectedFailedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Failed to Get Doctor Account" {
		t.Errorf("expected message body %s but got %s", "Failed to Get Doctor Account", expectedBody.ErrorMessage)
	}
}
