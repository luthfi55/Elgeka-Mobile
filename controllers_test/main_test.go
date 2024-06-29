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
	expiresTime, _ := time.Parse(time.RFC1123, "Sat, 22 Jun 2024 10:02:25 GMT")
	return &http.Cookie{
		Name:     "Authorization",
		Value:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTkwNTA1NDUsInN1YiI6IjE0YzlhNDQzLTAzZTUtNGJhNi05NjY0LTBmODIwYjE5ZDhhYiJ9.A9zLQeWYlOZwztxHhB_ZBU3qZQUsWpQjvpI3ucus7k8",
		Path:     "/",
		HttpOnly: true,
		Expires:  expiresTime,
	}
}

func DoctorCookieConfiguration() *http.Cookie {
	expiresTime, _ := time.Parse(time.RFC1123, "Sat, 22 Jun 2024 10:02:25 GMT")
	return &http.Cookie{
		Name:     "Authorization",
		Value:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjA2ODM4MDEsInN1YiI6IjVmYWUwYjVkLTAxZDAtNDhmYi05ZTIzLWNhMWU5YWNhMmYyMCJ9.OJO2SW3W7MDRH64_ysnYTjX3Ol1vpg1s0NUfgaXsRhc",
		Path:     "/",
		HttpOnly: true,
		Expires:  expiresTime,
	}
}

func WebsiteBearierTokenConfiguration() *http.Cookie {
	return &http.Cookie{
		Name:     "Authorization",
		Value:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImlkIjoiNGUyZDY4NDgtM2FlMy00NjdjLTk5NzQtZTVkMTdhYWJhMGU3IiwidXNlcm5hbWUiOiJwZW5ndXJ1c2VsZ2VrYSIsImZ1bGxfbmFtZSI6IlBlbmd1cnVzIFV0YW1hIEVMR0VLQSBKQUJBUiIsImlzX2FjdGl2ZSI6dHJ1ZSwic3VwZXJVc2VyIjp0cnVlfSwiaWF0IjoxNzE4MDkwOTE5LCJleHAiOjE3MTgxMTI1MTl9.uZKwT15qe6Q6xytlbLtFQNysSoKd4VR0MI5diYqqge8",
		Path:     "/",
		HttpOnly: true,
	}
}
