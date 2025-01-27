package repl

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mozarting/lox/internal/runner"
)

func RunPrompt() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		runner.Run(line)
	}
}
