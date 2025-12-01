package entities

type FileUpload struct{
	FileName string `json:"fileName"`
	FileData string `json:"fileData"`
}
