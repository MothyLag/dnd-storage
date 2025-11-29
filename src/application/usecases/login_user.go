package usecases

import (
	"dnd-storage/src/application/ports"
	"dnd-storage/src/domain/entities"
	"errors"
)

type LoginUser struct{
	jwtSecret string
	userRepo ports.UserRepository
	userService ports.UserService
	keyService ports.KeyService
	jwtService ports.JWTService
}

func NewLoginUser(jwtSecret string,ur ports.UserRepository,us ports.UserService,ks ports.KeyService, js ports.JWTService) *LoginUser{
	return &LoginUser{
		jwtSecret: jwtSecret,
		userRepo: ur,
		userService: us,
		keyService: ks,
		jwtService: js,
	}
}

func (lu *LoginUser) Execute(loginAttempt entities.LoginUser) (string,error){
	if err := lu.userService.ValidateLoginAttempt(loginAttempt); err != nil{
		return "",err
	}
	user,err := lu.userRepo.GetUserByUserName(loginAttempt.UserName)	
	if err != nil{
		return "",errors.New("Invalid Credentials")
	}
	if !lu.keyService.ValidateKeyPair(loginAttempt.Password,user.Password){
		return "",errors.New("Invalid Credentials")	
	}
	return lu.jwtService.GenerateToken(user,lu.jwtSecret)
}

