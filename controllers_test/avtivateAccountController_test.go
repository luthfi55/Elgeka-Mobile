package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"os/signal"
	"syscall"
	"testing"

	"elgeka-mobile/controllers"
	"elgeka-mobile/initializers"

	"github.com/gin-gonic/gin"
)

type successExpectedOtpResponse struct {
	Message string `json:"Message"`
	OtpData []struct {
		Email string `json:"Email"`
	} `json:"OtpData"`
	Link []struct {
		Name string `json:"Name"`
		Link string `json:"Link"`
	} `json:"Link"`
}
type failedExpectedOtpResponse struct {
	ErrorMessage string `json:"ErrorMessage"`
	OtpData      []struct {
		Email string `json:"Email"`
	} `json:"OtpData"`
	Link []struct {
		Name string `json:"Name"`
		Link string `json:"Link"`
	} `json:"Link"`
}

func TestSendEmailOtp_Success(t *testing.T) {

	router := gin.Default()

	router.POST("/api/user/email_otp/:user_id", controllers.SendEmailOtp)

	user_id := "8d6f5d0e-ba9a-4c66-8633-62c861fb9c0e"

	req, err := http.NewRequest("POST", "/api/user/email_otp/"+user_id, bytes.NewBuffer(nil))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	var expectedBody successExpectedOtpResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Send Email OTP Successfully" {
		t.Errorf("expected message body %s but got %s", "Send Email OTP Successfully", expectedBody.Message)
	}

}

func TestSendEmailOtp_Failed(t *testing.T) {

	router := gin.Default()

	router.POST("/api/user/email_otp/:user_id", controllers.SendEmailOtp)

	user_id := "89dff9eb-fe50-40a6-8775-27d7b5997325"

	req, err := http.NewRequest("POST", "/api/user/email_otp/"+user_id, bytes.NewBuffer(nil))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status code %d but got %d", http.StatusNotFound, rec.Code)
	}

	var expectedBody failedExpectedOtpResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "User Not Found" {
		t.Errorf("expected message body %s but got %s", "User Not Found", expectedBody.ErrorMessage)
	}

}

func TestSendWhatsappOtp_Success(t *testing.T) {
	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, os.Interrupt, syscall.SIGTERM)

	go func() {
		initializers.ConnectToWhatsapp()
	}()

	router := gin.Default()

	router.POST("/api/user/whatsapp_otp/:user_id", controllers.SendWhatsappOtp)

	go func() {
		<-shutdownSignal

		initializers.DisconnectWhatsapp()

		os.Exit(0)
	}()

	user_id := "8d6f5d0e-ba9a-4c66-8633-62c861fb9c0e"

	req, err := http.NewRequest("POST", "/api/user/whatsapp_otp/"+user_id, bytes.NewBuffer(nil))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	var expectedBody successExpectedOtpResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Send Whatsapp OTP Successfully" {
		t.Errorf("expected message body %s but got %s", "Send Whatsapp OTP Successfully", expectedBody.Message)
	}

}

func TestSendWhatsappOtp_Failed(t *testing.T) {
	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, os.Interrupt, syscall.SIGTERM)

	go func() {
		initializers.ConnectToWhatsapp()
	}()

	router := gin.Default()

	router.POST("/api/user/whatsapp_otp/:user_id", controllers.SendWhatsappOtp)

	go func() {
		<-shutdownSignal

		initializers.DisconnectWhatsapp()

		os.Exit(0)
	}()

	user_id := "8d2baaf9-ba97-4237-a9ab-1095fe2f4745"

	req, err := http.NewRequest("POST", "/api/user/whatsapp_otp/"+user_id, bytes.NewBuffer(nil))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status code %d but got %d", http.StatusNotFound, rec.Code)
	}

	var expectedBody failedExpectedOtpResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "User Not Found" {
		t.Errorf("expected message body %s but got %s", "User Not Found", expectedBody.ErrorMessage)
	}

}

