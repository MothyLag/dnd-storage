package usecases

import (
	"dnd-storage/src/application/ports"
	"dnd-storage/src/domain/entities"
)

type LookCharacter struct{
	charRepo ports.CharacterRepository
}

func NewLookCharacter(charRepo ports.CharacterRepository) *LookCharacter{
	return &LookCharacter{charRepo: charRepo}
}

func (uc *LookCharacter) Execute(id string) (*entities.Character,error){
	return uc.charRepo.FindCharacterById(id)
}
