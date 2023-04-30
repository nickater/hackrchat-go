package chat_ui

import (
	"github.com/manifoldco/promptui"
)

func NewChatMenu() (string, error) {
	prompt := promptui.Prompt{
		Label: "Enter Recipient's Username",
	}

	result, err := prompt.Run()

	if err != nil {
		panic(err)
	}

	return result, nil

}
