package secretsdir

import (
	"errors"
	"os"

	config_secrets "github.com/UiP9AV6Y/go-config-secrets"
)

const (
	// EnvSystemdCredentialsDirectory is the environment variable
	// containing the filesystem location for credentials provided
	// by Systemd as documented in https://systemd.io/CREDENTIALS/
	EnvSystemdCredentialsDirectory = "CREDENTIALS_DIRECTORY"
)

var (
	ErrSystemdCredentialsDirectoryNotDefined = errors.New("environment variable CREDENTIALS_DIRECTORY is not defined or empty")
)

func NewSystemd() (config_secrets.SecretManager, error) {
	dir := os.Getenv(EnvSystemdCredentialsDirectory)
	if dir == "" {
		return nil, ErrSystemdCredentialsDirectoryNotDefined
	}

	return New(dir)
}
