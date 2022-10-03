package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestImplementations(t *testing.T) {

	t.Run("SignedInUser",
		scenario(
			&SignedInUser{
				id:    "123",
				email: "example@example.com",
			}),
	)
	t.Run("AnonymousUser", scenario(&AnonymousUser{}))
	t.Run("PluginInstance", scenario(&PluginInstance{id: "abcdef"}))
}

func scenario(user any) func(t *testing.T) {
	ctx := context.Background()
	ctx = setIdentity(ctx, user)
	access := user.(AccessProvider)

	return func(t *testing.T) {
		t.Run("HasAccess", func(t *testing.T) {
			require.Equal(t, access.HasAccess("user:id"), assertAccess(ctx, "user:id") == nil)
			require.Equal(t, access.HasAccess("user:email"), assertAccess(ctx, "user:email") == nil)
			require.Equal(t, access.HasAccess("time"), assertAccess(ctx, "time") == nil)
		})

		t.Run("GetID", func(t *testing.T) {
			id, err := getID(ctx)
			expected, ok := user.(IdentityProvider)
			if ok && access.HasAccess("user:id") {
				require.NoError(t, err)
				require.Equal(t, expected.GetID(), id)
			} else {
				require.Error(t, err)
				require.Empty(t, id)
			}
		})

		t.Run("GetEmail", func(t *testing.T) {
			email, err := getEmail(ctx)
			expected, ok := user.(EmailProvider)
			if ok && access.HasAccess("user:email") {
				require.NoError(t, err)
				require.Equal(t, expected.GetEmail(), email)
			} else {
				require.Error(t, err)
				require.Empty(t, email)
			}
		})

	}
}
