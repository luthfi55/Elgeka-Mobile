package initializers

import "elgeka-mobile/models"

func SyncDatabase() {
	DB.AutoMigrate(
		&models.User{},
		&models.Doctor{},
	)
}
