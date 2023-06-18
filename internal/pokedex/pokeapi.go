package pokedex

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	//baseURL is the base URL of the PokeAPI
	baseURL = "https://pokeapi.co/api/v2/"
)

// type http client private
type myHttpClient struct {
	client *http.Client
}

func NewMyHttpClient() *myHttpClient {
	return &myHttpClient{
		client: &http.Client{},
	}
}

// RespShallowLocations -
type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// Function http.get request from PokeAPI
func (c *myHttpClient) ListLocations() (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	//Create new Request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}
	req.Header.Set("Accept", "application/json")

	//Send request
	resp, err := c.client.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	//Read response
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}
	//Unmarshal response
	var shallowLocations RespShallowLocations
	err = json.Unmarshal(data, &shallowLocations)
	if err != nil {
		return RespShallowLocations{}, err
	}
	return shallowLocations, nil
}
