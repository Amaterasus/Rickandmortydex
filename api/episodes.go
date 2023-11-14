package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Episode struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Aired string `json:"air_date"`
	Episode string `json:"episode"`
	Characters []string `json:"characters"`
	Url string `json:"url"`
	Created string `json:"created"`
}

type EpisodesResponse struct {
	MetaData Info `json:"info"`
	Results []Episode `json:"results"`
}

func GetEpisodes(page string) (episodes []Episode, err error) {
	url := fmt.Sprintf("https://rickandmortyapi.com/api/episode?page=%v", page)

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

	episodesResponse := EpisodesResponse{}
	
	err = json.Unmarshal(body, &episodesResponse)

	if err != nil {
		return
	}

	episodes = episodesResponse.Results

	return 
}

func GetSpecifiedEpisodes(ids string) (episodes []Episode, err error) {
	url := fmt.Sprintf("https://rickandmortyapi.com/api/episode/%v", ids)

	log.Println(url)

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

	err = json.Unmarshal(body, &episodes)

	if err != nil {
		return
	}

	return
}

func GetSpecifiedEpisode(id string) (episode Episode, err error) {
	url := fmt.Sprintf("https://rickandmortyapi.com/api/episode/%v", id)

	log.Println(url)

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

	err = json.Unmarshal(body, &episode)

	if err != nil {
		return
	}

	return
}
