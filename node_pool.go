package symbiosis

import (
	"fmt"
)

type NodePool struct {
	ID              string `json:"id"`
	NodeTypeName    string `json:"nodeTypeName"`
	ClusterName     string `json:"clusterName"`
	DesiredQuantity int    `json:"desiredQuantity"`
}

type NodePoolService struct {
	client *Client
}

type NodePoolInput struct {
	ClusterName  string `json:"clusterName"`
	NodeTypeName string `json:"nodeTypeName"`
	Quantity     int    `json:"quantity"`
}

type NodePoolUpdateInput struct {
	Quantity int `json:"quantity"`
}

func (n *NodePoolService) Describe(id string) (*NodePool, error) {
	var nodePool *NodePool

	err := n.client.
		Call(fmt.Sprintf("rest/v1/node-pool/%s", id),
			"Get",
			&nodePool)

	if err != nil {
		return nil, err
	}

	return nodePool, nil

}

func (n *NodePoolService) Create(input *NodePoolInput) (*NodePool, error) {
	var nodePool *NodePool

	err := n.client.
		Call("rest/v1/node-pool",
			"Post",
			&nodePool,
			WithBody(input))

	if err != nil {
		return nil, err
	}

	return nodePool, nil
}

func (n *NodePoolService) Update(nodePoolId string, input *NodePoolUpdateInput) error {
	err := n.client.
		Call(fmt.Sprintf("rest/v1/node-pool/%s", nodePoolId),
			"Put",
			nil,
			WithBody(input))

	if err != nil {
		return err
	}

	return nil
}
