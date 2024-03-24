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
	expiresTime, _ := time.Parse(time.RFC1123, "Mon, 15 Apr 2024 17:00:20 GMT")
	return &http.Cookie{
		Name:     "Authorization",
		Value:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTMyMDA0MjAsInN1YiI6IjE0YzlhNDQzLTAzZTUtNGJhNi05NjY0LTBmODIwYjE5ZDhhYiJ9.L9z84gPX0l_O3GeyRi0ZAhMGxoWzXVV7k9fXw6KpEo4",
		Path:     "/",
		HttpOnly: true,
		Expires:  expiresTime,
	}
}
