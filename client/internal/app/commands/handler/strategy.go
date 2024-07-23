package strategy

import (
	"chat_cli/internal/app/commands"
	"chat_cli/internal/app/components"
	"chat_cli/internal/app/models"
	"chat_cli/internal/app/screens"
	"fmt"
)

type ExecutorStrategy interface {
	Init()
	Execute(shape string, args ...string) error
	Render() components.Component
}

type BaseStrategy struct {
	registry Registry
}

func NewBaseStrategy() *BaseStrategy {
	return &BaseStrategy{
		registry: NewCommandRegistry(),
	}
}

func (s *BaseStrategy) Execute(shape string, args ...string) error {
	cmd, ok := s.registry.Get(shape)
	if !ok {
		return fmt.Errorf("command is not registered: %v", shape)
	}
	return cmd.Execute(args...)
}

type ChatsExecutorStrategy struct {
	screen screens.ChatScreen
	chats  []models.ChatInfo
	*BaseStrategy
}

func NewChatsExecutorStrategy() *ChatsExecutorStrategy {
	return &ChatsExecutorStrategy{
		chats:        make([]models.ChatInfo, 0),
		BaseStrategy: NewBaseStrategy(),
	}
}

func (s *ChatsExecutorStrategy) Init() {
	s.registry.Register("/open", commands.NewOpenChatCommand())
	s.registry.Register("/join", commands.NewOpenChatCommand())
	s.registry.Register("/delete", commands.NewOpenChatCommand())

	// TODO: fetch chats
	s.screen.Chats = []models.ChatInfo{
		{ID: 128, Title: "ANiMe"},
		{ID: 125, Title: "ANiMe 2"},
		{ID: 129, Title: "ANiMe 3"},
	}
}

func (s *ChatsExecutorStrategy) Render() components.Component {
	return &screens.ChatScreen{
		Chats: s.chats,
	}
}
