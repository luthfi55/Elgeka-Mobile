package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"elgeka-mobile/controllers"
	"elgeka-mobile/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables(".env")
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func connectToWhatsappWithRetry() {
	for {
		initializers.ConnectToWhatsapp()
		log.Println("Connected to WhatsApp")
		time.Sleep(5 * time.Second)
	}
}

func main() {
	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, os.Interrupt, syscall.SIGTERM)

	go connectToWhatsappWithRetry()
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin == "https://stirring-pixie-ed5c9b.netlify.app" || origin == "https://elgeka-community-hub.netlify.app" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	})

	controllers.LoginController(r)
	controllers.RegisterController(r)
	controllers.ActivateAccountController(r)
	controllers.HealthStatusController(r)
	controllers.UserProfileController(r)
	controllers.MedicineController(r)
	controllers.SymptompController(r)
	controllers.DoctorProfileController(r)
	controllers.DoctorChartController(r)
	controllers.UserTreatmentController(r)

	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "8080"
	// }

	go func() {
		<-shutdownSignal
		os.Exit(0)
	}()

	// r.Run("0.0.0.0:" + port)
	r.Run(os.Getenv("PORT"))

}
