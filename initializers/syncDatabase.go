package initializers

import (
	"elgeka-mobile/models"
)

func SyncDatabase() {
	DB.AutoMigrate(
		&models.User{},
		&models.UserInformation{},
		&models.Doctor{},
		&models.BCR_ABL{},
		&models.Leukocytes{},
		&models.PotentialHydrogen{},
		&models.Hemoglobin{},
		&models.HeartRate{},
		&models.BloodPressure{},
		&models.UserPersonalDoctor{},
		&models.Medicine{},
		&models.MedicineSchedule{},
		// &models.SymptomQuestion{},
		&models.SymptomAnswer{},
		&models.UserTreatment{},
	)

	// SeedData(DB)
}

// func SeedData(db *gorm.DB) {
// 	oralSymptoms := []models.SymptomQuestion{
// 		{ID: "OL-01", Type: "Oral", Question: "Seberapa parah tingkat kekeringan mulut yang anda rasakan?"},
// 		{ID: "OL-02", Type: "Oral", Question: "Seberapa parah kesulitan menelan yang anda rasakan?"},
// 		{ID: "OL-03", Type: "Oral", Question: "Seberapa parah luka di mulut atau tenggorokan yang anda rasakan?"},
// 		{ID: "OL-04", Type: "Oral", Question: "Seberapa besar aktivitas yang terganggu oleh sariawan atau luka di mulut?"},
// 		{ID: "OL-05", Type: "Oral", Question: "Seberapa parah kulit pecah-pecah di sudut mulut yang anda rasakan?"},
// 		{ID: "OL-06", Type: "Oral", Question: "Apakah anda mengalami perubahan suara?"},
// 		{ID: "OL-07", Type: "Oral", Question: "Seberapa parah suara serak yang anda rasakan?"},
// 		{ID: "OL-08", Type: "Oral", Question: "Seberapa parah masalah dalam merasakan makanan atau minuman yang anda rasakan?"},
// 	}

// 	digestiveSymptoms := []models.SymptomQuestion{
// 		{ID: "DG-01", Type: "Digestive", Question: "Seberapa parah penurunan nafsu makan yang anda rasakan?"},
// 		{ID: "DG-02", Type: "Digestive", Question: "Seberapa besar aktivitas yang terganggu oleh nafsu makan yang menurun?"},
// 		{ID: "DG-03", Type: "Digestive", Question: "Seberapa parah mual yang anda rasakan?"},
// 		{ID: "DG-04", Type: "Digestive", Question: "Seberapa sering anda mengalami mual?"},
// 		{ID: "DG-05", Type: "Digestive", Question: "Seberapa parah muntah yang anda alami?"},
// 		{ID: "DG-06", Type: "Digestive", Question: "Seberapa sering anda mengalami muntah?"},
// 		{ID: "DG-07", Type: "Digestive", Question: "Seberapa parah sakit maag yang anda rasakan?"},
// 		{ID: "DG-08", Type: "Digestive", Question: "Seberapa sering anda mengalami sakit maag?"},
// 		{ID: "DG-09", Type: "Digestive", Question: "Apakah anda mengalami peningkatan buang angin atau gas (perut kembung)?"},
// 		{ID: "DG-10", Type: "Digestive", Question: "Seberapa parah kembung pada perut yang anda rasakan?"},
// 		{ID: "DG-11", Type: "Digestive", Question: "Seberapa sering anda mengalami kembung pada perut?"},
// 		{ID: "DG-12", Type: "Digestive", Question: "Seberapa parah cegukan yang anda rasakan?"},
// 		{ID: "DG-13", Type: "Digestive", Question: "Seberapa sering anda mengalami cegukan?"},
// 		{ID: "DG-14", Type: "Digestive", Question: "Seberapa parah sembelit yang anda rasakan?"},
// 		{ID: "DG-15", Type: "Digestive", Question: "Seberapa sering anda mengalami mencret?"},
// 		{ID: "DG-16", Type: "Digestive", Question: "Seberapa parah nyeri di perut yang anda rasakan?"},
// 		{ID: "DG-17", Type: "Digestive", Question: "Seberapa sering anda mengalami nyeri di perut?"},
// 		{ID: "DG-18", Type: "Digestive", Question: "Seberapa besar aktivitas yang terganggu oleh nyeri di perut?"},
// 		{ID: "DG-19", Type: "Digestive", Question: "Seberapa besar aktivitas yang terganggu oleh kehilangan kendali buang air besar?"},
// 		{ID: "DG-20", Type: "Digestive", Question: "Seberapa sering anda kehilangan kendali buang air besar?"},
// 	}

