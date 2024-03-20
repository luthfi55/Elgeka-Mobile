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

type successExpectedForgotPasswordResponse struct {
	Message string `json:"Message"`
	Data    []struct {
		ID      string `json:"ID"`
		Email   string `json:"Email"`
		OtpCode string `json:"OtpCode"`
	} `json:"Data"`
	Link []struct {
		Name string `json:"Name"`
		Link string `json:"Link"`
	} `json:"Link"`
}

type failedExpectedForgotPasswordResponse struct {
	ErrorMessage string `json:"ErrorMessage"`
	Data         []struct {
		Email string `json:"Email"`
	} `json:"Data"`
	Link []struct {
		Name string `json:"Name"`
		Link string `json:"Link"`
	} `json:"Link"`
}

func TestLogin_Success(t *testing.T) {

	router := gin.Default()

	router.POST("/api/user/login", controllers.UserLogin)

	reqBody := []byte(`{
		"Email": "angga515151@gmail.com",
		"Password": "L12345678*"
	}`)
	req, err := http.NewRequest("POST", "/api/user/login", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	type ExpectedResponse struct {
		Message string `json:"Message"`
		Data    []struct {
			Name  string `json:"Name"`
			Email string `json:"Email"`
		} `json:"Data"`
		Link []struct {
			Name string `json:"Name"`
			Link string `json:"Link"`
		} `json:"Link"`
	}

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	var expectedBody ExpectedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Login Success" {
		t.Errorf("expected message %s but got %s", "Login Success", expectedBody.Message)
	}
}

func TestLogin_Failed(t *testing.T) {

	router := gin.Default()

	router.POST("/api/user/login", controllers.UserLogin)

	reqBody := []byte(`{
		"Email": "test@gmail.com",
		"Password": "tess"
	}`)
	req, err := http.NewRequest("POST", "/api/user/login", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	type ExpectedResponse struct {
		ErrorMessage string `json:"ErrorMessage"`
		Data         []struct {
			Email string `json:"Name"`
		} `json:"Data"`
		Link []struct {
			Name string `json:"Name"`
			Link string `json:"Link"`
		} `json:"Link"`
	}

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d but got %d", http.StatusBadRequest, rec.Code)
	}

	var expectedBody ExpectedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Invalid email or password" {
		t.Errorf("expected message %s but got %s", "Invalid email or password", expectedBody.ErrorMessage)
	}
}
func TestValidateController(t *testing.T) {

	r := gin.New()
	r.GET("/api/user/validate", controllers.Validate)

	req, err := http.NewRequest("GET", "/api/user/validate", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTEzMDQ3ODgsInN1YiI6IjVmYWUwYjVkLTAxZDAtNDhmYi05ZTIzLWNhMWU5YWNhMmYyMCJ9.SWasyrWm5fKGe2x7eh3yhebsdo7OT136ELbl13pBrkY"})

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	if message, ok := response["message"]; !ok {
		t.Error("Expected 'message' key in response but not found")
	} else {
		expectedMessage := "your_expected_message_here"
		if message != expectedMessage {
			t.Errorf("Expected message %q but got %q", expectedMessage, message)
		}
	}
}

func TestForgotPassword_Success(t *testing.T) {
	router := gin.Default()
	router.POST("/api/user/forgot_password", controllers.ForgotPassword)

	reqBody := []byte(`{
		"Email": "angga515151@gmail.com"	
	}`)

	req, err := http.NewRequest("POST", "/api/user/forgot_password", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	var expectedBody successExpectedForgotPasswordResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success to Send Otp Code" {
		t.Errorf("expected message %s but got %s", "Success to Send Otp Code", expectedBody.Message)
	}
}

func TestForgotPassword_Failed(t *testing.T) {
	router := gin.Default()
	router.POST("/api/user/forgot_password", controllers.ForgotPassword)

	reqBody := []byte(`{
		"Email": "ffsdfasf@gmail.com"
	}`)

	req, err := http.NewRequest("POST", "/api/user/forgot_password", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d but got %d", http.StatusInternalServerError, rec.Code)
	}

	var expectedBody failedExpectedForgotPasswordResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Email Not Found" {
		t.Errorf("expected message %s but got %s", "Email Not Found", expectedBody.ErrorMessage)
	}
}

