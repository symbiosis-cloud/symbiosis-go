package symbiosis

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const clusterJSON = `
{
	"id": "42d712ab-d05b-46d6-99e2-fef2059e09ac",
	"name": "test",
	"kubeVersion": "1.23.5",
	"apiServerEndpoint": "does.not.matter",
	"state": "ACTIVE",
	"nodes": [
	  {
		"id": "19fc97a5-30de-4f35-8b28-a14bf640b557",
		"name": "random-0",
		"nodeType": {
		  "id": "2f8d2c39-23cd-4623-b816-9489a26c1b8d",
		  "name": "general-int-1",
		  "memoryMi": 1024,
		  "storageGi": 1,
		  "vcpu": 1,
		  "product": {
			"productCosts": [
			  {
				"currency": "USD",
				"unitCost": 5
			  }
			]
		  }
		},
		"region": {
		  "id": "9d420d32-31f8-4a4a-b790-d75ea58ffc08",
		  "name": "netherlands-1"
		},
		"privateIPv4Address": "10.0.0.1",
		"state": "ACTIVE"
	  }]
}
`

const serviceAccountJson = `
{
	"id": "test",
	"kubeConfig": "test",
	"serviceAccountToken": "test",
	"clusterCertificateAuthority": "test"
}
`

const userServiceAccountJson = `
[
	{
		"id": "test",
		"createdAt": "2022-06-20T13:39:21.242323Z",
		"subjectId": "test",
		"apiKeyId": "test",
		"type": "api-key"
	}
]
`

const clusterListJSON = `{ "content": [` + clusterJSON + `], ` + sortableJSON + ` }`
const nodeListJSON = `{ "content": [` + nodeJSON + `], ` + sortableJSON + ` }`

func TestListCluster(t *testing.T) {

	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/cluster?size=10&page=0"

	var fakeClusterList *ClusterList
	json.Unmarshal([]byte(clusterListJSON), &fakeClusterList)

	responder := httpmock.NewStringResponder(200, clusterListJSON)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	clusterList, err := c.Cluster.List(10, 0)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, fakeClusterList, clusterList)

	// test resty failure
	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	_, err = c.Cluster.List(10, 0)
	assert.Error(t, err)
}

func TestDescribeCluster(t *testing.T) {
	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/cluster/test"

	var fakeCluster *Cluster
	json.Unmarshal([]byte(clusterJSON), &fakeCluster)

	responder := httpmock.NewStringResponder(200, clusterJSON)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	cluster, err := c.Cluster.Describe("test")

	assert.Nil(t, err)
	assert.Equal(t, fakeCluster, cluster)

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	_, err = c.Cluster.Describe("test")
	assert.Error(t, err)

}

func TestCreateCluster(t *testing.T) {
	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/cluster"

	var fakeCluster *Cluster
	json.Unmarshal([]byte(clusterJSON), &fakeCluster)

	responder := httpmock.NewStringResponder(200, clusterJSON)
	httpmock.RegisterResponder("POST", fakeURL, responder)

	ClusterInput := &ClusterInput{
		Name: "test",
		Nodes: []ClusterNodeInput{
			{
				Quantity: 1,
				NodeType: "general-int-1",
			},
		},
		KubeVersion: "1.23.5",
		Region:      "germany-1",
		Configuration: ClusterConfigurationInput{
			EnableNginxIngress: false,
		},
	}

	cluster, err := c.Cluster.Create(ClusterInput)

	assert.Nil(t, err)
	assert.Equal(t, fakeCluster, cluster)
	assert.Equal(t, ClusterInput.Name, fakeCluster.Name)

	for _, node := range fakeCluster.Nodes {
		assert.Equal(t, node.NodeType.Name, ClusterInput.Nodes[0].NodeType)
	}

	assert.Equal(t, len(fakeCluster.Nodes), len(ClusterInput.Nodes))

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("POST", fakeURL, responder)

	_, err = c.Cluster.Create(ClusterInput)
	assert.Error(t, err)
}

func TestDeleteCluster(t *testing.T) {
	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/cluster/test"

	responder := httpmock.NewStringResponder(200, "")
	httpmock.RegisterResponder("DELETE", fakeURL, responder)

	err := c.Cluster.Delete("test")

	assert.Nil(t, err)

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("DELETE", fakeURL, responder)

	err = c.Cluster.Delete("test")
	assert.Error(t, err)
}

func TestListNodes(t *testing.T) {
	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/cluster/test/node"

	var fakeNodeList *NodeList
	json.Unmarshal([]byte(nodeListJSON), &fakeNodeList)

	responder := httpmock.NewStringResponder(200, nodeListJSON)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	nodeList, err := c.Cluster.ListNodes("test")

	assert.Nil(t, err)
	assert.Equal(t, fakeNodeList, nodeList)

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	_, err = c.Cluster.ListNodes("test")
	assert.Error(t, err)
}

