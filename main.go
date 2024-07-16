package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Should be: %s <filename>\n", os.Args[0])
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error opening file: %s\n", os.Args[1])
		os.Exit(1)
	}

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading from file: %s\n", os.Args[1])
		os.Exit(1)
	}

	file.Close()

	sContent := string(content)

	lines := strings.Split(sContent, "\n")

	for index, line := range lines {
		lines[index] = strings.Trim(line, "; \t")
	}

	fmt.Println(lines)
}
