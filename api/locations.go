package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Location struct {
	Id int `json:"id"`
	Name string `json:"name"`
	LocationType string `json:"type"`
	Dimension string `json:"dimension"`
	Residents []string `json:"residents"`
	Url string `json:"url"`
	Created string `json:"created"`
}

type LocationsResponse struct {
	MetaData Info `json:"info"`
	Results []Location `json:"results"`
}

func GetLocations(page string) (locations []Location, err error) {
	url := fmt.Sprintf("https://rickandmortyapi.com/api/location?page=%v", page)

	resp, err := http.Get(url)

	if err != nil {
		return
	}

	if resp.Body == nil {
		err = errors.New("No response body found")
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	locationsResponse := LocationsResponse{}
	
	err = json.Unmarshal(body, &locationsResponse)

	if err != nil {
		return
	}

	locations = locationsResponse.Results

	return 
}
