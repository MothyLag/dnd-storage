package controllers

import (
	"dnd-storage/src/application/usecases"
	"dnd-storage/src/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppController struct{
	CreateApp *usecases.CreateApp
	UpdateApp *usecases.UpdateApp
}

func NewAppController(CreateApp *usecases.CreateApp,UpdateApp *usecases.UpdateApp) *AppController{
	return &AppController{
		CreateApp: CreateApp,
		UpdateApp: UpdateApp,
	}
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

func (ac *AppController) UpdateAppHandler(c *gin.Context){
	var input entities.AppClient
	id := c.Param("id")
	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid Json:"+err.Error()})
		return
	}
	if err := ac.UpdateApp.Execute(input,id); err !=nil{
		c.JSON(http.StatusConflict,gin.H{"error":"Internal Error"+err.Error()})
		return
	}
	c.JSON(http.StatusAccepted,gin.H{
		"message":"App Successfully Updated",
	})
}
