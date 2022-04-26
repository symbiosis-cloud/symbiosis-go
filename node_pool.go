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

func (c *Client) DescribeNodePool(id string) (*NodePool, error) {
	var result *NodePool

	resp, err := c.symbiosisAPI.R().
		SetResult(&result).
		ForceContentType("application/json").
		Get(fmt.Sprintf("rest/v1/node-pool/%v", id))

	if err != nil {
		return nil, err
	}

	validated, err := c.ValidateResponse(resp, result)

	if err != nil {
		return nil, err
	}

	if validated == nil {
		return nil, nil
	}

	return validated.(*NodePool), nil

}
