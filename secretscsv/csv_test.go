package secretscsv_test

import (
	"context"
	"testing"

	"gotest.tools/v3/assert"

	"github.com/UiP9AV6Y/go-config-secrets/secretscsv"
)

func TestNew(t *testing.T) {
	tests := map[string]struct {
		wantError string
		wantValue string
	}{
		"key404": {
			wantError: "no such secret available",
		},
		"blank": {},
		"": {
			wantValue: "blank",
		},
		"secret": {
			wantValue: "value",
		},
	}

	for name, test := range tests {
		do := func(t *testing.T) {
			subject := secretscsv.New("testdata/secrets.csv", nil)

			got, err := subject.Fetch(context.TODO(), name)
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
