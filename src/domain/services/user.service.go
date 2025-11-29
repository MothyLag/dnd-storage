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

func (s *UserService)ValidateLoginAttempt(lu entities.LoginUser) error {
	if !lu.IsValidAttempt(){
		return errors.New("Invalid Attempt please fill all the fields")
	}
	return nil
}
