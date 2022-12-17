package symbiosis

import (
	"fmt"
)

type SecretService interface {
	Create(project string, secretKey string, input Secret) error
	GetSecretsByProject(project string) (SecretCollection, error)
	GetSecretsByProjectAndEnvironment(project string, environment ProjectEnvironment) (map[string]string, error)
}

type ProjectEnvironment string

const (
	ENVIRONMENT_DEVELOPMENT ProjectEnvironment = "development"
	ENVIRONMENT_PREVIEW     ProjectEnvironment = "preview"
	ENVIRONMENT_PRODUCTION  ProjectEnvironment = "production"
)

type Secret struct {
	DevelopmentValue string `json:"developmentValue"`
	PreviewValue     string `json:"previewValue"`
	ProductionValue  string `json:"productionValue"`
}

type SecretCollection map[string]*Secret

type SecretServiceClient struct {
	client *Client
}

func (n *SecretServiceClient) Create(project string, secretKey string, input Secret) error {

	err := n.client.Call(
		fmt.Sprintf("/rest/v1/project/%s/secret/%s", project, secretKey),
		"Put",
		nil,
		WithBody(input),
	)

	if err != nil {
		return err
	}

	return nil
}

func (n *SecretServiceClient) GetSecretsByProject(project string) (SecretCollection, error) {
	var secrets *SecretCollection

	err := n.client.
		Call(fmt.Sprintf("/rest/v1/project/%s/secret", project),
			"Get",
			&secrets)

	if err != nil {
		return nil, err
	}

	return *secrets, nil
}

func (n *SecretServiceClient) GetSecretsByProjectAndEnvironment(project string, environment ProjectEnvironment) (map[string]string, error) {
	var secrets *map[string]string

	err := n.client.
		Call(fmt.Sprintf("/rest/v1/project/%s/secret?environment=%s", project, environment),
			"Get",
			&secrets)

	if err != nil {
		return nil, err
	}

	return *secrets, nil
}
