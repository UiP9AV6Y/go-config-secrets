package secretsio

import (
	"bytes"
	"context"
	"io"

	config_secrets "github.com/UiP9AV6Y/go-config-secrets"
)

// New creates a secret manager which delegates the data retrieval to
// the given provider. The resulting data is stripped of any leading and
// trailing whitespaces.
func New(rp ReaderProvider) config_secrets.SecretManager {
	fetch := func(ctx context.Context, ref string) (string, error) {
		r, err := rp.Open(ctx, ref)
		if err != nil {
			return "", err
		}
		defer r.Close()

		b, err := io.ReadAll(r)
		if err != nil {
			return "", err
		}

		return string(bytes.TrimSpace(b)), nil
	}

	return config_secrets.SecretManagerFunc(fetch)
}
