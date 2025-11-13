package config_secrets

import (
	"context"
	"fmt"
	"os"
	"sync/atomic"
)

const EnvPrefix = "SECRET_"

var defaultManager atomic.Value

func init() {
	// reimplementation of secretsfunc to avoid cyclic dependencies
	fetch := func(_ context.Context, ref string) (string, error) {
		k := EnvPrefix + ref
		if v := os.Getenv(k); v != "" {
			return v, nil
		}

		return "", fmt.Errorf("no environment secret %q available", k)
	}

	defaultManager.Store(SecretManagerFunc(fetch))
}

// Default returns the default SecretManager.
func Default() SecretManager {
	return defaultManager.Load().(SecretManager)
}

// SetDefault makes sec the default SecretManager.
func SetDefault(sec SecretManager) {
	defaultManager.Store(sec)
}
