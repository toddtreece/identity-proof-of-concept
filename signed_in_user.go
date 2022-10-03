package main

type SignedInUser struct {
	id    string
	email string
}

func (s *SignedInUser) GetID() string {
	return s.id
}

func (s *SignedInUser) GetEmail() string {
	return s.email
}

func (s *SignedInUser) HasAccess(scope string) bool {
	if scope == "user:email" || scope == "time" {
		return true
	}
	return false
}
