package openweathermap

import (
	"encoding/json"
	"net/url"
)

// https://openweathermap.org/forecast5

// ByCityName ...
func (client *Client) ByCityName(name string, opt *Option) (*ForecastResponse, error) {

	if opt == nil {
		opt = &DefaultOption
	}

	u := client.BaseURL + "/data/2.5/forecast"

	// ?q=Tokyo&mode=json&apikey=1fb791ae4335504a8f367791bd4679d2&units=metric"
	v := url.Values{
		"apikey": {client.APIKey},
		"q":      {"Tokyo"},
		"mode":   {"json"},
	}

	res, err := client.HTTPClient.Get(u + "?" + v.Encode())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	response := new(ForecastResponse)
	if err := json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, err
	}

	return response, nil
}