func TestEmailRefreshOtpCode_Success(t *testing.T) {

	router := gin.Default()

	router.POST("/api/user/email_refresh_code/:user_id", controllers.RefreshOtpCode)

	user_id := "8d6f5d0e-ba9a-4c66-8633-62c861fb9c0e"

	req, err := http.NewRequest("POST", "/api/user/email_refresh_code/"+user_id, bytes.NewBuffer(nil))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	var expectedBody successExpectedOtpResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Refresh OTP Successfully" {
		t.Errorf("expected message body %s but got %s", "Refresh OTP Successfully", expectedBody.Message)
	}
}

func TestEmailRefreshOtpCode_Failed(t *testing.T) {

	router := gin.Default()

	router.POST("/api/user/email_refresh_code/:user_id", controllers.RefreshOtpCode)

	user_id := "89dff9eb-fe50-40a6-8775-27d7b5997325"

	req, err := http.NewRequest("POST", "/api/user/email_refresh_code/"+user_id, bytes.NewBuffer(nil))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status code %d but got %d", http.StatusNotFound, rec.Code)
	}

	var expectedBody failedExpectedOtpResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "User Not Found" {
		t.Errorf("expected message body %s but got %s", "User Not Found", expectedBody.ErrorMessage)
	}
}

func TestWhatsappRefreshOtpCode_Success(t *testing.T) {
	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, os.Interrupt, syscall.SIGTERM)

	go func() {
		initializers.ConnectToWhatsapp()
	}()

	router := gin.Default()

	router.POST("/api/user/whatsapp_refresh_code/:user_id", controllers.RefreshWhatsappOtpCode)

	go func() {
		<-shutdownSignal

		initializers.DisconnectWhatsapp()

		os.Exit(0)
	}()

	user_id := "8d6f5d0e-ba9a-4c66-8633-62c861fb9c0e"

	req, err := http.NewRequest("POST", "/api/user/whatsapp_refresh_code/"+user_id, bytes.NewBuffer(nil))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	var expectedBody successExpectedOtpResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Refresh Whatsapp OTP Successfully" {
		t.Errorf("expected message body %s but got %s", "Refresh Whatsapp OTP Successfully", expectedBody.Message)
	}

}

func TestWhatsappRefreshOtpCode_Failed(t *testing.T) {
	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, os.Interrupt, syscall.SIGTERM)

	go func() {
		initializers.ConnectToWhatsapp()
	}()

	router := gin.Default()

	router.POST("/api/user/whatsapp_refresh_code/:user_id", controllers.RefreshWhatsappOtpCode)

	go func() {
		<-shutdownSignal

		initializers.DisconnectWhatsapp()

		os.Exit(0)
	}()

	user_id := "89dff9eb-fe50-40a6-8775-27d7b5997325"

	req, err := http.NewRequest("POST", "/api/user/whatsapp_refresh_code/"+user_id, bytes.NewBuffer(nil))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status code %d but got %d", http.StatusNotFound, rec.Code)
	}

	var expectedBody failedExpectedOtpResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "User Not Found" {
		t.Errorf("expected message body %s but got %s", "User Not Found", expectedBody.ErrorMessage)
	}

}

func TestActivateUser_Success(t *testing.T) {
	router := gin.Default()

	router.POST("/api/user/activate/:user_id", controllers.Activate)

	reqBody := []byte(`{
			"OtpCode": "6344"
		}`)

	userID := "8d6f5d0e-ba9a-4c66-8633-62c861fb9c0e"

	req, err := http.NewRequest("POST", "/api/user/activate/"+userID, bytes.NewBuffer(reqBody))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	var expectedBody successExpectedOtpResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "User Activated Successfully" {
		t.Errorf("expected message body %s but got %s", "User Activated Successfully", expectedBody.Message)
	}
}

func TestActivateUser_Failed(t *testing.T) {
	router := gin.Default()

	router.POST("/api/user/activate/:user_id", controllers.Activate)

	reqBody := []byte(`{
			"OtpCode": "3380"
		}`)

	userID := "8d6f5d0e-ba9a-4c66-8633-62c861fb9c0e"

	req, err := http.NewRequest("POST", "/api/user/activate/"+userID, bytes.NewBuffer(reqBody))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("expected status code %d but got %d", http.StatusUnauthorized, rec.Code)
	}

	var expectedBody failedExpectedOtpResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Incorrect OTP code" {
		t.Errorf("expected message body %s but got %s", "Incorrect OTP code", expectedBody.ErrorMessage)
	}
}

