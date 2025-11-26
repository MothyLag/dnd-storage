package usecases

import (
	"dnd-storage/src/application/ports"
	"dnd-storage/src/domain/entities"
	"dnd-storage/src/domain/services"
)

type CreateApp struct{
	repo ports.AppRepository
}		

func NewCreateApp(repo ports.AppRepository) *CreateApp{
	return &CreateApp{repo: repo}
}

func (c *CreateApp) Execute(app entities.AppClient) (string,string,error){
	if err := services.ValidateApp(app); err != nil{
		return "","",err	
	}
	apiKey,apiSecret,err := services.GenerateKeyPair()
	hashed,err := services.HashPassword(apiSecret)
	if err != nil {
		return "","",err
	}
	app.Apikey = apiKey
	app.AppSecret = hashed
	return apiKey,apiSecret,c.repo.Save(app)
}
