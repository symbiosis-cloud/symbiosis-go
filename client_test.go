package symbiosis

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestNewClientFromAPIKeyEmptyValues(t *testing.T) {
	_, err := NewClientFromAPIKey("")

	assert.ErrorContains(t, err, "No apiKey given")
}
func TestClientOption(t *testing.T) {

	client, _ := NewClientFromAPIKey("apiKey", WithTimeout(time.Second*11), WithEndpoint("https://someplace"))

	assert.Equal(t, time.Second*11, client.httpClient.GetClient().Timeout)
	assert.Equal(t, "https://someplace", client.httpClient.HostURL)
}

func TestNewClientFromAPIKeyResty(t *testing.T) {
	c, err := NewClientFromAPIKey("apiKey")

	assert.Nil(t, err)

	assert.IsType(t, &resty.Client{}, c.httpClient)
}

func TestValidateResponse(t *testing.T) {

	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	responseMap := map[int]string{
		404: `{ "timestamp": "2022-04-28T11:21:38.930+00:00", "status": 404, "error": "Not Found", "message": "404 NOT_FOUND", "path": "/rest/v1/node/x" }`,
		401: `{"status":401,"error":"Unauthorized","path":null,"message":"Token invalid"`,
		405: `{ "timestamp": "2022-04-28T09:29:26.884+00:00", "status": 405, "error": "Method Not Allowed", "message": "Request method 'GET' not supported", "path": "/rest/v1/node-pool" }`,
		200: `{ "content": {} }`,
		201: `{ "content": {} }`,
		999: "",
	}
	fakeURL := "http://does.not.matter"

	for statusCode, body := range responseMap {
		responder := httpmock.NewStringResponder(statusCode, body)
		httpmock.RegisterResponder("GET", fakeURL, responder)

		resp, err := c.httpClient.R().Get(fakeURL)

		if err != nil {
			t.Error(err)
		}

		err = c.ValidateResponse(resp)

		switch statusCode {
		case 401:
			assert.Error(t, &AuthError{StatusCode: 405})
			assert.ErrorContains(t, err, "Authentication failed")
			break
		case 200, 201:
			assert.Nil(t, err)
			break
		case 405, 400, 500:
			var genericError *GenericError
			json.Unmarshal([]byte(body), &genericError)

			fakeError := &GenericError{genericError.Status, genericError.ErrorType, genericError.Message, genericError.Path}

			assert.Equal(t, fakeError, genericError)
			assert.Equal(t, fakeError.Error(), "Symbiosis: Request method 'GET' not supported (type=Method Not Allowed, path=/rest/v1/node-pool)")
			assert.Error(t, fakeError, err)
			break
		case 404:
			fakeError := &NotFoundError{404, fakeURL, "GET"}
			assert.Error(t, fakeError, err)
			assert.Equal(t, fakeError.Error(), fmt.Sprintf("Symbiosis: GET %s. 404 not found", fakeURL))
			break
		default:
			assert.ErrorContains(t, err, "Unexpected error occurred")
		}
	}

	// test validation failure
	responder := httpmock.NewStringResponder(403, "")
	httpmock.RegisterResponder("GET", fakeURL, responder)

	err := c.Call(fakeURL, "Get", nil)

	assert.IsType(t, &GenericError{}, err)
}
