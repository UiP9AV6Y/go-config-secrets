package secretsio

import (
	"context"
	"io"
)

// ReaderProvider is used to describe system which prepare a resource for reading.
// It is the responsibility of the caller to close the provided reader after use.
type ReaderProvider interface {
	Open(context.Context, string) (io.ReadCloser, error)
}

// ReaderProviderFunc is an adapter to allow the use of ordinary functions as [io.ReadCloser] provider.
// If f is a function with the appropriate signature, ReaderProviderFunc(f) is a ReaderProvider that calls f.
type ReaderProviderFunc func(context.Context, string) (io.ReadCloser, error)

// Open calls f(ctx, ref).
func (f ReaderProviderFunc) Open(ctx context.Context, ref string) (io.ReadCloser, error) {
	return f(ctx, ref)
}
