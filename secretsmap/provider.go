package secretsmap

import (
	"context"
	"fmt"

	config_secrets "github.com/UiP9AV6Y/go-config-secrets"
)

// NewProvider is a secret manager which proxies requests to the relevant provider based on
// the requested reference.
func NewProvider(refs map[string]config_secrets.SecretProvider) config_secrets.SecretManager {
	fetch := func(ctx context.Context, ref string) (string, error) {
		if v, ok := refs[ref]; ok {
			return v.Fetch(ctx)
		}

		return "", fmt.Errorf("secret map does not contain a provider for %q", ref)
	}

	return config_secrets.SecretManagerFunc(fetch)
}