func TestSendWhatsappOtpDoctor_Success(t *testing.T) {
	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, os.Interrupt, syscall.SIGTERM)

	go func() {
		initializers.ConnectToWhatsapp()
	}()

	router := gin.Default()

	router.POST("/api/doctor/whatsapp_otp/:doctor_id", controllers.SendDoctorWhatsappOtp)

	go func() {
		<-shutdownSignal

		initializers.DisconnectWhatsapp()

		os.Exit(0)
	}()

	doctor_id := "e7d8712b-a456-4014-ad14-2b93c7bc1a6f"

	req, err := http.NewRequest("POST", "/api/doctor/whatsapp_otp/"+doctor_id, bytes.NewBuffer(nil))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	var expectedBody successExpectedOtpResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Send Whatsapp OTP Successfully" {
		t.Errorf("expected message body %s but got %s", "Send Whatsapp OTP Successfully", expectedBody.Message)
	}

}

func TestSendWhatsappOtpDoctor_Failed(t *testing.T) {
	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, os.Interrupt, syscall.SIGTERM)

	go func() {
		initializers.ConnectToWhatsapp()
	}()

	router := gin.Default()

	router.POST("/api/doctor/whatsapp_otp/:doctor_id", controllers.SendDoctorWhatsappOtp)

	go func() {
		<-shutdownSignal

		initializers.DisconnectWhatsapp()

		os.Exit(0)
	}()

	doctor_id := "8d2baaf9-ba97-4237-a9ab-1095fe2f4745"

	req, err := http.NewRequest("POST", "/api/doctor/whatsapp_otp/"+doctor_id, bytes.NewBuffer(nil))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status code %d but got %d", http.StatusNotFound, rec.Code)
	}

	var expectedBody failedExpectedOtpResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Doctor Not Found" {
		t.Errorf("expected message body %s but got %s", "Doctor Not Found", expectedBody.ErrorMessage)
	}

}

func TestSendEmailOtpDoctor_Success(t *testing.T) {

	router := gin.Default()

	router.POST("/api/doctor/email_otp/:doctor_id", controllers.SendDoctorEmailOtp)

	doctor_id := "e7d8712b-a456-4014-ad14-2b93c7bc1a6f"

	req, err := http.NewRequest("POST", "/api/doctor/email_otp/"+doctor_id, bytes.NewBuffer(nil))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	var expectedBody successExpectedOtpResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Send Email OTP Successfully" {
		t.Errorf("expected message body %s but got %s", "Send Email OTP Successfully", expectedBody.Message)
	}

}

func TestSendEmailOtpDoctor_Failed(t *testing.T) {

	router := gin.Default()

	router.POST("/api/doctor/email_otp/:doctor_id", controllers.SendDoctorEmailOtp)

	doctor_id := "89dff9eb-fe50-40a6-8775-27d7b5997325"

	req, err := http.NewRequest("POST", "/api/doctor/email_otp/"+doctor_id, bytes.NewBuffer(nil))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status code %d but got %d", http.StatusNotFound, rec.Code)
	}

	var expectedBody failedExpectedOtpResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Doctor Not Found" {
		t.Errorf("expected message body %s but got %s", "Doctor Not Found", expectedBody.ErrorMessage)
	}

}

func TestWhatsappRefreshOtpCodeDoctor_Success(t *testing.T) {
	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, os.Interrupt, syscall.SIGTERM)

	go func() {
		initializers.ConnectToWhatsapp()
	}()

	router := gin.Default()

	router.POST("/api/doctor/whatsapp_refresh_code/:doctor_id", controllers.RefreshDoctorWhatsappOtpCode)

	go func() {
		<-shutdownSignal

		initializers.DisconnectWhatsapp()

		os.Exit(0)
	}()

	doctor_id := "e7d8712b-a456-4014-ad14-2b93c7bc1a6f"

	req, err := http.NewRequest("POST", "/api/doctor/whatsapp_refresh_code/"+doctor_id, bytes.NewBuffer(nil))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	var expectedBody successExpectedOtpResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Refresh Whatsapp OTP Successfully" {
		t.Errorf("expected message body %s but got %s", "Refresh Whatsapp OTP Successfully", expectedBody.Message)
	}

}

