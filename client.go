package symbiosis

import (
	"errors"
	_ "fmt"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	symbiosisAPI *resty.Client
}

func NewClient(endpoint string, apiKey string) (*Client, error) {

	if endpoint == "" {
		return nil, errors.New("No endpoint given")
	}

	if apiKey == "" {
		return nil, errors.New("No apiKey given")
	}

	client := resty.New().
		SetHostURL(endpoint).
		SetHeader("X-Auth-ApiKey", apiKey).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	return &Client{client}, nil
}
