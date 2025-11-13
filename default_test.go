package config_secrets_test

import (
	"context"
	"testing"

	"gotest.tools/v3/assert"

	config_secrets "github.com/UiP9AV6Y/go-config-secrets"
)

func TestDefault(t *testing.T) {
	subject := config_secrets.Default()
	assert.Assert(t, subject != nil)

	haveRef := "TEST_KEY"
	wantValue := "value"
	t.Setenv(config_secrets.EnvPrefix+haveRef, wantValue)

	got, err := subject.Fetch(context.TODO(), haveRef)
	assert.NilError(t, err)
	assert.Equal(t, got, wantValue)

	got, err = subject.Fetch(context.TODO(), haveRef+haveRef)
	assert.Assert(t, err != nil)
	assert.Equal(t, got, "")
}
