package main

type AnonymousUser struct{}

func (u *AnonymousUser) HasAccess(scope string) bool {
	if scope == "time" {
		return true
	}
	return false
}
