package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (LocationArea, error) {
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		locationAreaResp := LocationArea{}
		err := json.Unmarshal(val, &locationAreaResp)
		if err != nil {
			return LocationArea{}, err
		}
		return locationAreaResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationAreaResp := LocationArea{}

	err = json.Unmarshal(data, &locationAreaResp)
	if err != nil {
		return LocationArea{}, err
	}
	c.cache.Add(url, data)
	return locationAreaResp, nil
}
