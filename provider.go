package config_secrets

import (
	"context"
)

// SecretProvider represents a value holder which can deliver its data on demand.
// It is loosly modeled after [github.com/prometheus/common/config#SecretReader].
type SecretProvider interface {
	Fetch(context.Context) (string, error)
}

// SecretProviderFunc is an adapter to allow the use of ordinary functions as secret providers.
// If f is a function with the appropriate signature, SecretProviderFunc(f) is a SecretProvider that calls f.
type SecretProviderFunc func(context.Context) (string, error)

// Fetch calls f(ctx).
func (f SecretProviderFunc) Fetch(ctx context.Context) (string, error) {
	return f(ctx)
}
