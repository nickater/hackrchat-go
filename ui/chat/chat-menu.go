package chat_ui

import (
	"errors"

	"github.com/manifoldco/promptui"
)

func ChatMenu(choices []string) (string, error) {
	prompt := promptui.Select{
		Label: "Select",
		Items: choices,
	}

	_, result, err := prompt.Run()

	if err != nil {
		panic(err)
	}

	switch result {
	case "New Chat":
		return "new", nil
	case "Existing Chat":
		return "existing", nil
	case "Exit":
		return "exit", nil
	default:
		return "", errors.New("invalid selection")
	}
}
