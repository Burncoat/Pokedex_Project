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
	callback 	func(*config, ...string) error
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
		"map": {
			name: "map",
			description: "Lists 20 locations at a time. Additional 'map' commands list the next 20.",
			callback: commandMap,
		},
		"mapb": {
			name: "map back",
			description: "Lists 20 locations at a time. Additional 'mapb' commands list the previous 20.",
			callback: commandMapB,
		},
		"explore": {
			name: "explore {location_area}",
			description: "Lists pokemon encountered in the area",
			callback: commandExplore,
		},
		"catch": {
			name: "catch {pokemon_name}",
			description: "Attempts to catch a pokemon and add it to your pokedex",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect {pokemon_name}",
			description: "Gives information on previously caught pokemon",
			callback: commandInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "Lists all pokemon you've caught",
			callback: commandPokedex,
		},
	}
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
		for {
			fmt.Print("Pokedex > ")
			scanner.Scan()
			text := cleanInput(scanner.Text())
			if len(text) == 0 {
				continue
			}
			commandName := text[0]
			args := []string{}
			if len(text) > 1 {
				args = text[1:]
			}


			command, exists := getCommands()[commandName]
			if exists {
				err := command.callback(cfg, args...)
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

