package controller

import (
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
		_, err := service.LoginUser(credentials)

		if err != nil {
			return err
		}

		util.SaveUsernameToConfigFile(credentials.Username)
		util.DisplayMessageAndWait("Login successful!", 2)
	case 1:
		credentials := ui.RegisterMenu()

		_, err := service.RegisterUser(credentials)

		if err != nil {
			return err
		}

		util.SaveUsernameToConfigFile(credentials.Username)
		util.DisplayMessageAndWait("Registration successful!", 2)
	case 2:
		return nil
	default:
		return nil
	}

	return nil
}