func TestRefreshForgotPassword_Success(t *testing.T) {
	router := gin.Default()
	user_id := "89dff9eb-fe50-40a6-8775-27d7b5997326"
	router.POST("/api/user/refresh_code/forgot_password/:user_id", controllers.RefreshForgotPasswordOtp)

	req, err := http.NewRequest("POST", "/api/user/refresh_code/forgot_password/"+user_id, nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	var expectedBody successExpectedForgotPasswordResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success to Send Otp Code" {
		t.Errorf("expected message body %s but got %s", "Success to Send Otp Code", expectedBody.Message)
	}
}

func TestRefreshForgotPassword_Failed(t *testing.T) {
	router := gin.Default()
	user_id := "89dff9eb-fe50-40a6-8775-27d7b5997327"
	router.POST("/api/user/refresh_code/forgot_password/:user_id", controllers.RefreshForgotPasswordOtp)

	req, err := http.NewRequest("POST", "/api/user/refresh_code/forgot_password/"+user_id, nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status code %d but got %d", http.StatusNotFound, rec.Code)
	}

	var expectedBody failedExpectedForgotPasswordResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "User Not Found" {
		t.Errorf("expected message body %s but got %s", "User Not Found", expectedBody.ErrorMessage)
	}
}

func TestCheckOtp_Success(t *testing.T) {
	router := gin.Default()
	router.POST("/api/user/check_otp/:user_id", controllers.CheckOtp)

	user_id := "89dff9eb-fe50-40a6-8775-27d7b5997326"
	reqBody := []byte(`{
		"OtpCode": "3242"
	}`)
	req, err := http.NewRequest("POST", "/api/user/check_otp/"+user_id, bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	var expectedBody successExpectedForgotPasswordResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Check Otp Successfully" {
		t.Errorf("expected message body %s but got %s", "Check Otp Successfully", expectedBody.Message)
	}
}

func TestCheckOtp_Failed(t *testing.T) {
	router := gin.Default()
	router.POST("/api/user/check_otp/:user_id", controllers.CheckOtp)

	user_id := "89dff9eb-fe50-40a6-8775-27d7b5997326"
	reqBody := []byte(`{
		"OtpCode": "3241"
	}`)
	req, err := http.NewRequest("POST", "/api/user/check_otp/"+user_id, bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("expected status code %d but got %d", http.StatusUnauthorized, rec.Code)
	}

	var expectedBody failedExpectedForgotPasswordResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Incorrect OTP code" {
		t.Errorf("expected message body %s but got %s", "Incorrect OTP code", expectedBody.ErrorMessage)
	}
}

func TestChangePassword_Success(t *testing.T) {
	router := gin.Default()
	router.POST("/api/user/change_password/:user_id/:otp_code", controllers.ChangePassword)

	user_id := "89dff9eb-fe50-40a6-8775-27d7b5997326"
	otp_code := "JDJhJDEwJG5lWnI4bnBiZ2JjaE1IRXNDUUc5UnVrL0J3MDBrczU5enphNDRKaFU2WE9odGI0bGthZ1ND"
	reqBody := []byte(`{
		"Password": "AL12345678*",
		"PasswordConfirmation": "AL12345678*"
	}`)
	req, err := http.NewRequest("POST", "/api/user/change_password/"+user_id+"/"+otp_code, bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	var expectedBody successExpectedForgotPasswordResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Update Password Successfully" {
		t.Errorf("expected message body %s but got %s", "Update Password Successfully", expectedBody.Message)
	}
}

func TestChangePassword_Failed(t *testing.T) {
	router := gin.Default()
	router.POST("/api/user/change_password/:user_id/:otp_code", controllers.ChangePassword)

	user_id := "89dff9eb-fe50-40a6-8775-27d7b5997326"
	otp_code := "JDJhJDEwJG5lWnI4bnBiZ2JjaE1IRXNDUUc5UnVrL0J3MDBrczU5enphNDRKaFU2WE9odGI0bGthZ1Nz"
	reqBody := []byte(`{
		"Password": "AL12345678*",
		"PasswordConfirmation": "AL12345678*"
	}`)
	req, err := http.NewRequest("POST", "/api/user/change_password/"+user_id+"/"+otp_code, bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d but got %d", http.StatusBadRequest, rec.Code)
	}

	var expectedBody failedExpectedForgotPasswordResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Incorect Otp Code" {
		t.Errorf("expected message body %s but got %s", "Incorect Otp Code", expectedBody.ErrorMessage)
	}
}
