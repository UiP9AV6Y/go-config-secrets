package secretscache_test

import (
	"context"
	"testing"
	"time"

	"gotest.tools/v3/assert"

	config_secrets "github.com/UiP9AV6Y/go-config-secrets"
	"github.com/UiP9AV6Y/go-config-secrets/secretscache"
)

func TestNew(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	var gotValue string
	var gotErr error
	mock := secretsManagerMock("v")
	subject := secretscache.New(mock, 1*time.Second)
	ctx := context.Background()

	gotValue, gotErr = subject.Fetch(ctx, "miss")
	assert.ErrorIs(t, gotErr, config_secrets.ErrNotFound)

	gotValue, gotErr = subject.Fetch(ctx, "match")
	assert.NilError(t, gotErr)
	assert.Equal(t, gotValue, "vv")

	gotValue, gotErr = subject.Fetch(ctx, "miss")
	assert.ErrorIs(t, gotErr, config_secrets.ErrNotFound)

	gotValue, gotErr = subject.Fetch(ctx, "match")
	assert.NilError(t, gotErr)
	assert.Equal(t, gotValue, "vv")

	time.Sleep(2 * time.Second)

	gotValue, gotErr = subject.Fetch(ctx, "miss")
	assert.ErrorIs(t, gotErr, config_secrets.ErrNotFound)

	gotValue, gotErr = subject.Fetch(ctx, "match")
	assert.NilError(t, gotErr)
	assert.Equal(t, gotValue, "vvv")
}

func secretsManagerMock(value string) config_secrets.SecretManager {
	fetch := func(_ context.Context, ref string) (string, error) {
		if ref != "match" {
			return "", config_secrets.NewSecretNotFoundError(ref)
		}

		value = value + "v"

		return value, nil
	}

	return config_secrets.SecretManagerFunc(fetch)
}
