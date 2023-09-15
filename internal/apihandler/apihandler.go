package apihandler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

const baseUrl = "https://pokeapi.co/api/v2/"

func RequestLocationAreas(pageUrl *string) (PokeLocationResponse, error) {

	client := http.Client{
		Timeout: time.Second * 5,
	}

	endpoint := "location-area/"
	fullUrl := baseUrl + endpoint

	if pageUrl != nil {
		fullUrl = *pageUrl
	}

	res, err := client.Get(fullUrl)
	if err != nil {
		log.Fatal(err)

	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if res.StatusCode > 399 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)

	}

	if err != nil {
		log.Fatal(err)
	}

	pokeAreasResp := PokeLocationResponse{}
	err = json.Unmarshal(body, &pokeAreasResp)
	if err != nil {
		log.Fatal(err)
	}
	return pokeAreasResp, nil
}
