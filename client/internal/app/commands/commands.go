package commands

import (
	"fmt"
)

type Command interface {
	Help() string
	Execute(args ...string) error
}

type OpenChatCommand struct{}

func NewOpenChatCommand() *OpenChatCommand {
	return &OpenChatCommand{}
}

func (c *OpenChatCommand) Help() string {
	return "/open <id>, where <id> is int"
}

func (c *OpenChatCommand) Execute(args ...string) error {
	fmt.Printf("Open chat command executed with args: %v\n", args)
	return nil
}
