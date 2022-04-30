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
