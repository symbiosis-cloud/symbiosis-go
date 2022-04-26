package symbiosis

import (
	// "reflect"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewClientEmptyValues(t *testing.T) {
	_, err := NewClient("", "123")

	assert.ErrorContains(t, err, "No endpoint given")

	_, err = NewClient("https://nowhere", "")

	assert.ErrorContains(t, err, "No apiKey given")
}

func TestNewClientResty(t *testing.T) {
	c, err := NewClient("endpoint", "apiKey")

	assert.Nil(t, err)

	assert.IsType(t, &resty.Client{}, c.symbiosisAPI)
}
