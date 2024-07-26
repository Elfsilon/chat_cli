package models

import (
	"chat_cli/internal/app/gen/chat"
	"fmt"
)

type ChatInfo struct {
	ID    int64
	Title string
}

func (c *ChatInfo) String() string {
	return fmt.Sprintf("id: %03v | name: %v", c.ID, c.Title)
}

type ChatEvent struct {
	UserID    int64
	Author    string
	Text      string
	ColorCode int32
	Type      chat.EventType
}
