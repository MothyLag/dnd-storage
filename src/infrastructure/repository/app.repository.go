package repository

import (
	"context"
	"dnd-storage/src/domain/entities"

	"go.mongodb.org/mongo-driver/mongo"
)

type AppMongoRepository struct{
	collection *mongo.Collection
}

func NewAppMongoRepository(db *mongo.Database) *AppMongoRepository{
	return &AppMongoRepository{collection: db.Collection("apps")}
}

func (repo *AppMongoRepository)Save(app entities.AppClient) error{
	_,err := repo.collection.InsertOne(context.Background(),app)
	return err
}
