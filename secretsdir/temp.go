package secretsdir

import (
	"os"

	config_secrets "github.com/UiP9AV6Y/go-config-secrets"
)

func NewTemp() (config_secrets.SecretManager, error) {
	return New(os.TempDir())
}
