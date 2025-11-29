package usecases

import (
	"dnd-storage/src/application/ports"
	"dnd-storage/src/domain/entities"
	"dnd-storage/src/domain/services"
)

type CreateApp struct{
	repo ports.AppRepository
	appService ports.AppServices
	keyService ports.KeyService
}		

func NewCreateApp(repo ports.AppRepository,appService ports.AppServices,keyService ports.KeyService) *CreateApp{
	return &CreateApp{
		repo: repo,
		appService: appService,
		keyService: keyService,
	}
}

func (c *CreateApp) Execute(app entities.AppClient) (string,string,error){
	if err := c.appService.ValidateApp(app); err != nil{
		return "","",err	
	}
	apiKey,apiSecret,err := c.keyService.GenerateKeyPair()
	hashed,err := services.HashPassword(apiSecret)
	if err != nil {
		return "","",err
	}
	app.Apikey = apiKey
	app.AppSecret = hashed
	return apiKey,apiSecret,c.repo.Save(app)
}
