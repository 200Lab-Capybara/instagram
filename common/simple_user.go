package common

import "github.com/google/uuid"

const RequesterKey = "RequesterKey"

type SimpleUser struct {
	UserId    uuid.UUID `json:"user_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Role      string    `json:"role"`
}

type Requester interface {
	UserId() uuid.UUID
	FirstName() string
	LastName() string
	Role() string
	Status() string
	GetSimpleUser() *SimpleUser
}

type requesterData struct {
	userId    uuid.UUID
	firstName string
	lastName  string
	role      string
	status    string
}

func (r *requesterData) UserId() uuid.UUID {
	return r.userId
}
func (r *requesterData) FirstName() string { return r.firstName }
func (r *requesterData) LastName() string  { return r.lastName }
func (r *requesterData) Role() string      { return r.role }
func (r *requesterData) Status() string    { return r.status }
func (r *requesterData) GetSimpleUser() *SimpleUser {
	return &SimpleUser{
		UserId:    r.userId,
		FirstName: r.firstName,
		LastName:  r.lastName,
		Role:      r.role,
	}
}

func NewRequester(sub uuid.UUID, firstName, lastName, role, status string) Requester {
	return &requesterData{
		userId:    sub,
		firstName: firstName,
		lastName:  lastName,
		role:      role,
		status:    status,
	}
}
