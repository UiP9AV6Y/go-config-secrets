package config_secrets

import (
	"context"
)

// RefSecret implements [github.com/prometheus/common/config#SecretReader]
// and delegates the secret resolution to the held manager.
type RefSecret struct {
	ref string
	mgr SecretManager
}

// NewRefSecret creates a new RefSecret instance.
// Using nil as manager will result in the [Default] being used.
func NewRefSecret(ref string, mgr SecretManager) *RefSecret {
	if mgr == nil {
		mgr = Default()
	}

	result := &RefSecret{
		ref: ref,
		mgr: mgr,
	}

	return result
}

// Fetch implements [github.com/prometheus/common/config#SecretReader]
func (s *RefSecret) Fetch(ctx context.Context) (string, error) {
	return s.mgr.Fetch(ctx, s.ref)
}

// Description implements [github.com/prometheus/common/config#SecretReader]
func (s *RefSecret) Description() string {
	return "ref " + s.ref
}

// Immutable implements [github.com/prometheus/common/config#SecretReader]
func (*RefSecret) Immutable() bool {
	return false
}
