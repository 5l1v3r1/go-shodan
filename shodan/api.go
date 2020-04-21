package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Profile struct to hold Shodan account information
type Profile struct {
	Member       bool    `json: "member"`
	Credits      int     `json: "credits"`
	Display_name string  `json: "display_name"`
	Created      float64 `json: "created"`
}

func (p *Client) Profile() (*Profile, error) {
	req, err := http.Get(fmt.Sprintf("%s/api-info?key=%s", BaseURL, p.apiKey))
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	var ret Profile
	if err := json.NewDecoder(req.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
