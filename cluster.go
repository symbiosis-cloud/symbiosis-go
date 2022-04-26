package symbiosis

import (
	"fmt"
)

type Cluster struct {
	Name      string
	State     string
	NodePools []NodePool
	Nodes     []Node
}

type ClusterResult struct {
	Clusters []*Cluster `json:"content"`
}

func (c *Client) ListClusters(maxSize int) (*ClusterResult, error) {

	// TODO handle paging
	var result *ClusterResult
	resp, err := c.symbiosisAPI.R().
		SetResult(&result).
		ForceContentType("application/json").
		Get(fmt.Sprintf("rest/v1/cluster?size=%d&page=0", maxSize))

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

	return validated.(*ClusterResult), nil
}

func (c *Client) DescribeCluster(name string) (*Cluster, error) {
	var result *Cluster

	resp, err := c.symbiosisAPI.R().
		SetResult(&result).
		ForceContentType("application/json").
		Get(fmt.Sprintf("rest/v1/cluster/%s", name))

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

	return validated.(*Cluster), nil
}
