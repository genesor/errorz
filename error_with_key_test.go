package errorz

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestErrorWithKey(t *testing.T) {
	t.Run("AsErrorz", func(t *testing.T) {
		cause := errors.New("cause")
		err, isErrorz := AsErrorz(NewErrorWithKey("code", "key",
			http.StatusBadRequest, cause))
		require.True(t, isErrorz)

		assert.EqualError(t, err, "key: cause")
		assert.Equal(t, "key", err.Key)
		assert.Equal(t, "code", err.Code)
		assert.Equal(t, cause, err.Cause)
		assert.Equal(t, http.StatusBadRequest, err.HTTPCode)
	})

	t.Run("no code", func(t *testing.T) {
		err := NewErrorWithKey("", "error with key", http.StatusBadRequest, nil)
		require.EqualError(t, err, "error with key")

		res := errors.Is(err, &ErrorWithKey{})
		require.True(t, res)

		res = Is[ErrorWithKey](err)
		require.True(t, res)

		err2 := fmt.Errorf("wrap2: %w", fmt.Errorf("wrap1: %w", err))
		res = errors.Is(err2, &ErrorWithKey{})
		require.True(t, res)

		res = Is[ErrorWithKey](err2)
		require.True(t, res)

		err3, ok := As[ErrorWithKey](err)
		require.True(t, ok)
		require.Equal(t, "error with key", err3.Key)

		err4, ok := As[ErrorWithKey](err2)
		require.True(t, ok)
		require.Equal(t, "error with key", err4.Key)

		err5, ok := As[ErrorWithKey](errors.New("not_same"))
		require.False(t, ok)
		require.Nil(t, err5)

		err6 := errors.New("base error")
		err7 := NewErrorWithKey("", "error with key", http.StatusBadRequest, err6)
		require.EqualError(t, err7, "error with key: base error")

		require.True(t, errors.Is(err7, &ErrorWithKey{}))
		require.True(t, errors.Is(err7, &ErrorWithKey{}))
		require.True(t, Is[ErrorWithKey](err7))

		err8, ok := As[ErrorWithKey](err7)
		require.True(t, ok)
		require.Equal(t, "error with key", err8.Key)

		err9 := new(ErrorWithKey)
		require.True(t, errors.As(err7, err9))
		require.Equal(t, "error with key", err9.Key)
	})

	t.Run("with code", func(t *testing.T) {
		err, ok := As[ErrorWithKey](NewErrorWithKey("awesome_code", "error",
			http.StatusBadRequest, nil))
		require.True(t, ok)
		require.Equal(t, "awesome_code", err.Code)
	})
}