func TestWhatsappRefreshOtpCodeDoctor_Failed(t *testing.T) {
	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, os.Interrupt, syscall.SIGTERM)

	go func() {
		initializers.ConnectToWhatsapp()
	}()

	router := gin.Default()

	router.POST("/api/doctor/whatsapp_refresh_code/:doctor_id", controllers.RefreshDoctorWhatsappOtpCode)

	go func() {
		<-shutdownSignal

		initializers.DisconnectWhatsapp()

		os.Exit(0)
	}()

	doctor_id := "89dff9eb-fe50-40a6-8775-27d7b5997325"

	req, err := http.NewRequest("POST", "/api/doctor/whatsapp_refresh_code/"+doctor_id, bytes.NewBuffer(nil))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status code %d but got %d", http.StatusNotFound, rec.Code)
	}

	var expectedBody failedExpectedOtpResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Doctor Not Found" {
		t.Errorf("expected message body %s but got %s", "Doctor Not Found", expectedBody.ErrorMessage)
	}

}

func TestListInactiveDoctor_Success(t *testing.T) {
	router := gin.Default()

	router.GET("/api/doctor/list_inactive", controllers.ListInactiveDoctor)

	req, err := http.NewRequest("GET", "/api/doctor/list_inactive", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.AddCookie(WebsiteBearierTokenConfiguration())
	bearerToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImlkIjoiNGUyZDY4NDgtM2FlMy00NjdjLTk5NzQtZTVkMTdhYWJhMGU3IiwidXNlcm5hbWUiOiJwZW5ndXJ1c2VsZ2VrYSIsImZ1bGxfbmFtZSI6IlBlbmd1cnVzIFV0YW1hIEVMR0VLQSBKQUJBUiIsImlzX2FjdGl2ZSI6dHJ1ZSwic3VwZXJVc2VyIjp0cnVlfSwiaWF0IjoxNzE4MDg1MjY0LCJleHAiOjE3MTgxMDY4NjR9.bC2sYI3q5mBYUuqpd87TSC0t3nnVp8AuVuvnrssFosc"
	req.Header.Set("Authorization", "Bearer "+bearerToken)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	type successExpectedResponse struct {
		Message string `json:"Message"`
		Data    []struct {
			ID   string `json:"ID"`
			Name string `json:"Name"`
		} `json:"Data"`
		Link []struct {
			Name string `json:"Name"`
			Link string `json:"Link"`
		} `json:"Link"`
	}

	var expectedBody successExpectedResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Get Data Successfully" {
		t.Errorf("expected message body %s but got %s", "Get Data Successfully", expectedBody.Message)
	}
}

func TestListInactiveDoctor_Failed(t *testing.T) {
	router := gin.Default()

	router.GET("/api/doctor/list_inactive", controllers.ListInactiveDoctor)

	req, err := http.NewRequest("GET", "/api/doctor/list_inactive", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d but got %d", http.StatusBadRequest, rec.Code)
	}

	var expectedBody failedExpectedOtpResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Token is required" {
		t.Errorf("expected message body %s but got %s", "Token is required", expectedBody.ErrorMessage)
	}
}

func TestEmailRefreshCodeDoctor_Success(t *testing.T) {
	router := gin.Default()
	router.POST("/api/doctor/email_refresh_code/:doctor_id", controllers.RefreshDoctorEmailOtpCode)

	DoctorID := "9b33a7c3-1d73-498d-ba96-53733ff7cbf4"
	req, err := http.NewRequest("POST", "/api/doctor/email_refresh_code/"+DoctorID, nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	var expectedBody successExpectedOtpResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Refresh Email OTP Successfully" {
		t.Errorf("expected message body %s but got %s", "Refresh Email OTP Successfully", expectedBody.Message)
	}
}

