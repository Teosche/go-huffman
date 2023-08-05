package routers

import (
	"go-huffman/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	huffmanController := controllers.NewHuffmanController()

	nsHuffman := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/huffman",
			beego.NSInclude(
				huffmanController,
			),
		),
	)
	beego.AddNamespace(nsHuffman)
}
