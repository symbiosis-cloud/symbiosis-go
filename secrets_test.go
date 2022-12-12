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
		"developmentValue": "secret221",
		"previewValue": "secret222",
		"productionValue": "secret223"
	},
	"super2": {
		"developmentValue": "secret222",
		"previewValue": "secret222",
		"productionValue": "secret222"
	},
	"super3": {
		"developmentValue": "dfsdscf",
		"previewValue": "dfsdscf",
		"productionValue": "dfsdscf"
	}
}
`

const previewValuesJSON = `
{
	"super3": "super-3-dev-secret"
}
`

func TestCreateSecret(t *testing.T) {
	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/project/test/secret/test"

	responder := httpmock.NewStringResponder(200, "")
	httpmock.RegisterResponder("PUT", fakeURL, responder)

	err := c.Secret.Create("test", "test", Secret{
		DevelopmentValue: "123",
		PreviewValue:     "123",
		ProductionValue:  "123",
	})

	assert.Nil(t, err)

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("PUT", fakeURL, responder)

	err = c.Secret.Create("test", "test", Secret{
		DevelopmentValue: "123",
		PreviewValue:     "123",
		ProductionValue:  "123",
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
	assert.Equal(t, secrets["super"].DevelopmentValue, "secret221")
	assert.Equal(t, secrets["super"].PreviewValue, "secret222")
	assert.Equal(t, secrets["super"].ProductionValue, "secret223")

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	_, err = c.Secret.GetSecretsByProject("test")
	assert.Error(t, err)
}

func TestGetSecretsByProjectAndEnvironment(t *testing.T) {
	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/project/test/secret?environment=preview"

	var fakeSecrets map[string]string
	json.Unmarshal([]byte(previewValuesJSON), &fakeSecrets)

	responder := httpmock.NewStringResponder(200, previewValuesJSON)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	secrets, err := c.Secret.GetSecretsByProjectAndEnvironment("test", ENVIRONMENT_PREVIEW)

	assert.Nil(t, err)
	assert.Equal(t, fakeSecrets, secrets)
	assert.Equal(t, len(secrets), 1)
	assert.Equal(t, secrets["super3"], "super-3-dev-secret")

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	_, err = c.Secret.GetSecretsByProjectAndEnvironment("test", ENVIRONMENT_PREVIEW)
	assert.Error(t, err)
}
