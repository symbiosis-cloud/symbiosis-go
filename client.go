package symbiosis

import (
	"errors"

	_ "fmt"
	"github.com/go-resty/resty/v2"
	"time"
)

type Client struct {
	symbiosisAPI *resty.Client
}

const (
	MethodGet  = "Get"
	MethodPost = "Post"
	MethodPath = "Patch"
)

type ClientOption func(c *resty.Client)

func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *resty.Client) {
		c.SetTimeout(timeout)
	}
}

func NewClient(endpoint string, apiKey string, opts ...ClientOption) (*Client, error) {

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

	for _, opt := range opts {
		opt(client)
	}

	apiClient := &Client{client}

	return apiClient, nil
}

func (c *Client) ValidateResponse(resp *resty.Response, result interface{}) (interface{}, error) {

	statusCode := resp.StatusCode()

	switch statusCode {
	case 401:
		return result, &AuthError{
			StatusCode: resp.StatusCode(),
			Err:        errors.New("Authentication failed"),
		}
		break
	case 201:
	case 200:
		return result, nil
		break
	case 404:
		return nil, &NotFoundError{404}
		break
	}

	symbiosisErr := resp.Error().(*GenericError)
	return nil, symbiosisErr
}
