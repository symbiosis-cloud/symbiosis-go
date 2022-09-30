package symbiosis

import (
	"encoding/json"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const nodePoolJSON = `
{
  "Id": "test",
  "name": "test",
  "clusterName": "test",
  "nodeTypeName": "general-1",
  "quantity": 2,
  "labels": [
	{
		"key": "hello",
		"value": "world"
	},
	{
		"key": "test",
		"value": "123"
	}
	],
	"taints": [
		{
			"key": "type",
			"value": "encoding",
			"effect": "NoSchedule"
		}
	],
	"nodes": [
		{
			"id": "test",
			"nodePoolId": "test",
			"name": "general-1-wripks",
			"nodeType": {
				"id": "test",
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
			"region": {
				"id": "fdcabb92-3557-48e6-a7f5-984251295303",
				"name": "germany-1"
			},
			"privateIPv4Address": "10.128.0.2",
			"state": "ACTIVE",
			"kubeState": "READY",
			"kubeVersion": "1.23.9",
			"priority": 1,
			"createdAt": "2022-09-01T12:47:51.109104Z"
		}
	],
	"autoscaling": {
		"enabled": true,
		"minSize": 1,
		"maxSize": 10
	}
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
	assert.Equal(t, fakeNodePool.Name, "test")
	assert.Equal(t, fakeNodePool.Labels[0].Key, "hello")
	assert.Equal(t, fakeNodePool.Labels[0].Value, "world")
	assert.Equal(t, fakeNodePool.Taints[0].Effect, EFFECT_NO_SCHEDULE)
	assert.Equal(t, fakeNodePool.Taints[0].Key, "type")
	assert.Equal(t, fakeNodePool.Taints[0].Value, "encoding")
	assert.Equal(t, fakeNodePool.Nodes[0].Name, "general-1-wripks")
	assert.Equal(t, fakeNodePool.Autoscaling, AutoscalingSettings{true, 1, 10})

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
		Autoscaling: AutoscalingSettings{
			Enabled: true,
			MinSize: 1,
			MaxSize: 10,
		},
	}
	err := c.NodePool.Update("test", input)

	assert.Nil(t, err)

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("PUT", fakeURL, responder)

	err = c.NodePool.Update("test", input)
	assert.Error(t, err)
}
func TestDeleteNodePool(t *testing.T) {
	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/node-pool/test"

	responder := httpmock.NewStringResponder(200, "")
	httpmock.RegisterResponder("DELETE", fakeURL, responder)

	err := c.NodePool.Delete("test")

	assert.Nil(t, err)

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("DELETE", fakeURL, responder)

	err = c.NodePool.Delete("test")
	assert.Error(t, err)
}
