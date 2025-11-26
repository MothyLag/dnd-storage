package repository

import (
	"context"
	"dnd-storage/src/domain/entities"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongoRepository struct {
	collection *mongo.Collection
}

func NewUserMongoRepository(db *mongo.Database) *UserMongoRepository{
	return &UserMongoRepository{collection:db.Collection("users")}
}

func (r *UserMongoRepository) Save(u entities.User) error{
	_,err := r.collection.InsertOne(context.Background(), u)
	return err
}
