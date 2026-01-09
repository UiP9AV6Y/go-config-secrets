package secretsmap

import (
	"context"
	"fmt"

	config_secrets "github.com/UiP9AV6Y/go-config-secrets"
)

// New is a secret manager which provides lookup results from the given map.
func New(refs map[string]string) config_secrets.SecretManager {
	fetch := func(_ context.Context, ref string) (string, error) {
		if v, ok := refs[ref]; ok {
			return v, nil
		}

		return "", config_secrets.NewSecretNotFoundError(ref)
	}

	return config_secrets.SecretManagerFunc(fetch)
}
