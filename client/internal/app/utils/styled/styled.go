package styled

import (
	"fmt"
	"math/rand"
	"os"
)

type Stylable string
type Style string

const (
	Red     Style = "\033[31m"
	Green   Style = "\033[32m"
	Yellow  Style = "\033[33m"
	Blue    Style = "\033[34m"
	Magenta Style = "\033[35m"
	Cyan    Style = "\033[36m"
	Clear   Style = "\033[0m"
)

var colors = []Style{Red, Green, Yellow, Blue, Magenta, Cyan}

func New(s string) Stylable {
	return Stylable(s)
}

func RndColorCode() int32 {
	return int32(rand.Intn(len(colors)))
}

func ColorFromCode(code int32) Style {
	return colors[code]
}

func logStyled(message string, style Style, a ...any) error {
	message = fmt.Sprintf(message, a...)
	fmt.Print(Stylable(message).WithStyle(style).Build())
	return nil
}

func Infof(message string, a ...any) error {
	return logStyled(message, Blue, a...)
}

func Errorf(message string, a ...any) error {
	return logStyled(message, Red, a...)
}

func Fatalf(message string, a ...any) {
	logStyled(message, Red, a...)
	fmt.Println()
	os.Exit(1)
}

func Successf(message string, a ...any) error {
	return logStyled(message, Green, a...)
}

func (s Stylable) WithStyle(style Style) Stylable {
	return Stylable(fmt.Sprintf("%v%v%v", style, s, Clear))
}

func (s Stylable) Build() string {
	return string(s)
}
