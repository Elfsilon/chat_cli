package strategy

import "chat_cli/internal/app/commands"

type Registry interface {
	Get(shape string) (commands.Command, bool)
	Register(shape string, command commands.Command)
}

type CommandRegistry struct {
	cmd map[string]commands.Command
}

func NewCommandRegistry() *CommandRegistry {
	return &CommandRegistry{
		cmd: make(map[string]commands.Command, 0),
	}
}

func (r *CommandRegistry) Get(shape string) (commands.Command, bool) {
	cmd, ok := r.cmd[shape]
	return cmd, ok
}

func (r *CommandRegistry) Register(shape string, command commands.Command) {
	r.cmd[shape] = command
}
