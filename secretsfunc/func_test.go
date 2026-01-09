package secretsfunc_test

import (
	"context"
	"testing"

	"gotest.tools/v3/assert"

	"github.com/UiP9AV6Y/go-config-secrets/secretsfunc"
)

func TestNew(t *testing.T) {
	tests := map[string]struct {
		haveRef string

		wantError string
		wantValue string
	}{
		"404": {
			haveRef:   "404",
			wantError: `secret lookup did not yield a value for "TEST_404"`,
		},
		"blank": {
			wantValue: "value02",
		},
		"regular usecase": {
			haveRef:   "SECRET_KEY",
			wantValue: "value01",
		},
	}

	t.Setenv("TEST_SECRET_KEY", "value01")
	t.Setenv("TEST_", "value02")

	for name, test := range tests {
		do := func(t *testing.T) {
			subject := secretsfunc.NewEnv("TEST_")

			got, err := subject.Fetch(context.TODO(), test.haveRef)
			if test.wantError != "" {
				assert.ErrorContains(t, err, test.wantError)
			} else {
				assert.NilError(t, err)
				assert.Equal(t, got, test.wantValue)
			}
		}

		t.Run(name, do)
	}
}
