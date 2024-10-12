package errorz

// ErrorWithKey is a struct compatible error to be extended or used in other errors
type ErrorWithKey struct {
	Key      string
	HTTPCode int
	Code     string
	Cause    error
}

// NewErrorWithKey creates a new ErrorWithKey
func NewErrorWithKey(code, key string, httpCode int, cause error) *ErrorWithKey {
	return &ErrorWithKey{
		Key:      key,
		HTTPCode: httpCode,
		Code:     code,
		Cause:    cause,
	}
}

// Error ...
func (e ErrorWithKey) Error() string {
	s := e.Key
	if e.Cause != nil {
		s += ": " + e.Cause.Error()
	}
	return s
}

// Is is used by the standard "errors" package to identify an error as ErrorWithKey
func (e *ErrorWithKey) Is(err error) bool {
	_, ok := err.(*ErrorWithKey)
	return ok
}

// As is used by the standard "errors" package to identify an error as ErrorWithKey
func (e *ErrorWithKey) As(err interface{}) bool {
	err2, ok := err.(*ErrorWithKey)
	if ok {
		*err2 = *e
	}
	return ok
}
