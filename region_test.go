package symbiosis

import (
	"encoding/json"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const regionJSON = `
[
	{
		"id": "xxxxxxx-3557-48e6-a7f5-984251295303",
		"name": "germany-1"
	}
]
`

func TestListRegions(t *testing.T) {
	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/region"

	var fakeRegions []Region
	json.Unmarshal([]byte(regionJSON), &fakeRegions)

	responder := httpmock.NewStringResponder(200, regionJSON)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	regions, err := c.Region.List()

	assert.Nil(t, err)
	assert.Equal(t, fakeRegions, regions)
	assert.Equal(t, regions[0].Name, "germany-1")
	assert.Equal(t, regions[0].ID, "xxxxxxx-3557-48e6-a7f5-984251295303")

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	_, err = c.Region.List()
	assert.Error(t, err)
}
