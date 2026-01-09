package secretsdir

import (
	"bytes"
	"context"
	"os"

	config_secrets "github.com/UiP9AV6Y/go-config-secrets"
)

type sec struct {
	root *os.Root
}

func New(dir string) (config_secrets.SecretManager, error) {
	root, err := os.OpenRoot(dir)
	if err != nil {
		return nil, err
	}

	result := &sec{
		root: root,
	}

	return result, nil
}

func (s *sec) Fetch(ctx context.Context, ref string) (string, error) {
	b, err := s.root.ReadFile(ref)
	if err != nil {
		return "", config_secrets.NewSecretError(ref, err)
	}

	return string(bytes.TrimSpace(b)), nil
}
