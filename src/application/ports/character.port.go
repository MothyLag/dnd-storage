package ports

import "dnd-storage/src/domain/entities"

type CharacterRepository interface{
	Save(character entities.Character) error
	FindCharactersByUserID(userId string) ([]entities.Character,error)
	FindCharacterById(id string)(*entities.Character,error)
}

type CharacterService interface{
	ValidateNewCharacter(c entities.Character) error
}
