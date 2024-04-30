package controllers_test

import (
	"elgeka-mobile/controllers"
	"elgeka-mobile/middleware"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

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

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("expected status code %d but got %d", http.StatusUnauthorized, rec.Code)
	}
}

func TestDoctorPatientAccept_Success(t *testing.T) {
	router := gin.Default()
	acceptance_id := "c6fa5126-ba8a-4a14-b406-b80c06486767"

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
	acceptance_id := "c6fa5126-ba8a-4a14-b406-b80c06486767"

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

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("expected status code %d but got %d", http.StatusUnauthorized, rec.Code)
	}
}

func TestDoctorPatientProfile_Success(t *testing.T) {
	router := gin.Default()
	acceptance_id := "3c297684-c68b-4828-a3d6-a734c246d685"

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
