package pokeapi


import (
		"encoding/json"
		"net/http"
		"io/ioutil"
		"errors"
		"github.com/John-1005/pokedex/internal/pokecache"
)



type Client struct {
	BaseURL string
	cache *pokecache.Cache
}



func NewClient(cache *pokecache.Cache) Client {
	return Client{
		cache: cache, 
		BaseURL: "https://pokeapi.co/api/v2",

	}
}


func (c *Client) GetLocationAreas(url string) (LocationAreaResponse, error) {

	if url == "" {
		url = c.BaseURL + "/location-area"

	}

	cachedData, found := c.cache.Get(url) 
	if found {
		var response LocationAreaResponse
		err := json.Unmarshal(cachedData, &response)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		return response, nil
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

	c.cache.Add(url, body)

	var locResponse LocationAreaResponse
	err = json.Unmarshal(body, &locResponse)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	
	return locResponse, nil

}
