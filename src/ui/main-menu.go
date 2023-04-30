package ui

import (
	"errors"

	"github.com/manifoldco/promptui"
)

func MainMenu(choices []string) (string, error) {
	prompt := promptui.Select{
		Label: "Select",
		Items: choices,
	}

	_, result, err := prompt.Run()

	if err != nil {
		panic(err)
	}

	switch result {
	case choices[0]:
		return choices[0], nil
	case choices[1]:
		return choices[1], nil
	case choices[2]:
		return choices[2], nil
	default:
		return "", errors.New("invalid selection")
	}
}
