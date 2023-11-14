package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Character struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Status string `json:"status"`
	Species string `json:"species"`
	Subspecies string `json:"type"`
	Gender string `json:"gender"` 
	Origin struct {
		Name string `json:"name"`
		Link string `json:"url"`
	} `json:"origin"`
	Location struct {
		Name string `json:"name"`
		Link string `json:"url"`
	} `json:"location"`
	Image string `json:"image"`
	Episodes []string `json:"episode"`
	Url string `json:"url"`
}

type CharactersResponse struct {
	MetaData Info `json:"info"`
	Results []Character `json:"results"`
}

func GetCharacters(page string) (characters []Character, err error) {
	url := fmt.Sprintf("https://rickandmortyapi.com/api/character?page=%v", page)

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

	charactersResponse := CharactersResponse{}
	
	err = json.Unmarshal(body, &charactersResponse)

	if err != nil {
		return
	}

	characters = charactersResponse.Results

	return 
}

func GetCharacter(id string) (character Character, episodes []Episode, err error) {
	url := fmt.Sprintf("https://rickandmortyapi.com/api/character/%v", id)

	resp, err := http.Get(url)

	if err != nil {
		log.Println("Error in the get request")
		return
	}

	if resp.Body == nil {
		log.Println("No response body found")
		err = errors.New("No response body found")
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("error reading body")
		return
	}
	
	err = json.Unmarshal(body, &character)

	if err != nil {
		log.Println("error unmarshalling json")
		return
	}

	ids := sliceIds(character.Episodes)

	// Lazy implementation to handle when the api only has 1 episode and doesn't return an array should be done better but realistically I would prefer this fix in the api than in this logic
	if len(character.Episodes) == 1 {
		var episode Episode
		episode, err = GetSpecifiedEpisode(ids)

		if err != nil {
			log.Println("error getting episodes")
			return 
		}
		episodes = append(episodes, episode)

		return
	}

	episodes, err = GetSpecifiedEpisodes(ids)

	if err != nil {
		log.Println("error getting episodes")
		return
	}

	return 
}