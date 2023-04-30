package controller

import (
	"fmt"

	"github.com/nickater/hackrchat-cli/src/service"
	"github.com/nickater/hackrchat-cli/src/ui"
	"github.com/nickater/hackrchat-cli/src/util"
)

func AuthController() error {
	choices := []string{"Login", "Register", "Exit"}
	choices_index := ui.AuthMenu(choices)

	switch choices_index {
	case 0:
		credentials := ui.LoginMenu()

		accessToken, err := service.LoginUser(credentials)

		if err != nil {
			return err
		}

		util.SaveUsernameToConfigFile(credentials.Username)
		fmt.Println(accessToken)
	case 1:
		credentials := ui.RegisterMenu()

		access_token, err := service.RegisterUser(credentials)

		if err != nil {
			return err
		}

		util.SaveUsernameToConfigFile(credentials.Username)
		fmt.Println(access_token)
	case 2:
		return nil
	default:
		return nil
	}

	util.ClearConsole()
	return nil
}
