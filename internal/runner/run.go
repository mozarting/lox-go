package runner

import (
	"fmt"

	"github.com/mozarting/lox/internal/lexer"
)

func Run(input string) {
	l := lexer.New(input)
	tokens := l.ScanTokens()
	for _, t := range tokens {
		fmt.Printf("Type: %s Lexeme: %s Literal: %s Line: %d\n", t.Type, t.Lexeme, t.Literal, t.Line)
	}
}
