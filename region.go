package symbiosis

type RegionService interface {
	List() ([]Region, error)
}

type RegionServiceClient struct {
	client *Client
}

func (n *RegionServiceClient) List() ([]Region, error) {

	var regions []Region

	err := n.client.
		Call("rest/v1/region",
			"Get",
			&regions)

	if err != nil {
		return nil, err
	}

	return regions, nil
}
