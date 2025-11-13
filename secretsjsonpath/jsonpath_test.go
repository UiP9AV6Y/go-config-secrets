package secretsjsonpath_test

import (
	"context"
	"testing"

	"gotest.tools/v3/assert"

	"github.com/UiP9AV6Y/go-config-secrets/secretsio"
	"github.com/UiP9AV6Y/go-config-secrets/secretsjsonpath"
)

func TestNew(t *testing.T) {
	tests := map[string]struct {
		haveQuery string

		wantError string
		wantValue string
	}{
		"404.json": {
			haveQuery: "$.value",
			wantError: "no such file or directory",
		},
		"blank.txt": {
			haveQuery: "$.value",
			wantError: "EOF",
		},
		"empty.json": {
			haveQuery: "$.value",
			wantError: "JSON path query yielded 0 results",
		},
		"too_many.json": {
			haveQuery: "$.secrets[*].value",
			wantError: "JSON path query yielded 2 results",
		},
		"malformed.json": {
			haveQuery: "$.value",
			wantError: `invalid character 'b' after object key`,
		},
		"float64.json": {
			haveQuery: "$.value",
			wantError: "JSON path query yielded float64 instead of string",
		},
		"string.json": {
			haveQuery: "$.value",
			wantValue: "test",
		},
	}

	fixtures, err := secretsio.NewDirReaderProvider("testdata")
	assert.NilError(t, err)

	for name, test := range tests {
		do := func(t *testing.T) {
			subject, err := secretsjsonpath.New(test.haveQuery, fixtures)
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
