package repository

import (
	"context"
	"dnd-storage/src/domain/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo *CharacterMongoRepository) Save(character entities.Character) error{
	newDoc,err := Character{}.FromEntity(character)
	if err != nil{
		return err
	}
	_,err = repo.collection.InsertOne(context.Background(),newDoc)
	return err
}

func (repo *CharacterMongoRepository) FindCharactersByUserID(userId string) ([]entities.Character,error){
	ownerId,err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil,err
	}
	filter := bson.M{"ownerId":ownerId}
	cursor,err := repo.collection.Find(context.TODO(),filter)
	if err != nil{
		return nil,err
	}
	defer cursor.Close(context.TODO())

	var docs []Character
	if err = cursor.All(context.TODO(),&docs)	; err !=nil{
		return nil,err
	}

	characters := make([]entities.Character,0,len(docs))
	for _,doc := range docs{
		characters = append(characters, doc.toCharacter())
	}

	return characters,nil
}

func (repo *CharacterMongoRepository) FindCharacterById(id string)(*entities.Character,error){
	var character Character
	userId,err := primitive.ObjectIDFromHex(id)
	if err != nil{
		return nil,err
	}
	filter := bson.M{"_id":userId}
	response := repo.collection.FindOne(context.Background(),filter)
	err = response.Decode(&character)
	if err != nil{
		return nil,err
	}
	entity := character.toCharacter()
	return &entity,nil
}
