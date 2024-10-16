// Code generated by http://github.com/genesor/errorz (v1.0.0). DO NOT EDIT.

package errorz

import (
	"fmt"
	"net/http"
)

// ForbiddenResourceError is used when the resource cannot be accessed by authenticated user.
type ForbiddenResourceError struct {
	*ErrorWithKey
}

// NewForbiddenResourceError is used when the resource cannot be accessed by authenticated user.
func NewForbiddenResourceError(code, key string) error {
	return WrapWithForbiddenResourceError(nil, code, key)
}

// NewForbiddenResourceErrorf is used when the resource cannot be accessed by authenticated user.
func NewForbiddenResourceErrorf(code, key string, args ...any) error {
	return WrapWithForbiddenResourceErrorf(nil, code, key, args...)
}

// WrapWithForbiddenResourceError is used when the resource cannot be accessed by authenticated user.
func WrapWithForbiddenResourceError(err error, code, key string) error {
	return NewError[ForbiddenResourceError](code, key, http.StatusForbidden, err)
}

// WrapWithForbiddenResourceErrorf is used when the resource cannot be accessed by authenticated user.
func WrapWithForbiddenResourceErrorf(err error, code, key string, args ...any) error {
	return NewError[ForbiddenResourceError](code, fmt.Sprintf(key, args...), http.StatusForbidden, err)
}

// IsForbiddenResourceError identifies an error as an ForbiddenResourceError.
func IsForbiddenResourceError(err error) bool {
	return Is[ForbiddenResourceError](err)
}

// AsForbiddenResourceError tries to cast an error as an ForbiddenResourceError.
func AsForbiddenResourceError(err error) (*ForbiddenResourceError, bool) {
	return As[ForbiddenResourceError](err)
}

// Is is used by the standard "errors" package to identify an error as ForbiddenResourceError.
func (e *ForbiddenResourceError) Is(err error) bool {
	_, ok := err.(*ForbiddenResourceError)
	return ok
}

// As is used by the standard "errors" package to identify an error as ForbiddenResourceError.
func (e *ForbiddenResourceError) As(err interface{}) bool {
	err2, ok := err.(*ForbiddenResourceError)
	if ok {
		*err2 = *e
	}
	return ok
}

// Unwrap is used by the standard "errors" package to dive into the error chain.
func (e *ForbiddenResourceError) Unwrap() error {
	return e.ErrorWithKey
}
