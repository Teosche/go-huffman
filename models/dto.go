package models

type RequestDTO struct {
	Request string `json:"request"`
}

type ResponseDTO struct {
	CodeMap map[string]string `json:"codeMap"`
}

type ErrorResponseDTO struct {
	Message string `json:"message"`
}
