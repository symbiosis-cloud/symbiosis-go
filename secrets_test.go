package symbiosis

import (
	"encoding/json"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const secretsJSON = `
{
	"super": {
		"value": "secret222",
		"isDevelopmentSecret": true,
		"isPreviewSecret": false,
		"isProductionSecret": true
	},
	"super2": {
		"value": "secret222",
		"isDevelopmentSecret": true,
		"isPreviewSecret": false,
		"isProductionSecret": true
	},
	"super3": {
		"value": "dfsdscf",
		"isDevelopmentSecret": true,
		"isPreviewSecret": true,
		"isProductionSecret": true
	}
}
`

const previewSecretsJSON = `
{
	"super3": {
		"value": "dfsdscf",
		"isDevelopmentSecret": false,
		"isPreviewSecret": true,
		"isProductionSecret": false
	}
}
`

func TestCreateSecret(t *testing.T) {
	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/project/test/secret/test"

	responder := httpmock.NewStringResponder(200, "")
	httpmock.RegisterResponder("PUT", fakeURL, responder)

	err := c.Secret.Create("test", "test", Secret{
		Value:               "123",
		IsDevelopmentSecret: false,
		IsPreviewSecret:     true,
		IsProductionSecret:  false,
	})

	assert.Nil(t, err)

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("PUT", fakeURL, responder)

	err = c.Secret.Create("test", "test", Secret{
		Value:               "123",
		IsDevelopmentSecret: false,
		IsPreviewSecret:     true,
		IsProductionSecret:  false,
	})
	assert.Error(t, err)
}

func TestGetSecretsByProject(t *testing.T) {
	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/project/test/secret"

	var fakeSecrets SecretCollection
	json.Unmarshal([]byte(secretsJSON), &fakeSecrets)

	responder := httpmock.NewStringResponder(200, secretsJSON)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	secrets, err := c.Secret.GetSecretsByProject("test")

	assert.Nil(t, err)
	assert.Equal(t, fakeSecrets, secrets)
	assert.Equal(t, secrets["super"].Value, "secret222")
	assert.Equal(t, secrets["super"].IsDevelopmentSecret, true)
	assert.Equal(t, secrets["super"].IsPreviewSecret, false)
	assert.Equal(t, secrets["super"].IsProductionSecret, true)

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	_, err = c.Secret.GetSecretsByProject("test")
	assert.Error(t, err)
}

func TestGetSecretsByProjectAndEnvironment(t *testing.T) {
	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/project/test/secret"

	var fakeSecrets SecretCollection
	json.Unmarshal([]byte(previewSecretsJSON), &fakeSecrets)

	responder := httpmock.NewStringResponder(200, previewSecretsJSON)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	secrets, err := c.Secret.GetSecretsByProjectAndEnvironment("test", ENVIRONMENT_PREVIEW)

	assert.Nil(t, err)
	assert.Equal(t, len(secrets), 1)
	assert.Equal(t, fakeSecrets, secrets)
	assert.Equal(t, secrets["super3"].IsPreviewSecret, true)

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	_, err = c.Secret.GetSecretsByProjectAndEnvironment("test", ENVIRONMENT_PREVIEW)
	assert.Error(t, err)
}
