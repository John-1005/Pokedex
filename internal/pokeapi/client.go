package pokeapi


import (
		"encoding/json"
		"net/http"
		"io/ioutil"
		"errors"
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
