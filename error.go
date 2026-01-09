package config_secrets

import (
	"errors"
)

// ErrNotFound is a generic error indicating a lookup error
var ErrNotFound = errors.New("no such secret available")

// SecretError records an error and the associated secret lookup that caused it.
type SecretError struct {
	Ref string // Lookup key causing the error
	Err error  // The actual error
}

// NewSecretError is a convenience function for
// creating [SecretError] instances.
func NewSecretError(ref string, err error) *SecretError {
	result := &SecretError{
		Ref: ref,
		Err: err,
	}

	return result
}

// NewSecretNotFoundError calls [NewSecretError] with the provided
// lookup key and [ErrNotFound]
func NewSecretNotFoundError(ref string) *SecretError {
	return NewSecretError(ref, ErrNotFound)
}

func (e *SecretError) Error() string { return e.Ref + ": " + e.Err.Error() }

func (e *SecretError) Unwrap() error { return e.Err }
