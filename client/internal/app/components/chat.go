package components

import (
	"chat_cli/internal/app/models"
	"fmt"
)

type TitledComponent struct {
	Title string
	Child Component
}

func (c *TitledComponent) Render() {
	fmt.Printf("%v:\n", c.Title)
	c.Child.Render()
}

type ChatComponent struct {
	ID     int
	Title  string
	Unread int
}

func (c *ChatComponent) Render() {
	uw := ""
	if c.Unread > 0 {
		uw = fmt.Sprintf("(+%v new)", c.Unread)
	}
	fmt.Printf("% 2v: %v %v\n", c.ID, c.Title, uw)
}

type ChatListComponent struct {
	Chats []models.ChatInfo
}

func (c *ChatListComponent) Render() {
	items := make([]Component, len(c.Chats))
	for i, chat := range c.Chats {
		item := &ChatComponent{
			ID:     chat.ID,
			Title:  chat.Title,
			Unread: 0,
		}
		items[i] = item
	}
	list := &ListView{children: items}
	list.Render()
}
