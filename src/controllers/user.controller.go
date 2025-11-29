package controllers

import (
	"dnd-storage/src/application/usecases"
	"dnd-storage/src/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{
	createUser *usecases.CreateUser
	loginUser *usecases.LoginUser
}

func NewUserController(createUser *usecases.CreateUser,loginUser *usecases.LoginUser) *UserController{
	return &UserController{createUser: createUser,loginUser: loginUser}
}

func (uc *UserController) CreateUserHandler(c *gin.Context){
	var input entities.User
	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"invalid Json"})
		return
	}

	if err := uc.createUser.ExecuteCreateUser(input);err != nil{
		c.JSON(http.StatusConflict,gin.H{"error":err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{
		"message":"User successfully created",
		"user": input,
	})
}

func (uc *UserController) LoginUserHandler(c *gin.Context){
	var input entities.LoginUser
	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid request"})
		return
	}
	token,err := uc.loginUser.Execute(input)
	if err != nil{
		c.JSON(http.StatusUnauthorized,gin.H{"error":"Invalid credentials"})
		return
	}

	c.JSON(http.StatusAccepted,gin.H{
		"message":"Successfully authenticated",
		"token":token,
	})
}
