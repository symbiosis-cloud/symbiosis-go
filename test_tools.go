package symbiosis

import (
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
)

func GetMockedClient() *Client {
	httpClient := resty.New()
	httpmock.ActivateNonDefault(httpClient.GetClient())

	client, _ := NewClient("apiKey")
	client.httpClient = httpClient

	return client
}
