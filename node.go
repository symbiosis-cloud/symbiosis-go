package symbiosis

import (
	"fmt"
)

type ProductCost struct {
	Currency string  `json:"currency"`
	UnitCost float32 `json:"unitCost"`
}

type Product struct {
	ProductCosts []*ProductCost `json:"productCosts"`
}

type Region struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type NodeType struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	MemoryMi  int      `json:"memoryMi"`
	StorageGi int      `json:"storageGi"`
	Vcpu      int      `json:"vcpu"`
	Product   *Product `json:"product"`
}

type Node struct {
	ID                 string    `json:"id"`
	Name               string    `json:"name"`
	NodeType           *NodeType `json:"nodeType"`
	Region             *Region   `json:"region"`
	PrivateIPv4Address string    `json:"privateIPv4Address"`
	State              string    `json:"state"`
	client             *Client
}

type NodeService struct {
	client *Client
}

func (n *NodeService) Describe(name string) (*Node, error) {

	var node *Node

	err := n.client.
		Call(fmt.Sprintf("rest/v1/node/%s", name),
			"Get",
			&node)

	if err != nil {
		return nil, err
	}

	return node, nil
}

func (n *NodeService) Recycle(name string) error {

	err := n.client.
		Call(fmt.Sprintf("rest/v1/node/%v/recycle", name),
			"Put",
			nil)

	if err != nil {
		return err
	}

	return nil
}
