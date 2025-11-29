package services

import (
	"dnd-storage/src/domain/entities"
	"errors"
)

type AppService struct{}

func NewAppService() *AppService{
	return &AppService{}
}

func (s *AppService)ValidateApp(app entities.AppClient) error{
	if !app.IsValid(){
		return errors.New("Invalid App Name")
	}
	return nil
}
