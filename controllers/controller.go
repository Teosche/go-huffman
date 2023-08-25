package controllers

import (
	"encoding/json"
	"fmt"
	"go-huffman/huffman"
	"go-huffman/models"
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
)

type HuffmanController struct {
	beego.Controller
	HuffmanService huffman.Service
}

func NewHuffmanController() *HuffmanController {
	return &HuffmanController{
		HuffmanService: huffman.NewService(),
	}
}

// HuffmanCompression compresses the input string using Huffman coding
// @router /huffmanCompression [post]
// @Param request body models.RequestDTO true "The string to compress"
// @Success 200 {object} models.ResponseDTO
// @Failure 400 {object} ErrorResponse
func (c *HuffmanController) HuffmanCompression() {
	var request models.RequestDTO

	decoder := json.NewDecoder(c.Ctx.Request.Body)
	if err := decoder.Decode(&request); err != nil {
		c.handleErrorResponse(http.StatusBadRequest, err.Error())
		return
	}

	response := c.HuffmanService.HuffmanCompression(request)

	c.Data["json"] = response

	if err := c.ServeJSON(); err != nil {
		c.handleErrorResponse(http.StatusInternalServerError, err.Error())
		return
	}
}

func (c *HuffmanController) handleErrorResponse(status int, message string) {
	c.Ctx.Output.SetStatus(status)
	c.Data["json"] = models.ErrorResponseDTO{Message: message}
	err := c.ServeJSON()

	if err != nil {
		fmt.Println("An error occurred:", err)
		return
	}
}
