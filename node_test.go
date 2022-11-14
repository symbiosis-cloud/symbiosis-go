package symbiosis

import (
	"encoding/json"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
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

var typesJson = `
[
	{
		"id": "xxxxxx-2023-4e81-8263-20d57e8461b4",
		"name": "general-1",
		"memoryMi": 2048,
		"storageGi": 20,
		"vcpu": 1,
		"product": {
			"productCosts": [
				{
					"currency": "USD",
					"unitCost": 6.00
				}
			]
		}
	},
	{
		"id": "xxxxxx-38cf-4586-b5a6-a534f71e532d",
		"name": "general-2",
		"memoryMi": 4096,
		"storageGi": 30,
		"vcpu": 2,
		"product": {
			"productCosts": [
				{
					"currency": "USD",
					"unitCost": 12.00
				}
			]
		}
	}
]
`

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

func TestNodeTypes(t *testing.T) {
	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/node-type"

	var fakeTypes []NodeType
	json.Unmarshal([]byte(typesJson), &fakeTypes)

	responder := httpmock.NewStringResponder(200, typesJson)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	types, err := c.Node.Types()

	assert.Nil(t, err)
	assert.Equal(t, fakeTypes, types)
	assert.Equal(t, types[0].Name, "general-1")

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	_, err = c.Node.Types()
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

func TestDeleteNode(t *testing.T) {
	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/node/test"

	responder := httpmock.NewStringResponder(200, "")
	httpmock.RegisterResponder("DELETE", fakeURL, responder)

	err := c.Node.Delete("test")

	assert.Nil(t, err)

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("DELETE", fakeURL, responder)

	err = c.Node.Delete("test")
	assert.Error(t, err)
}
