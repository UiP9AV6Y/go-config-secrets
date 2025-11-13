package secretsdir

import (
	"errors"
	"os"

	config_secrets "github.com/UiP9AV6Y/go-config-secrets"
)

var (
	ErrContainerCredentialsDirectoryNotDefined = errors.New("none of the known container secret directories exists")

	ContainerCredentialsDirectories = []string{
		"/run/secrets",
		"/var/run/secrets/kubernetes.io",
		"/usr/share/rhel/secrets",
	}
)

func NewContainer() (config_secrets.SecretManager, error) {
	for _, dir := range ContainerCredentialsDirectories {
		if fi, err := os.Stat(dir); err == nil && fi.IsDir() {
			return New(dir)
		}
	}

	return nil, ErrContainerCredentialsDirectoryNotDefined
}
