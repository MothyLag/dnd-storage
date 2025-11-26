package ports

import "dnd-storage/src/domain/entities"

type AppRepository interface{
	Save(app entities.AppClient) error
}
