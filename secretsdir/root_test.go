package secretsdir_test

import (
	"context"
	"testing"

	"gotest.tools/v3/assert"

	"github.com/UiP9AV6Y/go-config-secrets/secretsdir"
)

func TestNew(t *testing.T) {
	tests := map[string]struct {
		wantError string
		wantValue string
	}{
		"404.json": {
			wantError: "no such file or directory",
		},
		"blank.txt": {},
		"secret.txt": {
			wantValue: "test",
		},
	}

	for name, test := range tests {
		do := func(t *testing.T) {
			subject, err := secretsdir.New("testdata")
			assert.NilError(t, err)

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
