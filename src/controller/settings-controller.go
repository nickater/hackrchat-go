package controller

import (
	"fmt"

	settings_ui "github.com/nickater/hackrchat-cli/src/ui/settings"
	"github.com/nickater/hackrchat-cli/src/util"
)

func SettingsController() error {
	controller_running := true
	for controller_running {
		remember_me := "Remember Me"
		update_username := "Update Username"
		update_password := "Reset Password"
		back := "Back"
		choices := []string{remember_me, update_username, update_password, back}
		main_menu_selection_index, err := settings_ui.SettingsMenu(choices)

		if err != nil {
			panic(err)
		}

		switch main_menu_selection_index {
		case 0:
			err := rememberMeController()

			if err != nil {
				return err
			}
		case 1:

		case 2:
			fmt.Println(update_password)
		case 3:
			controller_running = false
		default:
			controller_running = false
		}
	}
	util.ClearConsole()
	return nil
}

func rememberMeController() error {
	controller_running := true
	for controller_running {
		remember_me_result, err := settings_ui.UpdateRememberMe()

		if err != nil {
			return err
		}
		util.SaveRememberMeToConfigFile(remember_me_result)

		fmt.Println(remember_me_result)
	}

	util.ClearConsole()
	return nil
}

func updateUsernameController() error {
	controller_running := true
	for controller_running {
		update_username_result, err := settings_ui.UpdateUsername()
		if err != nil {
			return err
		}

		fmt.Println(update_username_result)
	}

	util.ClearConsole()
	return nil
}

func updatePasswordController() error {
	controller_running := true
	for controller_running {
		update_password_result, err := settings_ui.UpdatePassword()
		if err != nil {
			return err
		}

		fmt.Println(update_password_result)
	}
	util.ClearConsole()
	return nil
}
