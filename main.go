package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

func main() {
	examples := [2]string{"hello", "utest"}
	bold := color.New(color.FgCyan, color.Bold).Printf

	if len(os.Args) < 2 {
		fmt.Println("Please specify a demo to run: ")
		bold("  %s\n", "go run main.go <example>")
		fmt.Println("Available examples:")
		for _, example := range examples {
			fmt.Println("  - " + example)
		}
		return
	}

	switch os.Args[1] {
	case "hello":
		hello()
	default:
		fmt.Println("Unknown demo:", os.Args[1])
	}
}
