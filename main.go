package main

import (
	"github.com/cellargalaxy/wx-gateway/controller"
)

func main() {
	err := controller.Controller()
	if err != nil {
		panic(err)
	}
}
