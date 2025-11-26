package services

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateKeyPair()(string,string,error){
	keyBytes := make([]byte,16)
	secretBytes := make([]byte,32)

	if _,err := rand.Read(keyBytes); err != nil{
		return "","",nil
	}

	if _,err := rand.Read(secretBytes); err != nil{
		return "","",nil
	}

	apikey := hex.EncodeToString(keyBytes)
	apiSecret := hex.EncodeToString(secretBytes)

	return apikey,apiSecret,nil
}
