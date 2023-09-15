package apihandler

import (
	"io"
	"log"
	"net/http"
	"time"
)

func Request(path string) []byte {

	const url = "https://pokeapi.co/api/v2/"

	client := http.Client{
		Timeout: time.Second * 5,
	}

	res, err := client.Get(url + path)
	if err != nil {
		log.Fatal(err)

	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)

	}

	if err != nil {
		log.Fatal(err)
	}

	return body
}
