package symbiosis

import (
	"fmt"
)

type Cluster struct {
	Name      string
	State     string
	NodePools []*NodePool
	Nodes     []*Node
	client    *Client
}

type ClusterResult struct {
	Clusters []*Cluster `json:"content"`
}

func (c *Client) ListClusters(maxSize int, page int) (*ClusterResult, error) {

	// TODO handle paging
	var result *ClusterResult
	resp, err := c.symbiosisAPI.R().
		SetResult(&result).
		ForceContentType("application/json").
		Get(fmt.Sprintf("rest/v1/cluster?size=%d&page=%s", maxSize, page))

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

	if len(result.Clusters) > 0 {
		for _, cluster := range result.Clusters {
			cluster.populateClient(c)
		}
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

	result.populateClient(c)

	return validated.(*Cluster), nil
}

func (c *Cluster) populateClient(client *Client) {
	c.client = client

	if len(c.Nodes) > 0 {
		for _, node := range c.Nodes {
			node.client = client
		}
	}
}
