package main

import (
	"fmt"
	"log"

	"github.com/Burncoat/pokedex_project/internal/pokeapi"
)

func commandMap() error {
	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
}
