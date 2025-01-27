package file

import (
	"fmt"
	"os"

	"github.com/mozarting/lox/internal/runner"
)

func ReadFile(f string) {
	file_content, err := os.ReadFile(f)
	if err != nil {
		fmt.Println("ERROR: Failed to read file, ", err)
		return
	}
	fmt.Println(string(file_content))
	runner.Run(string(file_content))
}
