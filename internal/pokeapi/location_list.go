package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocation, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationList := RespShallowLocation{}
		err := json.Unmarshal(val, &locationList)

		if err != nil {
			return RespShallowLocation{}, err
		}
		return locationList, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocation{}, err
	}

	list_data, err := io.ReadAll(resp.Body)

	if err != nil {
		return RespShallowLocation{}, err
	}

	locationList := RespShallowLocation{}

	err = json.Unmarshal(list_data, &locationList)
	if err != nil {
		return RespShallowLocation{}, err
	}
	return locationList, nil
}
