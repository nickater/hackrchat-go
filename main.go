package main

import (
	"github.com/nickater/hackrchat-cli/src/controller"
	"github.com/nickater/hackrchat-cli/src/util"
)

func main() {
	util.ClearConsole()
	authResult := controller.AuthController()

	if authResult != nil {
		panic(authResult)
	}

	controller.MainController()
}
