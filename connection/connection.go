package connection

import (
	"errors"
	"net/http"
	"os"
	"time"
)

// AlphaVantageConnection is an interface for connecting to the AlphaVantage API
type AlphaVantageConnection interface {
	Request(request AlphaVantageRequest) (interface{}, error)
}

type alphaVantageConnection struct {
	client *http.Client
}

func loadAPIKey() (string, error) {
	key, ok := os.LookupEnv("ALPHA_VANTAGE_API_KEY")
	if !ok {
		return "", errors.New("ALPHA_VANTAGE_API_KEY not set")
	}
	return key, nil
}

// Build a request url from the query function and parameters
func makeURL(queryFunction QueryFunction, params map[string]string) (string, error) {
	urlString := API_BASE_URL + "?function=" + string(queryFunction)
	for key, value := range params {
		urlString += "&" + key + "=" + value
	}
	apiKey, err := loadAPIKey()
	if err != nil {
		return "", err
	}
	urlString += "&apikey=" + apiKey

	return urlString, nil
}

func (c *alphaVantageConnection) Request(request AlphaVantageRequest) (interface{}, error) {
	url, err := makeURL(request.QueryFunction(), request.QueryParams())
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	return request.ParseResponse(resp)
}

// NewAlphaVantageConnection creates a new AlphaVantageConnection
func NewAlphaVantageConnection() AlphaVantageConnection {
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	return &alphaVantageConnection{
		client: client,
	}
}
