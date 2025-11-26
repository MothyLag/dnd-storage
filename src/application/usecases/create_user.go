package usecases

import (
	ports "dnd-storage/src/application/ports"
	entities "dnd-storage/src/domain/entities"
	"dnd-storage/src/domain/services"
)

type CreateUser struct{
	repo ports.UserRepository
}

func NewCreateUser(repo ports.UserRepository) *CreateUser{
	return &CreateUser{repo: repo}	
}

func (c *CreateUser) ExecuteCreateUser(u entities.User) error{
	if err := services.ValidateUser(u); err != nil{
		return err
	}
	hashed,err := services.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashed
	return c.repo.Save(u)
}
