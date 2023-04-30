package chat_ui

import (
	"github.com/manifoldco/promptui"
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

	return result, nil

}
