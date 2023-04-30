package ui

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/nickater/hackrchat-cli/model"
)

func AuthMenu() model.Credential {
	prompt := promptui.Select{
		Label: "Select",
		Items: []string{"Login", "Register", "Exit"},
	}

	index, result, err := prompt.Run()

	fmt.Println(index)
	if err != nil {
		panic(err)
	}

	switch result {
	case "Login":
		return LoginMenu()
	case "Register":
		return RegisterMenu()
	case "Exit":
		panic("Exit")
	default:
		panic("Exit")
	}
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
