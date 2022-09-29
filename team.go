package symbiosis

import (
	"encoding/json"
	"errors"
	"fmt"
)

type TeamService interface {
	GetMemberByEmail(email string) (*TeamMember, error)
	GetInvitationByEmail(email string) (*TeamMember, error)
	InviteMembers(emails []string, role string) ([]*TeamMember, error)
	DeleteMember(email string) error
	ChangeRole(email string, role string) error
}

type TeamMember struct {
	Email  string `json:"email"`
	TeamId string `json:"teamId"`
	Role   string `json:"role"`
}

type Invite struct {
	Emails []string `json:"emails"`
	Role   string   `json:"role"`
}

const (

	// Has full access, can create teams
	RoleOwner = "OWNER"

	// Has full access within a team, can invite other team members
	RoleAdmin = "ADMIN"

	// Can manage resources in clusters but not create them
	RoleMember = "MEMBER"
)

type TeamServiceClient struct {
	client *Client
}

func GetValidRoles() map[string]bool {
	return map[string]bool{RoleOwner: true, RoleAdmin: true, RoleMember: true}
}

func ValidateRole(role string) error {
	validRoles := GetValidRoles()

	if _, ok := validRoles[role]; !ok {
		return errors.New("Invalid role given")
	}

	return nil
}

func (t *TeamServiceClient) GetMemberByEmail(email string) (*TeamMember, error) {
	var member *TeamMember

	err := t.client.
		Call(fmt.Sprintf("rest/v1/team/member/%s", email),
			"Get",
			&member)

	if err != nil {
		return nil, err
	}

	return member, nil
}

func (t *TeamServiceClient) GetInvitationByEmail(email string) (*TeamMember, error) {

	var member *TeamMember

	err := t.client.
		Call(fmt.Sprintf("rest/v1/team/member/invite/%s", email),
			"Get",
			&member)

	if err != nil {
		return nil, err
	}

	return member, nil
}

func (t *TeamServiceClient) InviteMembers(emails []string, role string) ([]*TeamMember, error) {
	err := ValidateRole(role)

	if err != nil {
		return nil, err
	}

	var teamMembers []*TeamMember

	body, _ := json.Marshal(Invite{emails, role})

	err = t.client.
		Call(fmt.Sprintf("rest/v1/team/member/invite"),
			"Post",
			&teamMembers,
			WithBody(body))

	if err != nil {
		return nil, err
	}

	return teamMembers, nil
}

func (t *TeamServiceClient) DeleteMember(email string) error {
	err := t.client.
		Call(fmt.Sprintf("rest/v1/team/member/%s", email),
			"Delete",
			nil)

	if err != nil {
		return err
	}

	return nil

}

func (t *TeamServiceClient) ChangeRole(email string, role string) error {

	err := ValidateRole(role)

	if err != nil {
		return err
	}

	err = t.client.
		Call(fmt.Sprintf("rest/v1/team/member/%s", email),
			"Put",
			nil,
			WithBody([]byte(fmt.Sprintf(`{"role":"%s"}`, role))))

	if err != nil {
		return err
	}

	return nil
}
