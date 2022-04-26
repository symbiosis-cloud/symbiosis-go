package symbiosis

import (
	"fmt"
)

type Node struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	NodeType struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		MemoryMi  int    `json:"memoryMi"`
		StorageGi int    `json:"storageGi"`
		Vcpu      int    `json:"vcpu"`
		Product   struct {
			ProductCosts []struct {
				Currency string  `json:"currency"`
				UnitCost float32 `json:"unitCost"`
			} `json:"productCosts"`
		} `json:"product"`
	} `json:"nodeType"`
	Region struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"region"`
	PrivateIPv4Address string `json:"privateIPv4Address"`
	State              string `json:"state"`
	client             *Client
}

func (c *Client) DescribeNode(name string) (*Node, error) {
	var result *Node

	resp, err := c.symbiosisAPI.R().
		SetResult(&result).
		ForceContentType("application/json").
		Get(fmt.Sprintf("rest/v1/node/%v", name))

	if err != nil {
		return nil, err
	}

	result.client = c
	validated, err := c.ValidateResponse(resp, result)

	if err != nil {
		return nil, err
	}

	if validated == nil {
		return nil, nil
	}

	return validated.(*Node), nil
}

func (n *Node) Recycle() error {
	c := n.client

	resp, err := c.symbiosisAPI.R().
		ForceContentType("application/json").
		Put(fmt.Sprintf("rest/v1/node/%v/recycle", n.Name))

	if err != nil {
		return err
	}

	_, err = c.ValidateResponse(resp, nil)

	if err != nil {
		return err
	}
	return nil
}
