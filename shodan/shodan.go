package shodan

// BaseURL defines the Shodan URL
const BaseURL = "https://api.shodan.io"

// Client holds the API key
type Client struct {
	apiKey string
}

// New function is to initialize Client
func New(apiKey string) *Client {
	return &Client{apiKey: apiKey}
}
