package symbiosis

import (
	"fmt"
	"time"
)

type ApiKey struct {
	ID          string    `json:"id"`
	Token       string    `json:"token"`
	SubjectID   string    `json:"subjectId"`
	Role        UserRole  `json:"role"`
	Description string    `json:"description"`
	LastUsedAt  time.Time `json:"lastUsedAt"`
}

type ApiKeyService interface {
	Create(input ApiKeyInput) (*ApiKey, error)
	List() (ApiKeyCollection, error)
	Delete(id string) error
}

type ApiKeyInput struct {
	Role        UserRole `json:"role"`
	Description string   `json:"description"`
}

type ApiKeyCollection []*ApiKey

type ApiKeyServiceClient struct {
	client *Client
}

func (n *ApiKeyServiceClient) Create(input ApiKeyInput) (*ApiKey, error) {
	var apiKey *ApiKey

	err := n.client.Call(
		"/rest/v1/api-keys",
		"Post",
		&apiKey,
		WithBody(input),
	)

	if err != nil {
		return nil, err
	}

	return apiKey, nil
}

func (n *ApiKeyServiceClient) List() (ApiKeyCollection, error) {
	var apiKeys ApiKeyCollection

	err := n.client.
		Call("/rest/v1/api-keys",
			"Get",
			&apiKeys)

	if err != nil {
		return nil, err
	}

	return apiKeys, nil
}

func (n *ApiKeyServiceClient) Delete(id string) error {
	err := n.client.
		Call(fmt.Sprintf("/rest/v1/api-keys/%s", id),
			"Delete",
			nil)

	if err != nil {
		return err
	}

	return nil
}
