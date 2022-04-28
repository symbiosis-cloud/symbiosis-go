package symbiosis

import (
	"fmt"
)

type NodePool struct {
	ID              string
	ClusterName     string
	NodeTypeName    string
	IsMaster        bool
	DesiredQuantity int
}

type NodePoolService struct {
	client *Client
}

func (n *NodePoolService) Describe(id string) (*NodePool, error) {
	var result *NodePool

	resp, err := n.client.httpClient.R().
		SetResult(&result).
		ForceContentType("application/json").
		Get(fmt.Sprintf("rest/v1/node-pool/%s", id))

	if err != nil {
		return nil, err
	}

	validated, err := n.client.ValidateResponse(resp, result)

	if err != nil {
		return nil, err
	}

	if validated == nil {
		return nil, nil
	}

	return validated.(*NodePool), nil

}