func TestEmailRefreshCodeDoctor_Failed(t *testing.T) {
	router := gin.Default()
	router.POST("/api/doctor/email_refresh_code/:doctor_id", controllers.RefreshDoctorEmailOtpCode)

	DoctorID := "08e28077-3927-4c70-9d30-2c8ce71b40c1"
	req, err := http.NewRequest("POST", "/api/doctor/email_refresh_code/"+DoctorID, nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status code %d but got %d", http.StatusNotFound, rec.Code)
	}

	var expectedBody failedExpectedOtpResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Doctor Not Found" {
		t.Errorf("expected message body %s but got %s", "Doctor Not Found", expectedBody.ErrorMessage)
	}
}

func TestActivateEmailDoctor_Success(t *testing.T) {

	router := gin.Default()

	router.POST("/api/doctor/activate_otp/:doctor_id", controllers.ActivateOtpDoctor)

	reqBody := []byte(`{
		"OtpCode": "0923"
	}`)

	DoctorID := "9b33a7c3-1d73-498d-ba96-53733ff7cbf4"
	req, err := http.NewRequest("POST", "/api/doctor/activate_otp/"+DoctorID, bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	var expectedBody successExpectedOtpResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Doctor Otp Activated Successfully" {
		t.Errorf("expected message body %s but got %s", "Doctor Otp Activated Successfully", expectedBody.Message)
	}

}

func TestActivateEmailDoctor_Failed(t *testing.T) {

	router := gin.Default()

	router.POST("/api/doctor/activate_otp/:doctor_id", controllers.ActivateOtpDoctor)

	reqBody := []byte(`{
			"OtpCode": "8136"
	}`)

	DoctorID := "9b33a7c3-1d73-498d-ba96-53733ff7cbf4"
	req, err := http.NewRequest("POST", "/api/doctor/activate_otp/"+DoctorID, bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("expected status code %d but got %d", http.StatusUnauthorized, rec.Code)
	}

	var expectedBody failedExpectedOtpResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Incorrect OTP code" {
		t.Errorf("expected message body %s but got %s", "Incorrect OTP code", expectedBody.ErrorMessage)
	}

}

func TestActivateDoctorAccount_Success(t *testing.T) {
	router := gin.Default()

	router.POST("/api/doctor/activate_account/:doctor_id", controllers.ActivateDoctor)

	DoctorID := "e6438b13-4ff2-4f25-9d97-5ba9379002c7"
	req, err := http.NewRequest("POST", "/api/doctor/activate_account/"+DoctorID, nil)

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.AddCookie(WebsiteBearierTokenConfiguration())
	bearerToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImlkIjoiNGUyZDY4NDgtM2FlMy00NjdjLTk5NzQtZTVkMTdhYWJhMGU3IiwidXNlcm5hbWUiOiJwZW5ndXJ1c2VsZ2VrYSIsImZ1bGxfbmFtZSI6IlBlbmd1cnVzIFV0YW1hIEVMR0VLQSBKQUJBUiIsImlzX2FjdGl2ZSI6dHJ1ZSwic3VwZXJVc2VyIjp0cnVlfSwiaWF0IjoxNzE4MDg1MjY0LCJleHAiOjE3MTgxMDY4NjR9.bC2sYI3q5mBYUuqpd87TSC0t3nnVp8AuVuvnrssFosc"
	req.Header.Set("Authorization", "Bearer "+bearerToken)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	var expectedBody successExpectedOtpResponse

	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)

	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Doctor Activated Successfully" {
		t.Errorf("expected message body %s but got %s", "Doctor Activated Successfully", expectedBody.Message)
	}

}

