package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name 		string
	description string
	callback 	func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
	}
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
		for {
			fmt.Print("Pokedex > ")
			scanner.Scan()
			text := cleanInput(scanner.Text())
			if len(text) == 0 {
				continue
			}
			commandName := text[0]
			command, exists := getCommands()[commandName]
			if exists {
				err := command.callback()
				if err != nil {
					fmt.Println(err)
				}
				continue
			} else {
				fmt.Println("Unknown command")
				continue
			}
		}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	final := strings.Fields(lowered)
	return final
}

