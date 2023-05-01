package chat_ui

import (
	"github.com/manifoldco/promptui"
	"github.com/nickater/hackrchat-cli/src/util"
)

func NewChatMenu() (string, error) {
	prompt := promptui.Prompt{
		Label: "Enter Recipient's Username",
	}

	result, err := prompt.Run()

	if err != nil {
		panic(err)
	}

	util.ClearConsole()

	return result, nil
}
