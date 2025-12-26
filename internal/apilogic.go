package internal

import (
	"encoding/json"
	"io"
	"net/http"
)

type Maps struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetMapHelper(URL string) (Maps, error) {
	var mapData Maps

	if len(URL) == 0 {
		URL = "https://pokeapi.co/api/v2/location-area/"
	}

	res, err := http.Get(URL)
	if err != nil {
		return mapData, err
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if err := json.Unmarshal(body, &mapData); err != nil {
		return mapData, err
	}

	if res.StatusCode > 299 {
		return mapData, err
	}
	if err != nil {
		return mapData, err
	}
	return mapData, nil
}
