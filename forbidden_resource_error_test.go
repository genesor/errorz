// Code generated by http://github.com/genesor/errorz (v1.0.0). DO NOT EDIT.

package errorz

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestForbiddenResourceError(t *testing.T) {
	t.Run("AsErrorz", func(t *testing.T) {
		err, isErrorz := AsErrorz(NewForbiddenResourceError("code", "key"))
		require.True(t, isErrorz)

		assert.EqualError(t, err, "key")
		assert.Equal(t, "key", err.Key)
		assert.Equal(t, "code", err.Code)
		assert.Nil(t, err.Cause)
		assert.Equal(t, http.StatusForbidden, err.HTTPCode)
	})

	t.Run("no code", func(t *testing.T) {
		err := NewForbiddenResourceError("", "some error content")
		require.EqualError(t, err, "some error content")

		res := errors.Is(err, &ForbiddenResourceError{})
		require.True(t, res)

		err2 := fmt.Errorf("wrap2: %w", fmt.Errorf("wrap1: %w", err))
		res = errors.Is(err2, &ForbiddenResourceError{})
		require.True(t, res)

		res = IsForbiddenResourceError(err2)
		require.True(t, res)

		err3, ok := AsForbiddenResourceError(err)
		require.True(t, ok)
		require.Equal(t, "some error content", err3.Key)

		err4, ok := AsForbiddenResourceError(err2)
		require.True(t, ok)
		require.Equal(t, "some error content", err4.Key)

		err5, ok := AsForbiddenResourceError(fmt.Errorf("not a ForbiddenResourceError"))
		require.False(t, ok)
		require.Nil(t, err5)

		err6 := errors.New("base error")
		err7 := WrapWithForbiddenResourceError(err6, "", "some error content")
		require.EqualError(t, err7, "some error content: base error")

		require.True(t, errors.Is(err7, &ForbiddenResourceError{}))
		require.True(t, errors.Is(err7, &ForbiddenResourceError{}))
		require.True(t, IsForbiddenResourceError(err7))

		err8, ok := AsForbiddenResourceError(err7)
		require.True(t, ok)
		require.Equal(t, "some error content", err8.Key)

		err9 := new(ForbiddenResourceError)
		require.True(t, errors.As(err7, err9))
		require.Equal(t, "some error content", err9.Key)
	})

	t.Run("with code", func(t *testing.T) {
		err, ok := AsForbiddenResourceError(NewForbiddenResourceError("awesome_code", "error"))
		require.True(t, ok)
		require.Equal(t, "awesome_code", err.Code)
	})

	t.Run("with formatting", func(t *testing.T) {
		err, ok := AsForbiddenResourceError(NewForbiddenResourceErrorf("", "error %d", 1))
		require.True(t, ok)
		require.Empty(t, err.Code)
		require.Equal(t, "error 1", err.Key)

		err2, ok := AsForbiddenResourceError(WrapWithForbiddenResourceErrorf(err, "", "error %d", 2))
		require.True(t, ok)
		require.Empty(t, err2.Code)
		require.Equal(t, "error 2", err2.Key)
		require.Equal(t, err, err2.Cause)
	})
}
