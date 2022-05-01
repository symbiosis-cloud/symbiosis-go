package symbiosis

import (
	"encoding/json"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

const teamMemberJSON = `
{
	"email": "test@tca0.nl",
	"teamId": "uuid",
	"role": "MEMBER"
}`

const teamMembersJSON = `[ ` + teamMemberJSON + ` ]`

func TestGetValidRoles(t *testing.T) {
	validRoles := map[string]bool{RoleOwner: true, RoleAdmin: true, RoleMember: true}

	assert.Equal(t, validRoles, GetValidRoles())
}

func TestGetMemberByEmail(t *testing.T) {
	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/team/member/test@tca0.nl"

	var fakeTeamMember *TeamMember
	json.Unmarshal([]byte(teamMemberJSON), &fakeTeamMember)

	responder := httpmock.NewStringResponder(200, teamMemberJSON)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	teamMember, err := c.Team.GetMemberByEmail("test@tca0.nl")

	assert.Nil(t, err)
	assert.Equal(t, fakeTeamMember, teamMember)

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	_, err = c.Team.GetMemberByEmail("test@tca0.nl")
	assert.Error(t, err)
}

func TestGetInvitationByEmail(t *testing.T) {
	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/team/member/invite/test@tca0.nl"

	var fakeTeamMember *TeamMember
	json.Unmarshal([]byte(teamMemberJSON), &fakeTeamMember)

	responder := httpmock.NewStringResponder(200, teamMemberJSON)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	teamMember, err := c.Team.GetInvitationByEmail("test@tca0.nl")

	assert.Nil(t, err)
	assert.Equal(t, fakeTeamMember, teamMember)

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("GET", fakeURL, responder)

	_, err = c.Team.GetInvitationByEmail("test@tca0.nl")
	assert.Error(t, err)
}

func TestInviteMembers(t *testing.T) {
	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/team/member/invite"

	var fakeTeamMembers []*TeamMember
	json.Unmarshal([]byte(teamMembersJSON), &fakeTeamMembers)

	responder := httpmock.NewStringResponder(200, teamMembersJSON)
	httpmock.RegisterResponder("POST", fakeURL, responder)

	teamMembers, err := c.Team.InviteMembers([]string{"test@tca0.nl"}, RoleAdmin)

	assert.Nil(t, err)
	assert.Equal(t, fakeTeamMembers, teamMembers)

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("POST", fakeURL, responder)

	_, err = c.Team.InviteMembers([]string{"test@tca0.nl"}, RoleAdmin)
	assert.Error(t, err)

	// test invalid role
	_, err = c.Team.InviteMembers([]string{"test@tca0.nl"}, "Invalid")
	assert.ErrorContains(t, err, "Invalid role given")
}

func TestDeleteMember(t *testing.T) {
	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/team/member/test@tca0.nl"

	responder := httpmock.NewStringResponder(200, "")
	httpmock.RegisterResponder("DELETE", fakeURL, responder)

	err := c.Team.DeleteMember("test@tca0.nl")

	assert.Nil(t, err)

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("DELETE", fakeURL, responder)

	err = c.Team.DeleteMember("test@tca0.nl")
	assert.Error(t, err)
}

func TestChangeRole(t *testing.T) {
	c := getMocketClient()
	defer httpmock.DeactivateAndReset()

	fakeURL := "/rest/v1/team/member/test@tca0.nl"

	responder := httpmock.NewStringResponder(200, "false")
	httpmock.RegisterResponder("PUT", fakeURL, responder)

	err := c.Team.ChangeRole("test@tca0.nl", RoleAdmin)

	assert.Nil(t, err)

	responder = httpmock.NewErrorResponder(assert.AnError)
	httpmock.RegisterResponder("PUT", fakeURL, responder)

	err = c.Team.ChangeRole("test@tca0.nl", RoleAdmin)
	assert.Error(t, err)

	// test invalid role
	err = c.Team.ChangeRole("test@tca0.nl", "Invalid")
	assert.ErrorContains(t, err, "Invalid role given")
}
