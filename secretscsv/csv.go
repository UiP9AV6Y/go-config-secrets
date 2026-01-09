package secretscsv

import (
	"context"
	"encoding/csv"
	"io"
	"os"

	config_secrets "github.com/UiP9AV6Y/go-config-secrets"
	"github.com/UiP9AV6Y/go-config-secrets/secretsio"
)

// Options affect parsing and extraction of information
// from parsed CSV data.
type Options struct {
	// Column index containing secret keys/references.
	// Numbering starts at 1.
	RefFieldColumn int
	// Column index containing secret values.
	// Numbering starts at 1.
	ValueFieldColumn int

	// Remove leding spaces from fields.
	// See [csv#Reader] for details.
	TrimLeadingSpace bool

	// Comma is the field delimiter.
	// See [csv#Reader] for details.
	Comma rune
	// Comment, if not 0, is the comment character.
	// See [csv#Reader] for details.
	Comment rune
}

// ColumnIndices returns the slice indices for [csv.Reader#Read] results;
// secret ref and value respectively.
func (o *Options) ColumnIndices() (ki int, vi int) {
	ki = 0
	vi = 1

	if o == nil {
		return
	}

	if o.RefFieldColumn > 0 {
		ki = o.RefFieldColumn - 1
	}

	if o.ValueFieldColumn > 0 {
		vi = o.ValueFieldColumn - 1
	}

	return
}

// ApplyReader modified the provided reader according to previously
// configured options.
func (o *Options) ApplyReader(r *csv.Reader) {
	if o == nil {
		return
	}

	if o.Comma != 0 {
		r.Comma = o.Comma
	}

	if o.Comment != 0 {
		r.Comment = o.Comment
	}

	r.TrimLeadingSpace = o.TrimLeadingSpace

	return
}

type sec struct {
	rdr  secretsio.ReaderProvider
	opts *Options
}

// New creates a [config_secrets.SecretManager] instance using the
// provided CSV as backing source. The file is parsed with each
// request. Optional options affect the parsing and extraction behaviour.
func New(path string, opts *Options) config_secrets.SecretManager {
	fetch := func(_ context.Context, ref string) (string, error) {
		rc, err := os.Open(path)
		if err != nil {
			return "", config_secrets.NewSecretError(ref, err)
		}
		defer rc.Close()

		rdr := csv.NewReader(rc)
		ki, vi := opts.ColumnIndices()

		opts.ApplyReader(rdr)
		rdr.ReuseRecord = true

		for {
			record, err := rdr.Read()
			if err == io.EOF {
				break
			}

			if err != nil {
				return "", config_secrets.NewSecretError(ref, err)
			} else if ki < len(record) && vi < len(record) && record[ki] == ref {
				return record[vi], nil
			}
		}

		return "", config_secrets.NewSecretNotFoundError(ref)
	}

	return config_secrets.SecretManagerFunc(fetch)
}
