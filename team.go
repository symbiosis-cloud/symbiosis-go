package symbiosis

import (
	"encoding/json"
	"errors"
	"fmt"
)

type TeamMember struct {
	Email  string `json:"email"`
	TeamId string `json:"teamId"`
	Role   string `json:"role"`
	client *Client
}

type Invite struct {
	Emails []string `json:"emails"`
	Role   string   `json:"role"`
}

const (
	RoleCluster = "CLUSTER"
	RoleOwner   = "OWNER"
	RoleAdmin   = "ADMIN"
	RoleMember  = "MEMBER"
)

func GetValidRoles() map[string]bool {
	return map[string]bool{RoleCluster: true, RoleOwner: true, RoleAdmin: true, RoleMember: true}
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

	result.client = c

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

func (c *Client) InviteTeamMembers(emails []string, role string) ([]*TeamMember, error) {
	validRoles := GetValidRoles()

	if _, ok := validRoles[role]; !ok {
		return nil, errors.New("Invalid role given")
	}

	var result []*TeamMember

	body, err := json.Marshal(Invite{emails, role})

	if err != nil {
		return nil, errors.New("Failed to create invite")
	}

	resp, err := c.symbiosisAPI.R().
		SetResult(&result).
		ForceContentType("application/json").
		SetBody(body).
		Post("rest/v1/team/member/invite")

	if err != nil {
		return nil, err
	}

	validated, err := c.ValidateResponse(resp, result)

	if err != nil {
		return nil, err
	}

	return validated.([]*TeamMember), nil
}

func (t *TeamMember) Delete() error {
	resp, err := t.client.symbiosisAPI.R().
		ForceContentType("application/json").
		Delete(fmt.Sprintf("rest/v1/team/member/%s", t.Email))

	if err != nil {
		return err
	}

	_, err = t.client.ValidateResponse(resp, nil)

	if err != nil {
		return err
	}

	return nil
}

func (t *TeamMember) ChangeRole(role string) (*TeamMember, error) {
	var result *TeamMember

	resp, err := t.client.symbiosisAPI.R().
		SetResult(&result).
		ForceContentType("application/json").
		SetBody([]byte(fmt.Sprintf(`{"role":"%s"}`, role))).
		Put(fmt.Sprintf("rest/v1/team/member/%s", t.Email))

	if err != nil {
		return nil, err
	}

	validated, err := t.client.ValidateResponse(resp, result)

	if err != nil {
		return nil, err
	}

	return validated.(*TeamMember), nil
}
