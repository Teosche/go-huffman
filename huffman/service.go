package huffman

import "go-huffman/models"

type Service interface {
	HuffmanCompression(request models.RequestDTO) models.ResponseDTO
}
