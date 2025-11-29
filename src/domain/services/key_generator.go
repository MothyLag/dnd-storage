package services

import (
	"crypto/rand"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

type KeyService struct{}

func NewKeyService() *KeyService{
	return &KeyService{}
}

func (s *KeyService)GenerateKeyPair()(string,string,error){
	keyBytes := make([]byte,16)
	secretBytes := make([]byte,32)

	if _,err := rand.Read(keyBytes); err != nil{
		return "","",err
	}

	if _,err := rand.Read(secretBytes); err != nil{
		return "","",err
	}

	apikey := hex.EncodeToString(keyBytes)
	apiSecret := hex.EncodeToString(secretBytes)

	return apikey,apiSecret,nil
}

func (s *KeyService)ValidateKeyPair(apiSecretRequest,apiSecretStored string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(apiSecretStored),[]byte(apiSecretRequest))	
	return err == nil
}
