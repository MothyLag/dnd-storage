package services

import (
	"dnd-storage/src/domain/entities"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct{
	
}

func NewJWTService() *JWTService{
	return &JWTService{}
}

func (s *JWTService) GenerateToken(user entities.User,jwtSecret string) (string,error){
	claims := jwt.MapClaims{
		"id": user.ID,
		"username":user.UserName,
		"role":user.Role,
		"exp":time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenStr,err := token.SignedString([]byte(jwtSecret))
	if err != nil{
		return "",fmt.Errorf("Error Signing token: %s",err)
	}
	return tokenStr,nil
}