func TestCreateServiceAccountForSelf(t *testing.T) {
	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/cluster/test/user-service-account"

	var fakeServiceAccount *ServiceAccount
	json.Unmarshal([]byte(serviceAccountJson), &fakeServiceAccount)

	responder := httpmock.NewStringResponder(200, serviceAccountJson)
	httpmock.RegisterResponder("POST", fakeURL, responder)

	serviceAccount, err := c.Cluster.CreateServiceAccountForSelf("test")

	assert.Nil(t, err)
	assert.Equal(t, fakeServiceAccount, serviceAccount)
	assert.Equal(t, "test", serviceAccount.ClusterCertificateAuthority)
	assert.Equal(t, "test", serviceAccount.ID)
	assert.Equal(t, "test", serviceAccount.KubeConfig)
	assert.Equal(t, "test", serviceAccount.ServiceAccountToken)

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("POST", fakeURL, responder)

	_, err = c.Cluster.CreateServiceAccountForSelf("test")
	assert.Error(t, err)
}

func TestCreateServiceAccount(t *testing.T) {
	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/cluster/test/user-service-account/test"

	var fakeServiceAccount *ServiceAccount
	json.Unmarshal([]byte(serviceAccountJson), &fakeServiceAccount)

	responder := httpmock.NewStringResponder(200, serviceAccountJson)
	httpmock.RegisterResponder("POST", fakeURL, responder)

	serviceAccount, err := c.Cluster.CreateServiceAccount("test", "test")

	assert.Nil(t, err)
	assert.Equal(t, fakeServiceAccount, serviceAccount)
	assert.Equal(t, "test", serviceAccount.ClusterCertificateAuthority)
	assert.Equal(t, "test", serviceAccount.ID)
	assert.Equal(t, "test", serviceAccount.KubeConfig)
	assert.Equal(t, "test", serviceAccount.ServiceAccountToken)

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("POST", fakeURL, responder)

	_, err = c.Cluster.CreateServiceAccount("test", "test")
	assert.Error(t, err)
}

func TestGetServiceAccount(t *testing.T) {
	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/cluster/test/user-service-account/test"

	var fakeServiceAccount *ServiceAccount
	json.Unmarshal([]byte(serviceAccountJson), &fakeServiceAccount)

	responder := httpmock.NewStringResponder(200, serviceAccountJson)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	serviceAccount, err := c.Cluster.GetServiceAccount("test", "test")

	assert.Nil(t, err)
	assert.Equal(t, fakeServiceAccount, serviceAccount)
	assert.Equal(t, "test", serviceAccount.ClusterCertificateAuthority)
	assert.Equal(t, "test", serviceAccount.ID)
	assert.Equal(t, "test", serviceAccount.KubeConfig)
	assert.Equal(t, "test", serviceAccount.ServiceAccountToken)

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	_, err = c.Cluster.GetServiceAccount("test", "test")
	assert.Error(t, err)
}

func TestListUserServiceAccount(t *testing.T) {
	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/cluster/test/user-service-account"

	var fakeUserServiceAccount []*UserServiceAccount
	json.Unmarshal([]byte(userServiceAccountJson), &fakeUserServiceAccount)

	responder := httpmock.NewStringResponder(200, userServiceAccountJson)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	userServiceAccounts, err := c.Cluster.ListUserServiceAccounts("test")

	assert.Nil(t, err)

	assert.Equal(t, fakeUserServiceAccount, userServiceAccounts)

	for _, u := range userServiceAccounts {
		assert.Equal(t, "test", u.APIKeyID)
		assert.Equal(t, time.Date(2022, time.June, 20, 13, 39, 21, 242323000, time.UTC), u.CreatedAt)
		assert.Equal(t, "test", u.ID)
		assert.Equal(t, "test", u.SubjectID)
		assert.Equal(t, "api-key", u.Type)
	}

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	_, err = c.Cluster.ListUserServiceAccounts("test")
	assert.Error(t, err)
}

func TestDeleteServiceAccount(t *testing.T) {
	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/cluster/test/user-service-account/test"

	responder := httpmock.NewStringResponder(200, "")
	httpmock.RegisterResponder("DELETE", fakeURL, responder)

	err := c.Cluster.DeleteServiceAccount("test", "test")

	assert.Nil(t, err)

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("DELETE", fakeURL, responder)

	err = c.Cluster.DeleteServiceAccount("test", "test")
	assert.Error(t, err)
}
