package main

import (
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/msg-gateway/controller"
)

func init() {
	util.InitLog("msg_gateway.log")
}

func main() {
	err := controller.Controller()
	if err != nil {
		panic(err)
	}
}
