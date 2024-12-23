package main

import (
	"fmt"
)

func Utest() {
	fmt.Println("Run the following command to see unit tests at work:")
	fmt.Printf("  %s\n", Command())
}

func Command() string {
	return "go test -v"
}
