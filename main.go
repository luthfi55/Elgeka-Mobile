package main

import (
	"os"

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
	r := gin.Default()

	controllers.LoginController(r)
	controllers.RegisterController(r)
	controllers.ActivateAccountController(r)

	r.Run(os.Getenv("PORT"))
}
