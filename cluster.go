package symbiosis

import (
	"fmt"
	"time"
)

type ClusterConfigurationInput struct {
	EnableNginxIngress bool `json:"nginxIngress"`
}

type ClusterInput struct {
	Name          string                    `json:"name"`
	KubeVersion   string                    `json:"kubeVersion"`
	Region        string                    `json:"regionName"`
	Nodes         []ClusterNodeInput        `json:"nodes"`
	Configuration ClusterConfigurationInput `json:"configuration"`
}

type ClusterNodeInput struct {
	NodeType string `json:"nodeTypeName"`
	Quantity int    `json:"quantity"`
}

type Cluster struct {
	ID                string      `json:"id"`
	Name              string      `json:"name"`
	KubeVersion       string      `json:"kubeVersion"`
	APIServerEndpoint string      `json:"apiServerEndpoint"`
	State             string      `json:"state"`
	Nodes             []*Node     `json:"nodes"`
	NodePools         []*NodePool `json:"nodePools"`
	CreatedAt         time.Time   `json:"createdAt"`
}

type ClusterList struct {
	Clusters []*Cluster `json:"content"`
	*SortAndPageable
}

type NodeList struct {
	Nodes []*Node `json:"content"`
	*SortAndPageable
}

type ClusterService struct {
	client *Client
}

func (c *ClusterService) List(maxSize int, page int) (*ClusterList, error) {

	// TODO handle paging
	var clusterList *ClusterList

	err := c.client.
		Call(fmt.Sprintf("rest/v1/cluster?size=%d&page=%d", maxSize, page),
			"Get",
			&clusterList)

	if err != nil {
		return nil, err
	}

	return clusterList, nil
}

func (c *ClusterService) Describe(clusterName string) (*Cluster, error) {
	var cluster *Cluster

	err := c.client.
		Call(fmt.Sprintf("rest/v1/cluster/%s", clusterName),
			"Get",
			&cluster)

	if err != nil {
		return nil, err
	}

	return cluster, nil
}

func (c *ClusterService) Create(ClusterInput *ClusterInput) (*Cluster, error) {
	var cluster *Cluster

	err := c.client.
		Call(fmt.Sprintf("rest/v1/cluster"),
			"Post",
			&cluster,
			WithBody(ClusterInput))

	if err != nil {
		return nil, err
	}

	return cluster, nil
}

func (c *ClusterService) Delete(clusterName string) error {

	err := c.client.
		Call(fmt.Sprintf("rest/v1/cluster/%s", clusterName),
			"Delete",
			nil)

	if err != nil {
		return err
	}

	return nil

}

func (c *ClusterService) ListNodes(clusterName string) (*NodeList, error) {
	var nodeList *NodeList

	err := c.client.
		Call(fmt.Sprintf("rest/v1/cluster/%s/node", clusterName),
			"Get",
			&nodeList)

	if err != nil {
		return nil, err
	}

	return nodeList, nil
}
