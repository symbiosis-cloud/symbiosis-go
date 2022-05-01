package symbiosis

import (
	"encoding/json"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

const nodeJSON = `
{
	"id": "aaaaaaa-362b-4e49-afd4-d138239c71b1",
	"name": "test",
	"nodeType": {
		"id": "aaaaaaa-2023-4e81-8263-10d57e8461b4",
		"name": "general-int-1",
		"memoryMi": 4096,
		"storageGi": 20,
		"vcpu": 2,
		"product": {
			"productCosts": [
				{
					"currency": "USD",
					"unitCost": 5.00
				}
			]
		}
	},
	"region": {
		"id": "aaaaaaa-fca9-40b8-be77-055764f273e4",
		"name": "dev-1"
	},
	"privateIPv4Address": "10.128.0.5",
	"state": "ACTIVE",
	"kubeVersion": "1.23.5",
	"priority": 1
}`

func TestDescribeNode(t *testing.T) {
	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/node/test"

	var fakeNode *Node
	json.Unmarshal([]byte(nodeJSON), &fakeNode)

	responder := httpmock.NewStringResponder(200, nodeJSON)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	node, err := c.Node.Describe("test")

	assert.Nil(t, err)
	assert.Equal(t, fakeNode, node)

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	_, err = c.Node.Describe("test")
	assert.Error(t, err)
}

func TestRecycleNode(t *testing.T) {
	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/node/test/recycle"

	responder := httpmock.NewStringResponder(200, "")
	httpmock.RegisterResponder("PUT", fakeURL, responder)

	err := c.Node.Recycle("test")

	assert.Nil(t, err)

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("PUT", fakeURL, responder)

	err = c.Node.Recycle("test")
	assert.Error(t, err)
}
