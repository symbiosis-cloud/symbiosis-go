package symbiosis

import (
	"fmt"
)

type TeamMember struct {
	Email string
	Role  string
}

func (c *Client) DescribeTeamMember(email string) (*TeamMember, error) {
	var result *TeamMember

	resp, err := c.symbiosisAPI.R().
		SetResult(&result).
		ForceContentType("application/json").
		Get(fmt.Sprintf("rest/v1/team/member/%s", email))

	if err != nil {
		return nil, err
	}

	validated, err := c.ValidateResponse(resp, result)

	if err != nil {
		return nil, err
	}

	return validated.(*TeamMember), nil
}

func (c *Client) DescribeTeamMemberInvitation(email string) (*TeamMember, error) {
	var result *TeamMember

	resp, err := c.symbiosisAPI.R().
		SetResult(&result).
		ForceContentType("application/json").
		Get(fmt.Sprintf("rest/v1/team/member/invite/%s", email))

	if err != nil {
		return nil, err
	}

	validated, err := c.ValidateResponse(resp, result)

	if err != nil {
		return nil, err
	}

	return validated.(*TeamMember), nil
}
