package symbiosis

import (
	"encoding/json"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

const clusterJSON = `
{
	"id": "42d712ab-d05b-46d6-99e2-fef2059e09ac",
	"name": "test",
	"kubeVersion": "1.23.5",
	"apiServerEndpoint": "does.not.matter",
	"state": "ACTIVE",
	"nodes": [
	  {
		"id": "19fc97a5-30de-4f35-8b28-a14bf640b557",
		"name": "random-0",
		"nodeType": {
		  "id": "2f8d2c39-23cd-4623-b816-9489a26c1b8d",
		  "name": "general-1",
		  "memoryMi": 1024,
		  "storageGi": 1,
		  "vcpu": 1,
		  "product": {
			"productCosts": [
			  {
				"currency": "USD",
				"unitCost": 5
			  }
			]
		  }
		},
		"region": {
		  "id": "9d420d32-31f8-4a4a-b790-d75ea58ffc08",
		  "name": "netherlands-1"
		},
		"privateIPv4Address": "10.0.0.1",
		"state": "ACTIVE"
	  }]
}
`

const clusterListJSON = `{ "content": [` + clusterJSON + `], ` + sortableJSON + ` }`
const nodeListJSON = `{ "content": [` + nodeJSON + `], ` + sortableJSON + ` }`

func TestListCluster(t *testing.T) {

	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/cluster?size=10&page=0"

	var fakeClusterList *ClusterList
	json.Unmarshal([]byte(clusterListJSON), &fakeClusterList)

	responder := httpmock.NewStringResponder(200, clusterListJSON)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	clusterList, err := c.Cluster.List(10, 0)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, fakeClusterList, clusterList)

	// test resty failure
	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	_, err = c.Cluster.List(10, 0)
	assert.Error(t, err)
}

func TestDescribeCluster(t *testing.T) {
	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/cluster/test"

	var fakeCluster *Cluster
	json.Unmarshal([]byte(clusterJSON), &fakeCluster)

	responder := httpmock.NewStringResponder(200, clusterJSON)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	cluster, err := c.Cluster.Describe("test")

	assert.Nil(t, err)
	assert.Equal(t, fakeCluster, cluster)

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	_, err = c.Cluster.Describe("test")
	assert.Error(t, err)
}

func TestDeleteCluster(t *testing.T) {
	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/cluster/test"

	responder := httpmock.NewStringResponder(200, "")
	httpmock.RegisterResponder("DELETE", fakeURL, responder)

	err := c.Cluster.Delete("test")

	assert.Nil(t, err)

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("DELETE", fakeURL, responder)

	err = c.Cluster.Delete("test")
	assert.Error(t, err)
}

func TestListNodes(t *testing.T) {
	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/cluster/test/node"

	var fakeNodeList *NodeList
	json.Unmarshal([]byte(nodeListJSON), &fakeNodeList)

	responder := httpmock.NewStringResponder(200, nodeListJSON)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	nodeList, err := c.Cluster.ListNodes("test")

	assert.Nil(t, err)
	assert.Equal(t, fakeNodeList, nodeList)

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	_, err = c.Cluster.ListNodes("test")
	assert.Error(t, err)
}
