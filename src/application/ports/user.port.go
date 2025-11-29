package ports

import entities "dnd-storage/src/domain/entities"

type UserRepository interface{
	Save(user entities.User) error
	GetUserByUserName(userName string) (entities.User,error)
}

type UserService interface{
	ValidateUser(user entities.User) error
	ValidateLoginAttempt(u entities.LoginUser) error
}
