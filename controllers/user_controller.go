package controllers

import (
	"net/http"
	"rest_api_mongodb/models"
	"rest_api_mongodb/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func New(userservice services.UserService) UserController {
	return UserController{
		UserService: userservice,
	}
}

func (uc *UserController) CreateUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = uc.UserService.CreateUser(&user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) GetUser(context *gin.Context) {
	username := context.Param("name")
	user, err := uc.UserService.GetUser(&username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch the data"})
		return
	}
	context.JSON(http.StatusOK, user)
}

func (uc *UserController) GetAll(context *gin.Context) {
	users, err := uc.UserService.GetAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch the data"})
		return
	}
	context.JSON(http.StatusOK, users)
}

func (uc *UserController) UpdateUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = uc.UserService.UpdateUser(&user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update the user"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "updated user"})
}
func (uc *UserController) DeleteUser(context *gin.Context) {
	userName := context.Param("name")
	err := uc.UserService.DeleteUser(&userName)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not Delete the data"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "deleted user"})
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userRoute := rg.Group("/user")
	userRoute.POST("/create", uc.CreateUser)
	userRoute.GET("/get/:name", uc.GetUser)
	userRoute.GET("/getall", uc.GetAll)
	userRoute.PATCH("/update", uc.UpdateUser)
	userRoute.DELETE("/delete/:name", uc.DeleteUser)
}
