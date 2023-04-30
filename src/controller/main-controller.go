package controller

import (
	"github.com/nickater/hackrchat-cli/src/ui"
	"github.com/nickater/hackrchat-cli/src/util"
)

func MainController() error {
	controller_running := true
	for controller_running {
		chat := "Chat"
		settings := "Settings"
		exit := "Exit"
		choices := []string{chat, settings, exit}
		main_menu_selection, err := ui.MainMenu(choices)

		if err != nil {
			return err
		}

		switch main_menu_selection {
		case chat:
			ChatController()
		case settings:
			SettingsController()
		case exit:
			controller_running = false
		}
	}
	util.ClearConsole()
	return nil
}
