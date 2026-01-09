package promconfig

import (
	"github.com/prometheus/common/config"

	config_secrets "github.com/UiP9AV6Y/go-config-secrets"
)

// NewSecretProvider creates a new secret provider using the first non-empty
// value as input or nil if no meaningful data is available. The secret
// manager is optional and only used if ref is defined.
func NewSecretProvider(value, file, ref string, mgr config_secrets.SecretManager) config_secrets.SecretProvider {
	switch {
	case value != "":
		return config.NewInlineSecret(value)
	case file != "":
		return config.NewFileSecret(file)
	case ref != "":
		return config_secrets.NewRefSecret(ref, mgr)
	default:
		return nil
	}
}

// NewSecretSecretProvider calls [NewSecretProvider]
func NewSecretSecretProvider(value config.Secret, file, ref string, mgr config_secrets.SecretManager) config_secrets.SecretProvider {
	return NewSecretProvider(string(value), file, ref, mgr)
}
