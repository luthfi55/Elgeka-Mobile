package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables(envFile string) error {
	// Load environment variables from the specified file
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
		return err
	}
	return nil
}
