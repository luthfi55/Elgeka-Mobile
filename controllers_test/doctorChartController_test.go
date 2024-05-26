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

func TestGenderChart_Success(t *testing.T) {
	router := gin.Default()

	router.GET("/api/doctor/patient/data/gender", middleware.RequireAuth, controllers.DataByGender)

	req, err := http.NewRequest("GET", "/api/doctor/patient/data/gender", nil)
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

	if expectedBody.Message != "Succes to Get Patient Data by Gender" {
		t.Errorf("expected message body %s but got %s", "Succes to Get Patient Data by Gender", expectedBody.Message)
	}
}

func TestGenderChart_Failed(t *testing.T) {
	router := gin.Default()

	router.GET("/api/doctor/patient/data/gender", middleware.RequireAuth, controllers.DataByGender)

	req, err := http.NewRequest("GET", "/api/doctor/patient/data/gender", nil)
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

func TestAgeChart_Success(t *testing.T) {
	router := gin.Default()

	router.GET("/api/doctor/patient/data/age", middleware.RequireAuth, controllers.DataByAge)

	req, err := http.NewRequest("GET", "/api/doctor/patient/data/age", nil)
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

	if expectedBody.Message != "Succes to Get Patient Data by Age" {
		t.Errorf("expected message body %s but got %s", "Succes to Get Patient Data by Age", expectedBody.Message)
	}
}

func TestAgeChart_Failed(t *testing.T) {
	router := gin.Default()

	router.GET("/api/doctor/patient/data/age", middleware.RequireAuth, controllers.DataByAge)

	req, err := http.NewRequest("GET", "/api/doctor/patient/data/age", nil)
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

func TestDiagnosisDateChart_Success(t *testing.T) {
	router := gin.Default()

	router.GET("/api/doctor/patient/data/diagnosis_date", middleware.RequireAuth, controllers.DataByDiagnosisDate)

	req, err := http.NewRequest("GET", "/api/doctor/patient/data/diagnosis_date", nil)
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

	if expectedBody.Message != "Succes to Get Patient Data by Diagnonsis Date" {
		t.Errorf("expected message body %s but got %s", "Succes to Get Patient Data by Diagnonsis Date", expectedBody.Message)
	}
}

func TestDiagnosisDateChart_Failed(t *testing.T) {
	router := gin.Default()

	router.GET("/api/doctor/patient/data/diagnosis_date", middleware.RequireAuth, controllers.DataByDiagnosisDate)

	req, err := http.NewRequest("GET", "/api/doctor/patient/data/diagnosis_date", nil)
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

func TestTreatmentChart_Success(t *testing.T) {
	router := gin.Default()

	router.GET("/api/doctor/patient/data/treatment", middleware.RequireAuth, controllers.DataByTreatment)

	req, err := http.NewRequest("GET", "/api/doctor/patient/data/treatment", nil)
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

	if expectedBody.Message != "Success to Get Patient Data by Treatment" {
		t.Errorf("expected message body %s but got %s", "Success to Get Patient Data by Treatment", expectedBody.Message)
	}
}

func TestTreatmentChart_Failed(t *testing.T) {
	router := gin.Default()

	router.GET("/api/doctor/patient/data/treatment", middleware.RequireAuth, controllers.DataByTreatment)

	req, err := http.NewRequest("GET", "/api/doctor/patient/data/treatment", nil)
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

func TestGetSymptompUserData_Success(t *testing.T) {
	router := gin.Default()

	router.GET("/api/doctor/patient/data/symptom/:type/:user_id", middleware.RequireAuth, controllers.GetSymptomUserData)

	data_type := "Oral"
	user_id := "14c9a443-03e5-4ba6-9664-0f820b19d8ab"
	req, err := http.NewRequest("GET", "/api/doctor/patient/data/symptom/"+data_type+"/"+user_id, nil)
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

func TestGetSymptompUserData_Failed(t *testing.T) {
	router := gin.Default()

	router.GET("/api/doctor/patient/data/symptom/:type/:user_id", middleware.RequireAuth, controllers.GetSymptomUserData)

	data_type := "Oral"
	user_id := "14c9a443-03e5-4ba6-9664-0f820b19d8pl"
	req, err := http.NewRequest("GET", "/api/doctor/patient/data/symptom/"+data_type+"/"+user_id, nil)
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

	// Debugging: Print the response body
	// t.Logf("Response Body: %s", rec.Body.String())

	var expectedBody failedExpectedOtpResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatalf("failed to unmarshal response body: %v\nResponse body: %s", err, rec.Body.String())
	}

	if expectedBody.ErrorMessage != "Failed to Get Symptom User Data" {
		t.Errorf("expected message body %s but got %s", "Failed to Get Symptom User Data", expectedBody.ErrorMessage)
	}
}
