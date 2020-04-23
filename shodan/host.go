package host

type HostInfo struct {
	Os        string `json: "os"`
	Timestamp `json: "timestamp"`
	Isp       string   `json: "isp"`
	Asn       string   `json: "asn"`
	Hostnames []string `json: "hostnames"`
	IP        string   `json: "ip"`
	Domains   []string `json: "domains"`
	Org       `json: "org"`
	Data      `json: "data"`
	Port      `json: "port"`
	Ip_string `json: "ip_string"`
}

type HostLocation struct {
	City          string  `json: "city"`
	Region_code   int     `json: "region_code"`
	Area_code     int     `json: "area_code"`
	Longitude     float64 `json: "longitude"`
	Country_code3 string  `json: "country_code3"`
	Country_name  string  `json: "country_name"`
	Postal_code   int     `json: "postal_code"`
	Dma_code      int     `json: "dma_code"`
	Country_code  string  `json: "country_code"`
	Latitude      int     `json: "latitude"`
}

type HostSearch struct {
	Matches []HostInfo `json: "matches"`
}

func (s *Client) HostSearch(q string) (*HostSearch, error) {
	res, err := http.Get(fmt.Sprintf("%s/shodan/host/search?key=%s&query=%s", BaseURL, s.apiKey, q))
}
if err != nil {
	fmt.Println(err)
}
defer res.Body.Close()

var ret HostSearch
if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
	return nil, err
}
return &ret, nil
}