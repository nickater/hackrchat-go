package settings_ui

import (
	"github.com/manifoldco/promptui"
)

func SettingsMenu(choices []string) (int, error) {
	prompt := promptui.Select{
		Label: "Select",
		Items: choices,
	}

	index, _, err := prompt.Run()

	if err != nil {
		panic(err)
	}

	return index, nil
}
