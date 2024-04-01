package initializers

import (
	"elgeka-mobile/models"

	"gorm.io/gorm"
)

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
		&models.UserPersonalDoctor{},
		&models.Medicine{},
		&models.SymptomQuestion{},
		&models.SymptomAnswer{},
	)

	SeedData(DB)
}

func SeedData(db *gorm.DB) {
	oralSymptoms := []models.SymptomQuestion{
		{ID: "OL-01", Type: "Oral", Question: "Berapa Tingkat Keparahan Mulut Kering Anda yang paling parah?", Option: "{'Tidak', 'Ringan', 'Sedang', 'Parah', 'Sangat Parah}"},
		{ID: "OL-02", Type: "Oral", Question: "Seberapa parahkah tingkat keparahan kesulitan menelan yang Anda alami?", Option: "{'Tidak', 'Ringan', 'Sedang', 'Parah', 'Sangat Parah}"},
		{ID: "OL-03", Type: "Oral", Question: "Seberapa parahkah tingkat keparahan luka di mulut atau tenggorokan Anda?", Option: "{'Tidak sama sekali', 'Sedikit', 'Agak', 'Cukup Sedikit', 'Sangat Parah}"},
		{ID: "OL-04", Type: "Oral", Question: "Seberapa besar sariawan atau luka di mulut mengganggu aktivitas Anda sehari-hari?", Option: "{'Tidak sama sekali', 'Sedikit', 'Agak', 'Cukup Sedikit', 'Sangat Parah}"},
	}

	for _, symptom := range oralSymptoms {
		db.Create(&symptom)
	}
}
