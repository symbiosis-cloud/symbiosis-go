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

type ClusterList struct {
	Clusters []*Cluster `json:"content"`
	SortAndPageable
}

type NodeList struct {
	Nodes []*Node `json:"content"`
	SortAndPageable
}

type ClusterService struct {
	client *Client
}

func (c *ClusterService) List(maxSize int, page int) (*ClusterList, error) {

	// TODO handle paging
	var result *ClusterList
	resp, err := c.client.httpClient.R().
		SetResult(&result).
		ForceContentType("application/json").
		Get(fmt.Sprintf("rest/v1/cluster?size=%d&page=%d", maxSize, page))

	if err != nil {
		return nil, err
	}

	validated, err := c.client.ValidateResponse(resp, result)

	if err != nil {
		return nil, err
	}

	if validated == nil {
		return nil, nil
	}

	return validated.(*ClusterList), nil
}

func (c *ClusterService) Describe(clusterName string) (*Cluster, error) {
	var result *Cluster

	resp, err := c.client.httpClient.R().
		SetResult(&result).
		ForceContentType("application/json").
		Get(fmt.Sprintf("rest/v1/cluster/%s", clusterName))

	if err != nil {
		return nil, err
	}

	validated, err := c.client.ValidateResponse(resp, result)

	if err != nil {
		return nil, err
	}

	if validated == nil {
		return nil, nil
	}

	return validated.(*Cluster), nil
}

func (c *ClusterService) ListNodes(clusterName string) (*NodeList, error) {
	var result *NodeList

	resp, err := c.client.httpClient.R().
		SetResult(&result).
		ForceContentType("application/json").
		Get(fmt.Sprintf("rest/v1/cluster/%s/node", clusterName))

	if err != nil {
		return nil, err
	}

	validated, err := c.client.ValidateResponse(resp, result)

	if err != nil {
		return nil, err
	}

	if validated == nil {
		return nil, nil
	}

	return validated.(*NodeList), nil
}
