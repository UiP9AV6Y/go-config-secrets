package secretstest

import (
	"context"

	config_secrets "github.com/UiP9AV6Y/go-config-secrets"
)

// NewStaticManager creates a [config_secrets.SecretManager] which
// always returns the given value as lookup result.
func NewStaticManager(v string) config_secrets.SecretManager {
	fetch := func(_ context.Context, _ string) (string, error) {
		return v, nil
	}

	return config_secrets.SecretManagerFunc(fetch)
}
