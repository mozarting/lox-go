package file

import (
	"fmt"
	"os"

	"github.com/mozarting/lox/internal/lexer"
)

func ReadFile(f string) {
	file_content, err := os.ReadFile(f)
	if err != nil {
		fmt.Println("ERROR: Failed to read file, ", err)
		return
	}
	fmt.Println(string(file_content))
	lexer.Run(string(file_content))
}
