package services

import (
	"dnd-storage/src/domain/entities"
	"errors"
)

func ValidateApp(app entities.AppClient) error{
	if !app.IsValid(){
		return errors.New("Invalid App Name")
	}
	return nil
}
