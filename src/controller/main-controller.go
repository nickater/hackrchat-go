package controller

import (
	"github.com/nickater/hackrchat-cli/src/ui"
)

func MainController() error {
	controller_running := true
	for controller_running {
		chat := "Chat"
		settings := "Settings"
		exit := "Exit"
		choices := []string{chat, settings, exit}
		main_menu_selection_index := ui.MainMenu(choices)

		switch main_menu_selection_index {
		case 0:
			ChatController()
		case 1:
			SettingsController()
		case 2:
			controller_running = false
		default:
			controller_running = false
		}
	}
	return nil
}
