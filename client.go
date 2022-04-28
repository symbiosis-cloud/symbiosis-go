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
	httpClient *resty.Client

	Team     *TeamService
	Cluster  *ClusterService
	NodePool *NodePoolService
	Node     *NodeService
}

type ClientOption func(c *resty.Client)

type SortAndPageable struct {
	Pageable struct {
		Sort struct {
			Sorted   bool `json:"sorted"`
			Unsorted bool `json:"unsorted"`
			Empty    bool `json:"empty"`
		} `json:"sort"`
		PageNumber int  `json:"pageNumber"`
		PageSize   int  `json:"pageSize"`
		Offset     int  `json:"offset"`
		Paged      bool `json:"paged"`
		Unpaged    bool `json:"unpaged"`
	} `json:"pageable"`
	TotalPages    int  `json:"totalPages"`
	TotalElements int  `json:"totalElements"`
	Last          bool `json:"last"`
	Sort          struct {
		Sorted   bool `json:"sorted"`
		Unsorted bool `json:"unsorted"`
		Empty    bool `json:"empty"`
	} `json:"sort"`
	NumberOfElements int  `json:"numberOfElements"`
	First            bool `json:"first"`
	Size             int  `json:"size"`
	Number           int  `json:"number"`
	Empty            bool `json:"empty"`
}

func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *resty.Client) {
		c.SetTimeout(timeout)
	}
}

func WithEndpoint(endpoint string) ClientOption {
	return func(c *resty.Client) {
		c.SetHostURL(endpoint)
	}
}

func NewClient(apiKey string, opts ...ClientOption) (*Client, error) {

	if apiKey == "" {
		return nil, errors.New("No apiKey given")
	}

	httpClient := resty.New().
		SetHostURL(apiEndpoint).
		SetHeader("X-Auth-ApiKey", apiKey).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetTimeout(time.Second * 10)

	for _, opt := range opts {
		opt(httpClient)
	}

	client := &Client{
		httpClient: httpClient,
	}
	client.Team = &TeamService{client}
	client.Cluster = &ClusterService{client}
	client.Node = &NodeService{client}
	client.NodePool = &NodePoolService{client}

	return client, nil
}

func (c *Client) ValidateResponse(resp *resty.Response, result interface{}) (interface{}, error) {

	statusCode := resp.StatusCode()

	switch statusCode {
	case 401:
		return result, &AuthError{
			StatusCode: resp.StatusCode(),
			Err:        errors.New("Authentication failed"),
		}
	case 201:
	case 200:
		return result, nil
	case 405:
	case 400:
	case 500:
		var genericError *GenericError
		json.Unmarshal(resp.Body(), &genericError)

		return nil, genericError
	case 404:
		return nil, &NotFoundError{404, resp.Request.URL, resp.Request.Method}
	}

	return nil, errors.New("Unexpected error occurred")
}
