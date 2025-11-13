package secretsio

import (
	"context"
	"io"
	"net/http"
	"strings"
)

// NewHTTPReaderProvider creates a reader provider which yields the HTTP response body
// of a remote resource as reader source. The provided URL template will be processed
// with each retrieval operation by replacing all instances of `$REF` with the requested
// secret reference. If the provided client is nil, [http.DefaultClient] will be used.
func NewHTTPReaderProvider(urlTpl string, c *http.Client) ReaderProvider {
	var client *http.Client
	if c != nil {
		client = c
	} else {
		client = http.DefaultClient
	}

	fetch := func(ctx context.Context, ref string) (io.ReadCloser, error) {
		url := strings.ReplaceAll(urlTpl, "$REF", ref)
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return nil, err
		}

		res, err := client.Do(req)
		if err != nil {
			return nil, err
		}

		return res.Body, nil
	}

	return ReaderProviderFunc(fetch)
}
