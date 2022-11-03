package symbiosis

import (
	"encoding/json"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const projectsJson = `
[
	{
		"id": "aaaaaaa-edb4-4e74-97d1-76287c25b898",
		"name": "test-example",
		"owner": "symbiosis",
		"repository": "cli-example",
		"configuration": {
			"productionBranch": "main"
		},
		"productionClusterId": null
	}
]
`

const projectJson = `
{
	"id": "aaaaaaa-edb4-4e74-97d1-76287c25b898",
	"name": "test-example",
	"owner": "symbiosis",
	"repository": "cli-example",
	"configuration": {
		"productionBranch": "main"
	},
	"productionClusterId": null
}`

func TestListProjects(t *testing.T) {
	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/project"

	var fakeProjects []*Project
	json.Unmarshal([]byte(projectsJson), &fakeProjects)

	responder := httpmock.NewStringResponder(200, projectsJson)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	projects, err := c.Project.List()

	assert.Nil(t, err)
	assert.Equal(t, projects, fakeProjects)
	assert.Equal(t, projects[0].ID, "aaaaaaa-edb4-4e74-97d1-76287c25b898")
	assert.Equal(t, projects[0].Name, "test-example")

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	_, err = c.Project.List()

	assert.Error(t, err)
}

func TestDescribeProject(t *testing.T) {
	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/project/test-example"

	var fakeProject *Project
	json.Unmarshal([]byte(projectJson), &fakeProject)

	responder := httpmock.NewStringResponder(200, projectJson)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	project, err := c.Project.Describe("test-example")

	assert.Nil(t, err)
	assert.Equal(t, fakeProject, project)
	assert.Equal(t, fakeProject.Name, "test-example")
	assert.Equal(t, fakeProject.ID, "aaaaaaa-edb4-4e74-97d1-76287c25b898")
	assert.Equal(t, fakeProject.Repository, "cli-example")

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	_, err = c.Project.Describe("test-example")

	assert.Error(t, err)
}
