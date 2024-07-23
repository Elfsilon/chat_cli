package screens

import (
	"chat_cli/internal/app/components"
	"chat_cli/internal/app/models"
)

type ChatScreen struct {
	Chats []models.ChatInfo
}

func (s *ChatScreen) Render() {
	view := &components.TitledComponent{
		Title: "Chats",
		Child: &components.ChatListComponent{
			Chats: s.Chats,
		},
	}
	view.Render()
}
