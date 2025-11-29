package ports

import "dnd-storage/src/domain/entities"

type AppRepository interface{
	Save(app entities.AppClient) error
	FindAppByApiKey(apiKey string) (entities.AppClient,error)
	Update(app entities.AppClient,id string) error
}

type AppServices interface{
	ValidateKeyPair(apiSecretRequest,apiSecretStored string) bool
}
