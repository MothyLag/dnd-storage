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
		"exp":time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenStr,err := token.SignedString([]byte(jwtSecret))
	if err != nil{
		return "",fmt.Errorf("Error Signing token: %s",err)
	}
	return tokenStr,nil
}

func (s *JWTService) ValidateToken(tokenStr,jwtSecret string) (jwt.MapClaims,error){
	token,err := jwt.Parse(tokenStr,func(t *jwt.Token)(any,error){
		if _,ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil,fmt.Errorf("Unexpected signing method")
		}
		return []byte(jwtSecret),nil
	})
	if err != nil{
		return nil,err
	}

	if claims,ok :=token.Claims.(jwt.MapClaims); ok && token.Valid{
		return claims,nil
	}

	return nil, fmt.Errorf("Invalid token")
}
