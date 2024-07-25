package console

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadLine() string {
	text, r := "", bufio.NewReader(os.Stdin)

	for {
		r.Buffered()

		str, err := r.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			continue
		}
		text = str
		break
	}
	return strings.ReplaceAll(text, "\n", "")
}
