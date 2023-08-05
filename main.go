package main

import (
	_ "go-huffman/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}
