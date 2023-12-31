package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {

	//Concatenating the base URL and the endpoint
	endpoint := "/location-area"
	fullURL := baseURL + endpoint


	//check if pageUrl is not nil
	if pageURL != nil {
		fullURL = *pageURL
	}

	// check if the data is in the cache
	data, ok := c.cache.Get(fullURL)
	if ok{
		//cache hit
		fmt.Println("cache hit")
		locationAreasResp := LocationAreasResp{}
		err := json.Unmarshal(data, &locationAreasResp)
		if err != nil {
			return LocationAreasResp{}, err
		}

		return locationAreasResp, nil
	}
	fmt.Println("cache miss")

	//Creating a Req
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	//Making the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}

	//Check the status code
	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("Bad status code: %v", resp.StatusCode)
	}

	//Read the body
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	//Unmarshal the data
	locationAreasResp := LocationAreasResp{}
	err = json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	c.cache.Add(fullURL, data)

	return locationAreasResp, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {

	//Concatenating the base URL and the endpoint
	endpoint := "/location-area/" + locationAreaName
	fullURL := baseURL + endpoint

	// check if the data is in the cache
	data, ok := c.cache.Get(fullURL)
	if ok{
		//cache hit
		fmt.Println("cache hit")
		locationArea := LocationArea{}
		err := json.Unmarshal(data, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}

		return locationArea, nil
	}
	fmt.Println("cache miss")

	//Creating a Req
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	//Making the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}

	//Check the status code
	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("Bad status code: %v", resp.StatusCode)
	}

	//Read the body
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	//Unmarshal the data
	locationArea := LocationArea{}
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(fullURL, data)

	return locationArea, nil
}