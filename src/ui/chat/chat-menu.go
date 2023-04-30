package chat_ui

import (
	"github.com/manifoldco/promptui"
)

func ChatMenu(choices []string) int {
	prompt := promptui.Select{
		Label: "Select",
		Items: choices,
	}

	index, _, err := prompt.Run()

	if err != nil {
		panic(err)
	}

	return index
}
