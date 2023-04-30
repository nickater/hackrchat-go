package controller

import (
	"fmt"

	settings_ui "github.com/nickater/hackrchat-cli/ui/settings"
	"github.com/nickater/hackrchat-cli/util"
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
			fmt.Println(remember_me)
		case 1:
			fmt.Println(update_username)
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
