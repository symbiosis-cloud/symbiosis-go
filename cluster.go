package symbiosis

import (
	"fmt"
)

type SymbiosisApiError struct {
	Status    int32
	ErrorType string `json:"error"`
	Message   string
	Path      string
}

type NodePool struct {
	Id              string
	ClusterName     string
	NodeTypeName    string
	IsMaster        bool
	DesiredQuantity int
}

type Cluster struct {
	Name      string
	State     string
	NodePools []NodePool
}

type ClusterResult struct {
	Clusters []*Cluster `json:"content"`
}

func (error *SymbiosisApiError) Error() string {
	return fmt.Sprintf("Symbiosis: %v (type=%v, path=%v)", error.Message, error.ErrorType, error.Path)
}

func (client *Client) ListClusters(maxSize int) (*ClusterResult, error) {
	api := client.symbiosisAPI
	// TODO handle paging

	var result *ClusterResult
	resp, err := api.R().
		SetResult(&result).
		ForceContentType("application/json").
		Get(fmt.Sprintf("rest/v1/cluster?size=%d&page=0", maxSize))

	if err != nil {
		return nil, err
	}
	if resp.StatusCode() == 404 {
		return result, nil
	}
	if resp.StatusCode() != 200 {
		symbiosisErr := resp.Error().(*SymbiosisApiError)
		return nil, symbiosisErr
	}

	return result, nil
}
