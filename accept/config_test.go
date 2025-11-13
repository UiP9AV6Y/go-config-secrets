package accept_test

import (
	"testing"

	"github.com/prometheus/common/config"
	"gotest.tools/v3/assert"

	config_secrets "github.com/UiP9AV6Y/go-config-secrets"
)

func TestSecretManager(t *testing.T) {
	var sec config.SecretManager = config_secrets.Default()

	assert.Assert(t, sec != nil)
}

func TestSecretReader(t *testing.T) {
	var r config.SecretReader = config_secrets.NewRefSecret("test", config_secrets.Default())

	assert.Assert(t, r != nil)
}
