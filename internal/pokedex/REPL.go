package pokedex

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
		"map": {
			name:        "map",
			description: "Prints the map",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Prints the map with back order",
			callback:    commandMapb,
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

// Function commandMap prints the map
func commandMap() error {
	client := NewMyHttpClient()
	locationResponse, err := client.ListLocations(nil)

	if err != nil {
		return err
	}
	for _, location := range locationResponse.Results {
		fmt.Println(location.Name)
	}
	client.next = locationResponse.Next
	return nil
}

// Function commandMapb prints the map with back order
func commandMapb() error {
	return nil
}

// Function replScanner scans the input and executes the command
func ReplScanner(scanner *bufio.Scanner, stdin *os.File) {
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
