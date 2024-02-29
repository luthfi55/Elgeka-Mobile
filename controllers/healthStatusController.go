package controllers

import (
	"elgeka-mobile/initializers"
	"elgeka-mobile/middleware"
	"elgeka-mobile/models"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func HealthStatusController(r *gin.Engine) {
	r.POST("api/user/health_status/bcr_abl", middleware.RequireAuth, CreateBcrAbl)
	r.GET("api/user/health_status/bcr_abl", middleware.RequireAuth, GetBcrAbl)
	r.PUT("api/user/health_status/bcr_abl/:bcr_abl_id", middleware.RequireAuth, UpdateBcrAbl)
	r.DELETE("api/user/health_status/bcr_abl/:bcr_abl_id", middleware.RequireAuth, DeleteBcrAbl)
}

func CreateBcrAbl(c *gin.Context) {
	var body models.BCR_ABL
	user, _ := c.Get("user")

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("%s %s", err.Field(), getValidationErrorTagMessage(err.Tag())))
		}
		errorMessage := strings.Join(validationErrors, ", ") // Join errors into a single string
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessage,
		})
		return
	}

	newUUID := uuid.New()
	bcr_abl := models.BCR_ABL{
		ID:     newUUID,
		UserID: user.(uuid.UUID),
		Data:   body.Data,
		Notes:  body.Notes,
		Date:   body.Date,
	}

	if err := initializers.DB.Create(&bcr_abl).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"User":  user,
		"Data":  body.Data,
		"Notes": body.Notes,
		"Date":  body.Date,
	})
}

func GetBcrAbl(c *gin.Context) {
	var bcr_abl []models.BCR_ABL
	user, _ := c.Get("user")

	initializers.DB.Where("user_id = ?", user).Find(&bcr_abl)

	if initializers.DB.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": initializers.DB.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"User": user,
		"Data": bcr_abl,
	})
}

func UpdateBcrAbl(c *gin.Context) {
	var body models.BCR_ABL
	bcr_abl_id := c.Param("bcr_abl_id")

	user, _ := c.Get("user")

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("%s %s", err.Field(), getValidationErrorTagMessage(err.Tag())))
		}
		errorMessage := strings.Join(validationErrors, ", ") // Join errors into a single string
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessage,
		})
		return
	}

	var bcr_abl models.BCR_ABL
	initializers.DB.First(&bcr_abl, "ID = ?", bcr_abl_id)

	if initializers.DB.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": initializers.DB.Error.Error(),
		})
		return
	}

	bcr_abl = models.BCR_ABL{
		ID:     uuid.MustParse(bcr_abl_id),
		UserID: user.(uuid.UUID),
		Data:   body.Data,
		Notes:  body.Notes,
		Date:   body.Date,
	}

	if err := initializers.DB.Save(&bcr_abl).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"User":  user,
		"Data":  body.Data,
		"Notes": body.Notes,
		"Date":  body.Date,
	})

}

func DeleteBcrAbl(c *gin.Context) {
	bcr_abl_id := c.Param("bcr_abl_id")

	var bcr_abl models.BCR_ABL
	initializers.DB.First(&bcr_abl, "ID = ?", bcr_abl_id)

	if initializers.DB.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": initializers.DB.Error.Error(),
		})
		return
	}

	if err := initializers.DB.Unscoped().Delete(&bcr_abl).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}
