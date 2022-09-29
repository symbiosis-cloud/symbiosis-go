package symbiosis

import (
	"fmt"
	"time"
)

type ClusterService interface {
	List(maxSize int, page int) (*ClusterList, error)
	Describe(clusterName string) (*Cluster, error)
	DescribeById(id string) (*Cluster, error)
	Create(input *ClusterInput) (*Cluster, error)
	Delete(clusterName string) error
	ListNodes(clusterName string) (*NodeList, error)
	CreateServiceAccountForSelf(clusterName string) (*ServiceAccount, error)
	GetServiceAccount(clusterName string, serviceAccountId string) (*ServiceAccount, error)
	CreateServiceAccount(clusterName string, subjectId string) (*ServiceAccount, error)
	DeleteServiceAccount(clusterName string, serviceAccountId string) error
	ListUserServiceAccounts(clusterName string) ([]*UserServiceAccount, error)
	GetIdentity(clusterName string) (*ClusterIdentity, error)
}

type ClusterInput struct {
	Name              string             `json:"name"`
	KubeVersion       string             `json:"kubeVersion"`
	Region            string             `json:"regionName"`
	Nodes             []ClusterNodeInput `json:"nodes"`
	IsHighlyAvailable bool               `json:"isHighlyAvailable"`
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
	IsHighlyAvailable bool        `json:"isHighlyAvailable"`
}

type ClusterList struct {
	Clusters []*Cluster `json:"content"`
	*SortAndPageable
}

type NodeList struct {
	Nodes []*Node `json:"content"`
	*SortAndPageable
}

type ClusterServiceClient struct {
	client *Client
}

type UserServiceAccount struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	SubjectID string    `json:"subjectId"`
	APIKeyID  string    `json:"apiKeyId"`
	Type      string    `json:"type"`
}

type ServiceAccountInput struct {
	SubjectId string `json:"subjectId"`
}

type ClusterIdentity struct {
	CertificatePem                 string `json:"certificatePem"`
	PrivateKeyPem                  string `json:"privateKeyPem"`
	ExpiresAtEpochSecond           int    `json:"expiresAtEpochSecond"`
	ClusterCertificateAuthorityPem string `json:"clusterCertificateAuthorityPem"`
}

func (c *ClusterServiceClient) List(maxSize int, page int) (*ClusterList, error) {

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

func (c *ClusterServiceClient) Describe(clusterName string) (*Cluster, error) {
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

func (c *ClusterServiceClient) DescribeById(id string) (*Cluster, error) {
	var cluster *Cluster

	err := c.client.
		Call(fmt.Sprintf("rest/v1/cluster/by-id/%s", id),
			"Get",
			&cluster)

	if err != nil {
		return nil, err
	}

	return cluster, nil
}

func (c *ClusterServiceClient) Create(input *ClusterInput) (*Cluster, error) {
	var cluster *Cluster

	err := c.client.
		Call("rest/v1/cluster",
			"Post",
			&cluster,
			WithBody(input))

	if err != nil {
		return nil, err
	}

	return cluster, nil
}

func (c *ClusterServiceClient) Delete(clusterName string) error {

	err := c.client.
		Call(fmt.Sprintf("rest/v1/cluster/%s", clusterName),
			"Delete",
			nil)

	if err != nil {
		return err
	}

	return nil

}

func (c *ClusterServiceClient) ListNodes(clusterName string) (*NodeList, error) {
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

type ServiceAccount struct {
	ID                          string `json:"id"`
	KubeConfig                  string `json:"kubeConfig"`
	ServiceAccountToken         string `json:"serviceAccountToken"`
	ClusterCertificateAuthority string `json:"clusterCertificateAuthority"`
}

func (c *ClusterServiceClient) CreateServiceAccountForSelf(clusterName string) (*ServiceAccount, error) {
	var serviceAccount *ServiceAccount

	err := c.client.
		Call(fmt.Sprintf("rest/v1/cluster/%s/user-service-account", clusterName),
			"Post",
			&serviceAccount)

	if err != nil {
		return nil, err
	}

	return serviceAccount, nil
}

func (c *ClusterServiceClient) CreateServiceAccount(clusterName string, subjectId string) (*ServiceAccount, error) {
	var serviceAccount *ServiceAccount

	err := c.client.
		Call(fmt.Sprintf("rest/v1/cluster/%s/user-service-account/%s", clusterName, subjectId),
			"Post",
			&serviceAccount)

	if err != nil {
		return nil, err
	}

	return serviceAccount, nil
}

func (c *ClusterServiceClient) GetServiceAccount(clusterName string, serviceAccountId string) (*ServiceAccount, error) {
	var serviceAccount *ServiceAccount

	err := c.client.
		Call(fmt.Sprintf("rest/v1/cluster/%s/user-service-account/%s", clusterName, serviceAccountId),
			"Get",
			&serviceAccount)

	if err != nil {
		return nil, err
	}

	return serviceAccount, nil
}

func (c *ClusterServiceClient) DeleteServiceAccount(clusterName string, serviceAccountId string) error {
	err := c.client.
		Call(fmt.Sprintf("rest/v1/cluster/%s/user-service-account/%s", clusterName, serviceAccountId),
			"Delete",
			nil)

	if err != nil {
		return err
	}
	return nil
}

func (c *ClusterServiceClient) ListUserServiceAccounts(clusterName string) ([]*UserServiceAccount, error) {
	var userServiceAccount []*UserServiceAccount

	err := c.client.
		Call(fmt.Sprintf("rest/v1/cluster/%s/user-service-account", clusterName),
			"Get",
			&userServiceAccount)

	if err != nil {
		return nil, err
	}

	return userServiceAccount, nil
}

func (c *ClusterServiceClient) GetIdentity(clusterName string) (*ClusterIdentity, error) {
	var clusterIdentity *ClusterIdentity

	err := c.client.
		Call(fmt.Sprintf("rest/v1/cluster/%s/identity", clusterName),
			"Get",
			&clusterIdentity)

	if err != nil {
		return nil, err
	}

	return clusterIdentity, nil
}
