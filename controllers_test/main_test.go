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
		Value:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTcwNTgxNjQsInN1YiI6IjVmYWUwYjVkLTAxZDAtNDhmYi05ZTIzLWNhMWU5YWNhMmYyMCJ9.j4kcnBLIyOPnXl5Ok1gZKhqsAysXs2MTsEmzN23sVi0",
		Path:     "/",
		HttpOnly: true,
		Expires:  expiresTime,
	}
}

func WebsiteBearierTokenConfiguration() *http.Cookie {
	return &http.Cookie{
		Name:     "Authorization",
		Value:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTg4MTEwMzYsInN1YiI6IjE0YzlhNDIzLTAzZTEtNGJhNi05NjY0LTBmODIwYjE5ZDhhYiJ9.mJ7Xmk5yJwj2ISKci06xl0Gapn3VSNTGTaRXOZcRAKY",
		Path:     "/",
		HttpOnly: true,
	}
}
