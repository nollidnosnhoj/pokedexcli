package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type PokedexLocationResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func fetch(url string) (*PokedexLocationResponse, error) {
	res, err := http.Get(url)
	if err != nil || res.StatusCode != 200 {
		return nil, errors.New("error fetching locations")
	}
	resposne := &PokedexLocationResponse{}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("error reading response body")
	}
	err = json.Unmarshal(data, resposne)
	if err != nil {
		return nil, errors.New("error parsing response body")
	}
	return resposne, nil
}