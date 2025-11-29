package entities

type AppClient struct{
	ID string `json:"id,omitempty"`
	AppName string `json:"app_name"`
	Apikey string `json:"api_key"`
	AppSecret string `json:"app_secret"`
	Rol string `json:"rol"`
}

func (app *AppClient) IsValid() bool{
	return app.AppName != "" && len(app.AppName) > 3 
}
