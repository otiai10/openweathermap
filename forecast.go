package openweathermap

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// https://openweathermap.org/forecast5

// ByCityName ...
func (client *Client) ByCityName(name string, opt *Option) (*ForecastResponse, error) {

	if opt == nil {
		opt = &DefaultOption
	}

	u, err := url.Parse(client.BaseURL + "/data/2.5/forecast")
	if err != nil {
		return nil, err
	}

	query := opt.Query()

	// ?q=Tokyo&mode=json&apikey=1fb791ae4335504a8f367791bd4679d2&units=metric"
	query.Add("apikey", client.APIKey)
	query.Add("q", "Tokyo")
	query.Add("mode", "json")

	u.RawQuery = query.Encode()

	fmt.Println(u.String())
	res, err := client.HTTPClient.Get(u.String())
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
