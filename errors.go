package errorz

import (
	errors "github.com/pkg/errors"
)

// An errorz compatible error:
// - Embeds *ErrorWithKey or is ErrorWithKey directly
// - Unwraps the ErrorWithKey if it is embedded
// - Implements `Is(err error) bool` and `As(err interface{}) bool`.

// Error is the type constraint that includes all errorz errors.
type Error interface {
	~struct{ *ErrorWithKey } | ErrorWithKey
}

// AsErrorz retrieves the underlying ErrorWithKey of an error when possible.
// Useful when the actual type does not matter (e.g. in tracing, logging, ...).
// Use `AsXXXError` / `As[X]` to retrieve the concrete error instead.
func AsErrorz(err error) (*ErrorWithKey, bool) {
	return As[ErrorWithKey](err)
}

// Is identifies an error as any errorz.Error.
// Will not succeed if `Is(err error) bool` is not implemented correctly.
func Is[E Error, PE interface {
	*E
	error
}](err error) bool {
	return errors.Is(err, PE(new(E)))
}

// As tries to cast an error as any errorz.Error.
// Will not succeed if `As(err interface{}) bool` is not implemented correctly.
func As[E Error](err error) (*E, bool) {
	err2 := new(E)
	if ok := errors.As(err, err2); !ok {
		return nil, false
	}
	return err2, true
}

// NewError creates an error that embeds an ErrorWithKey.
func NewError[E ~struct {
	*ErrorWithKey
}, PE interface {
	*E
	error
}](code, key string, httpCode int, err error) error {
	return errors.WithStack(PE(&E{
		ErrorWithKey: NewErrorWithKey(code, key, httpCode, err),
	}))
}
