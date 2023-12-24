package main

type command struct {
	name string
	description string
	run func(c *config) error
}

func getCommands() map[string]command {
	return map[string]command{
		"map": {
			name: "map",
			description: "Show the map",
			run: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Show previous map",
			run: commandMapb,
		},
		"help": {
			name: "help",
			description: "Show help",
			run: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exit the program",
			run: commandExit,
		},
	}
}