// 	respiratorySymptoms := []models.SymptomQuestion{
// 		{ID: "RP-01", Type: "Respiratory", Question: "Seberapa parah sesak napas yang anda rasakan?"},
// 		{ID: "RP-02", Type: "Respiratory", Question: "Seberapa besar aktivitas yang terganggu oleh sesak napas?"},
// 		{ID: "RP-03", Type: "Respiratory", Question: "Seberapa parah batuk yang anda rasakan?"},
// 		{ID: "RP-04", Type: "Respiratory", Question: "Seberapa besar aktivitas yang terganggu oleh mengi (suara seperti siulan saat bernapas)?"},
// 		{ID: "RP-05", Type: "Respiratory", Question: "Seberapa parah jantung berdebar (palpitasi) yang anda rasakan?"},
// 		{ID: "RP-06", Type: "Respiratory", Question: "Seberapa sering anda merasakan jantung berdebar (palpitasi)?"},
// 	}

// 	skinSymptoms := []models.SymptomQuestion{
// 		{ID: "SK-01", Type: "Skin", Question: "Apa anda mengalami ruam?"},
// 		{ID: "SK-02", Type: "Skin", Question: "Seberapa parah kulit kering yang anda rasakan?"},
// 		{ID: "SK-03", Type: "Skin", Question: "Seberapa parah kulit gatal yang anda rasakan?"},
// 		{ID: "SK-04", Type: "Skin", Question: "Apa anda mengalami kulit bentol?"},
// 		{ID: "SK-05", Type: "Skin", Question: "Seberapa parah Hand-Foot (ruam pada tangan atau kaki yang menyebabkan kulit pecah, mengelupas, kemerahan, atau nyeri) yang anda rasakan?"},
// 		{ID: "SK-06", Type: "Skin", Question: "Apa anda mengalami luka merah?"},
// 		{ID: "SK-07", Type: "Skin", Question: "Seberapa parah luka bakar yang anda rasakan?"},
// 		{ID: "SK-08", Type: "Skin", Question: "Apa anda mengalami penggelapan kulit yang tidak biasa?"},
// 		{ID: "SK-09", Type: "Skin", Question: "Apa anda mengalami stretch marks?"},
// 		{ID: "SK-10", Type: "Skin", Question: "Apa anda mudah mengalami memar (berwarna hitam atau biru)?"},
// 		{ID: "SK-11", Type: "Skin", Question: "Seberapa parah menggigil atau gemetar yang anda rasakan?"},
// 		{ID: "SK-12", Type: "Skin", Question: "Seberapa sering anda merasakan menggigil atau gemetar?"},
// 		{ID: "SK-13", Type: "Skin", Question: "Seberapa parah berkeringat tak terduga yang anda rasakan?"},
// 		{ID: "SK-14", Type: "Skin", Question: "Seberapa sering anda berkeringat secara tak terduga di siang atau malam hari (tidak berhubungan dengan sengatan panas atau muka memerah)?"},
// 		{ID: "SK-15", Type: "Skin", Question: "Apa anda mengalami pengurangan keringat tak terduga?"},
// 		{ID: "SK-16", Type: "Skin", Question: "Seberapa parah hot flashes yang anda rasakan?"},
// 		{ID: "SK-17", Type: "Skin", Question: "Seberapa sering anda merasakan hot flashes?"},
// 		{ID: "SK-18", Type: "Skin", Question: "Seberapa parah mimisan yang anda rasakan?"},
// 		{ID: "SK-19", Type: "Skin", Question: "Seberapa sering anda merasakan mimisan?"},
// 		{ID: "SK-20", Type: "Skin", Question: "Seberapa parah bau badan yang anda rasakan?"},
// 	}

