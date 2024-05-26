package controllers_test

import (
	"bytes"
	"elgeka-mobile/controllers"
	"elgeka-mobile/middleware"
	"elgeka-mobile/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

type successSubmitSymptomResponse struct {
	Message string               `json:"Message"`
	Data    models.SymptomAnswer `json:"Data"`
	Link    []struct {
		Name string `json:"Name"`
		Link string `json:"Link"`
	} `json:"Link"`
}

type failedSubmitSymptomResponse struct {
	ErrorMessage string   `json:"ErrorMessage"`
	Data         []string `json:"Data"`
	Link         []struct {
		Name string `json:"Name"`
		Link string `json:"Link"`
	} `json:"Link"`
}

func TestSubmitSymptom_Success(t *testing.T) {
	router := gin.Default()

	router.POST("/api/user/symptom/answer", middleware.RequireAuth, controllers.SubmitSymptom)

	reqBody := models.SymptomAnswer{
		Type:       "Oral",
		Answer:     "1,2,4,5,2,4,4,3",
		WordAnswer: "'Parah','Parah','Parah','Sedikit','Parah','Tidak','Parah','Sangat Parah''",
		Date:       "2024-03-11",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/api/user/symptom/answer", bytes.NewBuffer(reqJSON))
	req.AddCookie(CookieConfiguration())

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected status code %d but got %d", http.StatusCreated, rec.Code)
	}

	var expectedBody successSubmitSymptomResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.Message != "Success to Add Oral Symptom" {
		t.Errorf("expected message %s but got %s", "Success to Add Oral Symptom", expectedBody.Message)
	}
}

func TestSubmitSymptom_Failed(t *testing.T) {
	router := gin.Default()

	router.POST("/api/user/symptom/answer", middleware.RequireAuth, controllers.SubmitSymptom)

	reqBody := models.SymptomAnswer{
		Type:   "Oral",
		Answer: "'Parah','Parah','Parah','Sedikit','Parah','Tidak','Parah'",
		Date:   "2024-03-11",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/api/user/symptom/answer", bytes.NewBuffer(reqJSON))
	req.AddCookie(CookieConfiguration())

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d but got %d", http.StatusBadRequest, rec.Code)
	}

	var expectedBody failedSubmitSymptomResponse
	err = json.Unmarshal(rec.Body.Bytes(), &expectedBody)
	if err != nil {
		t.Fatal(err)
	}

	if expectedBody.ErrorMessage != "Number of elements in the answer array must be 8 for Oral Type, current elements = '7'" {
		t.Errorf("expected message %s but got %s", "Number of elements in the answer array must be 8 for Oral Type, current elements = '7'", expectedBody.ErrorMessage)
	}
}
