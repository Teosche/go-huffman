package huffman

import (
	"fmt"
	"go-huffman/models"
	"sort"
)

type serviceImpl struct {
}

func NewService() Service {
	return serviceImpl{}
}

func (s serviceImpl) HuffmanCompression(request models.RequestDTO) models.ResponseDTO {
	var huffmanResponse models.ResponseDTO

	huffmanNodes := s.setHuffmanNodes(request)
	huffmanTree := s.buildHuffmanTree(huffmanNodes)

	codeMap := make(map[string]string)
	s.buildHuffmanCode(huffmanTree, "", codeMap)

	huffmanResponse.CodeMap = codeMap

	return huffmanResponse
}

func (s serviceImpl) setHuffmanNodes(request models.RequestDTO) []*models.HuffmanNode {
	var huffmanNodes []*models.HuffmanNode

	frequencies := make(map[string]int)

	for _, char := range request.Request {
		frequencies[string(char)]++
	}

	for char, freq := range frequencies {
		node := &models.HuffmanNode{
			Char:      char,
			Frequency: freq,
		}
		huffmanNodes = append(huffmanNodes, node)
	}

	return huffmanNodes
}

func (s serviceImpl) buildHuffmanTree(huffmanNodes []*models.HuffmanNode) *models.HuffmanNode {
	for len(huffmanNodes) > 1 {
		sort.SliceStable(huffmanNodes, func(i, j int) bool {
			return huffmanNodes[i].Frequency < huffmanNodes[j].Frequency
		})
		lNode := huffmanNodes[0]
		rNode := huffmanNodes[1]

		newNode := &models.HuffmanNode{
			Frequency: lNode.Frequency + rNode.Frequency,
			Left:      lNode,
			Right:     rNode,
		}

		huffmanNodes = huffmanNodes[2:]

		huffmanNodes = append(huffmanNodes, newNode)
	}

	return huffmanNodes[0]
}

func (s serviceImpl) buildHuffmanCode(node *models.HuffmanNode, code string, codeMap map[string]string) {
	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil {
		codeMap[node.Char] = code
		fmt.Println(string(node.Char), code)

	}

	s.buildHuffmanCode(node.Left, code+"0", codeMap)
	s.buildHuffmanCode(node.Right, code+"1", codeMap)
}