// 	hairSymptoms := []models.SymptomQuestion{
// 		{ID: "HR-01", Type: "Hair", Question: "Apakah anda mengalami kerontokan rambut?"},
// 	}

// 	nailsSymptoms := []models.SymptomQuestion{
// 		{ID: "NS-01", Type: "Nails", Question: "Apakah anda mengalami kehiilangan kuku jari tangan atau kaki?"},
// 		{ID: "NS-02", Type: "Nails", Question: "Apakah anda mengalami benjol pada kuku jari tangan atau kaki?"},
// 		{ID: "NS-03", Type: "Nails", Question: "Apakah anda mengalami perubahan warna pada kuku jari tangan atau kaki?"},
// 	}

// 	swellingSymptoms := []models.SymptomQuestion{
// 		{ID: "SG-01", Type: "Nails", Question: "Seberapa parah pembengkakan pada lengan atau kaki yang anda rasakan?"},
// 		{ID: "SG-02", Type: "Nails", Question: "Seberapa sering anda mengalami pembengkakan pada lengan atau kaki?"},
// 		{ID: "SG-03", Type: "Nails", Question: "Seberapa besar aktivitas yang terganggu oleh pembengkakan pada lengan atau kaki?"},
// 	}

// 	sensesSymptoms := []models.SymptomQuestion{
// 		{ID: "SS-01", Type: "Senses", Question: "Apakah anda merasa lebih sensitif terhadap sinar matahari?"},
// 		{ID: "SS-02", Type: "Senses", Question: "Seberapa parah kesemutan atau mati rasa pada tangan atau kaki yang anda rasakan?"},
// 		{ID: "SS-03", Type: "Senses", Question: "Seberapa besar aktivitas yang terganggu oleh kesemutan atau mati rasa pada lengan atau kaki?"},
// 		{ID: "SS-04", Type: "Senses", Question: "Seberapa parah pengelihatan buram yang anda rasakan?"},
// 		{ID: "SS-05", Type: "Senses", Question: "Seberapa besar aktivitas yang terganggu oleh pengelihatan buram?"},
// 		{ID: "SS-06", Type: "Senses", Question: "Apakah anda merasa kilatan cahaya di depan mata?"},
// 		{ID: "SS-07", Type: "Senses", Question: "Apakah anda merasa bintik atau garis yang melayang di depan mata?"},
// 		{ID: "SS-08", Type: "Senses", Question: "Seberapa parah mata berair yang anda rasakan?"},
// 		{ID: "SS-09", Type: "Senses", Question: "Seberapa besar aktivitas yang terganggu oleh mata berair?"},
// 	}

// 	moodsSymptoms := []models.SymptomQuestion{
// 		{ID: "MD-01", Type: "Moods", Question: "Seberapa parah kecemasan yang anda rasakan?"},
// 		{ID: "MD-02", Type: "Moods", Question: "Seberapa sering anda mengalami kecemasan?"},
// 		{ID: "MD-03", Type: "Moods", Question: "Seberapa besar aktivitas yang terganggu oleh kecemasan?"},
// 		{ID: "MD-04", Type: "Moods", Question: "Seberapa parah kehilangan semangat yang anda rasakan?"},
// 		{ID: "MD-05", Type: "Moods", Question: "Seberapa sering anda mengalami kehilangan semangat?"},
// 		{ID: "MD-06", Type: "Moods", Question: "Seberapa besar aktivitas yang terganggu oleh kehilangan semangat?"},
// 		{ID: "MD-07", Type: "Moods", Question: "Seberapa parah kesedihan yang anda rasakan?"},
// 		{ID: "MD-08", Type: "Moods", Question: "Seberapa sering anda merasa sedih?"},
// 		{ID: "MD-09", Type: "Moods", Question: "Seberapa besar aktivitas yang terganggu oleh kesedihan?"},
// 	}

