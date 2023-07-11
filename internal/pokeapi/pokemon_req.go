package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {

	//Concatenating the base URL and the endpoint
	endpoint := "/pokemon/" + pokemonName
	fullURL := baseURL + endpoint

	// check if the data is in the cache
	data, ok := c.cache.Get(fullURL)
	if ok{
		//cache hit
		fmt.Println("cache hit")
		pokemon := Pokemon{}
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}

		return pokemon, nil
	}
	fmt.Println("cache miss")

	//Creating a Req
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	//Making the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	//Check the status code
	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("Bad status code: %v", resp.StatusCode)
	}

	//Read the body
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	//Unmarshal the data
	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullURL, data)

	return pokemon, nil
}