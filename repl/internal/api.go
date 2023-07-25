package internal

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationAreaAPIResponse struct {
	Count    int                  `json:"count"`
	Next     string               `json:"next"`
	Previous any                  `json:"previous"`
	Results  []LocationAreaResult `json:"results"`
}

type LocationAreaResult struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func GetInformation(url string) (LocationAreaAPIResponse, error) {
	res, err := http.Get(url)
	locationAreaAPIResponse := LocationAreaAPIResponse{}
	if err != nil {
		return locationAreaAPIResponse, err
	}
	if res.StatusCode > 299 {
		return locationAreaAPIResponse, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		return locationAreaAPIResponse, err
	}

	err = json.Unmarshal(body, &locationAreaAPIResponse)
	if err != nil {
		return locationAreaAPIResponse, err
	}

	return locationAreaAPIResponse, nil
}
