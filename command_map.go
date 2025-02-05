package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
	return getLocationResp(cfg, cfg.Next)
}

func commandMapb(cfg *config) error {
	if cfg.Previous == nil {
		fmt.Println("You're on the first page")
		return nil
	}
	return getLocationResp(cfg, cfg.Previous)
}

func getLocationResp(cfg *config, pageUrl *string) error {
	listLocations, err := cfg.pokeapiClient.ListLocations(pageUrl)
	if err != nil {
		return err
	}
	cfg.Next = listLocations.Next
	cfg.Previous = listLocations.Previous

	for _, location := range listLocations.Results {
		fmt.Println(location.Name)
	}
	return nil
}
