package strategy

import (
	"errors"
	"strings"
)

type Executor interface {
}

type CommandExecutor struct {
	strategy ExecutorStrategy
}

// func (e *CommandExecutor) ReadLine() string {
// 	text, r := "", bufio.NewReader(os.Stdin)
// 	for {
// 		str, err := r.ReadString('\n')
// 		if err != nil {
// 			fmt.Println(err)
// 			continue
// 		}
// 		text = str
// 		break
// 	}
// 	return strings.ReplaceAll(text, "\n", "")
// }

// func (e *CommandExecutor) WaitForCommand() error {
// 	raw := e.ReadLine()

// 	return cmd.Execute(args...)
// }

func (e *CommandExecutor) Execute(rawCmd string) error {
	chips := strings.Split(rawCmd, " ")

	if len(chips) <= 0 || len(chips[0]) < 2 || chips[0][0] != '/' {
		return errors.New("invalid command format")
	}

	shape, args := chips[0], chips[1:]
	return e.strategy.Execute(shape, args...)
}
