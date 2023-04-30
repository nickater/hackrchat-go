package settings_ui

import (
	general_ui "github.com/nickater/hackrchat-cli/src/ui/general"
)

func UpdateUsername() (string, error) {
	return general_ui.GetInput("Enter new username: ")
}
