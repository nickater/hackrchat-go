package controller

import (
	"fmt"

	"github.com/nickater/hackrchat-cli/service"
	"github.com/nickater/hackrchat-cli/ui"
	"github.com/nickater/hackrchat-cli/util"
)

func AuthController() error {
	credentials := ui.AuthMenu()

	accessToken, err := service.LoginUser(credentials)

	if err != nil {
		return err
	}

	fmt.Println(accessToken)

	util.SaveUsernameToConfigFile(credentials.Username)
	fmt.Println(credentials)
	util.ClearConsole()
	return nil
}
