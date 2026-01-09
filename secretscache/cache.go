package secretscache

import (
	"context"
	"time"

	"github.com/waylen888/lazymap"

	config_secrets "github.com/UiP9AV6Y/go-config-secrets"
)

func New(sec config_secrets.SecretManager, ttl time.Duration) config_secrets.SecretManager {
	cache := lazymap.New[string, string](ttl)
	fetch := func(ctx context.Context, ref string) (string, error) {
		return cache.LoadOrCtor(ctx, ref, sec.Fetch)
	}

	return config_secrets.SecretManagerFunc(fetch)
}
