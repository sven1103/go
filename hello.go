package main

import (
	"strings"
	"time"
)
import "fmt"

func Hello() {
	const defaultName string = "golander"
	var name string
	fmt.Print("Please enter your name [golander]: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		recover()
	}
	if strings.TrimSpace(name) == "" {
		name = defaultName
	}
	fmt.Printf("Hello %s! \n", name)
	fmt.Printf("It is now %s.", time.Now())
}
