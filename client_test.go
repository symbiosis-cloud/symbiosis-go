package symbiosis

import (
	// "reflect"
	"github.com/go-resty/resty/v2"
	_ "github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewClientEmptyValues(t *testing.T) {
	_, err := NewClient("")

	assert.ErrorContains(t, err, "No apiKey given")
}
func TestClientOption(t *testing.T) {
	client, _ := NewClient("apiKey", WithTimeout(time.Second*11), WithEndpoint("https://someplace"))

	assert.Equal(t, time.Second*11, client.httpClient.GetClient().Timeout)
	assert.Equal(t, "https://someplace", client.httpClient.HostURL)
}

func TestNewClientResty(t *testing.T) {
	c, err := NewClient("apiKey")

	assert.Nil(t, err)

	assert.IsType(t, &resty.Client{}, c.httpClient)
}
