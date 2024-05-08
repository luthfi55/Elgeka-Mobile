package controllers

import (
	"elgeka-mobile/initializers"
	"elgeka-mobile/middleware"
	"elgeka-mobile/models"
	symptomresponse "elgeka-mobile/response/SymptomResponse"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SymptompController(r *gin.Engine) {
	r.POST("api/user/symptom/answer", middleware.RequireAuth, SubmitSymptom)
}

func SubmitSymptom(c *gin.Context) {
	user, _ := c.Get("user")
	var body models.SymptomAnswer

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	var data models.SymptomAnswerType
	data.Type = body.Type
	data.Answer = body.Answer
	data.Date = body.Date

	newUUID := uuid.New()
	body.ID = newUUID
	body.UserID = user.(uuid.UUID)

	if body.Date == "" {
		symptomresponse.SymptomTypeNotFoundResponse(c, "Date Can't be Null", data, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	if body.Type == "Oral" {
		OralSymptom(c, body)
	} else if body.Type == "Digestive" {
		DigestiveSymptom(c, body)
	} else if body.Type == "Respiratory" {
		RespiratorySymptom(c, body)
	} else if body.Type == "Skin" {
		SkinSymptom(c, body)
	} else if body.Type == "Hair" {
		HairSymptom(c, body)
	} else if body.Type == "Nails" {
		NailsSymptom(c, body)
	} else if body.Type == "Swelling" {
		SwellingSymptom(c, body)
	} else if body.Type == "Senses" {
		SensesSymptom(c, body)
	} else if body.Type == "Moods" {
		MoodsSymptom(c, body)
	} else if body.Type == "Pain" {
		PainSymptom(c, body)
	} else if body.Type == "Cognitive" {
		CognitiveSymptom(c, body)
	} else if body.Type == "Urinary" {
		UrinarySymptom(c, body)
	} else if body.Type == "Genitals" {
		GenitalsSymptom(c, body)
	} else if body.Type == "Reproductive" {
		ReproductiveSymptom(c, body)
	} else if body.Type == "" {
		symptomresponse.SymptomTypeNotFoundResponse(c, "Type Can't be Null.", data, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	} else {
		symptomresponse.SymptomTypeNotFoundResponse(c, "Type "+data.Type+" Not Found in Symptom.", data, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

}

func OralSymptom(c *gin.Context, body models.SymptomAnswer) {

	answers := strings.Split(body.Answer, ",")

	word_answers := strings.Split(body.WordAnswer, ",")
	length := len(answers)

	if len(answers) != 8 {
		symptomresponse.SubmitSymptomFailedResponse(c, "Number of elements in the answer array must be 8 for Oral Type, current elements = '"+strconv.Itoa(length)+"'", answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	if len(word_answers) != 8 {
		symptomresponse.SubmitSymptomFailedResponse(c, "Number of elements in the word answer array must be 8 for Oral Type, current elements = '"+strconv.Itoa(len(word_answers))+"'", word_answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Create(&body).Error; err != nil {
		symptomresponse.SubmitSymptomFailedResponse(c, "Failed to Add Oral Symptom", answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	symptomresponse.SubmitSymptomSuccessResponse(c, "Success to Add Oral Symptom", body, "http://localhost:3000/api/user/symptom/answer", http.StatusCreated)
}

func DigestiveSymptom(c *gin.Context, body models.SymptomAnswer) {
	body.Answer = strings.Trim(body.Answer, "'")
	answers := strings.Split(body.Answer, ",")
	body.WordAnswer = strings.Trim(body.WordAnswer, "'")
	word_answers := strings.Split(body.WordAnswer, ",")
	length := len(answers)

	if len(answers) != 20 {
		symptomresponse.SubmitSymptomFailedResponse(c, "Number of elements in the answer array must be 20 for Digestive Type, current elements = '"+strconv.Itoa(length)+"'", answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	if len(word_answers) != 20 {
		symptomresponse.SubmitSymptomFailedResponse(c, "Number of elements in the word answer array must be 20 for Digestive Type, current elements = '"+strconv.Itoa(len(word_answers))+"'", word_answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Create(&body).Error; err != nil {
		symptomresponse.SubmitSymptomFailedResponse(c, "Failed to Add Digestive Symptom", answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	symptomresponse.SubmitSymptomSuccessResponse(c, "Success to Add Digestive Symptom", body, "http://localhost:3000/api/user/symptom/answer", http.StatusCreated)
}

func RespiratorySymptom(c *gin.Context, body models.SymptomAnswer) {
	body.Answer = strings.Trim(body.Answer, "'")
	answers := strings.Split(body.Answer, ",")
	body.WordAnswer = strings.Trim(body.WordAnswer, "'")
	word_answers := strings.Split(body.WordAnswer, ",")
	length := len(answers)

	if len(answers) != 6 {
		symptomresponse.SubmitSymptomFailedResponse(c, "Number of elements in the answer array must be 6 for Respiratory Type, current elements = '"+strconv.Itoa(length)+"'", answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	if len(word_answers) != 6 {
		symptomresponse.SubmitSymptomFailedResponse(c, "Number of elements in the word answer array must be 6 for Respiratory Type, current elements = '"+strconv.Itoa(len(word_answers))+"'", word_answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Create(&body).Error; err != nil {
		symptomresponse.SubmitSymptomFailedResponse(c, "Failed to Add Respiratory Symptom", answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	symptomresponse.SubmitSymptomSuccessResponse(c, "Success to Add Respiratory Symptom", body, "http://localhost:3000/api/user/symptom/answer", http.StatusCreated)
}

func SkinSymptom(c *gin.Context, body models.SymptomAnswer) {
	body.Answer = strings.Trim(body.Answer, "'")
	answers := strings.Split(body.Answer, ",")
	body.WordAnswer = strings.Trim(body.WordAnswer, "'")
	word_answers := strings.Split(body.WordAnswer, ",")
	length := len(answers)

	if len(answers) != 20 {
		symptomresponse.SubmitSymptomFailedResponse(c, "Number of elements in the answer array must be 20 for Skin Type, current elements = '"+strconv.Itoa(length)+"'", answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	if len(word_answers) != 20 {
		symptomresponse.SubmitSymptomFailedResponse(c, "Number of elements in the word answer array must be 20 for Skin Type, current elements = '"+strconv.Itoa(len(word_answers))+"'", word_answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Create(&body).Error; err != nil {
		symptomresponse.SubmitSymptomFailedResponse(c, "Failed to Add Skin Symptom", answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	symptomresponse.SubmitSymptomSuccessResponse(c, "Success to Add Skin Symptom", body, "http://localhost:3000/api/user/symptom/answer", http.StatusCreated)
}

func HairSymptom(c *gin.Context, body models.SymptomAnswer) {
	body.Answer = strings.Trim(body.Answer, "'")
	answers := strings.Split(body.Answer, ",")
	body.WordAnswer = strings.Trim(body.WordAnswer, "'")
	word_answers := strings.Split(body.WordAnswer, ",")
	length := len(answers)

	if len(answers) != 1 {
		symptomresponse.SubmitSymptomFailedResponse(c, "Number of elements in the answer array must be 1 for Hair Type, current elements = '"+strconv.Itoa(length)+"'", answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	if len(word_answers) != 1 {
		symptomresponse.SubmitSymptomFailedResponse(c, "Number of elements in the answer array must be 1 for Hair Type, current elements = '"+strconv.Itoa(len(word_answers))+"'", word_answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Create(&body).Error; err != nil {
		symptomresponse.SubmitSymptomFailedResponse(c, "Failed to Add Hair Symptom", answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	symptomresponse.SubmitSymptomSuccessResponse(c, "Success to Add Hair Symptom", body, "http://localhost:3000/api/user/symptom/answer", http.StatusCreated)
}

func NailsSymptom(c *gin.Context, body models.SymptomAnswer) {
	body.Answer = strings.Trim(body.Answer, "'")
	answers := strings.Split(body.Answer, ",")
	body.WordAnswer = strings.Trim(body.WordAnswer, "'")
	word_answers := strings.Split(body.WordAnswer, ",")
	length := len(answers)

	if len(answers) != 3 {
		symptomresponse.SubmitSymptomFailedResponse(c, "Number of elements in the answer array must be 3 for Nails Type, current elements = '"+strconv.Itoa(length)+"'", answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	if len(word_answers) != 3 {
		symptomresponse.SubmitSymptomFailedResponse(c, "Number of elements in the word answer array must be 3 for Nails Type, current elements = '"+strconv.Itoa(len(word_answers))+"'", word_answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Create(&body).Error; err != nil {
		symptomresponse.SubmitSymptomFailedResponse(c, "Failed to Add Nails Symptom", answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	symptomresponse.SubmitSymptomSuccessResponse(c, "Success to Add Nails Symptom", body, "http://localhost:3000/api/user/symptom/answer", http.StatusCreated)
}

func SwellingSymptom(c *gin.Context, body models.SymptomAnswer) {
	body.Answer = strings.Trim(body.Answer, "'")
	answers := strings.Split(body.Answer, ",")
	body.WordAnswer = strings.Trim(body.WordAnswer, "'")
	word_answers := strings.Split(body.WordAnswer, ",")
	length := len(answers)

	if len(answers) != 3 {
		symptomresponse.SubmitSymptomFailedResponse(c, "Number of elements in the answer array must be 3 for Swelling Type, current elements = '"+strconv.Itoa(length)+"'", answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	if len(word_answers) != 3 {
		symptomresponse.SubmitSymptomFailedResponse(c, "Number of elements in the word answer array must be 3 for Swelling Type, current elements = '"+strconv.Itoa(len(word_answers))+"'", word_answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Create(&body).Error; err != nil {
		symptomresponse.SubmitSymptomFailedResponse(c, "Failed to Add Swelling Symptom", answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	symptomresponse.SubmitSymptomSuccessResponse(c, "Success to Add Swelling Symptom", body, "http://localhost:3000/api/user/symptom/answer", http.StatusCreated)
}

func SensesSymptom(c *gin.Context, body models.SymptomAnswer) {
	body.Answer = strings.Trim(body.Answer, "'")
	answers := strings.Split(body.Answer, ",")
	body.WordAnswer = strings.Trim(body.WordAnswer, "'")
	word_answers := strings.Split(body.WordAnswer, ",")
	length := len(answers)

	if len(answers) != 9 {
		symptomresponse.SubmitSymptomFailedResponse(c, "Number of elements in the answer array must be 9 for Senses Type, current elements = '"+strconv.Itoa(length)+"'", answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	if len(word_answers) != 9 {
		symptomresponse.SubmitSymptomFailedResponse(c, "Number of elements in the word answer array must be 9 for Senses Type, current elements = '"+strconv.Itoa(len(word_answers))+"'", word_answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Create(&body).Error; err != nil {
		symptomresponse.SubmitSymptomFailedResponse(c, "Failed to Add Senses Symptom", answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	symptomresponse.SubmitSymptomSuccessResponse(c, "Success to Add Senses Symptom", body, "http://localhost:3000/api/user/symptom/answer", http.StatusCreated)
}

func MoodsSymptom(c *gin.Context, body models.SymptomAnswer) {
	body.Answer = strings.Trim(body.Answer, "'")
	answers := strings.Split(body.Answer, ",")
	body.WordAnswer = strings.Trim(body.WordAnswer, "'")
	word_answers := strings.Split(body.WordAnswer, ",")
	length := len(answers)

	if len(answers) != 9 {
		symptomresponse.SubmitSymptomFailedResponse(c, "Number of elements in the answer array must be 9 for Moods Type, current elements = '"+strconv.Itoa(length)+"'", answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	if len(word_answers) != 9 {
		symptomresponse.SubmitSymptomFailedResponse(c, "Number of elements in the word answer array must be 9 for Moods Type, current elements = '"+strconv.Itoa(len(word_answers))+"'", word_answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Create(&body).Error; err != nil {
		symptomresponse.SubmitSymptomFailedResponse(c, "Failed to Add Moods Symptom", answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	symptomresponse.SubmitSymptomSuccessResponse(c, "Success to Add Moods Symptom", body, "http://localhost:3000/api/user/symptom/answer", http.StatusCreated)
}

func PainSymptom(c *gin.Context, body models.SymptomAnswer) {
	body.Answer = strings.Trim(body.Answer, "'")
	answers := strings.Split(body.Answer, ",")
	body.WordAnswer = strings.Trim(body.WordAnswer, "'")
	word_answers := strings.Split(body.WordAnswer, ",")
	length := len(answers)

	if len(answers) != 13 {
		symptomresponse.SubmitSymptomFailedResponse(c, "Number of elements in the answer array must be 13 for Pain Type, current elements = '"+strconv.Itoa(length)+"'", answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	if len(word_answers) != 13 {
		symptomresponse.SubmitSymptomFailedResponse(c, "Number of elements in the word answer array must be 13 for Pain Type, current elements = '"+strconv.Itoa(len(word_answers))+"'", word_answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Create(&body).Error; err != nil {
		symptomresponse.SubmitSymptomFailedResponse(c, "Failed to Add Pain Symptom", answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	symptomresponse.SubmitSymptomSuccessResponse(c, "Success to Add Pain Symptom", body, "http://localhost:3000/api/user/symptom/answer", http.StatusCreated)
}

func CognitiveSymptom(c *gin.Context, body models.SymptomAnswer) {
	body.Answer = strings.Trim(body.Answer, "'")
	answers := strings.Split(body.Answer, ",")
	body.WordAnswer = strings.Trim(body.WordAnswer, "'")
	word_answers := strings.Split(body.WordAnswer, ",")
	length := len(answers)

	if len(answers) != 11 {
		symptomresponse.SubmitSymptomFailedResponse(c, "Number of elements in the answer array must be 11 for Cognitive Type, current elements = '"+strconv.Itoa(length)+"'", answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	if len(word_answers) != 11 {
		symptomresponse.SubmitSymptomFailedResponse(c, "Number of elements in the word answer array must be 11 for Cognitive Type, current elements = '"+strconv.Itoa(len(word_answers))+"'", word_answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Create(&body).Error; err != nil {
		symptomresponse.SubmitSymptomFailedResponse(c, "Failed to Add Cognitive Symptom", answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	symptomresponse.SubmitSymptomSuccessResponse(c, "Success to Add Cognitive Symptom", body, "http://localhost:3000/api/user/symptom/answer", http.StatusCreated)
}

func UrinarySymptom(c *gin.Context, body models.SymptomAnswer) {
	body.Answer = strings.Trim(body.Answer, "'")
	answers := strings.Split(body.Answer, ",")
	body.WordAnswer = strings.Trim(body.WordAnswer, "'")
	word_answers := strings.Split(body.WordAnswer, ",")
	length := len(answers)

	if len(answers) != 8 {
		symptomresponse.SubmitSymptomFailedResponse(c, "Number of elements in the answer array must be 8 for Urinary Type, current elements = '"+strconv.Itoa(length)+"'", answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	if len(word_answers) != 8 {
		symptomresponse.SubmitSymptomFailedResponse(c, "Number of elements in the word answer array must be 8 for Urinary Type, current elements = '"+strconv.Itoa(len(word_answers))+"'", word_answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Create(&body).Error; err != nil {
		symptomresponse.SubmitSymptomFailedResponse(c, "Failed to Add Urinary Symptom", answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	symptomresponse.SubmitSymptomSuccessResponse(c, "Success to Add Urinary Symptom", body, "http://localhost:3000/api/user/symptom/answer", http.StatusCreated)
}

func GenitalsSymptom(c *gin.Context, body models.SymptomAnswer) {
	body.Answer = strings.Trim(body.Answer, "'")
	answers := strings.Split(body.Answer, ",")
	body.WordAnswer = strings.Trim(body.WordAnswer, "'")
	word_answers := strings.Split(body.WordAnswer, ",")
	length := len(answers)

	if len(answers) != 6 {
		symptomresponse.SubmitSymptomFailedResponse(c, "Number of elements in the answer array must be 6 for Genitals Type, current elements = '"+strconv.Itoa(length)+"'", answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	if len(word_answers) != 6 {
		symptomresponse.SubmitSymptomFailedResponse(c, "Number of elements in the word answer array must be 6 for Genitals Type, current elements = '"+strconv.Itoa(len(word_answers))+"'", word_answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Create(&body).Error; err != nil {
		symptomresponse.SubmitSymptomFailedResponse(c, "Failed to Add Genitals Symptom", answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	symptomresponse.SubmitSymptomSuccessResponse(c, "Success to Add Genitals Symptom", body, "http://localhost:3000/api/user/symptom/answer", http.StatusCreated)
}

func ReproductiveSymptom(c *gin.Context, body models.SymptomAnswer) {
	body.Answer = strings.Trim(body.Answer, "'")
	answers := strings.Split(body.Answer, ",")
	body.WordAnswer = strings.Trim(body.WordAnswer, "'")
	word_answers := strings.Split(body.WordAnswer, ",")
	length := len(answers)

	if len(answers) != 5 {
		symptomresponse.SubmitSymptomFailedResponse(c, "Number of elements in the answer array must be 5 for Reproductive Type, current elements = '"+strconv.Itoa(length)+"'", answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	if len(word_answers) != 5 {
		symptomresponse.SubmitSymptomFailedResponse(c, "Number of elements in the word answer array must be 5 for Reproductive Type, current elements = '"+strconv.Itoa(len(word_answers))+"'", word_answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	if err := initializers.DB.Create(&body).Error; err != nil {
		symptomresponse.SubmitSymptomFailedResponse(c, "Failed to Add Reproductive Symptom", answers, "http://localhost:3000/api/user/symptom/answer", http.StatusBadRequest)
		return
	}

	symptomresponse.SubmitSymptomSuccessResponse(c, "Success to Add Reproductive Symptom", body, "http://localhost:3000/api/user/symptom/answer", http.StatusCreated)
}
