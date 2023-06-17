package main

import (
	"testing"
)

// TestGetCommandList tests the function getCommandList
func TestGetCommandList(t *testing.T) {
	commands := getCommandList()
	if len(commands) != 2 {
		t.Errorf("Expected 2 commands, got %d", len(commands))
	}
}

// TestCommandHelp tests the function commandHelp
func TestCommandHelp(t *testing.T) {
	err := commandHelp()
	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}
}
