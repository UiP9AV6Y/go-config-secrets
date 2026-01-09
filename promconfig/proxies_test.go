package promconfig_test

import (
	"context"
	"testing"

	"github.com/prometheus/common/config"
	"gotest.tools/v3/assert"

	"github.com/UiP9AV6Y/go-config-secrets/promconfig"
	"github.com/UiP9AV6Y/go-config-secrets/secretsio"
)

func TestNewAuthorizationCredentials(t *testing.T) {
	tests := map[string]struct {
		have *config.Authorization

		wantProvider bool
		wantSecret   string
	}{
		"all": {
			have: &config.Authorization{
				Credentials:     config.Secret("inline-value"),
				CredentialsFile: "file02",
				CredentialsRef:  "file03",
			},
			wantProvider: true,
			wantSecret:   "inline-value",
		},
		"file": {
			have: &config.Authorization{
				CredentialsFile: "file01",
			},
			wantProvider: true,
			wantSecret:   "file-one",
		},
		"nil": {},
		"ref": {
			have: &config.Authorization{
				CredentialsRef: "file02",
			},
			wantProvider: true,
			wantSecret:   "file-two",
		},
		"value": {
			have: &config.Authorization{
				Credentials: config.Secret("inline-value"),
			},
			wantProvider: true,
			wantSecret:   "inline-value",
		},
	}

	haveProvider, err := secretsio.NewDirReaderProvider("testdata")
	assert.NilError(t, err)
	haveMgr := secretsio.New(haveProvider)

	for name, test := range tests {
		do := func(t *testing.T) {
			if test.have != nil {
				test.have.SetDirectory("testdata")
			}

			gotProvider := promconfig.NewAuthorizationCredentials(test.have, haveMgr)
			if !test.wantProvider {
				assert.Assert(t, gotProvider == nil)
				return
			}

			assert.Assert(t, gotProvider != nil)
			gotSecret, gotErr := gotProvider.Fetch(context.TODO())
			assert.NilError(t, gotErr)
			assert.Equal(t, gotSecret, test.wantSecret)
		}

		t.Run(name, do)
	}
}

func TestNewBasicAuthUsername(t *testing.T) {
	tests := map[string]struct {
		have *config.BasicAuth

		wantProvider bool
		wantSecret   string
	}{
		"all": {
			have: &config.BasicAuth{
				Username:     "inline-value",
				UsernameFile: "file02",
				UsernameRef:  "file03",
			},
			wantProvider: true,
			wantSecret:   "inline-value",
		},
		"file": {
			have: &config.BasicAuth{
				UsernameFile: "file01",
			},
			wantProvider: true,
			wantSecret:   "file-one",
		},
		"nil": {},
		"ref": {
			have: &config.BasicAuth{
				UsernameRef: "file02",
			},
			wantProvider: true,
			wantSecret:   "file-two",
		},
		"value": {
			have: &config.BasicAuth{
				Username: "inline-value",
			},
			wantProvider: true,
			wantSecret:   "inline-value",
		},
	}

	haveProvider, err := secretsio.NewDirReaderProvider("testdata")
	assert.NilError(t, err)
	haveMgr := secretsio.New(haveProvider)

	for name, test := range tests {
		do := func(t *testing.T) {
			if test.have != nil {
				test.have.SetDirectory("testdata")
			}

			gotProvider := promconfig.NewBasicAuthUsername(test.have, haveMgr)
			if !test.wantProvider {
				assert.Assert(t, gotProvider == nil)
				return
			}

			assert.Assert(t, gotProvider != nil)
			gotSecret, gotErr := gotProvider.Fetch(context.TODO())
			assert.NilError(t, gotErr)
			assert.Equal(t, gotSecret, test.wantSecret)
		}

		t.Run(name, do)
	}
}

