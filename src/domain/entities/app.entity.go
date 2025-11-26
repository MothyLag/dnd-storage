package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type AppClient struct{
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	AppName string `bson:"appName" json:"app_name"`
	Apikey string `bson:"apiKey" json:"api_key"`
	AppSecret string `bson:"appSecret" json:"app_secret"`
}

func (app *AppClient) IsValid() bool{
	return app.AppName != "" && len(app.AppName) > 3 
}
