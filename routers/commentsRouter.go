package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["go-huffman/controllers:HuffmanController"] = append(beego.GlobalControllerRouter["go-huffman/controllers:HuffmanController"],
        beego.ControllerComments{
            Method: "HuffmanCompression",
            Router: `/huffmanCompression`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
