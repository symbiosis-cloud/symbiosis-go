package symbiosis

import (
	"encoding/json"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

const nodePoolJSON = `
{
  "Id": "test",
  "clusterName": "test",
  "nodeTypeName": "general-1",
  "quantity": 2
}`

func TestDescribeNodePool(t *testing.T) {
	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/node-pool/test"

	var fakeNodePool *NodePool
	json.Unmarshal([]byte(nodePoolJSON), &fakeNodePool)

	responder := httpmock.NewStringResponder(200, nodePoolJSON)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	nodePool, err := c.NodePool.Describe("test")

	assert.Nil(t, err)
	assert.Equal(t, fakeNodePool, nodePool)

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	_, err = c.NodePool.Describe("test")
	assert.Error(t, err)
}
func TestCreateNodePool(t *testing.T) {
	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/node-pool"

	var fakeNodePool *NodePool
	json.Unmarshal([]byte(nodePoolJSON), &fakeNodePool)

	responder := httpmock.NewStringResponder(200, nodePoolJSON)
	httpmock.RegisterResponder("POST", fakeURL, responder)

	input := &NodePoolInput{
		ClusterName:  "test",
		NodeTypeName: "general-1",
		Quantity:     1,
	}
	nodePool, err := c.NodePool.Create(input)

	assert.Nil(t, err)
	assert.Equal(t, fakeNodePool.ClusterName, input.ClusterName)
	assert.Equal(t, fakeNodePool.NodeTypeName, input.NodeTypeName)
	assert.Equal(t, fakeNodePool, nodePool)

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("POST", fakeURL, responder)

	_, err = c.NodePool.Create(input)
	assert.Error(t, err)
}
func TestUpdateNodePool(t *testing.T) {
	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/node-pool/test"

	responder := httpmock.NewStringResponder(200, "")
	httpmock.RegisterResponder("PUT", fakeURL, responder)

	input := &NodePoolUpdateInput{
		Quantity: 2,
	}
	err := c.NodePool.Update("test", input)

	assert.Nil(t, err)

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("PUT", fakeURL, responder)

	err = c.NodePool.Update("test", input)
	assert.Error(t, err)
}
