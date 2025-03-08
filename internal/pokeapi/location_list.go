package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespLocationsList, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := RespLocationsList{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespLocationsList{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationsList{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationsList{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespLocationsList{}, err
	}

	locationsResp := RespLocationsList{}

	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return RespLocationsList{}, err
	}
	c.cache.Add(url, data)
	return locationsResp, nil
}
