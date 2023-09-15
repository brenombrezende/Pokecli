package main

import (
	"fmt"

	"github.com/brenombrezende/pokecli/internal/apihandler"
)

type PokeLocation struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

const path = "location/"

func commandMap() error {

	locations := apihandler.Request(path)
	fmt.Println(string(locations))
	return nil
}

func commandMapb() error {
	fmt.Println("To be implemented!")
	return nil
}