func TestNewBasicAuthPassword(t *testing.T) {
	tests := map[string]struct {
		have *config.BasicAuth

		wantProvider bool
		wantSecret   string
	}{
		"all": {
			have: &config.BasicAuth{
				Password:     config.Secret("inline-value"),
				PasswordFile: "file02",
				PasswordRef:  "file03",
			},
			wantProvider: true,
			wantSecret:   "inline-value",
		},
		"file": {
			have: &config.BasicAuth{
				PasswordFile: "file01",
			},
			wantProvider: true,
			wantSecret:   "file-one",
		},
		"nil": {},
		"ref": {
			have: &config.BasicAuth{
				PasswordRef: "file02",
			},
			wantProvider: true,
			wantSecret:   "file-two",
		},
		"value": {
			have: &config.BasicAuth{
				Password: config.Secret("inline-value"),
			},
			wantProvider: true,
			wantSecret:   "inline-value",
		},
	}

	haveProvider, err := secretsio.NewDirReaderProvider("testdata")
	assert.NilError(t, err)
	haveMgr := secretsio.New(haveProvider)

	for name, test := range tests {
		do := func(t *testing.T) {
			if test.have != nil {
				test.have.SetDirectory("testdata")
			}

			gotProvider := promconfig.NewBasicAuthPassword(test.have, haveMgr)
			if !test.wantProvider {
				assert.Assert(t, gotProvider == nil)
				return
			}

			assert.Assert(t, gotProvider != nil)
			gotSecret, gotErr := gotProvider.Fetch(context.TODO())
			assert.NilError(t, gotErr)
			assert.Equal(t, gotSecret, test.wantSecret)
		}

		t.Run(name, do)
	}
}

func TestNewOAuth2ClientSecret(t *testing.T) {
	tests := map[string]struct {
		have *config.OAuth2

		wantProvider bool
		wantSecret   string
	}{
		"all": {
			have: &config.OAuth2{
				ClientSecret:     config.Secret("inline-value"),
				ClientSecretFile: "file02",
				ClientSecretRef:  "file03",
			},
			wantProvider: true,
			wantSecret:   "inline-value",
		},
		"file": {
			have: &config.OAuth2{
				ClientSecretFile: "file01",
			},
			wantProvider: true,
			wantSecret:   "file-one",
		},
		"nil": {},
		"ref": {
			have: &config.OAuth2{
				ClientSecretRef: "file02",
			},
			wantProvider: true,
			wantSecret:   "file-two",
		},
		"value": {
			have: &config.OAuth2{
				ClientSecret: config.Secret("inline-value"),
			},
			wantProvider: true,
			wantSecret:   "inline-value",
		},
	}

	haveProvider, err := secretsio.NewDirReaderProvider("testdata")
	assert.NilError(t, err)
	haveMgr := secretsio.New(haveProvider)

	for name, test := range tests {
		do := func(t *testing.T) {
			if test.have != nil {
				test.have.SetDirectory("testdata")
			}

			gotProvider := promconfig.NewOAuth2ClientSecret(test.have, haveMgr)
			if !test.wantProvider {
				assert.Assert(t, gotProvider == nil)
				return
			}

			assert.Assert(t, gotProvider != nil)
			gotSecret, gotErr := gotProvider.Fetch(context.TODO())
			assert.NilError(t, gotErr)
			assert.Equal(t, gotSecret, test.wantSecret)
		}

		t.Run(name, do)
	}
}

func TestNewOAuth2ClientCertificateKey(t *testing.T) {
	tests := map[string]struct {
		have *config.OAuth2

		wantProvider bool
		wantSecret   string
	}{
		"all": {
			have: &config.OAuth2{
				ClientCertificateKey:     config.Secret("inline-value"),
				ClientCertificateKeyFile: "file02",
				ClientCertificateKeyRef:  "file03",
			},
			wantProvider: true,
			wantSecret:   "inline-value",
		},
		"file": {
			have: &config.OAuth2{
				ClientCertificateKeyFile: "file01",
			},
			wantProvider: true,
			wantSecret:   "file-one",
		},
		"nil": {},
		"ref": {
			have: &config.OAuth2{
				ClientCertificateKeyRef: "file02",
			},
			wantProvider: true,
			wantSecret:   "file-two",
		},
		"value": {
			have: &config.OAuth2{
				ClientCertificateKey: config.Secret("inline-value"),
			},
			wantProvider: true,
			wantSecret:   "inline-value",
		},
	}

	haveProvider, err := secretsio.NewDirReaderProvider("testdata")
	assert.NilError(t, err)
	haveMgr := secretsio.New(haveProvider)

	for name, test := range tests {
		do := func(t *testing.T) {
			if test.have != nil {
				test.have.SetDirectory("testdata")
			}

			gotProvider := promconfig.NewOAuth2ClientCertificateKey(test.have, haveMgr)
			if !test.wantProvider {
				assert.Assert(t, gotProvider == nil)
				return
			}

			assert.Assert(t, gotProvider != nil)
			gotSecret, gotErr := gotProvider.Fetch(context.TODO())
			assert.NilError(t, gotErr)
			assert.Equal(t, gotSecret, test.wantSecret)
		}

		t.Run(name, do)
	}
}
