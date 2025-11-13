package main

import (
	"log"

	"github.com/prometheus/common/config"

	"github.com/UiP9AV6Y/go-config-secrets/secretsmap"
)

func main() {
	const ref = "example"

	sec := map[string]string{
		ref: "secret-authentication-bearer-token",
	}
	mgr := secretsmap.New(sec)

	auth := &config.Authorization{
		CredentialsRef: ref,
	}
	cfg := config.HTTPClientConfig{
		Authorization: auth,
	}

	_, err := config.NewClientFromConfig(cfg, "config-secrets-accept",
		config.WithSecretManager(mgr))
	if err != nil {
		log.Fatalf("unable to create HTTP client: %w", err)
	}

	log.Println("HTTP client with authorization is ready for use")
}
