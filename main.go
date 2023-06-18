package main

import (
	"bufio"
	"os"

	"github.com/ibnumasud/prokcil-pokedex/internal/pokedex"
)

// Function main is the entry point of the program
func main() {
	stdin := os.Stdin
	scanner := bufio.NewScanner(stdin)
	pokedex.ReplScanner(scanner, stdin)
}
