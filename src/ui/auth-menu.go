package ui

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/nickater/hackrchat-cli/src/model"
)

func AuthMenu(choices []string) int {
	prompt := promptui.Select{
		Label: "Select",
		Items: choices,
	}

	index, _, err := prompt.Run()

	fmt.Println(index)
	if err != nil {
		panic(err)
	}

	return index
}

func RegisterMenu() model.Credential {
	prompt := promptui.Prompt{
		Label: "Username",
	}
	username, err := prompt.Run()
	if err != nil {
		panic(err)
	}

	prompt = promptui.Prompt{
		Label: "Password",
		Mask:  '*',
	}

	password, err := prompt.Run()

	if err != nil {
		panic(err)
	}

	return model.Credential{Username: username, Password: password}
}

func LoginMenu() model.Credential {
	prompt := promptui.Prompt{
		Label: "Username",
	}
	username, err := prompt.Run()
	if err != nil {
		panic(err)
	}

	prompt = promptui.Prompt{
		Label: "Password",
		Mask:  '*',
	}

	password, err := prompt.Run()

	if err != nil {
		panic(err)
	}

	return model.Credential{Username: username, Password: password}
}
