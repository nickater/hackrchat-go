package controller

import (
	"fmt"

	chat_ui "github.com/nickater/hackrchat-cli/src/ui/chat"
)

func ChatController() error {
	controller_running := true
	for controller_running {
		new_chat := "New Chat"
		existing_chat := "Existing Chat"
		back := "Back"
		choices := []string{new_chat, existing_chat, back}
		choice_index := chat_ui.ChatMenu(choices)

		switch choice_index {
		case 0:
			err := newChatController()

			if err != nil {
				return err
			}

			// TODO: Search for user
			// TODO: Create new chat
		case 1:

			// TODO: Search for chat
			err := existingChatController()

			if err != nil {
				return err
			}

			// TODO: Load chat
		case 2:
			controller_running = false
		default:
			controller_running = false
		}

	}
	return nil
}

func newChatController() error {
	controller_running := true
	for controller_running {
		new_chat_result, err := chat_ui.NewChatMenu()
		if err != nil {
			return err
		}

		fmt.Println(new_chat_result)
	}
	return nil
}

func existingChatController() error {
	controller_running := true
	for controller_running {
		chat_names := append([]string{"the boyz", "stacy", "gus20k"}, "Back")

		existing_chat_result, err := chat_ui.ExistingChatMenu(chat_names)
		if err != nil {
			return err
		}

		fmt.Println(existing_chat_result)
	}
	return nil
}
