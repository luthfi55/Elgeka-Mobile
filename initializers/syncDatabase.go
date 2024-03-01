package initializers

import "elgeka-mobile/models"

func SyncDatabase() {
	DB.AutoMigrate(
		&models.User{},
		&models.Doctor{},
		&models.BCR_ABL{},
		&models.Leukocytes{},
		&models.PotentialHydrogen{},
		&models.Hemoglobin{},
		&models.HeartRate{},
		&models.BloodPressure{},
	)
}
