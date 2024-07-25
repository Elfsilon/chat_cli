package models

import "fmt"

type ChatInfo struct {
	ID    int64
	Title string
}

func (c *ChatInfo) String() string {
	return fmt.Sprintf("% 2v: %v", c.ID, c.Title)
}

type ChatMessage struct {
	Author string
	Text   string
}
