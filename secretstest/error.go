package secretstest

import (
	"context"

	config_secrets "github.com/UiP9AV6Y/go-config-secrets"
)

// NewErrorManager creates a [config_secrets.SecretManager] which
// always returns the given error as lookup result.
func NewErrorManager(err error) config_secrets.SecretManager {
	fetch := func(_ context.Context, _ string) (string, error) {
		return "", err
	}

	return config_secrets.SecretManagerFunc(fetch)
}
