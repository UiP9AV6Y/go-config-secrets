package promconfig

import (
	"github.com/prometheus/common/config"

	config_secrets "github.com/UiP9AV6Y/go-config-secrets"
)

// NewAuthorizationCredentials calls [NewSecretSecretProvider] using the provided config source.
// Only credentials related fields are used.
func NewAuthorizationCredentials(a *config.Authorization, mgr config_secrets.SecretManager) config_secrets.SecretProvider {
	if a == nil {
		return nil
	}

	return NewSecretSecretProvider(a.Credentials, a.CredentialsFile, a.CredentialsRef, mgr)
}

// NewBasicAuthUsername calls [NewSecretProvider] using the provided config source.
// Only username related fields are used.
func NewBasicAuthUsername(b *config.BasicAuth, mgr config_secrets.SecretManager) config_secrets.SecretProvider {
	if b == nil {
		return nil
	}

	return NewSecretProvider(b.Username, b.UsernameFile, b.UsernameRef, mgr)
}

// NewBasicAuthPassword calls [NewSecretSecretProvider] using the provided config source.
// Only password related fields are used.
func NewBasicAuthPassword(b *config.BasicAuth, mgr config_secrets.SecretManager) config_secrets.SecretProvider {
	if b == nil {
		return nil
	}

	return NewSecretSecretProvider(b.Password, b.PasswordFile, b.PasswordRef, mgr)
}

// NewOAuth2ClientSecret calls [NewSecretSecretProvider] using the provided config source.
// Only client secret related fields are used.
func NewOAuth2ClientSecret(o *config.OAuth2, mgr config_secrets.SecretManager) config_secrets.SecretProvider {
	if o == nil {
		return nil
	}

	return NewSecretSecretProvider(o.ClientSecret, o.ClientSecretFile, o.ClientSecretRef, mgr)
}

// NewOAuth2ClientCertificateKey calls [NewSecretSecretProvider] using the provided config source.
// Only client certificate key related fields are used.
func NewOAuth2ClientCertificateKey(o *config.OAuth2, mgr config_secrets.SecretManager) config_secrets.SecretProvider {
	if o == nil {
		return nil
	}

	return NewSecretSecretProvider(o.ClientCertificateKey, o.ClientCertificateKeyFile, o.ClientCertificateKeyRef, mgr)
}
