package symbiosis

import "fmt"

// /rest/v1/github/{owner}/{repository}/project/secret?environment=preview|development|production

type Project struct {
	ID                  string        `json:"id"`
	Name                string        `json:"name"`
	Owner               string        `json:"owner"`
	Repository          string        `json:"repository"`
	Configuration       Configuration `json:"configuration"`
	ProductionClusterID interface{}   `json:"productionClusterId"`
}
type Configuration struct {
	ProductionBranch string `json:"productionBranch"`
}

type ProjectService interface {
	List() ([]*Project, error)
	Describe(projectName string) (*Project, error)
}

type ProjectServiceClient struct {
	client *Client
}

func (n *ProjectServiceClient) List() ([]*Project, error) {

	var projects []*Project

	err := n.client.Call(
		"/rest/v1/project",
		"Get",
		&projects,
	)

	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (n *ProjectServiceClient) Describe(projectName string) (*Project, error) {
	var project *Project

	err := n.client.Call(
		fmt.Sprintf("/rest/v1/project/%s", projectName),
		"Get",
		&project,
	)

	if err != nil {
		return nil, err
	}

	return project, nil
}
