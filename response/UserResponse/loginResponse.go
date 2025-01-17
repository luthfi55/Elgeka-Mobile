package userresponse

import (
	"elgeka-mobile/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckAccountFailedResponse(c *gin.Context) {
	linkItem := models.LinkItem{
		Name: "Login",
		Link: "http://localhost:3000/api/user/login",
	}
	var UserId []models.UserIdData
	response := models.LoginUserFailedResponse{
		ErrorMessage: "Need  Login",
		Data:         UserId,
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(http.StatusBadRequest, response)
}

func LoginFailedResponse(c *gin.Context, message string, data models.UserIdData, link string, status int) {
	datas := models.UserIdData{
		ID: data.ID,
	}

	linkItem := models.LinkItem{
		Name: "Login",
		Link: link,
	}

	response := models.LoginUserFailedResponse{
		ErrorMessage: message,
		Data:         []models.UserIdData{datas},
		Link:         []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func LoginSuccessResponse(c *gin.Context, message string, data models.Login, link string, status int) {
	datas := models.Login{
		Name:  data.Name,
		Email: data.Email,
	}

	linkItem := models.LinkItem{
		Name: "Dashboard",
		Link: link,
	}

	response := models.LoginUserSuccessResponse{
		Message: message,
		Data:    []models.Login{datas},
		Link:    []models.LinkItem{linkItem},
	}

	c.JSON(status, response)
}

func LoginWebsiteSuccessResponse(c *gin.Context, message string, data models.Login, link string, token string, status int) {
	datas := models.Login{
		Name:  data.Name,
		Email: data.Email,
	}

	linkItem := models.LinkItem{
		Name: "Dashboard",
		Link: link,
	}

	response := models.LoginUserWebsiteSuccessResponse{
		Message: message,
		Data:    []models.Login{datas},
		Link:    []models.LinkItem{linkItem},
		Token:   token,
	}

	c.JSON(status, response)
}
