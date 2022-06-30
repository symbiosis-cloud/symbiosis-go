package symbiosis

import (
	"fmt"
)

type SchedulerEffect string

const (
	EFFECT_NO_SCHEDULE        SchedulerEffect = "NoSchedule"
	EFFECT_PREFER_NO_SCHEDULE SchedulerEffect = "PreferNoSchedule"
	EFFECT_NO_EXECUTE         SchedulerEffect = "NoExecute"
)

type NodePool struct {
	ID              string      `json:"id"`
	NodeTypeName    string      `json:"nodeTypeName"`
	ClusterName     string      `json:"clusterName"`
	DesiredQuantity int         `json:"desiredQuantity"`
	Labels          []NodeLabel `json:"labels"`
	Taints          []NodeTaint `json:"taints"`
}

type NodePoolService struct {
	client *Client
}

type NodePoolUpdateInput struct {
	Quantity int `json:"quantity"`
}

type NodePoolInput struct {
	Name         string      `json:"name"`
	ClusterName  string      `json:"clusterName"`
	NodeTypeName string      `json:"nodeTypeName"`
	Quantity     int         `json:"quantity"`
	Labels       []NodeLabel `json:"labels"`
	Taints       []NodeTaint `json:"taints"`
}

type NodeLabel struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type NodeTaint struct {
	Key    string          `json:"key"`
	Value  string          `json:"value"`
	Effect SchedulerEffect `json:"effect"`
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

func (n *NodePoolService) Delete(nodePoolId string) error {
	err := n.client.
		Call(fmt.Sprintf("rest/v1/node-pool/%s", nodePoolId),
			"Delete",
			nil)

	if err != nil {
		return err
	}

	return nil
}