// 	painSymptoms := []models.SymptomQuestion{
// 		{ID: "PN-01", Type: "Pain", Question: "Seberapa parah nyeri yang anda rasakan?"},
// 		{ID: "PN-02", Type: "Pain", Question: "Seberapa sering anda mengalami nyeri?"},
// 		{ID: "PN-03", Type: "Pain", Question: "Seberapa besar aktivitas yang terganggu oleh nyeri?"},
// 		{ID: "PN-04", Type: "Pain", Question: "Seberapa parah sakit kepala yang anda rasakan?"},
// 		{ID: "PN-05", Type: "Pain", Question: "Seberapa sering anda mengalami sakit kepala?"},
// 		{ID: "PN-06", Type: "Pain", Question: "Seberapa besar aktivitas yang terganggu oleh sakit kepala?"},
// 		{ID: "PN-07", Type: "Pain", Question: "Seberapa parah nyeri otot yang anda rasakan?"},
// 		{ID: "PN-08", Type: "Pain", Question: "Seberapa sering anda mengalami nyeri otot?"},
// 		{ID: "PN-09", Type: "Pain", Question: "Seberapa besar aktivitas yang terganggu oleh nyeri otot?"},
// 		{ID: "PN-10", Type: "Pain", Question: "Seberapa parah nyeri persendian yang anda rasakan?"},
// 		{ID: "PN-11", Type: "Pain", Question: "Seberapa sering anda mengalami nyeri persendian?"},
// 		{ID: "PN-12", Type: "Pain", Question: "Seberapa besar aktivitas yang terganggu oleh nyeri persendian?"},
// 		{ID: "PN-13", Type: "Pain", Question: "Apakah anda merasakan nyeri, pada penggunaan obat suntik atau infus?"},
// 	}

// 	cognitiveSymptoms := []models.SymptomQuestion{
// 		{ID: "CG-01", Type: "Cognitive", Question: "Seberapa parah pusing yang anda rasakan?"},
// 		{ID: "CG-02", Type: "Cognitive", Question: "Seberapa besar aktivitas yang terganggu oleh pusing?"},
// 		{ID: "CG-03", Type: "Cognitive", Question: "Seberapa parah telinga berdenging yang anda rasakan?"},
// 		{ID: "CG-04", Type: "Cognitive", Question: "Seberapa parah masalah konsentrasi yang anda rasakan?"},
// 		{ID: "CG-05", Type: "Cognitive", Question: "Seberapa besar aktivitas yang terganggu oleh masalah konsentrasi?"},
// 		{ID: "CG-06", Type: "Cognitive", Question: "Seberapa parah masalah ingatan yang anda rasakan?"},
// 		{ID: "CG-07", Type: "Cognitive", Question: "Seberapa besar aktivitas yang terganggu oleh masalah ingatan?"},
// 		{ID: "CG-08", Type: "Cognitive", Question: "Seberapa parah insomnia yang anda rasakan?"},
// 		{ID: "CG-09", Type: "Cognitive", Question: "Seberapa besar aktivitas yang terganggu oleh insomnia?"},
// 		{ID: "CG-10", Type: "Cognitive", Question: "Seberapa parah kelelahan yang anda rasakan?"},
// 		{ID: "CG-11", Type: "Cognitive", Question: "Seberapa besar aktivitas yang terganggu oleh kelelahan?"},
// 	}

