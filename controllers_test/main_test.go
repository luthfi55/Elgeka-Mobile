package controllers_test

import (
	"elgeka-mobile/initializers"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	// Setup
	envFile := "../.env" // Sesuaikan dengan path ke file .env Anda
	if err := initializers.LoadEnvVariables(envFile); err != nil {
		panic(err)
	}

	// Setup
	initializers.ConnectToDb()
	initializers.SyncDatabase()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
	// initializers.ConnectToWhatsapp()

	// go func() {
	// 	initializers.ConnectToWhatsapp()
	// }()

	// Run tests
	code := m.Run()

	os.Exit(code)
}

func CookieConfiguration() *http.Cookie {
	expiresTime, _ := time.Parse(time.RFC1123, "Tue, 24 Sep 2024 10:06:58 GMT")
	return &http.Cookie{
		Name:     "Authorization",
		Value:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjcxNzI0OTgsInN1YiI6IjE0YzlhNDQzLTAzZTUtNGJhNi05NjY0LTBmODIwYjE5ZDhhYiJ9.D7O2iFUpbd9OkpT-leVCx0_A_fGIgrHkp6JLtg6NW0w",
		Path:     "/",
		HttpOnly: true,
		Expires:  expiresTime,
	}
}

func DoctorCookieConfiguration() *http.Cookie {
	expiresTime, _ := time.Parse(time.RFC1123, "Tue, 13 Aug 2024 10:02:25 GMT")
	return &http.Cookie{
		Name:     "Authorization",
		Value:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjM1MjkyNDQsInN1YiI6IjVmYWUwYjVkLTAxZDAtNDhmYi05ZTIzLWNhMWU5YWNhMmYyMCJ9.k66qpE_pQuvvpBff-rMLJf6Ip9kdI06qsN34MLNQ2aM",
		Path:     "/",
		HttpOnly: true,
		Expires:  expiresTime,
	}
}

func WebsiteBearierTokenConfiguration() *http.Cookie {
	return &http.Cookie{
		Name:     "Authorization",
		Value:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImlkIjoiNGUyZDY4NDgtM2FlMy00NjdjLTk5NzQtZTVkMTdhYWJhMGU3IiwidXNlcm5hbWUiOiJwZW5ndXJ1c2VsZ2VrYSIsImZ1bGxfbmFtZSI6IlBlbmd1cnVzIFV0YW1hIEVMR0VLQSBKYXdhIEJhcmF0IiwiaXNfYWN0aXZlIjp0cnVlLCJzdXBlclVzZXIiOnRydWV9LCJpYXQiOjE3MjA5NDM5NzgsImV4cCI6MTcyMDk2NTU3OH0.F3eZSmWHPw6Udm_g3rXwJEYrBm_hvCWvHxJCSzyvVQI",
		Path:     "/",
		HttpOnly: true,
	}
}
