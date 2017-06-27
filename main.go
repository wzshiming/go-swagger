package main

import (
	"github.com/wzshiming/go-swagger/swagger"
	"github.com/wzshiming/go-swagger/swaggergen"
	"github.com/wzshiming/go-swagger/utils"
)

func main() {

	rootapi := &swagger.Swagger{}
	//	err := generate.GenerateHead(rootapi, `C:\gopath\src\wjs_api`)
	//	if err != nil {
	//		ffmt.Mark(err)
	//		return
	//	}
	swaggergen.GB(rootapi, "wjs_api/routers")
	//	err = generate.GenerateBody(rootapi, `wjs_api/controllers`)
	//	if err != nil {
	//		ffmt.Mark(err)
	//		return
	//	}
	utils.WriteFile(rootapi, ".")
}
