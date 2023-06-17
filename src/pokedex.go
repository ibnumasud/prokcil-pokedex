package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// Function GetCommandList returns a list of commands
func getCommandList() map[string]cliCommand {
	return map[string]cliCommand{

		"help": {
			name:        "help",
			description: "Prints this help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the program",
			callback:    commandExit,
		},
	}
}

// Function commandHelp prints the help message
func commandHelp() error {
	fmt.Println("Welcome to Pokedex")
	fmt.Println("Usage")
	fmt.Println()
	for _, cmd := range getCommandList() {
		fmt.Printf("%s\t%s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

// Function commandExit exits the program
func commandExit() error {
	os.Exit(0)
	return nil
}

// Function replScanner scans the input and executes the command
func replScanner(scanner *bufio.Scanner, stdin *os.File) {
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		commandName := scanner.Text()
		command, ok := getCommandList()[commandName]
		if !ok {
			fmt.Printf("Unknown command %s\n", commandName)
			continue
		} else {
			command.callback()
		}
	}
}

// Function main is the entry point of the program
func main() {
	stdin := os.Stdin
	scanner := bufio.NewScanner(stdin)
	replScanner(scanner, stdin)
}
