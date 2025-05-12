package pokeapi


import (
		"encoding/json"
		"net/http"
		"io/ioutil"
		"errors"
		"net/url"
)



type Client struct {
	BaseURL string
}



func NewClient() Client {
	return Client{
		BaseURL: "https://pokeapi.co/api/v2",

	}
}


func (c *Client) GetLocationAreas(url string) (LocationAreaResponse, error) {
	if url == "" {
		url = c.BaseURL + "/location-area"

	}

	rsp, err := http.Get(url)
	if err != nil {
		return LocationAreaResponse{}, err

	}
	
	defer rsp.Body.Close()
	

	if rsp.StatusCode != http.StatusOK {
		return LocationAreaResponse{}, errors.New("unexpected status code")

	}
	
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return LocationAreaResponse{}, err

	}
	var locResponse LocationAreaResponse
	err = json.Unmarshal(body, &locResponse)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	
	return locResponse, nil

}

func (c *Client) GetLocationArea(area string) (LocationAreaPokemon, error) {
	if area == "" {
		return LocationAreaPokemon{}, errors.New("area name cannot be empty")
	}
	areaURL := c.BaseURL + "/location-area/" + url.PathEscape(area)
	
	rsp, err := http.Get(areaURL)
	if err != nil {
		return LocationAreaPokemon{}, err
	}
	
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		return LocationAreaPokemon{}, errors.New("unexpected status code")
	}

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return LocationAreaPokemon{}, err
	}

	var locResponse LocationAreaPokemon
	err = json.Unmarshal(body, &locResponse)
	if err != nil {
		return LocationAreaPokemon{}, err
	}
	return locResponse, nil
}

func (c *Client) PokemonDetails(name string) (PokemonDetails, error) {
	if name == "" {
		return PokemonDetails{}, errors.New("pokemon name cannot be empty")
	}

	pokemonURL := c.BaseURL + "/pokemon/" + name 

	rsp, err := http.Get(pokemonURL)
	if err != nil {
		return PokemonDetails{}, err
	}
	
	defer rsp.Body.Close()
	
	if rsp.StatusCode != http.StatusOK {
		return PokemonDetails{}, errors.New("unexpected status code")
	}

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return PokemonDetails{}, err
	}

	var pokemonDetails PokemonDetails
	err = json.Unmarshal(body, &pokemonDetails)
	if err != nil {
		return PokemonDetails{}, err
	}
	return pokemonDetails, nil
}