// 	urinarySymptoms := []models.SymptomQuestion{
// 		{ID: "UR-01", Type: "Urinary", Question: "Seberapa parah sakit saat buang air kecil yang anda rasakan?"},
// 		{ID: "UR-02", Type: "Urinary", Question: "Seberapa sering anda mengalami rasa ingin buang air kecil secara mendadak?"},
// 		{ID: "UR-03", Type: "Urinary", Question: "Seberapa besar aktivitas yang terganggu oleh rasa ingin buang air kecil secara mendadak?"},
// 		{ID: "UR-04", Type: "Urinary", Question: "Apa ada waktu di mana anda harus sering buang air kecil?"},
// 		{ID: "UR-05", Type: "Urinary", Question: "Seberapa besar aktivitas yang terganggu oleh sering buang air kecil?"},
// 		{ID: "UR-06", Type: "Urinary", Question: "Apa anda mengalami perubahan warna urin?"},
// 		{ID: "UR-07", Type: "Urinary", Question: "Seberapa sering anda mengompol?"},
// 		{ID: "UR-08", Type: "Urinary", Question: "Seberapa besar aktivitas yang terganggu oleh mengompol?"},
// 	}

// 	genitalsSymptoms := []models.SymptomQuestion{
// 		{ID: "GL-01", Type: "Genitals", Question: "Apa anda mengalami periode menstruasi yang tidak normal?"},
// 		{ID: "GL-02", Type: "Genitals", Question: "Apa anda melewati periode menstruasi yang harusnya terjadi?"},
// 		{ID: "GL-03", Type: "Genitals", Question: "Apa anda mengalami keputihan yang tidak biasa?"},
// 		{ID: "GL-04", Type: "Genitals", Question: "Seberapa parah kekeringan vagina yang anda rasakan?"},
// 		{ID: "GL-05", Type: "Genitals", Question: "Seberapa parah sakit di vagina saat berhubungan badan yang anda rasakan?"},
// 		{ID: "GL-06", Type: "Genitals", Question: "Seberapa parah pembesaran atau pelembutan area payudara yang anda rasakan?"},
// 	}

// 	reproductiveSymptoms := []models.SymptomQuestion{
// 		{ID: "RD-01", Type: "Reproductive", Question: "Seberapa parah kesulitan ereksi yang anda rasakan?"},
// 		{ID: "RD-02", Type: "Reproductive", Question: "Seberapa sering anda mengalami masalah ejakulasi?"},
// 		{ID: "RD-03", Type: "Reproductive", Question: "Seberapa parah penurunan gairah seksual yang anda rasakan?"},
// 		{ID: "RD-04", Type: "Reproductive", Question: "Apa anda merasa membutuhkan terlalu banyak waktu untuk orgasme atau klimaks?"},
// 		{ID: "RD-05", Type: "Reproductive", Question: "Apa anda tidak bisa orgasme atau klimaks?"},
// 	}

// 	for _, symptom := range oralSymptoms {
// 		db.Create(&symptom)
// 	}

// 	for _, symptom := range digestiveSymptoms {
// 		db.Create(&symptom)
// 	}

// 	for _, symptom := range respiratorySymptoms {
// 		db.Create(&symptom)
// 	}

// 	for _, symptom := range skinSymptoms {
// 		db.Create(&symptom)
// 	}

// 	for _, symptom := range hairSymptoms {
// 		db.Create(&symptom)
// 	}

// 	for _, symptom := range nailsSymptoms {
// 		db.Create(&symptom)
// 	}

// 	for _, symptom := range swellingSymptoms {
// 		db.Create(&symptom)
// 	}

// 	for _, symptom := range sensesSymptoms {
// 		db.Create(&symptom)
// 	}

// 	for _, symptom := range moodsSymptoms {
// 		db.Create(&symptom)
// 	}

// 	for _, symptom := range painSymptoms {
// 		db.Create(&symptom)
// 	}

// 	for _, symptom := range cognitiveSymptoms {
// 		db.Create(&symptom)
// 	}

// 	for _, symptom := range urinarySymptoms {
// 		db.Create(&symptom)
// 	}

// 	for _, symptom := range genitalsSymptoms {
// 		db.Create(&symptom)
// 	}

// 	for _, symptom := range reproductiveSymptoms {
// 		db.Create(&symptom)
// 	}

// }
