package usecases

import (
	"dnd-storage/src/application/ports"
	"dnd-storage/src/domain/entities"
	"dnd-storage/src/domain/services"
)

type UpdateApp struct{
	repo ports.AppRepository
}

func NewUpdateApp(repo ports.AppRepository) *UpdateApp{
	return &UpdateApp{repo: repo}
}

func (ua *UpdateApp)Execute(updatedApp entities.AppClient,id string) error{
	if err := services.ValidateApp(updatedApp);err != nil{
		return err
	}

	return ua.repo.Update(updatedApp,id)
}
