package ports

import "dnd-storage/src/domain/entities"

type JWTService interface{
	GenerateToken(user entities.User,jwtSecret string) (string,error)
}
