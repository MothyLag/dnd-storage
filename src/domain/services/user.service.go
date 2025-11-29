package services

import (
	"dnd-storage/src/domain/entities"
	"errors"
)

type UserService struct{}

func NewUserService() *UserService{
	return &UserService{}
}

func (s *UserService)ValidateUser(u entities.User) error {
	if !u.IsValid() {
		return errors.New("Invalid User: review information and try again")
	}
	return nil
}
