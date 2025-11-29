package usecases

import (
	ports "dnd-storage/src/application/ports"
	entities "dnd-storage/src/domain/entities"
	"dnd-storage/src/domain/services"
)

type CreateUser struct{
	repo ports.UserRepository
	service ports.UserService
}

func NewCreateUser(repo ports.UserRepository,service ports.UserService) *CreateUser{
	return &CreateUser{
		repo: repo,
		service: service,
	}	
}

func (c *CreateUser) ExecuteCreateUser(u entities.User) error{
	if err := c.service.ValidateUser(u); err != nil{
		return err
	}
	hashed,err := services.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashed
	return c.repo.Save(u)
}
