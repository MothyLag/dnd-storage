package usecases

import (
	"dnd-storage/src/application/ports"
	"dnd-storage/src/domain/entities"
	"fmt"
)

type CreateCharacter struct{
	charRepo ports.CharacterRepository
	charService ports.CharacterService
}

func NewCharacterUseCase(charRepo ports.CharacterRepository, charService ports.CharacterService) *CreateCharacter{
	return &CreateCharacter{
		charRepo: charRepo,
		charService: charService,
	}
}

func (cc *CreateCharacter)Execute(character entities.Character,ownerId string) error{
	if err := cc.charService.ValidateNewCharacter(character);err != nil{
		return fmt.Errorf("Bad Request:%s",err.Error())
	}
	character.OwnerId = ownerId
	return cc.charRepo.Save(character)
}
