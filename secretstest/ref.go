package secretstest

import (
	"context"

	config_secrets "github.com/UiP9AV6Y/go-config-secrets"
)

// NewRefManager creates a [config_secrets.SecretManager] which
// returns the query value as lookup result.
func NewRefManager() config_secrets.SecretManager {
	fetch := func(_ context.Context, ref string) (string, error) {
		return ref, nil
	}

	return config_secrets.SecretManagerFunc(fetch)
}
