package secretsjsonpath

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/theory/jsonpath"

	config_secrets "github.com/UiP9AV6Y/go-config-secrets"
	"github.com/UiP9AV6Y/go-config-secrets/secretsio"
)

type sec struct {
	rdr  secretsio.ReaderProvider
	path *jsonpath.Path
}

func New(query string, r secretsio.ReaderProvider) (config_secrets.SecretManager, error) {
	p, err := jsonpath.Parse(query)
	if err != nil {
		return nil, err
	}

	result := &sec{
		rdr:  r,
		path: p,
	}

	return result, nil
}

func (s *sec) Fetch(ctx context.Context, ref string) (string, error) {
	rc, err := s.rdr.Open(ctx, ref)
	if err != nil {
		return "", err
	}
	defer rc.Close()

	var v any
	dec := json.NewDecoder(rc)
	err = dec.Decode(&v)
	if err != nil {
		return "", err
	}

	nodes := s.path.Select(v)
	if l := len(nodes); l != 1 {
		return "", fmt.Errorf("JSON path query yielded %d results", l)
	}

	val, ok := nodes[0].(string)
	if !ok {
		return "", fmt.Errorf("JSON path query yielded %T instead of string", nodes[0])
	}

	return val, nil
}
