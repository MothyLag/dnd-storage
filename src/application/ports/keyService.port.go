package ports

type KeyService interface{
	GenerateKeyPair()(string,string,error)
	ValidateKeyPair(apiSecretRequest,apiSecretStored string) bool
}
