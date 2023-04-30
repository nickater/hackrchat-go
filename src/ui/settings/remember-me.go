package settings_ui

import (
	general_ui "github.com/nickater/hackrchat-cli/src/ui/general"
)

func UpdateRememberMe() (bool, error) {
	return general_ui.GetConfirmation("Remember me? (y/n): ")
}
