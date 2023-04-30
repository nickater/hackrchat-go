package general_ui

import (
	"github.com/manifoldco/promptui"
)

func GetInput(label string) (string, error) {
	prompt := promptui.Prompt{
		Label: label,
	}

	result, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return result, nil
}

func GetPassword(label string) (string, error) {
	prompt := promptui.Prompt{
		Label: label,
		Mask:  '*',
	}

	result, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return result, nil
}

func GetConfirmation(label string) (bool, error) {
	prompt := promptui.Prompt{
		Label:     label,
		IsConfirm: true,
	}

	result, err := prompt.Run()

	if err != nil {
		return false, err
	}

	if result == "y" {
		return true, nil
	}

	return false, nil
}

func GetSelection(label string, choices []string) (int, error) {
	prompt := promptui.Select{
		Label: label,
		Items: choices,
	}

	index, _, err := prompt.Run()

	if err != nil {
		return -1, err
	}

	return index, nil
}
