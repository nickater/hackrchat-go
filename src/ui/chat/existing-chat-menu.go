package chat_ui

import (
	"github.com/manifoldco/promptui"
	"github.com/nickater/hackrchat-cli/src/util"
)

func ExistingChatMenu(chat_names []string) (string, error) {

	prompt := promptui.Select{
		Label: "Pick from existing chats",
		Items: chat_names,
	}

	_, result, err := prompt.Run()

	if err != nil {
		panic(err)
	}

	util.ClearConsole()
	return result, nil

}
