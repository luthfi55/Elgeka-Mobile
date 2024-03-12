package main

import (
	"os"
	"os/signal"
	"syscall"

	"elgeka-mobile/controllers"
	"elgeka-mobile/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, os.Interrupt, syscall.SIGTERM)
	go func() {
		initializers.ConnectToWhatsapp()
	}()
	r := gin.Default()

	controllers.LoginController(r)
	controllers.RegisterController(r)
	controllers.ActivateAccountController(r)
	controllers.HealthStatusController(r)
	controllers.UserProfileController(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	go func() {
		<-shutdownSignal

		initializers.DisconnectWhatsapp()

		os.Exit(0)
	}()

	r.Run("0.0.0.0:" + port)
}
