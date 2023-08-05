package models

type HuffmanNode struct {
	Char      string
	Frequency int
	Left      *HuffmanNode
	Right     *HuffmanNode
}