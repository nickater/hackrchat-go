package main

import (
	"github.com/nickater/hackrchat-cli/controller"
	"github.com/nickater/hackrchat-cli/util"
)

func main() {
	util.ClearConsole()
	authResult := controller.AuthController()

	if authResult != nil {
		panic(authResult)
	}

	controller.MainController()
}
