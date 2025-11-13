package config_secrets

import (
	"context"
)

// SecretManager describes a system capable of retrieving secret values.
type SecretManager interface {
	// Fetch retrieves the secret value associated with the given reference.
	Fetch(context.Context, string) (string, error)
}

// SecretManagerFunc is an adapter to allow the use of ordinary functions as secret managers.
// If f is a function with the appropriate signature, SecretManagerFunc(f) is a SecretManager that calls f.
type SecretManagerFunc func(context.Context, string) (string, error)

// Fetch calls f(ctx, ref).
func (f SecretManagerFunc) Fetch(ctx context.Context, ref string) (string, error) {
	return f(ctx, ref)
}
