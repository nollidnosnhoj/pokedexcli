package main

import "fmt"

func commandHelp(c *config) error {
	fmt.Println("Available commands:")
	for _, cmd := range getCommands() {
		fmt.Printf("  %s: %s\n", cmd.name, cmd.description)
	}
	return nil
}