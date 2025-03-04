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

	return locationsResp, nil
}
