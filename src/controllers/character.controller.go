package controllers

import (
	"dnd-storage/src/application/usecases"
	"dnd-storage/src/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Charactercontroller struct{
	createCharacter *usecases.CreateCharacter
	lookCharByUserId *usecases.LookCharacterByUser
	lookChar *usecases.LookCharacter
}

func NewCharacterController(createCharacter *usecases.CreateCharacter,lcbuid *usecases.LookCharacterByUser, lookChar *usecases.LookCharacter) *Charactercontroller{
	return &Charactercontroller{
		createCharacter: createCharacter,
		lookCharByUserId: lcbuid,
		lookChar: lookChar,
	}
}

func (charController *Charactercontroller) PostCharacter(c *gin.Context){
	var input entities.Character
	ownerId := c.GetString("userId")
	if err := c.ShouldBindJSON(&input);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid Json: "+err.Error()})
		return
	}

	if err := charController.createCharacter.Execute(input,ownerId); err !=nil{
		c.JSON(http.StatusConflict, gin.H{"error":"Error creating character: "+err.Error()})
		return
	}

	c.JSON(http.StatusCreated,gin.H{
		"message":"Character successfully created",
	})
}

func (charController *Charactercontroller) GetCharactersByUserId(c *gin.Context){
	userId := c.GetString("userId")
	characters,err := charController.lookCharByUserId.Execute(userId)
	if err != nil{
		c.JSON(http.StatusConflict, gin.H{"error":"Error creating character: "+err.Error()})
		return
	}

	c.JSON(http.StatusOK,characters)
}

func (charController *Charactercontroller) GetCharacterById(c *gin.Context){
	id := c.Param("id")
	character,err := charController.lookChar.Execute(id)
	if err !=nil{
		c.JSON(http.StatusConflict, gin.H{"error":"Error creating character: "+err.Error()})
		return
	}

	c.JSON(http.StatusOK,character)
}
