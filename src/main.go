package main

import (
	"bufio"
	"os"
)

// Function main is the entry point of the program
func main() {
	stdin := os.Stdin
	scanner := bufio.NewScanner(stdin)
	replScanner(scanner, stdin)
}
