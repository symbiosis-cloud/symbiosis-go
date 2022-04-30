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
