package secretsio

import (
	"context"
	"io"
	"os"
)

// NewDirReaderProvider reads secret references from the filesystem
// under the given directory. Lookups cannot escape the boundary as
// internally an [os.Root] is used for boundary checks.
func NewDirReaderProvider(dir string) (ReaderProvider, error) {
	root, err := os.OpenRoot(dir)
	if err != nil {
		return nil, err
	}

	fetch := func(_ context.Context, ref string) (io.ReadCloser, error) {
		return root.Open(ref)
	}

	return ReaderProviderFunc(fetch), nil
}
