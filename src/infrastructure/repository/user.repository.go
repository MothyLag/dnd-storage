package repository

import (
	"context"
	"dnd-storage/src/domain/entities"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)
type UserDocument struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	UserName string `bson:"userName,omitempty"`
	Password string `bson:"password,omitempty"`
	Role string `bson:"role,omitempty"`
}

type UserMongoRepository struct {
	collection *mongo.Collection
}

func NewUserMongoRepository(db *mongo.Database) *UserMongoRepository{
	return &UserMongoRepository{collection:db.Collection("users")}
}

func (r *UserMongoRepository) Save(u entities.User) error{
	newDoc,err := usertoDocument(u)
	if err != nil{
		return fmt.Errorf("Error Saving User:%s",err.Error())
	}
	_,err = r.collection.InsertOne(context.Background(), newDoc)
	return err
}

func usertoDocument(user entities.User) (UserDocument,error){
	var oid primitive.ObjectID
	var err error

	if user.ID != ""{
		oid,err = primitive.ObjectIDFromHex(user.ID)	
		if err != nil{
			return UserDocument{},err
		}
	}

	return UserDocument{
		ID: oid,
		UserName: user.UserName,
		Password: user.Password,
	},nil
}

func userToDomain(doc UserDocument) entities.User{
	return entities.User{
		ID: doc.ID.Hex(),
		UserName: doc.UserName,
		Password: doc.Password,
	}
}
