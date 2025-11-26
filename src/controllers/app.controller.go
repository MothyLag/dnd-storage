package controllers

import (
	"dnd-storage/src/application/usecases"
	"dnd-storage/src/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppController struct{
	CreateApp *usecases.CreateApp
}

func NewAppController(CreateApp *usecases.CreateApp) *AppController{
	return &AppController{CreateApp: CreateApp}
}

func (ac *AppController) CreateAppHandler(c *gin.Context){
	var input entities.AppClient
	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid Json"})
		return
	}
	apiKey,apiSecret,err := ac.CreateApp.Execute(input)
	if err != nil{
		c.JSON(http.StatusConflict,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"message":"App successfully created",
		"api_name":input.AppName,
		"api_key":apiKey,
		"api_secret":apiSecret,
	})
}
