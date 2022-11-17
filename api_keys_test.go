package symbiosis

import (
	"encoding/json"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const apiKeyJSON = `
[
	{
		"id": "jyduxxxxxxxbdt",
		"token": "a0jyduhjjaykagxbdt****************",
		"subjectId": "xxxxxxx-040e-401f-a4e5-xxxxxxxxxxx",
		"role": "ADMIN",
		"lastUsedAt": "2022-06-01T18:36:37.646083Z"
	},
	{
		"id": "pdxxxxxxxxxxxe",
		"token": "a0pdxxhusgrasbfxxe****************",
		"subjectId": "xxxxxxx-1f16-4e98-893b-xxxxxxxxxxx",
		"role": "ADMIN",
		"lastUsedAt": null
	}
]
`

func TestCreateApiKey(t *testing.T) {
	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/api-keys"

	responder := httpmock.NewStringResponder(200, `{
		"id": "pdxxxxxxxxxxxe",
		"token": "a0pdxxhusgrasbfxxe****************",
		"subjectId": "xxxxxxx-1f16-4e98-893b-xxxxxxxxxxx",
		"role": "MEMBER",
		"lastUsedAt": null
	}`)
	httpmock.RegisterResponder("POST", fakeURL, responder)

	apiKey, err := c.ApiKeys.Create(ApiKeyInput{
		Role:        ROLE_MEMBER,
		Description: "hello world",
	})

	assert.Nil(t, err)
	assert.Equal(t, apiKey.ID, "pdxxxxxxxxxxxe")
	assert.Equal(t, apiKey.Role, ROLE_MEMBER)

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("POST", fakeURL, responder)

	_, err = c.ApiKeys.Create(ApiKeyInput{
		Role:        ROLE_MEMBER,
		Description: "hello world",
	})
	assert.Error(t, err)
}

func TestListApiKeys(t *testing.T) {
	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/api-keys"

	var fakeApiKeys ApiKeyCollection
	json.Unmarshal([]byte(apiKeyJSON), &fakeApiKeys)

	responder := httpmock.NewStringResponder(200, apiKeyJSON)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	apiKeys, err := c.ApiKeys.List()

	assert.Nil(t, err)
	assert.Equal(t, fakeApiKeys, apiKeys)

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	_, err = c.ApiKeys.List()
	assert.Error(t, err)
}

func TestDeleteApiKey(t *testing.T) {
	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/api-keys/test123"

	responder := httpmock.NewStringResponder(200, "")
	httpmock.RegisterResponder("DELETE", fakeURL, responder)

	err := c.ApiKeys.Delete("test123")

	assert.Nil(t, err)

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("DELETE", fakeURL, responder)

	err = c.ApiKeys.Delete("test123")
	assert.Error(t, err)
}
