package repository

import (
	"context"
	"dnd-storage/src/domain/entities"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AppClientDocument struct{
	ID primitive.ObjectID `bson:"_id,omitempty"`
	AppName string `bson:"appName,omitempty"`
	Apikey string `bson:"apiKey,omitempty"`
	AppSecret string `bson:"appSecret,omitempty"`
	Rol string `bson:"rol,omitempty"`
}

type AppMongoRepository struct{
	collection *mongo.Collection
}

func NewAppMongoRepository(db *mongo.Database) *AppMongoRepository{
	return &AppMongoRepository{collection: db.Collection("apps")}
}

func (repo *AppMongoRepository)Save(app entities.AppClient) error{
	doc,err := appToDocument(app)
	if err != nil{
		return fmt.Errorf("Error Creating document: %s",err.Error())
	}
	_,err = repo.collection.InsertOne(context.Background(),doc)
	return err
}

func (repo *AppMongoRepository)FindAppByApiKey(apiKey string) (entities.AppClient,error){
	var app AppClientDocument 
	result := repo.collection.FindOne(context.Background(),bson.M{"apiKey":apiKey})
	err := result.Decode(&app)
	if err != nil{
		if errors.Is(err,mongo.ErrNoDocuments){
			return entities.AppClient{},fmt.Errorf("App Not Found: %s",apiKey)
		}
		return entities.AppClient{},fmt.Errorf("Error Decoding Api Key: %s",err.Error())
	}
	return appToDomain(app),nil
}

func (repo *AppMongoRepository) Update(app entities.AppClient,id string) error{

	oid,err := primitive.ObjectIDFromHex(id)
	if err != nil{
		return fmt.Errorf("invalid id: %w",err)
	}
	doc,err := appToDocument(app)
	if err != nil{
		return fmt.Errorf("Error Updating App:%s",err.Error())
	}
	// Marshal to bytes[] document
	docMapBytes,err := bson.Marshal(doc)
	if err != nil{
		return fmt.Errorf("Error Updating App:%s",err.Error())
	}
	//UnMarshal bytes[] into dynamic Set object
	var docMap bson.M
	delete(docMap,"_id")
	if err = bson.Unmarshal(docMapBytes,&docMap); err != nil{
		return fmt.Errorf("Error Updating App:%s",err.Error())
	}

	filter := bson.M{"_id":oid}
	update := bson.M{"$set":docMap}

	_,err = repo.collection.UpdateOne(context.Background(),filter,update)
	if err != nil{
		return fmt.Errorf("Error updating App:%s",err.Error())
	}
	return nil
} 

func appToDomain(doc AppClientDocument) entities.AppClient{
	return entities.AppClient{
		ID: doc.ID.Hex(),
		Apikey: doc.Apikey,
		AppName: doc.AppName,
		Rol: doc.Rol,
	}
}

func appToDocument(app entities.AppClient)(AppClientDocument,error){
	var oid primitive.ObjectID
	var err error

	if app.ID != ""{
		oid,err = primitive.ObjectIDFromHex(app.ID)	
		if err != nil{
			return AppClientDocument{},err
		}
	}

	return AppClientDocument{
		ID: oid,
		Apikey: app.Apikey,
		AppName: app.AppName,
		AppSecret: app.AppSecret,
		Rol: app.Rol,
	},nil
}
