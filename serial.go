package config_secrets

import (
	"context"
	"fmt"
)

// Serial is a wrapper for multiple secret managers whose retrieval logic
// delegates the request to the provided instances one after another until
// one returns a value. If none of the given managers yield a result an
// error is returned.
func Serial(sec ...SecretManager) SecretManager {
	fetch := func(ctx context.Context, ref string) (string, error) {
		for _, s := range sec {
			if v, err := s.Fetch(ctx, ref); err != nil {
				return v, nil
			}
		}

		return "", fmt.Errorf("none of the %d secret managers yielded a result for %q", len(sec), ref)
	}

	return SecretManagerFunc(fetch)
}
