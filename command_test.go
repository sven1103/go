package main

import "testing"

func TestCommand(t *testing.T) {
	answer := Command()
	if answer != "go test -v" {
		t.Errorf("Command failed - expected 'go test -v', got '%s'", answer)
	}
}
