package services

import (
	"dnd-storage/src/domain/entities"
	"errors"
)

type CharacterService struct{}

func NewCharacterService() *CharacterService{
	return &CharacterService{}
}

func (CharacterService)ValidateNewCharacter(c entities.Character) error{
	if c.Name == ""{
		return errors.New("Name is Mandatory")
	}
	return nil
}
