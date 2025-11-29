package usecases

import (
	"dnd-storage/src/application/ports"
	"dnd-storage/src/domain/entities"
)

type UpdateApp struct{
	repo ports.AppRepository
	service ports.AppServices
}

func NewUpdateApp(repo ports.AppRepository, service ports.AppServices) *UpdateApp{
	return &UpdateApp{
		repo: repo,
		service: service,
	}
}

func (ua *UpdateApp)Execute(updatedApp entities.AppClient,id string) error{
	if err := ua.service.ValidateApp(updatedApp);err != nil{
		return err
	}

	return ua.repo.Update(updatedApp,id)
}
