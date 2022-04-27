package symbiosis

import (
	"encoding/json"
	"errors"

	"github.com/go-resty/resty/v2"
	"time"
)

const (
	apiEndpoint = "https://api.symbiosis.host"
)

type Client struct {
	symbiosisAPI *resty.Client
}

type ClientOption func(c *resty.Client)

func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *resty.Client) {
		c.SetTimeout(timeout)
	}
}

func WithAlternativeEndpoint(endpoint string) ClientOption {
	return func(c *resty.Client) {
		c.SetHostURL(endpoint)
	}
}

func NewClient(apiKey string, opts ...ClientOption) (*Client, error) {

	if apiKey == "" {
		return nil, errors.New("No apiKey given")
	}

	client := resty.New().
		SetHostURL(apiEndpoint).
		SetHeader("X-Auth-ApiKey", apiKey).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetTimeout(time.Second * 10)

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
	case 400:
		var badRequest *GenericError
		json.Unmarshal(resp.Body(), &badRequest)

		return nil, badRequest
	case 404:
		return nil, &NotFoundError{404, resp.Request.URL, resp.Request.Method}
		break
	}

	symbiosisErr := resp.Error().(*GenericError)
	return nil, symbiosisErr
}
