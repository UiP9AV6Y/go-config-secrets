package secretsfunc

import (
	"context"
	"fmt"
	"os"

	config_secrets "github.com/UiP9AV6Y/go-config-secrets"
)

// New creates a [config_secrets.SecretManager] using the provided
// function as value provider. Empty strings are considered lookup
// failures and are reported as such; all other values are returned
// as-is. The given prefix is concatenated with the requested reference
// to form the input parameter to the lookup function.
func New(prefix string, f func(string) string) config_secrets.SecretManager {
	fetch := func(_ context.Context, ref string) (string, error) {
		k := prefix + ref
		if v := f(k); v != "" {
			return v, nil
		}

		return "", config_secrets.NewSecretNotFoundError(k)
	}

	return config_secrets.SecretManagerFunc(fetch)
}

// NewEnv uses [os.Getenv] as the lookup function.
func NewEnv(prefix string) config_secrets.SecretManager {
	return New(prefix, os.Getenv)
}
