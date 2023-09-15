package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/brenombrezende/pokecli/internal/apihandler"
)

func commandMap(cfg *config) error {

	resp, err := apihandler.RequestLocationAreas(cfg.nextLocationsURL)
	if err != nil {
		log.Fatal(err)
	}

	for _, value := range resp.Results {
		fmt.Printf("- %s\n", string(value.Name))
	}

	cfg.nextLocationsURL = resp.Next
	cfg.prevLocationsURL = resp.Previous

	return nil
}

func commandMapb(cfg *config) error {

	if cfg.prevLocationsURL == nil {
		return errors.New("you are on the first page")
	}

	resp, err := apihandler.RequestLocationAreas(cfg.prevLocationsURL)
	if err != nil {
		log.Fatal(err)
	}
	for _, value := range resp.Results {
		fmt.Printf("- %s\n", string(value.Name))
	}

	cfg.nextLocationsURL = resp.Next
	cfg.prevLocationsURL = resp.Previous

	return nil
}
