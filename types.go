package main

type IdentityProvider interface {
	GetID() string
}

type EmailProvider interface {
	GetEmail() string
}

type AccessProvider interface {
	HasAccess(scope string) bool
}
