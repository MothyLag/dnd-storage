package services

import (
	"dnd-storage/src/domain/entities"
	"errors"
)

func ValidateUser(u entities.User) error {
	if !u.IsValid() {
		return errors.New("Invalid User: review information and try again")
	}
	return nil
}
