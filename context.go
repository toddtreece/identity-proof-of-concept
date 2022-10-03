package main

import (
	"context"
	"fmt"
	"time"
)

type identityKey struct{}

func setIdentity(ctx context.Context, v any) context.Context {
	return context.WithValue(ctx, identityKey{}, v)
}

func getIdentity(ctx context.Context) any {
	return ctx.Value(identityKey{})
}

// helper functions
func assertAccess(ctx context.Context, scope string) error {
	identity, ok := getIdentity(ctx).(AccessProvider)
	if !ok {
		return fmt.Errorf("AccessProvider not implemented")
	}

	if !identity.HasAccess(scope) {
		return fmt.Errorf("unauthorized attempt to access (%s)", scope)
	}
	return nil
}

func getEmail(ctx context.Context) (string, error) {
	if err := assertAccess(ctx, "user:email"); err != nil {
		return "", err
	}

	if identity, ok := getIdentity(ctx).(EmailProvider); ok {
		return identity.GetEmail(), nil
	}

	return "", fmt.Errorf("email address unavailable for identity")
}

func getID(ctx context.Context) (string, error) {
	if err := assertAccess(ctx, "user:id"); err != nil {
		return "", err
	}

	if identity, ok := getIdentity(ctx).(IdentityProvider); ok {
		return identity.GetID(), nil
	}

	return "", fmt.Errorf("identity unavailable")
}

func getTime(ctx context.Context) (time.Time, error) {
	if err := assertAccess(ctx, "time"); err != nil {
		return time.Unix(0, 0), err
	}
	return time.Now(), nil
}
