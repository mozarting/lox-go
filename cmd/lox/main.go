package main

import (
	"fmt"
	"os"

	"github.com/mozarting/lox/internal/file"
	"github.com/mozarting/lox/internal/repl"
)

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		fmt.Println("Usage: lox [script]")
	} else if len(args) == 1 {
		file.ReadFile(args[0])
	} else {
		repl.RunPrompt()
	}
}
