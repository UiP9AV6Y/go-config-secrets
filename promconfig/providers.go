package promconfig

import (
	"github.com/prometheus/common/config"

	config_secrets "github.com/UiP9AV6Y/go-config-secrets"
)

// SecretProviders describes various methods of retrieving a secret;
// only the first non-empty field is used as source.
type SecretProviders struct {
	Value config.Secret `yaml:"value,omitempty" json:"value,omitempty"`
	File  string        `yaml:"file,omitempty"  json:"file,omitempty"`
	Ref   string        `yaml:"ref,omitempty"   json:"ref,omitempty"`
}

// IsZero returns true if p is nil or none of the fields have been set.
func (p *SecretProviders) IsZero() bool {
	return p == nil || (p.Value == "" && p.File == "" && p.Ref == "")
}

func (p *SecretProviders) SetDirectory(dir string) {
	if p == nil {
		return
	}

	p.File = config.JoinDir(dir, p.File)
}

// SecretProvider calls [NewSecretSecretProvider] with the internal values
func (p *SecretProviders) SecretProvider(mgr config_secrets.SecretManager) config_secrets.SecretProvider {
	if p == nil {
		return nil
	}

	return NewSecretSecretProvider(p.Value, p.File, p.Ref, mgr)
}
