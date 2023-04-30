package settings_ui

import (
	general_ui "github.com/nickater/hackrchat-cli/src/ui/general"
)

func UpdatePassword() (string, error) {
	return general_ui.GetPassword("Enter new password: ")
}
