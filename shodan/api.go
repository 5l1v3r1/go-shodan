package shodan

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ProfileInfo struct {
	QueryCredits int    `json:"query_credits"`
	ScanCredits  int    `json:"scan_credits"`
	Telnet       bool   `json:"telnet"`
	Plan         string `json:"plan"`
	Https        bool   `json:"https"`
	Unlocked     bool   `json:"unlocked"`
}

func (s *Client) ProfileInfo() (*ProfileInfo, error) {
	res, err := http.Get(fmt.Sprintf(BaseURL+"/api-info?key=%s", s.apiKey))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var ret ProfileInfo
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
