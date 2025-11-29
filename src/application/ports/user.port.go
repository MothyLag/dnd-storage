package ports

import entities "dnd-storage/src/domain/entities"

type UserRepository interface{
	Save(user entities.User) error
}

type UserService interface{
	ValidateUser(user entities.User) error
}
