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

type NodeService struct {
	client *Client
}

func (n *NodeService) Describe(name string) (*Node, error) {
	var result *Node

	resp, err := n.client.httpClient.R().
		SetResult(&result).
		ForceContentType("application/json").
		Get(fmt.Sprintf("rest/v1/node/%s", name))

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

	return validated.(*Node), nil
}

func (n *NodeService) Recycle(name string) error {

	resp, err := n.client.httpClient.R().
		ForceContentType("application/json").
		Put(fmt.Sprintf("rest/v1/node/%v/recycle", name))

	if err != nil {
		return err
	}

	_, err = n.client.ValidateResponse(resp, nil)

	if err != nil {
		return err
	}
	return nil
}
