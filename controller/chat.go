package controller

import (
	"fmt"

	chat_ui "github.com/nickater/hackrchat-cli/ui/chat"
	"github.com/nickater/hackrchat-cli/util"
)

func ChatController() error {
	controller_running := true
	for controller_running {
		new_chat := "New Chat"
		existing_chat := "Existing Chat"
		back := "Back"
		choices := []string{new_chat, existing_chat, back}
		result, err := chat_ui.ChatMenu(choices)

		if err != nil {
			return err
		}
		switch result {
		case new_chat:
			new_chat_result, err := chat_ui.NewChatMenu()
			if err != nil {
				return err
			}

			fmt.Println(new_chat_result)
			// TODO: Search for user
			// TODO: Create new chat
		case existing_chat:
			// TODO: Search for chat
			chat_names := []string{"the boyz", "stacy", "gus20k"}

			existing_chat_result, err := chat_ui.ExistingChatMenu(chat_names)
			if err != nil {
				return err
			}
			fmt.Println(existing_chat_result)
			// TODO: Load chat
		case back:
			controller_running = false
		}
	}
	util.ClearConsole()
	return nil
}
