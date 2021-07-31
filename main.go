package main

import (
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/wx-gateway/controller"
)

func main() {
	util.InitLog("msg_gateway.log")
	err := controller.Controller()
	if err != nil {
		panic(err)
	}
}
