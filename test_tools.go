package symbiosis

import (
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
)

const RegionJson = `{
	"id": "random-uuid",
	"name": "germany-1"
}`

const sortableJSON = `
"pageable": {
	"sort": {
		"sorted": false,
		"unsorted": true,
		"empty": true
	},
	"pageNumber": 0,
	"pageSize": 20,
	"offset": 0,
	"paged": true,
	"unpaged": false
},
"totalPages": 1,
"totalElements": 1,
"last": true,
"sort": {
	"sorted": false,
	"unsorted": true,
	"empty": true
},
"numberOfElements": 1,
"first": true,
"size": 20,
"number": 0,
"empty": false`

func getMocketClient() *Client {
	httpClient := resty.New()
	httpmock.ActivateNonDefault(httpClient.GetClient())

	client, _ := NewClientFromAPIKey("apiKey")
	client.httpClient = httpClient

	return client
}