func TestActivateDoctorAccount_Failed(t *testing.T) {
	router := gin.Default()

	router.POST("/api/doctor/activate_account/:doctor_id", controllers.ActivateDoctor)

	DoctorID := "08e28077-3927-4c70-9d30-2c8ce71b40c1"
	req, err := http.NewRequest("POST", "/api/doctor/activate_account/"+DoctorID, nil)

	if err != nil {
		t.Fatal(err)
	}

	req.AddCookie(WebsiteBearierTokenConfiguration())
	bearerToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImlkIjoiNGUyZDY4NDgtM2FlMy00NjdjLTk5NzQtZTVkMTdhYWJhMGU3IiwidXNlcm5hbWUiOiJwZW5ndXJ1c2VsZ2VrYSIsImZ1bGxfbmFtZSI6IlBlbmd1cnVzIFV0YW1hIEVMR0VLQSBKQUJBUiIsImlzX2FjdGl2ZSI6dHJ1ZSwic3VwZXJVc2VyIjp0cnVlfSwiaWF0IjoxNzE4MDg1MjY0LCJleHAiOjE3MTgxMDY4NjR9.bC2sYI3q5mBYUuqpd87TSC0t3nnVp8AuVuvnrssFosc"
	req.Header.Set("Authorization", "Bearer "+bearerToken)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status code %d but got %d", http.StatusNotFound, rec.Code)
	}

	var expectedBody failedExpectedOtpResponse

	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)

	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Doctor Not Found" {
		t.Errorf("expected message body %s but got %s", "Doctor Not Found", expectedBody.ErrorMessage)
	}

}

func TestRejectDoctorAccount_Success(t *testing.T) {
	router := gin.Default()

	router.POST("/api/doctor/reject_activation/:doctor_id", controllers.RejectDoctor)

	DoctorID := "99a49bf8-11e1-4ffc-a374-3fb943e9637a"
	req, err := http.NewRequest("POST", "/api/doctor/reject_activation/"+DoctorID, nil)

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.AddCookie(WebsiteBearierTokenConfiguration())
	bearerToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImlkIjoiNGUyZDY4NDgtM2FlMy00NjdjLTk5NzQtZTVkMTdhYWJhMGU3IiwidXNlcm5hbWUiOiJwZW5ndXJ1c2VsZ2VrYSIsImZ1bGxfbmFtZSI6IlBlbmd1cnVzIFV0YW1hIEVMR0VLQSBKQUJBUiIsImlzX2FjdGl2ZSI6dHJ1ZSwic3VwZXJVc2VyIjp0cnVlfSwiaWF0IjoxNzE4MDg1MjY0LCJleHAiOjE3MTgxMDY4NjR9.bC2sYI3q5mBYUuqpd87TSC0t3nnVp8AuVuvnrssFosc"
	req.Header.Set("Authorization", "Bearer "+bearerToken)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	var expectedBody successExpectedOtpResponse

	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)

	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Reject Doctor Successfully" {
		t.Errorf("expected message body %s but got %s", "Reject Doctor Successfully", expectedBody.Message)
	}
}

func TestRejectDoctorAccount_Failed(t *testing.T) {
	router := gin.Default()

	router.POST("/api/doctor/reject_activation/:doctor_id", controllers.RejectDoctor)

	DoctorID := "08e28077-3927-4c70-9d30-2c8ce71b40c1"
	req, err := http.NewRequest("POST", "/api/doctor/reject_activation/"+DoctorID, nil)

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.AddCookie(WebsiteBearierTokenConfiguration())
	bearerToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImlkIjoiNGUyZDY4NDgtM2FlMy00NjdjLTk5NzQtZTVkMTdhYWJhMGU3IiwidXNlcm5hbWUiOiJwZW5ndXJ1c2VsZ2VrYSIsImZ1bGxfbmFtZSI6IlBlbmd1cnVzIFV0YW1hIEVMR0VLQSBKQUJBUiIsImlzX2FjdGl2ZSI6dHJ1ZSwic3VwZXJVc2VyIjp0cnVlfSwiaWF0IjoxNzE4MDg1MjY0LCJleHAiOjE3MTgxMDY4NjR9.bC2sYI3q5mBYUuqpd87TSC0t3nnVp8AuVuvnrssFosc"
	req.Header.Set("Authorization", "Bearer "+bearerToken)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status code %d but got %d", http.StatusNotFound, rec.Code)
	}

	var expectedBody failedExpectedOtpResponse

	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)

	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Doctor Not Found" {
		t.Errorf("expected message body %s but got %s", "Doctor Not Found", expectedBody.ErrorMessage)
	}
}
