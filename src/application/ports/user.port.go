package ports

import entities "dnd-storage/src/domain/entities"

type UserRepository interface{
	Save(user entities.User) error
}
