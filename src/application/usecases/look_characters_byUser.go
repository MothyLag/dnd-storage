package usecases

import (
	"dnd-storage/src/application/ports"
	"dnd-storage/src/domain/entities"
)

type LookCharacterByUser struct{
	charRepo ports.CharacterRepository
}

func NewLookCharacterByUser(charRepo ports.CharacterRepository) *LookCharacterByUser{
	return &LookCharacterByUser{charRepo: charRepo}
}

func (uc *LookCharacterByUser) Execute(userId string) ([]entities.Character,error){
	return uc.charRepo.FindCharactersByUserID(userId)
}
