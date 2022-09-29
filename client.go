package symbiosis

import (
	"encoding/json"
	"errors"

	"github.com/go-resty/resty/v2"

	"reflect"
	"time"
)

const (
	APIEndpoint        = "https://api.symbiosis.host"
	StagingAPIEndpoint = "https://api.staging.symbiosis.host"
)

type SymbiosisClient interface {
	NewClientFromAPIKey(apiKey string, opts ...ClientOption) (*Client, error)
}

type Client struct {
	httpClient *resty.Client

	Team     TeamService
	Cluster  ClusterService
	NodePool NodePoolService
	Node     NodeService
}

type ClientOption func(c *resty.Client)
type CallOption func(req *resty.Request)

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

func WithBody(body interface{}) CallOption {
	return func(req *resty.Request) {
		req.SetBody(body)
	}
}

func newHttpClient(opts ...ClientOption) *resty.Client {
	httpClient := resty.New().
		SetHostURL(APIEndpoint).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetTimeout(time.Second * 90)

	for _, opt := range opts {
		opt(httpClient)
	}

	return httpClient

}

func NewClientFromAPIKey(apiKey string, opts ...ClientOption) (*Client, error) {

	if apiKey == "" {
		return nil, errors.New("No apiKey given")
	}

	httpClient := newHttpClient(opts...)
	httpClient.SetHeader("X-Auth-ApiKey", apiKey)

	client := &Client{
		httpClient: httpClient,
	}
	client.Team = &TeamServiceClient{client}
	client.Cluster = &ClusterServiceClient{client}
	client.Node = &NodeServiceClient{client}
	client.NodePool = &NodePoolServiceClient{client}

	return client, nil
}

func (c *Client) ValidateResponse(resp *resty.Response) error {

	statusCode := resp.StatusCode()

	switch statusCode {
	case 401:
		return &AuthError{
			StatusCode: resp.StatusCode(),
			Err:        errors.New("Authentication failed"),
		}
	case 201, 200:
		return nil
	case 405, 400, 403, 500:
		var GenericError *GenericError
		json.Unmarshal(resp.Body(), &GenericError)

		return GenericError
	case 404:
		return &NotFoundError{404, resp.Request.URL, resp.Request.Method}
	}

	return errors.New("Unexpected error occurred")
}

func (c *Client) Call(route string, method string, targetInterface interface{}, opts ...CallOption) error {

	req := c.httpClient.R()

	if targetInterface != nil {
		req.SetResult(&targetInterface)
	}

	req.ForceContentType("application/json")

	for _, opt := range opts {
		opt(req)
	}

	res := reflect.ValueOf(req).
		MethodByName(method).
		Call([]reflect.Value{reflect.ValueOf(route)})

	var resp *resty.Response
	if v := res[0].Interface(); v != nil {
		resp = v.(*resty.Response)
	}

	var err error
	if v := res[1].Interface(); v != nil {
		err = v.(error)
	}

	if err != nil {
		return err
	}

	err = c.ValidateResponse(resp)

	if err != nil {
		return err
	}

	return nil
}
